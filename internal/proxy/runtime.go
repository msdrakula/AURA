package proxy

import (
	"meb/internal/httpio"
	"meb/internal/models"
)

// Runtime is live intercept/settings state. It is not the SQLite session log.
type Runtime interface {
	ListenAddr() (string, int)
	VerifyUpstream() bool
	SetRunning(bool)
	Log(level, msg string)
	IncErrors()
	ForwardAll() int
	InterceptRequest(id string, req *httpio.Request) (action, editedRaw string)
	InterceptResponse(id string, req *httpio.Request, resp *httpio.Response) (action, editedRaw string)
}

// TransactionWriter persists a finished exchange. *storage.Store satisfies this.
type TransactionWriter interface {
	SaveTransaction(tx *models.HTTPTransaction) error
}

// FindingWriter persists analyzer hits. *storage.Store satisfies this.
type FindingWriter interface {
	SaveFinding(f *models.AnalysisFinding) error
}
