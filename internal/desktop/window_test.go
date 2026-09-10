package desktop

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestWaitReadyOK(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := WaitReady(ctx, srv.URL); err != nil {
		t.Fatal(err)
	}
}

func TestWaitReadyCancel(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	err := WaitReady(ctx, "http://127.0.0.1:1")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestIsRunning(t *testing.T) {
	if IsRunning("http://127.0.0.1:1") {
		t.Fatal("expected not running")
	}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/status" {
			t.Errorf("path %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()
	if !IsRunning(srv.URL) {
		t.Fatal("expected running")
	}
}
