package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"meb/internal/store"
)

func TestOriginMissingAllowed(t *testing.T) {
	h := New(Options{Runtime: store.New(), Log: zap.NewNop()})
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/scanner/scan", strings.NewReader(`{"raw":""}`))
	h.ServeHTTP(rec, req)
	require.Equal(t, 400, rec.Code, rec.Body.String())
}

func TestOriginLoopbackAllowed(t *testing.T) {
	h := New(Options{Runtime: store.New(), Log: zap.NewNop()})
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/scanner/scan", strings.NewReader(`{"raw":""}`))
	req.Header.Set("Origin", "http://127.0.0.1:1337")
	h.ServeHTTP(rec, req)
	require.Equal(t, 400, rec.Code, rec.Body.String())
}

func TestOriginEvilRejected(t *testing.T) {
	h := New(Options{Runtime: store.New(), Log: zap.NewNop()})
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/scanner/scan", strings.NewReader(`{"raw":"GET / HTTP/1.1\r\nHost: x\r\n\r\n"}`))
	req.Header.Set("Origin", "https://evil.example")
	h.ServeHTTP(rec, req)
	require.Equal(t, 403, rec.Code, rec.Body.String())
}

func TestOriginCrossSiteFetchRejected(t *testing.T) {
	h := New(Options{Runtime: store.New(), Log: zap.NewNop()})
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/scanner/scan", strings.NewReader(`{"raw":""}`))
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	h.ServeHTTP(rec, req)
	require.Equal(t, 403, rec.Code, rec.Body.String())
}

func TestOriginLocalhostAllowed(t *testing.T) {
	h := New(Options{Runtime: store.New(), Log: zap.NewNop()})
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/status", nil)
	req.Header.Set("Origin", "http://localhost:1337")
	h.ServeHTTP(rec, req)
	require.Equal(t, 200, rec.Code, rec.Body.String())
}

func TestOriginWrongPortRejected(t *testing.T) {
	h := New(Options{Runtime: store.New(), Log: zap.NewNop()})
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/proxy/stop", strings.NewReader(""))
	req.Header.Set("Origin", "http://127.0.0.1:8080")
	h.ServeHTTP(rec, req)
	require.Equal(t, 403, rec.Code, rec.Body.String())
}
