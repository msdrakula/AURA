package proxy

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"go.uber.org/zap"

	"meb/internal/debuglog"
	"meb/internal/httpio"
	"meb/internal/models"
)

func (s *Server) finish(id string, req *httpio.Request, resp *httpio.Response, started time.Time, comment string) {
	if s.shouldRecordHistory() && s.history != nil {
		tx := transactionFrom(id, req, resp, started, comment)
		if err := s.history.SaveTransaction(tx); err != nil {
			s.log.Error("save transaction", zap.Error(err), zap.String("id", id))
		}
	}
	if s.shouldRecordHistory() && resp != nil && s.analyzer != nil {
		s.analyzer.Persist(s.log, s.findings, id, req, resp)
	}
	logProxyFlow(id, req, resp, started, comment)
}

func (s *Server) shouldRecordHistory() bool {
	if s.rt == nil {
		return true
	}
	type historyGate interface {
		RecordHistory() bool
	}
	if g, ok := s.rt.(historyGate); ok {
		return g.RecordHistory()
	}
	return true
}

func logProxyFlow(id string, req *httpio.Request, resp *httpio.Response, started time.Time, comment string) {
	fields := map[string]any{"id": id, "ms": time.Since(started).Milliseconds()}
	if req != nil {
		fields["method"] = req.Method
		fields["host"] = req.Host
		fields["path"] = req.Path
		fields["scheme"] = req.Scheme
	}
	if resp != nil {
		fields["status"] = resp.Status
		fields["resp_bytes"] = len(resp.Body)
	}
	level := "info"
	if comment != "" {
		level = "warn"
		fields["comment"] = comment
	} else if resp != nil && resp.Status >= 500 {
		level = "error"
	} else if resp != nil && resp.Status >= 400 {
		level = "warn"
	}
	debuglog.Write(debuglog.Event{Level: level, Src: "proxy", Module: "proxy", Action: "flow", Fields: fields})
}

func transactionFrom(id string, req *httpio.Request, resp *httpio.Response, started time.Time, comment string) *models.HTTPTransaction {
	ts := started
	if ts.IsZero() {
		ts = time.Now()
	}
	tx := &models.HTTPTransaction{
		ID:        id,
		Timestamp: ts.UTC(),
		Comment:   comment,
		Request: models.HTTPRequest{
			Headers: map[string]string{},
		},
	}
	if req != nil {
		tx.Request.Method = req.Method
		tx.Request.URL = req.URL()
		tx.Request.Headers = headerMap(req.Headers)
		tx.Request.Body = append([]byte(nil), req.Body...)
	}
	if resp != nil {
		tx.Response = &models.HTTPResponse{
			Status:   resp.Status,
			Headers:  headerMap(resp.Headers),
			Body:     append([]byte(nil), resp.Body...),
			Duration: time.Since(started),
		}
	}
	return tx
}

func headerMap(hs []httpio.Header) map[string]string {
	if len(hs) == 0 {
		return map[string]string{}
	}
	out := make(map[string]string, len(hs))
	for _, h := range hs {
		out[h.Key] = h.Value
	}
	return out
}

func newID() string {
	var b [5]byte
	if _, err := rand.Read(b[:]); err != nil {
		return hex.EncodeToString([]byte(time.Now().Format("15040")))
	}
	return hex.EncodeToString(b[:])
}
