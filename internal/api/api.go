package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"

	"meb/internal/certs"
	"meb/internal/codec"
	"meb/internal/intel"
	"meb/internal/models"
	"meb/internal/proxy"
	"meb/internal/repeater"
	"meb/internal/store"
	"meb/internal/wordlist"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// History is the persistent session log shown in the UI.
type History interface {
	GetTransactions(limit, offset int) ([]models.HTTPTransaction, error)
	GetTransaction(id string) (*models.HTTPTransaction, error)
	ClearTransactions() error
}

// Saver persists a transaction to storage.
type Saver interface {
	SaveTransaction(tx *models.HTTPTransaction) error
}

// Options wires the HTTP UI to proxy intercept state and optional SQLite history.
type Options struct {
	Runtime      *store.Store
	Proxy        *proxy.Server
	History      History
	Findings     FindingStore
	Batch        BatchRunner
	BatchFactory BatchFactory
	Saver        Saver
	Callback     InteractionStore
	Organizer    OrganizerStore
	Intel        *intel.Engine
	Lists        *wordlist.Library
	Log          *zap.Logger
	WebDir       string
	DataDir      string
}

type Server struct {
	Store        *store.Store
	Proxy        *proxy.Server
	History      History
	Findings     FindingStore
	Batch        BatchRunner
	BatchFactory BatchFactory
	Saver        Saver
	Callback     InteractionStore
	Organizer    OrganizerStore
	Intel        *intel.Engine
	Lists        *wordlist.Library
	Log          *zap.Logger
	WebDir       string
	DataDir      string
}

// New builds the REST and WebSocket handler.
func New(opts Options) http.Handler {
	log := opts.Log
	if log == nil {
		log = zap.NewNop()
	}
	s := &Server{
		Store:        opts.Runtime,
		Proxy:        opts.Proxy,
		History:      opts.History,
		Findings:     opts.Findings,
		Batch:        opts.Batch,
		BatchFactory: opts.BatchFactory,
		Saver:        opts.Saver,
		Callback:     opts.Callback,
		Organizer:    opts.Organizer,
		Intel:        opts.Intel,
		Lists:        opts.Lists,
		Log:          log,
		WebDir:       opts.WebDir,
		DataDir:      opts.DataDir,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/status", s.status)
	mux.HandleFunc("POST /api/proxy/start", s.proxyStart)
	mux.HandleFunc("POST /api/proxy/stop", s.proxyStop)
	mux.HandleFunc("POST /api/proxy/launch-browser", s.proxyLaunchBrowser)
	mux.HandleFunc("PUT /api/settings", s.updateSettings)
	mux.HandleFunc("GET /api/ca.crt", s.caCert)
	mux.HandleFunc("GET /api/history", s.history)
	mux.HandleFunc("GET /api/history/{flow_id}", s.historyOne)
	mux.HandleFunc("DELETE /api/history", s.historyClear)
	mux.HandleFunc("GET /api/intercept", s.interceptQueue)
	mux.HandleFunc("POST /api/intercept/forward-all", s.interceptForwardAll)
	mux.HandleFunc("POST /api/intercept/drop-all", s.interceptDropAll)
	mux.HandleFunc("POST /api/intercept/{pending_id}", s.interceptDecide)
	mux.HandleFunc("POST /api/repeater", s.repeater)
	mux.HandleFunc("POST /api/codec", s.codec)
	mux.HandleFunc("POST /api/batch/execute", s.batchExecute)
	mux.HandleFunc("GET /api/callback/{request_id}", s.callbackInteractions)
	mux.HandleFunc("GET /api/findings", s.findingsList)
	mux.HandleFunc("GET /api/findings/{transaction_id}", s.findingsForTx)
	mux.HandleFunc("POST /api/sequencer/analyze", s.sequencerAnalyze)
	mux.HandleFunc("GET /api/organizer", s.organizerList)
	mux.HandleFunc("POST /api/organizer", s.organizerAdd)
	mux.HandleFunc("PATCH /api/organizer/{id}", s.organizerUpdate)
	mux.HandleFunc("DELETE /api/organizer/{id}", s.organizerDelete)
	mux.HandleFunc("POST /api/discover/run", s.discoverRun)
	mux.HandleFunc("POST /api/scanner/scan", s.scannerRun)
	mux.HandleFunc("GET /api/intel/catalog", s.intelCatalog)
	mux.HandleFunc("GET /api/intel/targets", s.intelList)
	mux.HandleFunc("POST /api/intel/targets", s.intelCreate)
	mux.HandleFunc("GET /api/intel/targets/{id}", s.intelGet)
	mux.HandleFunc("POST /api/intel/targets/{id}/stages/{stage}", s.intelRunStage)
	mux.HandleFunc("POST /api/intel/targets/{id}/ingest-history", s.intelIngest)
	mux.HandleFunc("POST /api/fuzz/run", s.fuzzRun)
	mux.HandleFunc("GET /api/wordlists", s.wordlists)
	mux.HandleFunc("GET /api/wordlists/preview", s.wordlistPreview)
	mux.HandleFunc("GET /api/logs", s.logs)
	mux.HandleFunc("POST /api/debug-log", s.debugIngest)
	mux.HandleFunc("GET /api/ws", s.ws)
	mux.HandleFunc("GET /{$}", s.index)
	mux.Handle("GET /static/", noStore(http.StripPrefix("/static/", http.FileServer(http.Dir(filepath.Join(opts.WebDir, "static"))))))
	return withDebug(mux)
}

func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}

