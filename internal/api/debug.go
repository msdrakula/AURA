package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime/debug"
	"strings"
	"time"

	"meb/internal/debuglog"
)

func (s *Server) debugIngest(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	var body struct {
		Events []debuglog.Event `json:"events"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeErr(w, 400, "invalid JSON")
		return
	}
	debuglog.Ingest(body.Events)
	writeJSON(w, 200, map[string]any{"ok": true, "n": len(body.Events)})
}

type statusRecorder struct {
	http.ResponseWriter
	status int
	body   bytes.Buffer
}

func (s *statusRecorder) WriteHeader(code int) {
	s.status = code
	s.ResponseWriter.WriteHeader(code)
}

func (s *statusRecorder) Write(p []byte) (int, error) {
	if s.status == 0 {
		s.status = http.StatusOK
	}
	if s.status >= 400 && s.body.Len() < 800 {
		remain := 800 - s.body.Len()
		if len(p) < remain {
			s.body.Write(p)
		} else {
			s.body.Write(p[:remain])
		}
	}
	return s.ResponseWriter.Write(p)
}

func withDebug(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/ws" || r.URL.Path == "/api/debug-log" {
			next.ServeHTTP(w, r)
			return
		}
		defer func() {
			if rec := recover(); rec != nil {
				debuglog.Write(debuglog.Event{
					Level:  "error",
					Src:    "api",
					Module: "http",
					Action: "panic",
					Msg:    fmt.Sprint(rec),
					Fields: map[string]any{
						"method": r.Method,
						"path":   r.URL.Path,
						"stack":  clipStack(debug.Stack()),
					},
				})
				http.Error(w, "internal error", http.StatusInternalServerError)
			}
		}()

		rec := &statusRecorder{ResponseWriter: w, status: 0}
		start := time.Now()
		next.ServeHTTP(rec, r)
		status := rec.status
		if status == 0 {
			status = http.StatusOK
		}
		if !shouldLogHTTP(r, status) {
			return
		}
		level := "info"
		if status >= 500 {
			level = "error"
		} else if status >= 400 {
			level = "warn"
		}
		fields := map[string]any{
			"method": r.Method,
			"path":   r.URL.Path,
			"status": status,
			"ms":     time.Since(start).Milliseconds(),
		}
		if rec.body.Len() > 0 {
			fields["body"] = rec.body.String()
		}
		debuglog.Write(debuglog.Event{
			Level:  level,
			Src:    "api",
			Module: moduleFromAPIPath(r.URL.Path),
			Action: r.Method + " " + r.URL.Path,
			Msg:    r.Method + " " + r.URL.Path,
			Fields: fields,
		})
	})
}

func shouldLogHTTP(r *http.Request, status int) bool {
	if status >= 400 {
		return true
	}
	p := r.URL.Path
	if strings.HasPrefix(p, "/static/") || p == "/" {
		return false
	}
	if r.Method == http.MethodGet {
		switch {
		case p == "/api/status", p == "/api/history", p == "/api/intercept",
			p == "/api/logs", p == "/api/ca.crt", p == "/api/wordlists", p == "/api/intel/catalog":
			return false
		case strings.HasPrefix(p, "/api/history/"), strings.HasPrefix(p, "/api/intel/targets"):
			return false
		}
	}
	return strings.HasPrefix(p, "/api/")
}

func moduleFromAPIPath(p string) string {
	p = strings.TrimPrefix(p, "/api/")
	if p == "" {
		return "http"
	}
	if i := strings.IndexByte(p, '/'); i > 0 {
		return p[:i]
	}
	return p
}

func clipStack(b []byte) string {
	s := string(b)
	if len(s) > 4000 {
		return s[:4000] + "…"
	}
	return s
}
