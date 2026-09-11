package api

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"go.uber.org/zap"
)

func (s *Server) proxyLaunchBrowser(w http.ResponseWriter, r *http.Request) {
	addr := s.Proxy.Addr()
	if addr == "" {
		writeErr(w, 400, "proxy is not running")
		return
	}

	spec, err := planBrowserLaunch(addr)
	if err != nil {
		s.Log.Warn("launch browser", zap.Error(err))
		writeErr(w, 500, err.Error())
		return
	}

	cmd := exec.Command(spec.binary, spec.args...)
	if err := cmd.Start(); err != nil {
		s.Log.Warn("launch browser", zap.Error(err))
		writeErr(w, 500, fmt.Sprintf("failed to launch browser: %v", err))
		return
	}
	_ = cmd.Process.Release()

	s.Log.Info("launched browser", zap.String("binary", spec.binary), zap.String("proxy", addr), zap.String("engine", spec.engine))
	writeJSON(w, 200, map[string]any{"ok": true, "binary": spec.binary, "proxy": addr, "engine": spec.engine})
}

type browserLaunch struct {
	binary  string
	args    []string
	engine  string
	profile string
}

func planBrowserLaunch(proxyAddr string) (browserLaunch, error) {
	binary, engine, err := findBrowser()
	if err != nil {
		return browserLaunch{}, err
	}
	profile, err := os.MkdirTemp("", "aura-browser-")
	if err != nil {
		return browserLaunch{}, fmt.Errorf("browser profile: %w", err)
	}
	if engine == "firefox" {
		host, port, err := splitProxyAddr(proxyAddr)
		if err != nil {
			return browserLaunch{}, err
		}
		if err := writeFirefoxProxyPrefs(profile, host, port); err != nil {
			return browserLaunch{}, err
		}
		return browserLaunch{binary: binary, args: firefoxArgs(profile), engine: engine, profile: profile}, nil
	}
	return browserLaunch{binary: binary, args: chromiumArgs(proxyAddr, profile), engine: engine, profile: profile}, nil
}

func chromiumArgs(proxyAddr, profile string) []string {
	return []string{
		"--proxy-server=http://" + proxyAddr,
		"--user-data-dir=" + profile,
		"--ignore-certificate-errors",
		"--no-first-run",
		"--no-default-browser-check",
	}
}

func firefoxArgs(profile string) []string {
	return []string{"-no-remote", "-profile", profile}
}

func writeFirefoxProxyPrefs(profile, host string, port int) error {
	if err := os.MkdirAll(profile, 0o700); err != nil {
		return err
	}
	js := fmt.Sprintf(
		"user_pref(\"network.proxy.type\", 1);\n"+
			"user_pref(\"network.proxy.http\", %q);\n"+
			"user_pref(\"network.proxy.http_port\", %d);\n"+
			"user_pref(\"network.proxy.ssl\", %q);\n"+
			"user_pref(\"network.proxy.ssl_port\", %d);\n"+
			"user_pref(\"network.proxy.share_proxy_settings\", true);\n"+
			"user_pref(\"network.proxy.no_proxies_on\", \"\");\n",
		host, port, host, port,
	)
	return os.WriteFile(filepath.Join(profile, "user.js"), []byte(js), 0o600)
}

func splitProxyAddr(addr string) (string, int, error) {
	host, portStr, err := net.SplitHostPort(addr)
	if err != nil {
		return "", 0, fmt.Errorf("proxy address: %w", err)
	}
	port, err := strconv.Atoi(portStr)
	if err != nil || port <= 0 || port > 65535 {
		return "", 0, fmt.Errorf("proxy port %q", portStr)
	}
	return host, port, nil
}

func findBrowser() (string, string, error) {
	type cand struct {
		name   string
		engine string
	}
	candidates := []cand{
		{"chromium", "chromium"},
		{"chromium-browser", "chromium"},
		{"google-chrome", "chromium"},
		{"google-chrome-stable", "chromium"},
		{"brave-browser", "chromium"},
		{"microsoft-edge", "chromium"},
		{"firefox", "firefox"},
	}
	if runtime.GOOS == "windows" {
		candidates = append(candidates,
			cand{`C:\Program Files\Google\Chrome\Application\chrome.exe`, "chromium"},
			cand{`C:\Program Files (x86)\Google\Chrome\Application\chrome.exe`, "chromium"},
		)
	}
	for _, c := range candidates {
		path, err := exec.LookPath(c.name)
		if err == nil {
			return path, c.engine, nil
		}
	}
	var names []string
	for _, c := range candidates {
		names = append(names, c.name)
	}
	return "", "", fmt.Errorf("no supported browser found (tried: %s)", strings.Join(names, ", "))
}
