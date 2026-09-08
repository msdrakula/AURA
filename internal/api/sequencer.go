package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"go.uber.org/zap"

	"meb/internal/sequencer"
)

// SequencerAnalyzeRequest is the body for POST /api/sequencer/analyze.
type SequencerAnalyzeRequest struct {
	Tokens string `json:"tokens"`
}

func (s *Server) sequencerAnalyze(w http.ResponseWriter, r *http.Request) {
	var body SequencerAnalyzeRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		s.Log.Warn("sequencer: decode body", zap.Error(err))
		writeErr(w, 400, "invalid JSON: "+err.Error())
		return
	}
	tokens := strings.Split(body.Tokens, "\n")
	clean := make([]string, 0, len(tokens))
	for _, t := range tokens {
		t = strings.TrimSpace(t)
		if t != "" {
			clean = append(clean, t)
		}
	}
	if len(clean) < 10 {
		writeErr(w, 400, "need at least 10 tokens for analysis")
		return
	}
	res := sequencer.Analyze(clean)
	writeJSON(w, 200, res)
}
