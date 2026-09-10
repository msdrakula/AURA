package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestScannerEmptyRawIs400(t *testing.T) {
	h := New(Options{Log: zap.NewNop()})
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/scanner/scan", strings.NewReader(`{"scheme":"http","target":"127.0.0.1:9"}`))
	h.ServeHTTP(rec, req)
	require.Equal(t, 400, rec.Code, rec.Body.String())
	require.Contains(t, rec.Body.String(), "raw")
}

func TestScannerInvalidJSONIs400(t *testing.T) {
	h := New(Options{Log: zap.NewNop()})
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/scanner/scan", strings.NewReader(`{`))
	h.ServeHTTP(rec, req)
	require.Equal(t, 400, rec.Code)
}
