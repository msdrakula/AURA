package api

import (
	"fmt"
	"net/http"
	"os/exec"
	"runtime"
	"strings"

	"go.uber.org/zap"
)

func (s *Server) proxyLaunchBrowser(w http.ResponseWriter, r *http.Request) {
	addr := s.Proxy.Addr()
	if addr == "" {
		writeErr(w, 400, "proxy is not running")
		return
	}

	binary, args, err := findBrowser()
	if err != nil {
		s.Log.Warn("launch browser", zap.Error(err))
		writeErr(w, 500, err.Error())
		return
	}

	args = append(args, "--proxy-server=http://"+addr,
		"--user-data-dir="+freshProfileDir(),
		"--ignore-certificate-errors",
		"--no-first-run",
		"--no-default-browser-check",
	)

	cmd := exec.Command(binary, args...)
	if err := cmd.Start(); err != nil {
		s.Log.Warn("launch browser", zap.Error(err))
		writeErr(w, 500, fmt.Sprintf("failed to launch browser: %v", err))
		return
	}
	_ = cmd.Process.Release()

	s.Log.Info("launched browser", zap.String("binary", binary), zap.String("proxy", addr))
	writeJSON(w, 200, map[string]any{"ok": true, "binary": binary, "proxy": addr})
}

type browserCandidate struct {
	name string
	args []string
}

func findBrowser() (string, []string, error) {
	candidates := []browserCandidate{
		{"chromium", nil},
		{"chromium-browser", nil},
		{"google-chrome", nil},
		{"google-chrome-stable", nil},
		{"brave-browser", nil},
		{"microsoft-edge", nil},
		{"firefox", []string{"--proxy-server"}},
	}
	if runtime.GOOS == "windows" {
		candidates = append(candidates,
			browserCandidate{`C:\Program Files\Google\Chrome\Application\chrome.exe`, nil},
			browserCandidate{`C:\Program Files (x86)\Google\Chrome\Application\chrome.exe`, nil},
		)
	}
	for _, c := range candidates {
		path, err := exec.LookPath(c.name)
		if err == nil {
			return path, c.args, nil
		}
	}
	var names []string
	for _, c := range candidates {
		names = append(names, c.name)
	}
	return "", nil, fmt.Errorf("no supported browser found (tried: %s)", strings.Join(names, ", "))
}

func freshProfileDir() string {
	return "/tmp/aura-browser-profile"
}
