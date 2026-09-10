// Package debuglog writes a temporary session journal for debugging AURA.
// Each process start creates data/debug/session-*.jsonl and points
// data/debug/latest.jsonl at it.
package debuglog

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
	"unicode/utf8"
)

const (
	maxMsg    = 2000
	maxField  = 500
	maxBytes  = 32 << 20
	maxEvents = 200
)

// Event is one JSONL record.
type Event struct {
	TS     string         `json:"ts"`
	Level  string         `json:"level"`
	Src    string         `json:"src"`
	Module string         `json:"module,omitempty"`
	Action string         `json:"action,omitempty"`
	Msg    string         `json:"msg,omitempty"`
	Err    string         `json:"err,omitempty"`
	Fields map[string]any `json:"fields,omitempty"`
}

type session struct {
	mu   sync.Mutex
	dir  string
	path string
	f    *os.File
	enc  *json.Encoder
	n    int64
}

var (
	globalMu sync.Mutex
	current  *session
)

// Open starts a new session file under dir. Safe to call once from main.
func Open(dir string) (string, error) {
	if err := os.MkdirAll(dir, 0o777); err != nil {
		dir = filepath.Join(os.TempDir(), "aura-debug")
		if err2 := os.MkdirAll(dir, 0o777); err2 != nil {
			return "", err
		}
	}
	_ = os.Chmod(dir, 0o777)
	name := fmt.Sprintf("session-%s.jsonl", time.Now().Format("20060102-150405"))
	path := filepath.Join(dir, name)
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err != nil {
		dir = filepath.Join(os.TempDir(), "aura-debug")
		_ = os.MkdirAll(dir, 0o777)
		path = filepath.Join(dir, name)
		f, err = os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
		if err != nil {
			return "", err
		}
	}
	s := &session{dir: dir, path: path, f: f, enc: json.NewEncoder(f)}
	s.enc.SetEscapeHTML(false)
	if st, err := f.Stat(); err == nil {
		s.n = st.Size()
	}
	globalMu.Lock()
	if current != nil {
		_ = current.closeLocked()
	}
	current = s
	globalMu.Unlock()
	writePointers(dir, path)
	Write(Event{
		Level:  "info",
		Src:    "go",
		Module: "debuglog",
		Action: "session_start",
		Msg:    "debug session open",
		Fields: map[string]any{"file": path, "pid": os.Getpid()},
	})
	return path, nil
}

// Path returns the active session file, or empty if logging is off.
func Path() string {
	globalMu.Lock()
	defer globalMu.Unlock()
	if current == nil {
		return ""
	}
	return current.path
}

// Close flushes and closes the session file.
func Close() {
	globalMu.Lock()
	defer globalMu.Unlock()
	if current == nil {
		return
	}
	WriteUnlocked(Event{
		Level:  "info",
		Src:    "go",
		Module: "debuglog",
		Action: "session_end",
		Msg:    "debug session close",
	})
	_ = current.closeLocked()
	current = nil
}

func (s *session) closeLocked() error {
	if s == nil || s.f == nil {
		return nil
	}
	err := s.f.Sync()
	cerr := s.f.Close()
	s.f = nil
	if err != nil {
		return err
	}
	return cerr
}

func writePointers(dir, path string) {
	_ = os.WriteFile(filepath.Join(dir, "CURRENT"), []byte(path+"\n"), 0o644)
	latest := filepath.Join(dir, "latest.jsonl")
	_ = os.Remove(latest)
	if err := os.Symlink(filepath.Base(path), latest); err != nil {
		_ = os.WriteFile(filepath.Join(dir, "latest.path"), []byte(path+"\n"), 0o644)
	}
}

// Write appends one event. No-op if Open was not called.
func Write(ev Event) {
	globalMu.Lock()
	defer globalMu.Unlock()
	WriteUnlocked(ev)
}

// WriteUnlocked requires globalMu held.
func WriteUnlocked(ev Event) {
	if current == nil || current.f == nil {
		return
	}
	if ev.TS == "" {
		ev.TS = time.Now().UTC().Format(time.RFC3339Nano)
	}
	if ev.Level == "" {
		ev.Level = "info"
	}
	ev.Msg = clip(ev.Msg, maxMsg)
	ev.Err = clip(ev.Err, maxMsg)
	ev.Action = clip(ev.Action, 200)
	ev.Module = clip(ev.Module, 80)
	ev.Fields = clipFields(ev.Fields)
	if err := current.enc.Encode(ev); err != nil {
		return
	}
	if info, err := current.f.Stat(); err == nil {
		current.n = info.Size()
	}
	if current.n >= maxBytes {
		_ = current.rotateLocked()
	}
}

func (s *session) rotateLocked() error {
	_ = s.f.Sync()
	_ = s.f.Close()
	name := fmt.Sprintf("session-%s.jsonl", time.Now().Format("20060102-150405"))
	path := filepath.Join(s.dir, name)
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		s.f = nil
		return err
	}
	s.path = path
	s.f = f
	s.enc = json.NewEncoder(f)
	s.enc.SetEscapeHTML(false)
	s.n = 0
	writePointers(s.dir, path)
	return nil
}

// Info is a convenience writer.
func Info(module, action, msg string, fields map[string]any) {
	Write(Event{Level: "info", Src: "go", Module: module, Action: action, Msg: msg, Fields: fields})
}

// Warn is a convenience writer.
func Warn(module, action, msg string, fields map[string]any) {
	Write(Event{Level: "warn", Src: "go", Module: module, Action: action, Msg: msg, Fields: fields})
}

// Error is a convenience writer.
func Error(module, action, msg string, err error, fields map[string]any) {
	ev := Event{Level: "error", Src: "go", Module: module, Action: action, Msg: msg, Fields: fields}
	if err != nil {
		ev.Err = err.Error()
	}
	Write(ev)
}

// Ingest writes UI-originated events (already timestamped if the client sent ts).
func Ingest(events []Event) {
	if len(events) > maxEvents {
		events = events[:maxEvents]
	}
	globalMu.Lock()
	defer globalMu.Unlock()
	for i := range events {
		ev := events[i]
		if ev.Src == "" {
			ev.Src = "ui"
		}
		WriteUnlocked(ev)
	}
}

func clip(s string, n int) string {
	if s == "" || len(s) <= n {
		return s
	}
	if n < 4 {
		return s[:n]
	}
	for n > 0 && !utf8.ValidString(s[:n]) {
		n--
	}
	return s[:n] + "…"
}

func clipFields(in map[string]any) map[string]any {
	if len(in) == 0 {
		return nil
	}
	out := make(map[string]any, len(in))
	for k, v := range in {
		switch t := v.(type) {
		case string:
			out[k] = clip(t, maxField)
		case error:
			out[k] = clip(t.Error(), maxField)
		default:
			out[k] = v
		}
	}
	return out
}
