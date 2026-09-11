package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"go.uber.org/zap"

	"meb/internal/discover"
	"meb/internal/scanner"
)

// DiscoverRequest is the body for POST /api/discover/run.
type DiscoverRequest struct {
	BaseURL      string   `json:"base_url"`
	Wordlist     []string `json:"wordlist"`
	WordlistPath string   `json:"wordlist_path"`
	Workers      int      `json:"workers"`
	RPS          int      `json:"rps"`
	Cookies      string   `json:"cookies"`
	Authorized   bool     `json:"authorized"`
}

func (s *Server) discoverRun(w http.ResponseWriter, r *http.Request) {
	var body DiscoverRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeErr(w, 400, "invalid JSON: "+err.Error())
		return
	}
	if body.BaseURL == "" {
		writeErr(w, 400, "base_url is required")
		return
	}
	if rejectUnlessAuthorized(w, body.Authorized) {
		return
	}
	words, err := s.resolveWords(body.Wordlist, body.WordlistPath, "dirs")
	if err != nil {
		writeErr(w, 400, err.Error())
		return
	}
	if len(words) == 0 {
		writeErr(w, 400, "wordlist is required")
		return
	}
	if body.Workers > 30 {
		body.Workers = 30
	}
	if body.RPS > 50 {
		body.RPS = 50
	}
	ctx, cancel := context.WithTimeout(r.Context(), 30*time.Minute)
	defer cancel()
	var onResult func(string, string)
	if len(words) <= 80 {
		onResult = func(reqRaw, respRaw string) { s.saveToolFlow(reqRaw, respRaw, "", "discover") }
	}
	s.Log.Info("discover: run", zap.String("base_url", body.BaseURL), zap.Int("words", len(words)), zap.Int("workers", body.Workers))
	results, err := discover.Run(ctx, discover.Options{
		BaseURL: body.BaseURL, Wordlist: words,
		Workers: body.Workers, RPS: body.RPS, Cookies: body.Cookies,
		Timeout:  10 * time.Second,
		OnResult: onResult,
	}, s.Log)
	if err != nil && ctx.Err() == nil {
		s.Log.Warn("discover: run", zap.Error(err))
		writeErr(w, 500, err.Error())
		return
	}
	if results == nil {
		results = []discover.Result{}
	}
	writeJSON(w, 200, results)
}

// ScanRequest is the body for POST /api/scanner/scan.
type ScanRequest struct {
	Raw        string `json:"raw"`
	Scheme     string `json:"scheme"`
	Target     string `json:"target"`
	Authorized bool   `json:"authorized"`
}

func (s *Server) scannerRun(w http.ResponseWriter, r *http.Request) {
	var body ScanRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeErr(w, 400, "invalid JSON: "+err.Error())
		return
	}
	if strings.TrimSpace(body.Raw) == "" {
		writeErr(w, 400, "raw is required")
		return
	}
	if rejectUnlessAuthorized(w, body.Authorized) {
		return
	}
	if body.Scheme == "" {
		body.Scheme = "https"
	}
	// Log the baseline scan request.
	s.saveToolFlow(body.Raw, "", body.Scheme, "scanner")
	s.Log.Info("scanner: run", zap.String("scheme", body.Scheme), zap.String("target", body.Target))
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Minute)
	defer cancel()
	res, err := scanner.Scan(ctx, scanner.Options{
		Raw: body.Raw, Scheme: body.Scheme, Target: body.Target, Timeout: 15 * time.Second,
	})
	if ctx.Err() != nil {
		writeErr(w, 408, "scan stopped or timed out")
		return
	}
	if err != nil {
		s.Log.Warn("scanner: run", zap.Error(err))
		if strings.Contains(strings.ToLower(err.Error()), "baseline") {
			writeErr(w, 502, err.Error())
			return
		}
		writeErr(w, 400, err.Error())
		return
	}
	if res == nil {
		writeJSON(w, 200, map[string]any{"findings": []any{}})
		return
	}
	writeJSON(w, 200, res)
}