func writeErr(w http.ResponseWriter, code int, msg string) {
	writeJSON(w, code, map[string]string{"detail": msg})
}

func (s *Server) status(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, 200, s.Store.Snapshot())
}

func (s *Server) proxyStart(w http.ResponseWriter, r *http.Request) {
	if err := s.Proxy.Start(context.Background()); err != nil {
		writeErr(w, 400, err.Error())
		return
	}
	writeJSON(w, 200, map[string]any{"ok": true, "proxy_running": s.Store.IsRunning()})
}

func (s *Server) proxyStop(w http.ResponseWriter, r *http.Request) {
	_ = s.Proxy.Stop(r.Context())
	writeJSON(w, 200, map[string]any{"ok": true, "proxy_running": s.Store.IsRunning()})
}

func (s *Server) updateSettings(w http.ResponseWriter, r *http.Request) {
	var body map[string]any
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeErr(w, 400, err.Error())
		return
	}
	restart := s.Store.ApplySettings(body)
	if restart && s.Store.IsRunning() {
		_ = s.Proxy.Stop(r.Context())
		if err := s.Proxy.Start(context.Background()); err != nil {
			writeErr(w, 400, err.Error())
			return
		}
	}
	pub := s.Store.Snapshot()["settings"]
	s.Store.Emit(map[string]any{"type": "settings", "settings": pub})
	writeJSON(w, 200, pub)
}

func (s *Server) caCert(w http.ResponseWriter, r *http.Request) {
	if _, _, err := certs.EnsureCA(s.DataDir); err != nil {
		writeErr(w, 500, err.Error())
		return
	}
	pem, err := os.ReadFile(certs.CAPath(s.DataDir))
	if err != nil {
		writeErr(w, 500, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/x-pem-file")
	w.Header().Set("Content-Disposition", "attachment; filename=aura-ca.crt")
	w.WriteHeader(200)
	_, _ = w.Write(pem)
}

func (s *Server) history(w http.ResponseWriter, r *http.Request) {
	q := strings.ToLower(strings.TrimSpace(r.URL.Query().Get("q")))
	limit := 400
	if v := r.URL.Query().Get("limit"); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			limit = n
		}
	}
	if limit < 1 {
		limit = 1
	}
	if limit > 1000 {
		limit = 1000
	}
	if s.History != nil {
		items, err := s.History.GetTransactions(limit, 0)
		if err != nil {
			writeErr(w, 500, err.Error())
			return
		}
		out := make([]map[string]any, 0, len(items))
		for _, tx := range items {
			sum := transactionSummary(tx)
			if q != "" {
				blob := strings.ToLower(fmt.Sprintf("%v %v %v", sum["url"], sum["method"], sum["status"]))
				if !strings.Contains(blob, q) {
					continue
				}
			}
			out = append(out, sum)
		}
		writeJSON(w, 200, map[string]any{"items": out})
		return
	}
	items := s.Store.History()
	out := make([]map[string]any, 0, len(items))
	for i := len(items) - 1; i >= 0; i-- {
		f := items[i]
		if q != "" {
			blob := strings.ToLower(f.Request.URL() + " " + f.Request.Method + " " + f.Error)
			if f.Response != nil {
				blob += " " + strconv.Itoa(f.Response.Status)
			}
			if !strings.Contains(blob, q) {
				continue
			}
		}
		out = append(out, f.Summary())
		if len(out) >= limit {
			break
		}
	}
	writeJSON(w, 200, map[string]any{"items": out})
}

