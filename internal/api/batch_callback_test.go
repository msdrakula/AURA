package api

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"meb/internal/batch"
	"meb/internal/callback"
	"meb/internal/models"
	"meb/internal/store"
)

func TestBatchExecuteEndpoint(t *testing.T) {
	t.Parallel()

	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	}))
	t.Cleanup(upstream.Close)

	target := strings.TrimPrefix(upstream.URL, "http://")
	eng := batch.NewEngine(5, 100).WithTimeout(5 * time.Second)

	runtime := store.New()
	handler := New(Options{
		Runtime: runtime,
		Batch:   eng,
		Log:     zap.NewNop(),
	})

	body := BatchExecuteRequest{
		TemplateRaw: "GET /test?id=§x§&name=§y§ HTTP/1.1\r\nHost: " + target + "\r\nConnection: close\r\n\r\n",
		Scheme:      "http",
		Target:      target,
		PayloadSets: [][]string{{"a", "b"}, {"1", "2"}},
	}
	raw, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/api/batch/execute", bytes.NewReader(raw))
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	require.Equal(t, 200, rec.Code, "body: %s", rec.Body.String())
	var results []models.BatchResult
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &results))
	require.Len(t, results, 4)
	for _, r := range results {
		assert.Equal(t, 200, r.StatusCode)
		assert.Empty(t, r.Error)
	}
}

func TestBatchExecuteValidation(t *testing.T) {
	t.Parallel()
	handler := New(Options{Runtime: store.New(), Batch: batch.NewEngine(1, 1), Log: zap.NewNop()})

	tests := []struct {
		name string
		body string
		want int
	}{
		{"empty template", `{"scheme":"http"}`, 400},
		{"bad scheme", `{"template_raw":"GET / HTTP/1.1\r\nHost: x\r\n\r\n","scheme":"ftp"}`, 400},
		{"invalid json", `{not json`, 400},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/api/batch/execute", strings.NewReader(tt.body))
			rec := httptest.NewRecorder()
			handler.ServeHTTP(rec, req)
			assert.Equal(t, tt.want, rec.Code)
		})
	}
}

func TestCallbackEndpoint(t *testing.T) {
	t.Parallel()

	cb := callback.NewServer(0, zap.NewNop())
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	startErr := make(chan error, 1)
	go func() { startErr <- cb.Start(ctx) }()
	t.Cleanup(func() {
		cancel()
		select {
		case <-startErr:
		case <-time.After(3 * time.Second):
		}
	})

	deadline := time.Now().Add(2 * time.Second)
	for cb.Addr() == "" && time.Now().Before(deadline) {
		time.Sleep(5 * time.Millisecond)
	}
	require.NotEmpty(t, cb.Addr())

	client := &http.Client{Timeout: 3 * time.Second}
	for _, b := range []string{`{"e":1}`, `{"e":2}`} {
		req, _ := http.NewRequestWithContext(ctx, http.MethodPost, "http://"+cb.Addr()+"/sess-42", strings.NewReader(b))
		resp, err := client.Do(req)
		require.NoError(t, err)
		_, _ = io.ReadAll(resp.Body)
		_ = resp.Body.Close()
	}

	handler := New(Options{Runtime: store.New(), Callback: cb, Log: zap.NewNop()})

	req := httptest.NewRequest(http.MethodGet, "/api/callback/sess-42", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	require.Equal(t, 200, rec.Code)
	var items []callback.Interaction
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &items))
	require.Len(t, items, 2)
	for _, it := range items {
		assert.Equal(t, "POST", it.Method)
		assert.Contains(t, it.URL, "/sess-42")
	}
}

func TestCallbackEndpointMissing(t *testing.T) {
	t.Parallel()
	handler := New(Options{Runtime: store.New(), Callback: callback.NewServer(0, zap.NewNop()), Log: zap.NewNop()})

	req := httptest.NewRequest(http.MethodGet, "/api/callback/nope", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	require.Equal(t, 200, rec.Code)
	var items []callback.Interaction
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &items))
	assert.Empty(t, items)
}
