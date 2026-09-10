package desktop

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"time"
)

type browserCandidate struct {
	name string
	args []string
}

// Start opens the UI in a dedicated Chromium/Chrome window (no tab bar, no URL bar).
// Closing that window exits the returned process; the caller should then shut down AURA.
func Start(uiURL string) (*exec.Cmd, error) {
	binary, extra, err := findChromium()
	if err != nil {
		return nil, err
	}
	args := append(append([]string{}, extra...),
		"--app="+uiURL,
		"--user-data-dir="+profileDir(),
		"--no-first-run",
		"--no-default-browser-check",
		"--disable-sync",
		"--disable-features=Translate,MediaRouter",
		"--class=aura",
		"--window-size=1440,900",
		"--window-position=60,40",
	)
	cmd := exec.Command(binary, args...)
	cmd.Stdout = nil
	cmd.Stderr = nil
	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("open UI window: %w", err)
	}
	return cmd, nil
}

// WaitReady polls the UI until it responds or ctx is done.
func WaitReady(ctx context.Context, uiURL string) error {
	client := &http.Client{Timeout: 400 * time.Millisecond}
	ticker := time.NewTicker(80 * time.Millisecond)
	defer ticker.Stop()
	var last error
	for {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, uiURL, nil)
		if err != nil {
			return err
		}
		resp, err := client.Do(req)
		if err == nil {
			_ = resp.Body.Close()
			if resp.StatusCode < 500 {
				return nil
			}
			last = fmt.Errorf("ui status %d", resp.StatusCode)
		} else {
			last = err
		}
		select {
		case <-ctx.Done():
			if last == nil {
				last = ctx.Err()
			}
			return last
		case <-ticker.C:
		}
	}
}

func findChromium() (string, []string, error) {
	candidates := []browserCandidate{
		{"chromium", nil},
		{"chromium-browser", nil},
		{"google-chrome", nil},
		{"google-chrome-stable", nil},
		{"brave-browser", nil},
		{"microsoft-edge", nil},
		{"microsoft-edge-stable", nil},
	}
	if runtime.GOOS == "windows" {
		candidates = append(candidates,
			browserCandidate{`C:\Program Files\Google\Chrome\Application\chrome.exe`, nil},
			browserCandidate{`C:\Program Files (x86)\Google\Chrome\Application\chrome.exe`, nil},
			browserCandidate{`C:\Program Files\Microsoft\Edge\Application\msedge.exe`, nil},
		)
	}
	if runtime.GOOS == "darwin" {
		candidates = append([]browserCandidate{
			{"/Applications/Google Chrome.app/Contents/MacOS/Google Chrome", nil},
			{"/Applications/Chromium.app/Contents/MacOS/Chromium", nil},
		}, candidates...)
	}
	for _, c := range candidates {
		path, err := exec.LookPath(c.name)
		if err == nil {
			return path, c.args, nil
		}
		if strings.Contains(c.name, string(os.PathSeparator)) {
			if st, err := os.Stat(c.name); err == nil && !st.IsDir() {
				return c.name, c.args, nil
			}
		}
	}
	var names []string
	for _, c := range candidates {
		names = append(names, c.name)
	}
	return "", nil, fmt.Errorf("no Chromium-based browser found for app window (tried: %s)", strings.Join(names, ", "))
}

func profileDir() string {
	if d, err := os.UserCacheDir(); err == nil && d != "" {
		return filepath.Join(d, "aura", "ui-profile")
	}
	return filepath.Join(os.TempDir(), "aura-ui-profile")
}

// IsRunning reports whether an AURA UI is already serving uiURL.
func IsRunning(base string) bool {
	u, err := url.Parse(base)
	if err != nil {
		return false
	}
	u.Path = "/api/status"
	u.RawQuery = ""
	client := &http.Client{Timeout: 400 * time.Millisecond}
	resp, err := client.Get(u.String())
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}

// Open starts the desktop UI. On Kali/Linux with CGO it uses a GTK/WebKit window
// on the calling (main) thread. If that is unavailable, it falls back to Chromium --app.
func Open(ctx context.Context, title, uiURL string) error {
	if nativeAvailable() {
		if err := RunNative(ctx, title, uiURL); err == nil {
			return nil
		} else if ctx.Err() != nil {
			return ctx.Err()
		}
	}
	cmd, err := Start(uiURL)
	if err != nil {
		return err
	}
	done := make(chan error, 1)
	go func() { done <- cmd.Wait() }()
	select {
	case <-ctx.Done():
		if cmd.Process != nil {
			_ = cmd.Process.Signal(syscall.SIGTERM)
		}
		return ctx.Err()
	case err := <-done:
		return err
	}
}
