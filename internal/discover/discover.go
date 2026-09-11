// Package discover performs content discovery (dirbusting) against a target.
package discover

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

// Result is one probe result.
type Result struct {
	Path       string        `json:"path"`
	StatusCode int           `json:"status_code"`
	Length     int           `json:"length"`
	Duration   time.Duration `json:"duration"`
	Error      string        `json:"error,omitempty"`
}

// Options configures a discovery run.
type Options struct {
	BaseURL    string   // e.g. https://example.com
	Wordlist   []string // paths to try (without leading slash)
	Workers    int
	RPS        int
	Timeout    time.Duration
	Headers    []Header
	Cookies    string
	OnResult   func(rawReq, respRaw string) // optional per-probe logger
	OnProgress func(done, total int, r Result)
}

// Header is a custom header to add to probes.
type Header struct {
	Name  string
	Value string
}

// Run executes a content discovery scan and streams results to out.
func Run(ctx context.Context, opts Options, log *zap.Logger) ([]Result, error) {
	if !strings.HasPrefix(opts.BaseURL, "http") {
		return nil, fmt.Errorf("base_url must start with http or https")
	}
	if opts.Workers < 1 {
		opts.Workers = 10
	}
	if opts.Workers > 30 {
		opts.Workers = 30
	}
	if opts.RPS < 1 {
		opts.RPS = 20
	}
	if opts.RPS > 50 {
		opts.RPS = 50
	}
	const maxWordlist = 50000
	if len(opts.Wordlist) > maxWordlist {
		opts.Wordlist = opts.Wordlist[:maxWordlist]
	}
	if opts.Timeout <= 0 {
		opts.Timeout = 10 * time.Second
	}
	base := strings.TrimRight(opts.BaseURL, "/")
	limiter := rate.NewLimiter(rate.Limit(opts.RPS), opts.RPS)

	client := &http.Client{Timeout: opts.Timeout, CheckRedirect: func(*http.Request, []*http.Request) error {
		return http.ErrUseLastResponse
	}}

	tasks := make(chan string, len(opts.Wordlist))
	var mu sync.Mutex
	results := make([]Result, 0, len(opts.Wordlist))
	var done atomic.Int32
	total := len(opts.Wordlist)
	note := func(r Result) {
		mu.Lock()
		results = append(results, r)
		mu.Unlock()
		n := int(done.Add(1))
		if opts.OnProgress != nil {
			opts.OnProgress(n, total, r)
		}
	}
	var wg sync.WaitGroup
	for i := 0; i < opts.Workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for w := range tasks {
				if err := limiter.Wait(ctx); err != nil {
					return
				}
				u := base + "/" + strings.TrimLeft(w, "/")
				parsed, err := url.Parse(u)
				if err != nil {
					note(Result{Path: w, Error: err.Error()})
					continue
				}
				req, err := http.NewRequestWithContext(ctx, "GET", parsed.String(), nil)
				if err != nil {
					note(Result{Path: w, Error: err.Error()})
					continue
				}
				req.Header.Set("User-Agent", "AURA-Discover/1.0")
				for _, h := range opts.Headers {
					req.Header.Set(h.Name, h.Value)
				}
				if opts.Cookies != "" {
					req.Header.Set("Cookie", opts.Cookies)
				}
				start := time.Now()
				resp, err := client.Do(req)
				dur := time.Since(start)
				if err != nil {
					note(Result{Path: w, Error: err.Error(), Duration: dur})
					if log != nil {
						log.Debug("discover error", zap.String("path", w), zap.Error(err))
					}
					continue
				}
				length := 0
				if cl := resp.Header.Get("Content-Length"); cl != "" {
					fmt.Sscanf(cl, "%d", &length)
				}
				if length == 0 {
					buf := make([]byte, 4096)
					for {
						n, _ := resp.Body.Read(buf)
						length += n
						if n < len(buf) {
							break
						}
					}
				}
				_ = resp.Body.Close()
				if opts.OnResult != nil {
					rawReq := req.Method + " " + req.URL.RequestURI() + " HTTP/1.1\r\nHost: " + req.URL.Host + "\r\n\r\n"
					opts.OnResult(rawReq, "")
				}
				note(Result{Path: w, StatusCode: resp.StatusCode, Length: length, Duration: dur})
			}
		}()
	}
enqueue:
	for _, w := range opts.Wordlist {
		select {
		case <-ctx.Done():
			break enqueue
		case tasks <- w:
		}
	}
	close(tasks)
	wg.Wait()
	return results, ctx.Err()
}
