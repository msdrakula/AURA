package store

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	"meb/internal/httpio"
)

type Settings struct {
	ListenHost         string `json:"listen_host"`
	ListenPort         int    `json:"listen_port"`
	InterceptRequests  bool   `json:"intercept_requests"`
	InterceptResponses bool   `json:"intercept_responses"`
	InterceptFilter    string `json:"intercept_filter"`
	RecordHistory      bool   `json:"record_history"`
	VerifyUpstream     bool   `json:"verify_upstream"`
	UIHost             string `json:"ui_host"`
	UIPort             int    `json:"ui_port"`
}

func (s Settings) Public() map[string]any {
	b, _ := json.Marshal(s)
	var m map[string]any
	_ = json.Unmarshal(b, &m)
	return m
}

type Stats struct {
	Flows       int `json:"flows"`
	Intercepted int `json:"intercepted"`
	BytesIn     int `json:"bytes_in"`
	BytesOut    int `json:"bytes_out"`
	Errors      int `json:"errors"`
}

type Flow struct {
	ID             string
	Created        float64
	Request        *httpio.Request
	Response       *httpio.Response
	Error          string
	DurationMS     *int
	Client         string
	Comment        string
	EditedRequest  string
	EditedResponse string
}

func (f *Flow) Summary() map[string]any {
	req := f.Request
	var status any
	var reason any
	var mime any
	respBytes := 0
	if f.Response != nil {
		status = f.Response.Status
		reason = f.Response.Reason
		mime = f.Response.Header("Content-Type")
		respBytes = len(f.Response.Body)
	}
	var err any
	if f.Error != "" {
		err = f.Error
	}
	return map[string]any{
		"id":          f.ID,
		"created":     f.Created,
		"method":      req.Method,
		"url":         req.URL(),
		"host":        req.Host,
		"path":        req.Path,
		"scheme":      req.Scheme,
		"status":      status,
		"reason":      reason,
		"req_bytes":   len(req.Body),
		"resp_bytes":  respBytes,
		"duration_ms": f.DurationMS,
		"error":       err,
		"mime":        mime,
	}
}

type Pending struct {
	ID        string
	Phase     string
	Flow      *Flow
	Action    string
	EditedRaw string
	done      chan struct{}
	once      sync.Once
}

func (p *Pending) complete(action, raw string) {
	p.Action = action
	p.EditedRaw = raw
	p.once.Do(func() { close(p.done) })
}

type LogItem struct {
	Ts      float64 `json:"ts"`
	Level   string  `json:"level"`
	Message string  `json:"message"`
}

type Store struct {
	mu           sync.Mutex
	Settings     Settings
	history      []*Flow
	historyByID  map[string]*Flow
	pending      map[string]*Pending
	pendingOrder []string
	Logs         []LogItem
	subs         map[chan any]struct{}
	ProxyRunning bool
	Stats        Stats
}

const (
	maxHistory = 4000
	maxLogs    = 500
)

func New() *Store {
	return &Store{
		Settings: Settings{
			ListenHost:     "127.0.0.1",
			ListenPort:     8080,
			RecordHistory:  true,
			VerifyUpstream: true,
			UIHost:         "127.0.0.1",
			UIPort:         1337,
		},
		historyByID: map[string]*Flow{},
		pending:     map[string]*Pending{},
		subs:        map[chan any]struct{}{},
	}
}

func nowTS() float64 {
	return float64(time.Now().UnixNano()) / 1e9
}

func (s *Store) Log(level, message string) {
	item := LogItem{Ts: nowTS(), Level: level, Message: message}
	s.mu.Lock()
	s.Logs = append([]LogItem{item}, s.Logs...)
	if len(s.Logs) > maxLogs {
		s.Logs = s.Logs[:maxLogs]
	}
	s.mu.Unlock()
	s.Emit(map[string]any{"type": "log", "item": item})
}

