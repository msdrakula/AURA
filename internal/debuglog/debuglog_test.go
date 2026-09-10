package debuglog

import (
	"bufio"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestSessionWriteAndLatest(t *testing.T) {
	dir := t.TempDir()
	path, err := Open(dir)
	require.NoError(t, err)
	defer Close()

	Write(Event{Level: "error", Src: "test", Module: "proxy", Action: "handshake", Msg: "tls failed", Err: "eof"})
	Info("intel", "stage_start", "dirs", map[string]any{"target": "lab.local"})
	Close()

	raw, err := os.ReadFile(path)
	require.NoError(t, err)
	require.Contains(t, string(raw), `"action":"handshake"`)
	require.Contains(t, string(raw), `"module":"intel"`)

	latest := filepath.Join(dir, "latest.jsonl")
	got, err := os.ReadFile(latest)
	require.NoError(t, err)
	require.Equal(t, raw, got)

	cur, err := os.ReadFile(filepath.Join(dir, "CURRENT"))
	require.NoError(t, err)
	require.Contains(t, string(cur), path)
}

func TestIngestAndZap(t *testing.T) {
	dir := t.TempDir()
	path, err := Open(dir)
	require.NoError(t, err)
	t.Cleanup(Close)

	Ingest([]Event{{Level: "info", Action: "click", Msg: "btnStart", Module: "ui"}})

	core := zapcore.NewNopCore()
	log := Wrap(zap.New(core))
	log.Error("scanner failed", zap.String("host", "app.lab.local"), zap.Error(os.ErrNotExist))
	require.NoError(t, log.Sync())
	Close()

	f, err := os.Open(path)
	require.NoError(t, err)
	defer f.Close()
	var sawUI, sawZap bool
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		var ev Event
		require.NoError(t, json.Unmarshal(sc.Bytes(), &ev))
		if ev.Action == "click" && ev.Src == "ui" {
			sawUI = true
		}
		if ev.Msg == "scanner failed" {
			sawZap = true
			require.Equal(t, "error", ev.Level)
			require.Contains(t, ev.Err, "file does not exist")
		}
	}
	require.NoError(t, sc.Err())
	require.True(t, sawUI)
	require.True(t, sawZap)
}

func TestWriteNoOpenIsNoop(t *testing.T) {
	Close()
	Write(Event{Msg: "should not panic"})
}
