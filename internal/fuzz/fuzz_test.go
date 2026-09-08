package fuzz

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestRunFUZZ(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/admin":
			w.WriteHeader(200)
			_, _ = io.WriteString(w, "ok-admin")
		default:
			w.WriteHeader(404)
			_, _ = io.WriteString(w, "no")
		}
	}))
	t.Cleanup(srv.Close)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	hits, err := Run(ctx, Options{
		URL:      srv.URL + "/FUZZ",
		Wordlist: []string{"admin", "missing"},
		Hide:     []int{404},
		Workers:  2,
		RPS:      20,
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(hits) != 1 || hits[0].Payload != "admin" || hits[0].StatusCode != 200 {
		t.Fatalf("hits=%+v", hits)
	}
}

func TestRunRequiresFUZZ(t *testing.T) {
	_, err := Run(context.Background(), Options{URL: "http://127.0.0.1/", Wordlist: []string{"a"}})
	if err == nil || !strings.Contains(err.Error(), "FUZZ") {
		t.Fatalf("err=%v", err)
	}
}
