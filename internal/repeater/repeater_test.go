package repeater

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseRawRequest(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		raw     string
		scheme  string
		target  string
		wantErr string
		check   func(t *testing.T, raw string, scheme, target string)
	}{
		{
			name:   "get with host",
			raw:    "GET /path HTTP/1.1\r\nHost: example.com\r\n\r\n",
			scheme: "https",
			check: func(t *testing.T, raw, scheme, target string) {
				req, err := ParseRawRequest(raw, scheme, target)
				require.NoError(t, err)
				assert.Equal(t, "GET", req.Method)
				assert.Equal(t, "/path", req.Path)
				assert.Equal(t, "example.com", req.Host)
				assert.Equal(t, 443, req.Port)
				assert.Equal(t, "0", req.Header("Content-Length"))
			},
		},
		{
			name:   "lf newlines and body",
			raw:    "POST /submit HTTP/1.1\nHost: 127.0.0.1:8080\n\nhello",
			scheme: "http",
			check: func(t *testing.T, raw, scheme, target string) {
				req, err := ParseRawRequest(raw, scheme, target)
				require.NoError(t, err)
				assert.Equal(t, "POST", req.Method)
				assert.Equal(t, "hello", string(req.Body))
				assert.Equal(t, 8080, req.Port)
				assert.Equal(t, "5", req.Header("Content-Length"))
			},
		},
		{
			name:   "target overrides host",
			raw:    "GET / HTTP/1.1\r\nHost: example.com\r\n\r\n",
			scheme: "http",
			target: "127.0.0.1:9999",
			check: func(t *testing.T, raw, scheme, target string) {
				req, err := ParseRawRequest(raw, scheme, target)
				require.NoError(t, err)
				assert.Equal(t, "127.0.0.1", req.Host)
				assert.Equal(t, 9999, req.Port)
				assert.Equal(t, "127.0.0.1:9999", req.Header("Host"))
			},
		},
		{
			name:    "missing separator",
			raw:     "GET / HTTP/1.1\r\nHost: example.com",
			scheme:  "http",
			wantErr: "missing header/body separator",
		},
		{
			name:    "missing host",
			raw:     "GET / HTTP/1.1\r\nUser-Agent: meb\r\n\r\n",
			scheme:  "http",
			wantErr: "Host header required",
		},
		{
			name:    "bad scheme",
			raw:     "GET / HTTP/1.1\r\nHost: example.com\r\n\r\n",
			scheme:  "ftp",
			wantErr: "scheme must be http or https",
		},
		{
			name:    "bad request line",
			raw:     "NOPE\r\n\r\n",
			scheme:  "http",
			wantErr: "malformed request line",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.wantErr != "" {
				_, err := ParseRawRequest(tt.raw, tt.scheme, tt.target)
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr)
				return
			}
			tt.check(t, tt.raw, tt.scheme, tt.target)
		})
	}
}

func TestSendHTTP(t *testing.T) {
	t.Parallel()
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/ping", r.URL.Path)
		w.Header().Set("Content-Type", "text/plain")
		_, _ = w.Write([]byte("pong"))
	}))
	t.Cleanup(upstream.Close)

	host := strings.TrimPrefix(upstream.URL, "http://")
	raw := fmt.Sprintf("GET /ping HTTP/1.1\r\nHost: %s\r\nAccept: */*\r\n\r\n", host)
	res := Send(context.Background(), Options{
		Raw:     raw,
		Scheme:  "http",
		Verify:  true,
		Timeout: 5 * time.Second,
	})
	require.True(t, res.OK, res.Error)
	assert.Equal(t, 200, res.Status)
	assert.Contains(t, res.ResponseRaw, "pong")
}

func TestSendCanceledContext(t *testing.T) {
	t.Parallel()
	started := make(chan struct{})
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		close(started)
		time.Sleep(2 * time.Second)
		_, _ = w.Write([]byte("slow"))
	}))
	t.Cleanup(upstream.Close)

	host := strings.TrimPrefix(upstream.URL, "http://")
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-started
		cancel()
	}()
	res := Send(ctx, Options{
		Raw:     fmt.Sprintf("GET / HTTP/1.1\r\nHost: %s\r\n\r\n", host),
		Scheme:  "http",
		Timeout: 5 * time.Second,
	})
	assert.False(t, res.OK)
	assert.NotEmpty(t, res.Error)
}
