package batch

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEngineExecute(t *testing.T) {
	t.Parallel()

	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/test", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	}))
	t.Cleanup(upstream.Close)

	parsed, err := url.Parse(upstream.URL)
	require.NoError(t, err)
	target := parsed.Host

	eng := NewEngine(5, 100)

	template := "GET /test?id=§x§&name=§y§ HTTP/1.1\r\nHost: " + target + "\r\nConnection: close\r\n\r\n"
	results, err := eng.Execute(context.Background(), template, "http", target, [][]string{
		{"a", "b"},
		{"1", "2"},
	})
	require.NoError(t, err)
	require.Len(t, results, 4)

	for _, res := range results {
		assert.Equal(t, 200, res.StatusCode, "payloads=%v err=%q", res.Payloads, res.Error)
		assert.Empty(t, res.Error)
		assert.Equal(t, 2, res.BodyLength, "body length for %v", res.Payloads)
		assert.Len(t, res.Payloads, 2)
	}

	combos := map[string]bool{}
	for _, res := range results {
		combos[strings.Join(res.Payloads, "|")] = true
	}
	assert.True(t, combos["a|1"])
	assert.True(t, combos["a|2"])
	assert.True(t, combos["b|1"])
	assert.True(t, combos["b|2"])
}

func TestEngineExecuteEmptyPayloads(t *testing.T) {
	t.Parallel()
	eng := NewEngine(2, 10)
	results, err := eng.Execute(context.Background(), "GET / HTTP/1.1\r\nHost: x\r\n\r\n", "http", "", nil)
	require.NoError(t, err)
	assert.Empty(t, results)
}

func TestExecuteCancelContext(t *testing.T) {
	t.Parallel()

	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(50 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	t.Cleanup(upstream.Close)

	parsed, err := url.Parse(upstream.URL)
	require.NoError(t, err)
	target := parsed.Host

	eng := NewEngine(2, 1).WithTimeout(5 * time.Second)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	template := "GET /?q=§x§ HTTP/1.1\r\nHost: " + target + "\r\n\r\n"
	_, err = eng.Execute(ctx, template, "http", target, [][]string{{"a", "b", "c", "d"}})
	require.Error(t, err)
	assert.ErrorIs(t, err, context.Canceled)
}
