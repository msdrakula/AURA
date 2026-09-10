package intel

import "time"

// ProgressEvent is one live line while a map stage runs.
type ProgressEvent struct {
	Type    string    `json:"type"`
	Stage   string    `json:"stage,omitempty"`
	Action  string    `json:"action,omitempty"`
	Msg     string    `json:"msg,omitempty"`
	URL     string    `json:"url,omitempty"`
	Host    string    `json:"host,omitempty"`
	Value   string    `json:"value,omitempty"`
	Err     string    `json:"err,omitempty"`
	N       int       `json:"n,omitempty"`
	Total   int       `json:"total,omitempty"`
	Status  int       `json:"status,omitempty"`
	Timeout int       `json:"timeout,omitempty"`
	TS      time.Time `json:"ts"`
}

func (e *Engine) emit(ev ProgressEvent) {
	if e == nil || e.Progress == nil {
		return
	}
	ev.Type = "intel_progress"
	if ev.TS.IsZero() {
		ev.TS = time.Now().UTC()
	}
	e.Progress(ev)
}

func (e *Engine) emitTick(stage string, n, total, every int, ev ProgressEvent) {
	if n != 1 && n != total && (every <= 0 || n%every != 0) {
		return
	}
	if ev.Action == "" {
		ev.Action = "tick"
	}
	ev.Stage = stage
	ev.N = n
	ev.Total = total
	e.emit(ev)
}
