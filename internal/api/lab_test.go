package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestActiveToolsRequireLabConfirm(t *testing.T) {
	h := New(Options{Log: zap.NewNop()})
	cases := []struct {
		name string
		path string
		body string
	}{
		{"discover", "/api/discover/run", `{"base_url":"https://lab.local","wordlist":["admin"]}`},
		{"fuzz", "/api/fuzz/run", `{"url":"https://lab.local/FUZZ","wordlist":["a"]}`},
		{"scanner", "/api/scanner/scan", "{\"raw\":\"GET / HTTP/1.1\\r\\nHost: lab.local\\r\\n\\r\\n\"}"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, tc.path, strings.NewReader(tc.body))
			h.ServeHTTP(rec, req)
			require.Equal(t, 400, rec.Code, rec.Body.String())
			require.Contains(t, rec.Body.String(), "confirmation")
		})
	}
}