func (s *Store) AddFlow(flow *Flow) {
	s.mu.Lock()
	if old, ok := s.historyByID[flow.ID]; ok {
		for _, h := range s.history {
			if h == old {
				s.mu.Unlock()
				return
			}
		}
	}
	if len(s.history) == maxHistory {
		dropped := s.history[0]
		delete(s.historyByID, dropped.ID)
		s.history = s.history[1:]
	}
	s.history = append(s.history, flow)
	s.historyByID[flow.ID] = flow
	s.Stats.Flows++
	s.Stats.BytesIn += len(flow.Request.Body)
	if flow.Response != nil {
		s.Stats.BytesOut += len(flow.Response.Body)
	}
	sum := flow.Summary()
	s.mu.Unlock()
	s.Emit(map[string]any{"type": "flow", "flow": sum})
}

func (s *Store) History() []*Flow {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]*Flow, len(s.history))
	copy(out, s.history)
	return out
}

func (s *Store) FlowByID(id string) *Flow {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.historyByID[id]
}

func (s *Store) ClearHistory() {
	s.mu.Lock()
	s.history = nil
	s.historyByID = map[string]*Flow{}
	s.mu.Unlock()
	s.Emit(map[string]any{"type": "history_cleared"})
}

func (s *Store) matchesFilter(flow *Flow) bool {
	needle := strings.TrimSpace(strings.ToLower(s.Settings.InterceptFilter))
	if needle == "" {
		return true
	}
	blob := strings.ToLower(flow.Request.Method + " " + flow.Request.URL())
	return strings.Contains(blob, needle)
}

func (s *Store) Intercept(flow *Flow, phase string) string {
	s.mu.Lock()
	enabled := s.Settings.InterceptRequests
	if phase == "response" {
		enabled = s.Settings.InterceptResponses
	}
	if !enabled || !s.matchesFilter(flow) {
		s.mu.Unlock()
		return "forward"
	}
	p := &Pending{
		ID:     shortID(),
		Phase:  phase,
		Flow:   flow,
		Action: "forward",
		done:   make(chan struct{}),
	}
	s.pending[p.ID] = p
	s.pendingOrder = append(s.pendingOrder, p.ID)
	s.Stats.Intercepted++
	view := PendingView(p)
	s.mu.Unlock()
	s.Emit(map[string]any{"type": "intercept", "item": view})
	<-p.done
	s.mu.Lock()
	delete(s.pending, p.ID)
	order := s.pendingOrder[:0]
	for _, id := range s.pendingOrder {
		if id != p.ID {
			order = append(order, id)
		}
	}
	s.pendingOrder = order
	edited := p.EditedRaw
	action := p.Action
	s.mu.Unlock()
	if edited != "" {
		if phase == "request" {
			flow.EditedRequest = edited
		} else {
			flow.EditedResponse = edited
		}
	}
	s.Emit(map[string]any{"type": "intercept_done", "id": p.ID})
	return action
}

func (s *Store) Decide(id, action, raw string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	p := s.pending[id]
	if p == nil {
		return false
	}
	p.complete(action, raw)
	return true
}

func (s *Store) ForwardAll() int {
	s.mu.Lock()
	ids := append([]string{}, s.pendingOrder...)
	s.mu.Unlock()
	for _, id := range ids {
		s.Decide(id, "forward", "")
	}
	return len(ids)
}

// DropAll drops every pending intercepted message.
func (s *Store) DropAll() int {
	s.mu.Lock()
	ids := append([]string{}, s.pendingOrder...)
	s.mu.Unlock()
	for _, id := range ids {
		s.Decide(id, "drop", "")
	}
	return len(ids)
}

func (s *Store) PendingList() []map[string]any {
	s.mu.Lock()
	defer s.mu.Unlock()
	var items []map[string]any
	for _, id := range s.pendingOrder {
		if p := s.pending[id]; p != nil {
			items = append(items, PendingView(p))
		}
	}
	if items == nil {
		items = []map[string]any{}
	}
	return items
}