func (s *Server) historyOne(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("flow_id")
	if s.History != nil {
		tx, err := s.History.GetTransaction(id)
		if err != nil {
			writeErr(w, 500, err.Error())
			return
		}
		if tx == nil {
			writeErr(w, 404, "flow not found")
			return
		}
		writeJSON(w, 200, transactionDetail(*tx))
		return
	}
	flow := s.Store.FlowByID(id)
	if flow == nil {
		writeErr(w, 404, "flow not found")
		return
	}
	writeJSON(w, 200, flowDetail(flow))
}

func (s *Server) historyClear(w http.ResponseWriter, r *http.Request) {
	if s.History != nil {
		if err := s.History.ClearTransactions(); err != nil {
			writeErr(w, 500, err.Error())
			return
		}
	}
	s.Store.ClearHistory()
	writeJSON(w, 200, map[string]any{"ok": true})
}

func (s *Server) interceptQueue(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, 200, map[string]any{"items": s.Store.PendingList()})
}

func (s *Server) interceptDecide(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Action string  `json:"action"`
		Raw    *string `json:"raw"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeErr(w, 400, err.Error())
		return
	}
	if body.Action != "forward" && body.Action != "drop" {
		writeErr(w, 400, "action must be forward or drop")
		return
	}
	raw := ""
	if body.Raw != nil {
		raw = *body.Raw
	}
	if !s.Store.Decide(r.PathValue("pending_id"), body.Action, raw) {
		writeErr(w, 404, "not in queue")
		return
	}
	writeJSON(w, 200, map[string]any{"ok": true})
}

func (s *Server) interceptForwardAll(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, 200, map[string]any{"ok": true, "count": s.Store.ForwardAll()})
}

func (s *Server) interceptDropAll(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, 200, map[string]any{"ok": true, "count": s.Store.DropAll()})
}

func (s *Server) repeater(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Raw    string  `json:"raw"`
		Scheme string  `json:"scheme"`
		Target *string `json:"target"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeErr(w, 400, err.Error())
		return
	}
	if body.Scheme == "" {
		body.Scheme = "https"
	}
	if body.Scheme != "http" && body.Scheme != "https" {
		writeErr(w, 400, "scheme must be http or https")
		return
	}
	target := ""
	if body.Target != nil {
		target = *body.Target
	}
	res := repeater.Send(r.Context(), repeater.Options{
		Raw:     body.Raw,
		Scheme:  body.Scheme,
		Target:  target,
		Verify:  s.Store.VerifyUpstream(),
		Timeout: 30 * time.Second,
	})
	if res.OK {
		s.saveToolFlow(body.Raw, res.ResponseRaw, body.Scheme, "repeater")
	}
	writeJSON(w, 200, res)
}

func (s *Server) codec(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Action string `json:"action"`
		Data   string `json:"data"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeErr(w, 400, err.Error())
		return
	}
	result, err := codec.Run(body.Action, body.Data)
	if err != nil {
		writeJSON(w, 200, map[string]any{"ok": false, "error": err.Error(), "result": ""})
		return
	}
	writeJSON(w, 200, map[string]any{"ok": true, "result": result})
}

func (s *Server) logs(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, 200, map[string]any{"items": s.Store.LogItems()})
}

func (s *Server) ws(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()
	ch := s.Store.Subscribe()
	defer s.Store.Unsubscribe(ch)
	hello := map[string]any{
		"type": "hello",
		"status": map[string]any{
			"proxy_running": s.Store.IsRunning(),
			"settings":      s.Store.Snapshot()["settings"],
			"stats":         s.Store.Snapshot()["stats"],
			"pending":       s.Store.PendingCount(),
		},
	}
	if err := conn.WriteJSON(hello); err != nil {
		return
	}
	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			if _, _, err := conn.ReadMessage(); err != nil {
				return
			}
		}
	}()
	for {
		select {
		case <-done:
			return
		case ev, ok := <-ch:
			if !ok {
				return
			}
			if err := conn.WriteJSON(ev); err != nil {
				return
			}
		}
	}
}

func noStore(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		h.ServeHTTP(w, r)
	})
}

func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-store")
	http.ServeFile(w, r, filepath.Join(s.WebDir, "index.html"))
}
