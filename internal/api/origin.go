package api

import (
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	defaultUIPort = 1337
	maxAPIBody    = 32 << 20
)

func (s *Server) originOK(r *http.Request) bool {
	if r == nil {
		return false
	}
	if strings.EqualFold(r.Header.Get("Sec-Fetch-Site"), "cross-site") {
		return false
	}
	origin := strings.TrimSpace(r.Header.Get("Origin"))
	if origin == "" {
		return true
	}
	u, err := url.Parse(origin)
	if err != nil || u.Host == "" {
		return false
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return false
	}
	host := strings.ToLower(u.Hostname())
	port := u.Port()
	if port == "" {
		if u.Scheme == "https" {
			port = "443"
		} else {
			port = "80"
		}
	}
	wantHost, wantPort := "127.0.0.1", defaultUIPort
	if s != nil && s.Store != nil {
		wantHost, wantPort = s.Store.UIListen()
	}
	if port != strconv.Itoa(wantPort) {
		return false
	}
	if loopbackHost(host) {
		return true
	}
	wantHost = strings.ToLower(strings.TrimSpace(wantHost))
	if wantHost == "" || wantHost == "0.0.0.0" || wantHost == "::" || wantHost == "[::]" {
		return false
	}
	return host == wantHost
}

func loopbackHost(h string) bool {
	if h == "localhost" || h == "127.0.0.1" || h == "::1" {
		return true
	}
	ip := net.ParseIP(h)
	return ip != nil && ip.IsLoopback()
}

func (s *Server) guard(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api/") && !s.originOK(r) {
			writeErr(w, http.StatusForbidden, "origin not allowed")
			return
		}
		switch r.Method {
		case http.MethodPost, http.MethodPut, http.MethodPatch:
			r.Body = http.MaxBytesReader(w, r.Body, maxAPIBody)
		}
		next.ServeHTTP(w, r)
	})
}
