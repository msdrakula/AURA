package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"go.uber.org/zap"

	"meb/internal/ffuf"
	"meb/internal/intel"
	"meb/internal/wordlist"
)

func (s *Server) intelCatalog(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, 200, map[string]any{"stages": intel.Catalog()})
}

func (s *Server) intelList(w http.ResponseWriter, _ *http.Request) {
	if s.Intel == nil {
		writeErr(w, 500, "intel is not configured")
		return
	}
	list, err := s.Intel.Store.ListTargets()
	if err != nil {
		writeErr(w, 500, err.Error())
		return
	}
	if list == nil {
		list = []intel.Target{}
	}
	writeJSON(w, 200, map[string]any{"items": list})
}

type intelCreateBody struct {
	Target     string `json:"target"`
	Authorized bool   `json:"authorized"`
}

func (s *Server) intelCreate(w http.ResponseWriter, r *http.Request) {
	if s.Intel == nil {
		writeErr(w, 500, "intel is not configured")
		return
	}
	var body intelCreateBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeErr(w, 400, "invalid JSON")
		return
	}
	t, err := s.Intel.EnsureTarget(body.Target, body.Authorized)
	if err != nil {
		writeErr(w, 400, err.Error())
		return
	}
	s.writeIntelSnapshot(w, t.ID)
}

func (s *Server) intelGet(w http.ResponseWriter, r *http.Request) {
	s.writeIntelSnapshot(w, r.PathValue("id"))
}

func (s *Server) intelRunStage(w http.ResponseWriter, r *http.Request) {
	if s.Intel == nil {
		writeErr(w, 500, "intel is not configured")
		return
	}
	id := r.PathValue("id")
	stage := r.PathValue("stage")
	var opt intel.StageOptions
	if r.Body != nil {
		_ = json.NewDecoder(r.Body).Decode(&opt)
	}
	res, err := s.Intel.RunStage(r.Context(), id, stage, opt)
	stopped := r.Context().Err() != nil
	if err != nil && !stopped && res.Added == 0 && res.Stage.Status != intel.StatusError {
		writeErr(w, 400, err.Error())
		return
	}
	t, stages, arts, m, snapErr := s.Intel.Snapshot(id)
	if snapErr != nil {
		writeErr(w, 500, snapErr.Error())
		return
	}
	errMsg := ""
	if !stopped {
		errMsg = errString(err)
	}
	writeJSON(w, 200, map[string]any{
		"run":       res,
		"target":    t,
		"stages":    stages,
		"catalog":   intel.Catalog(),
		"artifacts": arts,
		"map":       m,
		"error":     errMsg,
		"stopped":   stopped,
	})
}

func (s *Server) intelIngest(w http.ResponseWriter, r *http.Request) {
	if s.Intel == nil {
		writeErr(w, 500, "intel is not configured")
		return
	}
	if s.History == nil {
		writeErr(w, 400, "history is empty")
		return
	}
	txs, err := s.History.GetTransactions(500, 0)
	if err != nil {
		writeErr(w, 500, err.Error())
		return
	}
	urls := make([]string, 0, len(txs))
	for _, tx := range txs {
		if tx.Request.URL != "" {
			urls = append(urls, tx.Request.URL)
		}
	}
	n, err := s.Intel.IngestHistory(r.PathValue("id"), urls)
	if err != nil {
		writeErr(w, 400, err.Error())
		return
	}
	t, stages, arts, m, err := s.Intel.Snapshot(r.PathValue("id"))
	if err != nil {
		writeErr(w, 500, err.Error())
		return
	}
	writeJSON(w, 200, map[string]any{
		"ingested":  n,
		"target":    t,
		"stages":    stages,
		"catalog":   intel.Catalog(),
		"artifacts": arts,
		"map":       m,
	})
}

func (s *Server) writeIntelSnapshot(w http.ResponseWriter, id string) {
	if s.Intel == nil {
		writeErr(w, 500, "intel is not configured")
		return
	}
	t, stages, arts, m, err := s.Intel.Snapshot(id)
	if err != nil {
		writeErr(w, 404, err.Error())
		return
	}
	writeJSON(w, 200, map[string]any{
		"target":    t,
		"stages":    stages,
		"catalog":   intel.Catalog(),
		"artifacts": arts,
		"map":       m,
	})
}

