package debuglog

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type fileCore struct {
	fields []zapcore.Field
	min    zapcore.Level
}

// FileCore is a zap core that mirrors records into the session JSONL file.
func FileCore() zapcore.Core {
	return fileCore{min: zapcore.DebugLevel}
}

// Wrap tees an existing logger to the session file. Caller info is preserved.
func Wrap(base *zap.Logger) *zap.Logger {
	if base == nil {
		base = zap.NewNop()
	}
	return zap.New(zapcore.NewTee(base.Core(), FileCore()), zap.AddCaller())
}

func (c fileCore) Enabled(lvl zapcore.Level) bool { return lvl >= c.min }

func (c fileCore) With(fs []zapcore.Field) zapcore.Core {
	all := make([]zapcore.Field, 0, len(c.fields)+len(fs))
	all = append(all, c.fields...)
	all = append(all, fs...)
	return fileCore{fields: all, min: c.min}
}

func (c fileCore) Check(e zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if c.Enabled(e.Level) {
		return ce.AddCore(e, c)
	}
	return ce
}

func (c fileCore) Write(e zapcore.Entry, fs []zapcore.Field) error {
	enc := zapcore.NewMapObjectEncoder()
	for _, f := range c.fields {
		f.AddTo(enc)
	}
	for _, f := range fs {
		f.AddTo(enc)
	}
	fields := enc.Fields
	errMsg := ""
	if v, ok := fields["error"]; ok {
		errMsg = clip(stringify(v), maxMsg)
		delete(fields, "error")
	}
	module := e.LoggerName
	if module == "" && e.Caller.Defined {
		module = e.Caller.TrimmedPath()
	}
	Write(Event{
		TS:     e.Time.UTC().Format(time.RFC3339Nano),
		Level:  e.Level.String(),
		Src:    "go",
		Module: module,
		Action: "zap",
		Msg:    e.Message,
		Err:    errMsg,
		Fields: fields,
	})
	return nil
}

func (c fileCore) Sync() error {
	globalMu.Lock()
	defer globalMu.Unlock()
	if current == nil || current.f == nil {
		return nil
	}
	return current.f.Sync()
}

func stringify(v any) string {
	switch t := v.(type) {
	case string:
		return t
	case error:
		return t.Error()
	default:
		return ""
	}
}
