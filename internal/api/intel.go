package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"go.uber.org/zap"

	"meb/internal/fuzz"
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
	URL          string            `json:"url"`
	Method       string            `json:"method"`
	Headers      map[string]string `json:"headers"`
	Body         string            `json:"body"`
	Wordlist     []string          `json:"wordlist"`
	WordlistPath string            `json:"wordlist_path"`
	Workers      int               `json:"workers"`
	RPS          int               `json:"rps"`
	Hide         []int             `json:"hide"`
	WordlistName string            `json:"wordlist_name"`
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
	words, err := s.resolveWords(body.Wordlist, body.WordlistPath, body.WordlistName)
	if err != nil {
		writeErr(w, 400, err.Error())
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), 30*time.Minute)
	defer cancel()
	s.Log.Info("fuzz: run", zap.String("url", body.URL), zap.Int("words", len(words)), zap.Int("workers", body.Workers))
	hits, err := fuzz.Run(ctx, fuzz.Options{
		URL:      body.URL,
		Method:   body.Method,
		Headers:  body.Headers,
		Body:     body.Body,
		Wordlist: words,
		Workers:  body.Workers,
		RPS:      body.RPS,
		Hide:     body.Hide,
		Timeout:  8 * time.Second,
	})
	if err != nil && len(hits) == 0 && ctx.Err() == nil {
		writeErr(w, 400, err.Error())
		return
	}
	if hits == nil {
		hits = []fuzz.Hit{}
	}
	writeJSON(w, 200, map[string]any{
		"items":     hits,
		"count":     len(hits),
		"tried":     len(words),
		"truncated": len(hits) >= 8000,
		"error":     errString(err),
	})
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
