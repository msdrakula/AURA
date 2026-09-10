package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"go.uber.org/zap"

	"meb/internal/batch"
	"meb/internal/callback"
	"meb/internal/models"
)

// BatchRunner executes a batch of HTTP requests with payload variations.
type BatchRunner interface {
	ExecuteAttack(ctx context.Context, template, scheme, target, attackType string, payloadSets [][]string) ([]models.BatchResult, error)
}

// BatchFactory creates a fresh batch engine with a per-request result logger.
type BatchFactory func(workers, rps int, onResult func(rawReq, respRaw string)) BatchRunner

// InteractionStore exposes stored callback interactions by request ID.
type InteractionStore interface {
	GetInteractions(requestID string) []callback.Interaction
}

// BatchExecuteRequest is the JSON body for POST /api/batch/execute.
type BatchExecuteRequest struct {
	TemplateRaw string     `json:"template_raw"`
	Scheme      string     `json:"scheme"`
	Target      string     `json:"target"`
	AttackType  string     `json:"attack_type"`
	PayloadSets [][]string `json:"payload_sets"`
	Workers     int        `json:"workers"`
	RPS         int        `json:"rps"`
}

func (s *Server) batchExecute(w http.ResponseWriter, r *http.Request) {
	if s.Batch == nil {
		writeErr(w, 503, "batch engine not configured")
		return
	}
	var body BatchExecuteRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		s.Log.Warn("batch: decode body", zap.Error(err))
		writeErr(w, 400, "invalid JSON: "+err.Error())
		return
	}
	if body.TemplateRaw == "" {
		writeErr(w, 400, "template_raw is required")
		return
	}
	if body.Scheme != "http" && body.Scheme != "https" {
		writeErr(w, 400, "scheme must be http or https")
		return
	}
	if body.AttackType == "" {
		body.AttackType = "combo"
	}

	ctx, cancel := context.WithTimeout(r.Context(), 60*time.Second)
	defer cancel()

	workers, rps := body.Workers, body.RPS
	if workers < 1 {
		workers = 10
	}
	if rps < 1 {
		rps = 50
	}
	if workers > 30 {
		workers = 30
	}
	if rps > 50 {
		rps = 50
	}

	// Use a fresh engine with logging if a factory is available, else the shared one.
	runner := s.Batch
	if s.BatchFactory != nil {
		runner = s.BatchFactory(workers, rps, func(reqRaw, respRaw string) {
			s.saveToolFlow(reqRaw, respRaw, body.Scheme, "intruder")
		})
	}
	results, err := runner.ExecuteAttack(ctx, body.TemplateRaw, body.Scheme, body.Target, body.AttackType, body.PayloadSets)
	if err != nil && ctx.Err() == nil {
		s.Log.Warn("batch: execute", zap.Error(err))
		writeErr(w, 400, err.Error())
		return
	}
	if results == nil {
		results = []models.BatchResult{}
	}
	writeJSON(w, 200, results)
}

func (s *Server) callbackInteractions(w http.ResponseWriter, r *http.Request) {
	if s.Callback == nil {
		writeErr(w, 503, "callback server not configured")
		return
	}
	requestID := r.PathValue("request_id")
	if requestID == "" {
		writeErr(w, 400, "request_id is required")
		return
	}
	items := s.Callback.GetInteractions(requestID)
	if items == nil {
		items = []callback.Interaction{}
	}
	writeJSON(w, 200, items)
}

// ensure batch.Engine satisfies BatchRunner at compile time.
var _ BatchRunner = (*batch.Engine)(nil)
