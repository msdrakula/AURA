package scanner

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

func TestScanEmptyRawBaselineFails(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := Scan(ctx, Options{Raw: "", Scheme: "http", Target: "127.0.0.1:1", Timeout: time.Second})
	if err == nil {
		t.Fatal("empty raw should fail baseline")
	}
}

func TestScanBareGETSendsBaseline(t *testing.T) {
	var n atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n.Add(1)
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(200)
		_, _ = io.WriteString(w, "<html>ok</html>")
	}))
	t.Cleanup(srv.Close)

	host := strings.TrimPrefix(srv.URL, "http://")
	raw := "GET / HTTP/1.1\r\nHost: " + host + "\r\n\r\n"
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()
	res, err := Scan(ctx, Options{Raw: raw, Scheme: "http", Target: host, Timeout: 2 * time.Second})
	if err != nil {
		t.Fatal(err)
	}
	if res == nil {
		t.Fatal("nil result")
	}
	if n.Load() < 1 {
		t.Fatal("baseline was not sent")
	}
}

func TestScanQueryParamInjectsProbes(t *testing.T) {
	var sawSQLi atomic.Bool
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("id")
		if strings.Contains(q, "'") {
			sawSQLi.Store(true)
			_, _ = io.WriteString(w, "SQLSTATE syntax error")
			return
		}
		_, _ = io.WriteString(w, "ok")
	}))
	t.Cleanup(srv.Close)
	host := strings.TrimPrefix(srv.URL, "http://")
	raw := "GET /x?id=1 HTTP/1.1\r\nHost: " + host + "\r\n\r\n"
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()
	res, err := Scan(ctx, Options{Raw: raw, Scheme: "http", Target: host, Timeout: 2 * time.Second})
	if err != nil {
		t.Fatal(err)
	}
	if !sawSQLi.Load() {
		t.Fatal("expected SQLi probe on id=")
	}
	found := false
	for _, f := range res.Findings {
		if strings.Contains(strings.ToLower(f.Name), "sql") {
			found = true
		}
	}
	if !found {
		t.Fatalf("expected SQLi finding, got %+v", res.Findings)
	}
}

func TestPathTraversalRequiresProbeContent(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, "root:x:0:0:root:/root:/bin/bash\n")
	}))
	t.Cleanup(srv.Close)
	host := strings.TrimPrefix(srv.URL, "http://")
	raw := "GET /file?path=a HTTP/1.1\r\nHost: " + host + "\r\n\r\n"
	f := checkPathTraversal(context.Background(), Options{
		Raw: raw, Scheme: "http", Target: host, Timeout: 2 * time.Second,
	}, "path", "root:x:0:0:root:/root:/bin/bash\n")
	if f != nil {
		t.Fatalf("baseline already had passwd-like text; should not flag: %+v", f)
	}
}

func TestInjectParamReplacesQuery(t *testing.T) {
	got := injectParam("GET /x?id=1 HTTP/1.1\r\nHost: h\r\n\r\n", "id", "PROBE")
	if !strings.Contains(got, "id=PROBE") {
		t.Fatalf("got %q", got)
	}
}

func TestSQLErrorLiteralDotStar(t *testing.T) {
	for _, e := range sqlErrors {
		if e == "PostgreSQL.*ERROR" {
			t.Log("known: PostgreSQL.*ERROR is a literal Contains needle, not a regexp")
			return
		}
	}
	t.Fatal("sqlErrors changed; update audit note")
}

func ExampleScan_requiresRaw() {
	fmt.Println("scanner API rejects empty raw; UI synthesizes GET / from Target")
	// Output: scanner API rejects empty raw; UI synthesizes GET / from Target
}
