// Package fuzz is a FUZZ-keyword HTTP fuzzer (ffuf-style) for authorized lab targets.
package fuzz

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/time/rate"
)

const (
	maxWorkers = 30
	maxRPS     = 50
	maxHits    = 8000
)

// Hit is one completed probe.
type Hit struct {
	Payload    string        `json:"payload"`
	URL        string        `json:"url"`
	StatusCode int           `json:"status_code"`
	Length     int           `json:"length"`
	Duration   time.Duration `json:"duration"`
	Error      string        `json:"error,omitempty"`
}

// Options configure a FUZZ run.
type Options struct {
	URL        string
	Method     string
	Headers    map[string]string
	Body       string
	Wordlist   []string
	Workers    int
	RPS        int
	Timeout    time.Duration
	Hide       []int
	OnHit      func(Hit)
	OnProgress func(done, total int, hit Hit)
}

// Run replaces FUZZ in URL, headers, and body with each wordlist entry.
func Run(ctx context.Context, opts Options) ([]Hit, error) {
	if !strings.Contains(opts.URL, "FUZZ") && !strings.Contains(opts.Body, "FUZZ") && !headersHaveFUZZ(opts.Headers) {
		return nil, fmt.Errorf("template must contain FUZZ")
	}
	if len(opts.Wordlist) == 0 {
		return nil, fmt.Errorf("wordlist is required")
	}
	if opts.Method == "" {
		opts.Method = http.MethodGet
	}
	if opts.Workers < 1 {
		opts.Workers = 8
	}
	if opts.Workers > maxWorkers {
		opts.Workers = maxWorkers
	}
	if opts.RPS < 1 {
		opts.RPS = 12
	}
	if opts.RPS > maxRPS {
		opts.RPS = maxRPS
	}
	if opts.Timeout <= 0 {
		opts.Timeout = 10 * time.Second
	}
	hide := map[int]struct{}{}
	for _, c := range opts.Hide {
		hide[c] = struct{}{}
	}

	client := &http.Client{
		Timeout: opts.Timeout,
		CheckRedirect: func(*http.Request, []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	limiter := rate.NewLimiter(rate.Limit(opts.RPS), opts.RPS)
	tasks := make(chan string)
	var mu sync.Mutex
	hits := make([]Hit, 0, 32)
	var tried atomic.Int32
	total := len(opts.Wordlist)
	var wg sync.WaitGroup
	for i := 0; i < opts.Workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for word := range tasks {
				if err := limiter.Wait(ctx); err != nil {
					return
				}
				hit := probe(ctx, client, opts, word)
				n := int(tried.Add(1))
				if opts.OnProgress != nil {
					opts.OnProgress(n, total, hit)
				}
				if _, skip := hide[hit.StatusCode]; skip && hit.Error == "" {
					continue
				}
				mu.Lock()
				if len(hits) < maxHits {
					hits = append(hits, hit)
				}
				mu.Unlock()
				if opts.OnHit != nil {
					opts.OnHit(hit)
				}
			}
		}()
	}
	for _, w := range opts.Wordlist {
		select {
		case <-ctx.Done():
			close(tasks)
			wg.Wait()
			return hits, ctx.Err()
		case tasks <- w:
		}
	}
	close(tasks)
	wg.Wait()
	return hits, nil
}

func headersHaveFUZZ(h map[string]string) bool {
	for k, v := range h {
		if strings.Contains(k, "FUZZ") || strings.Contains(v, "FUZZ") {
			return true
		}
	}
	return false
}

func probe(ctx context.Context, client *http.Client, opts Options, word string) Hit {
	start := time.Now()
	u := strings.ReplaceAll(opts.URL, "FUZZ", word)
	body := strings.ReplaceAll(opts.Body, "FUZZ", word)
	req, err := http.NewRequestWithContext(ctx, opts.Method, u, strings.NewReader(body))
	if err != nil {
		return Hit{Payload: word, URL: u, Error: err.Error(), Duration: time.Since(start)}
	}
	for k, v := range opts.Headers {
		req.Header.Set(strings.ReplaceAll(k, "FUZZ", word), strings.ReplaceAll(v, "FUZZ", word))
	}
	if req.Header.Get("User-Agent") == "" {
		req.Header.Set("User-Agent", "AURA-fuzz/1.0")
	}
	resp, err := client.Do(req)
	dur := time.Since(start)
	if err != nil {
		return Hit{Payload: word, URL: u, Error: err.Error(), Duration: dur}
	}
	n, _ := io.Copy(io.Discard, io.LimitReader(resp.Body, 1<<20))
	_ = resp.Body.Close()
	return Hit{Payload: word, URL: u, StatusCode: resp.StatusCode, Length: int(n), Duration: dur}
}
