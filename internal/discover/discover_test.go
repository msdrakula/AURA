package discover

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"go.uber.org/zap"
)

func TestRunFindsPath(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/admin" {
			w.WriteHeader(200)
			_, _ = io.WriteString(w, "ok")
			return
		}
		w.WriteHeader(404)
	}))
	t.Cleanup(srv.Close)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := Run(ctx, Options{
		BaseURL:  srv.URL,
		Wordlist: []string{"admin", "missing"},
		Workers:  2,
		RPS:      20,
		Timeout:  2 * time.Second,
	}, zap.NewNop())
	if err != nil && ctx.Err() == nil {
		t.Fatal(err)
	}
	found := false
	for _, r := range res {
		if r.Path == "admin" && r.StatusCode == 200 {
			found = true
		}
	}
	if !found {
		t.Fatalf("admin not found: %+v", res)
	}
}

func TestRunCancelStopsEnqueue(t *testing.T) {
	var n atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n.Add(1)
		time.Sleep(40 * time.Millisecond)
		w.WriteHeader(404)
	}))
	t.Cleanup(srv.Close)

	words := make([]string, 400)
	for i := range words {
		words[i] = "p" + itoa(i)
	}
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(80 * time.Millisecond)
		cancel()
	}()
	start := time.Now()
	_, _ = Run(ctx, Options{
		BaseURL:  srv.URL,
		Wordlist: words,
		Workers:  2,
		RPS:      50,
		Timeout:  2 * time.Second,
	}, zap.NewNop())
	if time.Since(start) > 4*time.Second {
		t.Fatal("cancel did not return promptly")
	}
	if n.Load() > 80 {
		t.Fatalf("too many probes after cancel: %d", n.Load())
	}
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var b [12]byte
	i := len(b)
	for n > 0 {
		i--
		b[i] = byte('0' + n%10)
		n /= 10
	}
	return string(b[i:])
}