func (s *Store) PendingCount() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.pending)
}

func (s *Store) Snapshot() map[string]any {
	s.mu.Lock()
	defer s.mu.Unlock()
	return map[string]any{
		"proxy_running": s.ProxyRunning,
		"settings":      s.Settings,
		"stats":         s.Stats,
		"pending":       len(s.pending),
		"history":       len(s.history),
	}
}

func (s *Store) LogItems() []LogItem {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.Logs) == 0 {
		return []LogItem{}
	}
	out := make([]LogItem, len(s.Logs))
	copy(out, s.Logs)
	return out
}

func (s *Store) Subscribe() chan any {
	ch := make(chan any, 64)
	s.mu.Lock()
	s.subs[ch] = struct{}{}
	s.mu.Unlock()
	return ch
}

func (s *Store) Unsubscribe(ch chan any) {
	s.mu.Lock()
	delete(s.subs, ch)
	s.mu.Unlock()
}

func (s *Store) Emit(event any) {
	s.mu.Lock()
	subs := make([]chan any, 0, len(s.subs))
	for ch := range s.subs {
		subs = append(subs, ch)
	}
	s.mu.Unlock()
	for _, ch := range subs {
		select {
		case ch <- event:
		default:
		}
	}
}

func (s *Store) IncErrors() {
	s.mu.Lock()
	s.Stats.Errors++
	s.mu.Unlock()
}

func (s *Store) SetRunning(v bool) {
	s.mu.Lock()
	s.ProxyRunning = v
	s.mu.Unlock()
}

func (s *Store) IsRunning() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.ProxyRunning
}

func PendingView(p *Pending) map[string]any {
	flow := p.Flow
	req := flow.Request
	payload := map[string]any{
		"id":          p.ID,
		"phase":       p.Phase,
		"flow_id":     flow.ID,
		"method":      req.Method,
		"url":         req.URL(),
		"request_raw": httpio.DecodeRaw(req.Raw()),
	}
	if flow.Response != nil {
		payload["response_raw"] = httpio.DecodeRaw(flow.Response.Raw())
		payload["status"] = flow.Response.Status
	}
	return payload
}

func (s *Store) ListenAddr() (string, int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.Settings.ListenHost, s.Settings.ListenPort
}

func (s *Store) VerifyUpstream() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.Settings.VerifyUpstream
}

func (s *Store) RecordHistory() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.Settings.RecordHistory
}

func (s *Store) ApplySettings(patch map[string]any) (restart bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if v, ok := patch["listen_host"].(string); ok && v != s.Settings.ListenHost {
		s.Settings.ListenHost = v
		restart = true
	}
	if v, ok := asInt(patch["listen_port"]); ok && v != s.Settings.ListenPort {
		s.Settings.ListenPort = v
		restart = true
	}
	if v, ok := patch["intercept_requests"].(bool); ok {
		s.Settings.InterceptRequests = v
	}
	if v, ok := patch["intercept_responses"].(bool); ok {
		s.Settings.InterceptResponses = v
	}
	if v, ok := patch["intercept_filter"].(string); ok {
		s.Settings.InterceptFilter = v
	}
	if v, ok := patch["record_history"].(bool); ok {
		s.Settings.RecordHistory = v
	}
	if v, ok := patch["verify_upstream"].(bool); ok {
		s.Settings.VerifyUpstream = v
	}
	return restart
}

func asInt(v any) (int, bool) {
	switch n := v.(type) {
	case float64:
		return int(n), true
	case int:
		return n, true
	case json.Number:
		i, err := n.Int64()
		return int(i), err == nil
	default:
		return 0, false
	}
}

func NewID() string {
	return shortID()
}

func shortID() string {
	var b [5]byte
	if _, err := rand.Read(b[:]); err != nil {
		return fmt.Sprintf("%010d", time.Now().UnixNano()%1e10)
	}
	return hex.EncodeToString(b[:])
}