func errString(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

type fuzzBody struct {
	Mode         string            `json:"mode"` // "fuzz" (FUZZ keyword anywhere) or "paths" (dirbust: FUZZ appended to base URL)
	URL          string            `json:"url"`
	Method       string            `json:"method"`
	Headers      map[string]string `json:"headers"`
	Cookies      string            `json:"cookies"`
	Body         string            `json:"body"`
	Wordlist     []string          `json:"wordlist"`
	WordlistPath string            `json:"wordlist_path"`
	Workers      int               `json:"workers"`
	RPS          int               `json:"rps"`
	Timeout      int               `json:"timeout"`
	WordlistName string            `json:"wordlist_name"`
	Authorized   bool              `json:"authorized"`
}

func (s *Server) resolveWords(inline []string, path, kind string) ([]string, error) {
	if len(inline) > 0 {
		return inline, nil
	}
	if path != "" {
		if s.Lists == nil {
			return nil, fmt.Errorf("seclists is not installed")
		}
		return s.Lists.Load(path)
	}
	switch strings.ToLower(kind) {
	case "params":
		return s.Lists.LoadFirst(wordlist.Params(), wordlist.DefaultParams...), nil
	case "dns", "hosts":
		return s.Lists.LoadFirst(wordlist.HostPrefixes(), wordlist.DefaultDNS...), nil
	default:
		return s.Lists.LoadFirst(wordlist.Dirs(), wordlist.DefaultDirs...), nil
	}
}

func (s *Server) fuzzRun(w http.ResponseWriter, r *http.Request) {
	var body fuzzBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeErr(w, 400, "invalid JSON")
		return
	}
	if rejectUnlessAuthorized(w, body.Authorized) {
		return
	}
	if ffuf.LookPath() == "" {
		writeErr(w, 500, "ffuf binary not found — install: sudo apt install ffuf")
		return
	}
	url := strings.TrimSpace(body.URL)
	if url == "" {
		writeErr(w, 400, "url is required")
		return
	}
	if body.Mode == "paths" {
		url = strings.TrimRight(url, "/") + "/FUZZ"
	}
	kind := body.WordlistName
	if kind == "" {
		kind = "params"
		if body.Mode == "paths" {
			kind = "dirs"
		}
	}
	words, err := s.resolveWords(body.Wordlist, body.WordlistPath, kind)
	if err != nil {
		writeErr(w, 400, err.Error())
		return
	}
	if len(words) == 0 {
		writeErr(w, 400, "wordlist is required")
		return
	}
	headers := make([]string, 0, len(body.Headers))
	for k, v := range body.Headers {
		headers = append(headers, k+": "+v)
	}
	ctx, cancel := context.WithTimeout(r.Context(), 30*time.Minute)
	defer cancel()

	flusher, ok := w.(http.Flusher)
	if !ok {
		writeErr(w, 500, "streaming unsupported")
		return
	}
	wlLabel := body.WordlistPath
	if wlLabel == "" {
		wlLabel = fmt.Sprintf("custom list (%d words)", len(words))
	}
	s.Log.Info("fuzz: run", zap.String("mode", body.Mode), zap.String("url", url), zap.Int("words", len(words)), zap.Int("workers", body.Workers))

	started := time.Now()
	w.Header().Set("Content-Type", "application/x-ndjson; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.WriteHeader(200)
	emit := func(v map[string]any) {
		b, _ := json.Marshal(v)
		_, _ = w.Write(append(b, '\n'))
		flusher.Flush()
	}
	emit(map[string]any{
		"type":    "start",
		"mode":    body.Mode,
		"url":     url,
		"method":  mapDefault(body.Method, "GET"),
		"wordlist": wlLabel,
		"words":   len(words),
		"workers": body.Workers,
		"rps":     body.RPS,
		"timeout": body.Timeout,
	})

	results, st, runErr := ffuf.Run(ctx, ffuf.Options{
		URL:      url,
		Method:   body.Method,
		Headers:  headers,
		Cookies:  body.Cookies,
		Body:     body.Body,
		Wordlist: words,
		Workers:  body.Workers,
		RPS:      body.RPS,
		Timeout:  body.Timeout,
		OnResult: func(r ffuf.Result) {
			emit(map[string]any{
				"type":         "result",
				"payload":      r.Payload,
				"position":     r.Position,
				"status":       r.Status,
				"length":       r.Length,
				"words":        r.Words,
				"lines":        r.Lines,
				"url":          r.URL,
				"content_type": r.ContentType,
				"redirect":     r.Redirect,
				"duration_ms":  r.DurationMs,
			})
		},
	})
	emit(map[string]any{
		"type":        "done",
		"count":       len(results),
		"tried":       st.Tried,
		"errors":      st.Errors,
		"truncated":   len(results) >= 20000,
		"duration_ms": time.Since(started).Milliseconds(),
		"error":       errString(runErr),
	})
}

func mapDefault(v, def string) string {
	if v == "" {
		return def
	}
	return v
}

func (s *Server) wordlists(w http.ResponseWriter, _ *http.Request) {
	items := []wordlist.Entry{}
	root := ""
	if s.Lists != nil {
		items = s.Lists.Catalog()
		root = s.Lists.Root
		if items == nil {
			items = []wordlist.Entry{}
		}
	}
	writeJSON(w, 200, map[string]any{
		"root":    root,
		"count":   len(items),
		"license": "MIT",
		"source":  "https://github.com/danielmiessler/SecLists",
		"dirs":    wordlist.Dirs(),
		"params":  wordlist.Params(),
		"hosts":   wordlist.HostPrefixes(),
		"items":   items,
	})
}

func (s *Server) wordlistPreview(w http.ResponseWriter, r *http.Request) {
	if s.Lists == nil {
		writeErr(w, 404, "seclists is not installed")
		return
	}
	path := r.URL.Query().Get("path")
	words, err := s.Lists.Load(path)
	if err != nil {
		writeErr(w, 400, err.Error())
		return
	}
	preview := words
	if len(preview) > 40 {
		preview = preview[:40]
	}
	writeJSON(w, 200, map[string]any{"path": path, "preview": preview, "lines": len(words)})
}
