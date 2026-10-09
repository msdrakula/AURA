package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"go.uber.org/zap"

	"meb/internal/scanner"
)

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
