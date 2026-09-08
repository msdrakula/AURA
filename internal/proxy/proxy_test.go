package proxy

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"

	"meb/internal/analyzer"
	"meb/internal/certs"
	"meb/internal/httpio"
	"meb/internal/models"
	"meb/internal/storage"
	"meb/internal/store"
)

func TestApplyRequestEdit(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		raw    string
		scheme string
		method string
		path   string
		host   string
		port   int
		body   string
	}{
		{
			name:   "post with body",
			raw:    "POST /api HTTP/1.1\r\nHost: example.com\r\n\r\nhello",
			scheme: "https",
			method: "POST",
			path:   "/api",
			host:   "example.com",
			port:   443,
			body:   "hello",
		},
		{
			name:   "host with port",
			raw:    "GET /x HTTP/1.1\r\nHost: 127.0.0.1:9000\r\n\r\n",
			scheme: "http",
			method: "GET",
			path:   "/x",
			host:   "127.0.0.1",
			port:   9000,
			body:   "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			req := &httpio.Request{Method: "GET", Path: "/", Version: "HTTP/1.1", Scheme: tt.scheme}
			applyRequestRaw(req, tt.raw)
			assert.Equal(t, tt.method, req.Method)
			assert.Equal(t, tt.path, req.Path)
			assert.Equal(t, tt.host, req.Host)
			assert.Equal(t, tt.port, req.Port)
			assert.Equal(t, tt.body, string(req.Body))
			assert.Equal(t, fmt.Sprintf("%d", len(req.Body)), req.Header("Content-Length"))
		})
	}
}

func TestApplyResponseEdit(t *testing.T) {
	t.Parallel()
	resp := &httpio.Response{Version: "HTTP/1.1", Status: 200, Reason: "OK"}
	applyResponseRaw(resp, "HTTP/1.1 201 Created\r\nContent-Type: text/plain\r\n\r\nxyz")
	assert.Equal(t, 201, resp.Status)
	assert.Equal(t, "Created", resp.Reason)
	assert.Equal(t, "xyz", string(resp.Body))
	assert.Equal(t, "3", resp.Header("Content-Length"))
}

func TestProxyHTTPForward(t *testing.T) {
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/hello", r.URL.Path)
		w.Header().Set("Content-Type", "text/plain")
		_, _ = w.Write([]byte("world"))
	}))
	t.Cleanup(upstream.Close)

	st := store.New()
	st.Settings.ListenHost = "127.0.0.1"
	st.Settings.ListenPort = 0
	db := &storage.Store{}
	require.NoError(t, db.Init(filepath.Join(t.TempDir(), "session.db")))
	t.Cleanup(func() { _ = db.Close() })

	px := New(Options{
		Runtime:  st,
		History:  db,
		Findings: db,
		Analyzer: analyzer.DefaultEngine(),
		Certs:    certs.NewAuthority(t.TempDir()),
		Log:      zap.NewNop(),
	})
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)
	require.NoError(t, px.Start(ctx))
	t.Cleanup(func() { _ = px.Stop(context.Background()) })

	proxyURL, err := url.Parse("http://" + px.Addr())
	require.NoError(t, err)
	client := &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		},
	}
	resp, err := client.Get(upstream.URL + "/hello")
	require.NoError(t, err)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "world", string(body))

	hist, err := db.GetTransactions(10, 0)
	require.NoError(t, err)
	deadline := time.Now().Add(2 * time.Second)
	for len(hist) == 0 && time.Now().Before(deadline) {
		time.Sleep(10 * time.Millisecond)
		hist, err = db.GetTransactions(10, 0)
		require.NoError(t, err)
	}
	require.NotEmpty(t, hist)
	assert.Equal(t, "GET", hist[0].Request.Method)
	assert.Empty(t, st.History())
}

func TestProxySaveErrorsDoNotBreakForward(t *testing.T) {
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok"))
	}))
	t.Cleanup(upstream.Close)

	st := store.New()
	st.Settings.ListenHost = "127.0.0.1"
	st.Settings.ListenPort = 0
	core, logs := observer.New(zap.ErrorLevel)
	px := New(Options{
		Runtime:  st,
		History:  failWriter{},
		Findings: failWriter{},
		Analyzer: analyzer.DefaultEngine(),
		Certs:    certs.NewAuthority(t.TempDir()),
		Log:      zap.New(core),
	})
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)
	require.NoError(t, px.Start(ctx))
	t.Cleanup(func() { _ = px.Stop(context.Background()) })

	proxyURL, err := url.Parse("http://" + px.Addr())
	require.NoError(t, err)
	client := &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		},
	}
	resp, err := client.Get(upstream.URL + "/")
	require.NoError(t, err)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "ok", string(body))
	require.NotEmpty(t, logs.All())
}

func TestProxyStopUnbinds(t *testing.T) {
	st := store.New()
	st.Settings.ListenHost = "127.0.0.1"
	st.Settings.ListenPort = 0
	px := New(Options{Runtime: st, Certs: certs.NewAuthority(t.TempDir()), Log: zap.NewNop()})
	ctx := context.Background()
	require.NoError(t, px.Start(ctx))
	require.NotEmpty(t, px.Addr())
	require.NoError(t, px.Stop(ctx))
	assert.Empty(t, px.Addr())
	assert.False(t, st.IsRunning())
	require.NoError(t, px.Start(ctx))
	t.Cleanup(func() { _ = px.Stop(ctx) })
	assert.NotEmpty(t, px.Addr())
}

type failWriter struct{}

func (failWriter) SaveTransaction(*models.HTTPTransaction) error {
	return errors.New("db down")
}

func (failWriter) SaveFinding(*models.AnalysisFinding) error {
	return errors.New("db down")
}
