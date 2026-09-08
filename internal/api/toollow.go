package api

import (
	"crypto/rand"
	"encoding/hex"
	"strings"
	"time"

	"github.com/google/uuid"

	"meb/internal/httpio"
	"meb/internal/models"
)

// saveToolFlow builds an HTTPTransaction from a raw request/response and
// persists it to history with a source tag (e.g. "repeater", "intruder",
// "scanner", "discover"). This makes every tool-sent request appear in the
// Logger and HTTP history.
func (s *Server) saveToolFlow(reqRaw, respRaw, scheme, source string) {
	if s.History == nil || strings.TrimSpace(reqRaw) == "" {
		return
	}
	id := uuid.NewString()
	tx := &models.HTTPTransaction{
		ID:        id,
		Timestamp: time.Now().UTC(),
		Comment:   "tool:" + source,
		Tags:      []string{"tool:" + source},
		Request: models.HTTPRequest{
			Headers: map[string]string{},
		},
	}
	// Parse request
	head, body := httpio.SplitHeadBody([]byte(strings.ReplaceAll(reqRaw, "\n", "\r\n")))
	start, hdrs := httpio.ParseHeaders(head)
	parts := strings.Fields(start)
	if len(parts) >= 1 {
		tx.Request.Method = parts[0]
	}
	urlPath := "/"
	if len(parts) >= 2 {
		urlPath = parts[1]
	}
	host := ""
	for _, h := range hdrs {
		tx.Request.Headers[h.Key] = h.Value
		if strings.EqualFold(h.Key, "Host") {
			host = h.Value
		}
	}
	// Build a full URL so the Logger/Target can show the host.
	if host != "" {
		sch := scheme
		if sch == "" {
			sch = "https"
		}
		tx.Request.URL = sch + "://" + host + urlPath
	} else {
		tx.Request.URL = urlPath
	}
	tx.Request.Body = body
	// Parse response
	if respRaw != "" {
		rhead, rbody := httpio.SplitHeadBody([]byte(strings.ReplaceAll(respRaw, "\n", "\r\n")))
		rstart, rhdrs := httpio.ParseHeaders(rhead)
		rparts := strings.Fields(rstart)
		status := 0
		if len(rparts) >= 2 {
			var sc int
			for _, ch := range rparts[1] {
				if ch >= '0' && ch <= '9' {
					sc = sc*10 + int(ch-'0')
				} else {
					break
				}
			}
			status = sc
		}
		rh := map[string]string{}
		for _, h := range rhdrs {
			rh[h.Key] = h.Value
		}
		tx.Response = &models.HTTPResponse{
			Status:  status,
			Headers: rh,
			Body:    rbody,
		}
	}
	// Persist via the storage Store.
	if s.Saver != nil {
		_ = s.Saver.SaveTransaction(tx)
	}
	if s.Store != nil {
		s.Store.Emit(map[string]any{"type": "flow", "source": source})
	}
	_ = scheme
}

// randID returns a short random hex id.
func randID() string {
	var b [8]byte
	_, _ = rand.Read(b[:])
	return hex.EncodeToString(b[:])
}
