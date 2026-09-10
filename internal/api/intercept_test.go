package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"meb/internal/httpio"
	"meb/internal/store"
)

func TestInterceptDecideUnblocksProxy(t *testing.T) {
	st := store.New()
	st.Settings.InterceptRequests = true
	h := New(Options{Runtime: st, Log: zap.NewNop()})

	req := &httpio.Request{
		Method:  "GET",
		Path:    "/",
		Host:    "example.com",
		Scheme:  "http",
		Version: "HTTP/1.1",
	}
	done := make(chan string, 1)
	go func() {
		action, _ := st.InterceptRequest("flow1", req)
		done <- action
	}()

	var id string
	deadline := time.Now().Add(2 * time.Second)
	for time.Now().Before(deadline) {
		items := st.PendingList()
		if len(items) > 0 {
			id, _ = items[0]["id"].(string)
			if id != "" {
				break
			}
		}
		time.Sleep(10 * time.Millisecond)
	}
	require.NotEmpty(t, id, "pending intercept never appeared")

	rec := httptest.NewRecorder()
	httpReq := httptest.NewRequest(http.MethodPost, "/api/intercept/"+id, strings.NewReader(`{"action":"forward","raw":""}`))
	h.ServeHTTP(rec, httpReq)
	require.Equal(t, 200, rec.Code, rec.Body.String())

	select {
	case action := <-done:
		require.Equal(t, "forward", action)
	case <-time.After(2 * time.Second):
		t.Fatal("forward did not unblock intercept")
	}
}
