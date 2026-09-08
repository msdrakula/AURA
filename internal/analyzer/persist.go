package analyzer

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"go.uber.org/zap"

	"meb/internal/httpio"
	"meb/internal/models"
	"meb/internal/storage"
)

var _ FindingStore = (*storage.Store)(nil)

// FindingStore persists rule hits. *storage.Store satisfies this.
type FindingStore interface {
	SaveFinding(f *models.AnalysisFinding) error
}

// Persist runs rules and writes each hit as models.AnalysisFinding.
// Database errors are logged and never returned to the caller.
func (e *Engine) Persist(log *zap.Logger, db FindingStore, transactionID string, req *httpio.Request, resp *httpio.Response) {
	if e == nil || db == nil || transactionID == "" {
		return
	}
	if log == nil {
		log = zap.NewNop()
	}
	for _, hit := range e.Analyze(req, resp) {
		f := &models.AnalysisFinding{
			ID:            findingID(),
			TransactionID: transactionID,
			RuleName:      hit.Rule,
			Severity:      models.Severity(hit.Severity),
			Title:         hit.Message,
			Description:   hit.Message,
			Evidence:      hit.Evidence,
		}
		if err := db.SaveFinding(f); err != nil {
			log.Error("save finding",
				zap.Error(err),
				zap.String("transaction_id", transactionID),
				zap.String("rule", hit.Rule),
			)
			continue
		}
	}
}

func findingID() string {
	var b [8]byte
	if _, err := rand.Read(b[:]); err != nil {
		return hex.EncodeToString([]byte(time.Now().Format("15040599")))
	}
	return hex.EncodeToString(b[:])
}
