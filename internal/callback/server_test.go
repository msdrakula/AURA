package callback

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestServerLogsInteractions(t *testing.T) {
	t.Parallel()

	srv := NewServer(0, zap.NewNop())
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	startErr := make(chan error, 1)
	go func() { startErr <- srv.Start(ctx) }()
	t.Cleanup(func() {
		cancel()
		select {
		case <-startErr:
		case <-time.After(3 * time.Second):
		}
	})

	deadline := time.Now().Add(2 * time.Second)
	for srv.Addr() == "" && time.Now().Before(deadline) {
		time.Sleep(5 * time.Millisecond)
	}
	require.NotEmpty(t, srv.Addr(), "server did not start")

	client := &http.Client{Timeout: 3 * time.Second}
	for i, body := range []string{`{"event":"first"}`, `{"event":"second"}`} {
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, "http://"+srv.Addr()+"/test-session-99", strings.NewReader(body))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		resp, err := client.Do(req)
		require.NoError(t, err, "request %d", i)
		raw, _ := io.ReadAll(resp.Body)
		_ = resp.Body.Close()
		assert.Equal(t, 200, resp.StatusCode, "request %d", i)
		assert.Equal(t, "OK", string(raw), "request %d body", i)
	}

	got := srv.GetInteractions("test-session-99")
	require.Len(t, got, 2)
	for _, it := range got {
		assert.Equal(t, "POST", it.Method)
		assert.Contains(t, it.URL, "/test-session-99")
		assert.False(t, it.Timestamp.IsZero())
	}
	assert.Contains(t, got[0].Body, "first")
	assert.Contains(t, got[1].Body, "second")
	assert.NotEmpty(t, got[0].Headers["Content-Type"])

	assert.Nil(t, srv.GetInteractions("nonexistent"))
}

func TestExtractRequestID(t *testing.T) {
	t.Parallel()
	tests := []struct {
		path string
		want string
	}{
		{"/abc-123/path", "abc-123"},
		{"/abc-123", "abc-123"},
		{"/session-x/sub/deep", "session-x"},
	}
	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			assert.Equal(t, tt.want, extractRequestID(tt.path))
		})
	}
	t.Run("root generates uuid", func(t *testing.T) {
		got := extractRequestID("/")
		assert.NotEmpty(t, got)
		assert.Len(t, got, 36)
	})
}
