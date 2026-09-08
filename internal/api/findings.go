package api

import (
	"net/http"
	"strings"

	"meb/internal/models"
)

// FindingStore lists persisted analyzer findings.
type FindingStore interface {
	GetFindings() ([]models.AnalysisFinding, error)
	GetFindingsByTransaction(txID string) ([]models.AnalysisFinding, error)
}

func (s *Server) findingsList(w http.ResponseWriter, r *http.Request) {
	if s.Findings == nil {
		writeJSON(w, 200, map[string]any{"items": []any{}})
		return
	}
	items, err := s.Findings.GetFindings()
	if err != nil {
		writeErr(w, 500, err.Error())
		return
	}
	if items == nil {
		items = []models.AnalysisFinding{}
	}
	writeJSON(w, 200, map[string]any{"items": items})
}

func (s *Server) findingsForTx(w http.ResponseWriter, r *http.Request) {
	txID := r.PathValue("transaction_id")
	if strings.TrimSpace(txID) == "" {
		writeErr(w, 400, "transaction_id required")
		return
	}
	if s.Findings == nil {
		writeJSON(w, 200, map[string]any{"items": []any{}})
		return
	}
	items, err := s.Findings.GetFindingsByTransaction(txID)
	if err != nil {
		writeErr(w, 500, err.Error())
		return
	}
	if items == nil {
		items = []models.AnalysisFinding{}
	}
	writeJSON(w, 200, map[string]any{"items": items})
}
