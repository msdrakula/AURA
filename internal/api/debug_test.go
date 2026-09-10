package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"meb/internal/debuglog"
	"meb/internal/store"
)

func TestDebugLogIngest(t *testing.T) {
	dir := t.TempDir()
	path, err := debuglog.Open(dir)
	require.NoError(t, err)
	t.Cleanup(debuglog.Close)

	h := New(Options{Runtime: store.New(), Log: zap.NewNop()})
	body, _ := json.Marshal(map[string]any{
		"events": []map[string]any{{
			"level":  "error",
			"module": "ui",
			"action": "click",
			"msg":    "btnFuzzRun",
		}},
	})
	req := httptest.NewRequest(http.MethodPost, "/api/debug-log", bytes.NewReader(body))
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	require.Equal(t, 200, rec.Code)

	raw, err := os.ReadFile(path)
	require.NoError(t, err)
	require.Contains(t, string(raw), "btnFuzzRun")
	require.Contains(t, string(raw), `"src":"ui"`)
}

func TestDebugHTTPWarnOnClientError(t *testing.T) {
	dir := t.TempDir()
	path, err := debuglog.Open(dir)
	require.NoError(t, err)
	t.Cleanup(debuglog.Close)

	h := New(Options{Runtime: store.New(), Log: zap.NewNop()})
	req := httptest.NewRequest(http.MethodPost, "/api/codec", bytes.NewReader([]byte(`{`)))
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	require.Equal(t, 400, rec.Code)

	raw, err := os.ReadFile(path)
	require.NoError(t, err)
	require.Contains(t, string(raw), "/api/codec")
	require.Contains(t, string(raw), `"status":400`)
}
