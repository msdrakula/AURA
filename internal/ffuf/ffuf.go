// Package ffuf drives the ffuf binary (Fuzz Faster U Fool) and streams its
// JSONL results back to the caller. It powers the merged Fuzz tab.
package ffuf

import (
	"bufio"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

const (
	maxWords   = 50000
	maxResults = 20000
)

// LookPath resolves the ffuf binary.
func LookPath() string {
	bin, err := exec.LookPath("ffuf")
	if err != nil {
		return ""
	}
	return bin
}

// Result is one ffuf JSONL record (ffuf -json), normalized for the UI.
type Result struct {
	Payload     string `json:"payload"`
	Position    int    `json:"position"`
	Status      int    `json:"status"`
	Length      int    `json:"length"`
	Words       int    `json:"words"`
	Lines       int    `json:"lines"`
	URL         string `json:"url"`
	ContentType string `json:"content_type"`
	Redirect    string `json:"redirect"`
	DurationMs  int    `json:"duration_ms"`
}

// Stats summarizes a finished or stopped run.
type Stats struct {
	Tried   int `json:"tried"`
	Total   int `json:"total"`
	Errors  int `json:"errors"`
	Results int `json:"results"`
}

// Options configures one ffuf run.
type Options struct {
	URL        string   // must contain the Keyword
	Method     string   // GET when empty
	Headers    []string // raw "Name: Value" pairs
	Cookies    string   // "NAME1=VALUE1; NAME2=VALUE2"
	Body       string
	Wordlist   []string
	Workers    int    // -t threads
	RPS        int    // -rate, 0 leaves ffuf unthrottled
	Timeout    int    // per-request seconds, 10 when zero
	Keyword    string // FUZZ when empty
	OnResult   func(Result)
	OnProgress func(done, total int)
}

// Run executes ffuf once and returns the collected results.
// Progress lines on stderr drive OnProgress; ctx cancellation kills the process.
func Run(ctx context.Context, opts Options) ([]Result, Stats, error) {
	var st Stats
	if opts.Keyword == "" {
		opts.Keyword = "FUZZ"
	}
	if !strings.Contains(opts.URL, opts.Keyword) {
		return nil, st, fmt.Errorf("url must contain %s", opts.Keyword)
	}
	if len(opts.Wordlist) == 0 {
		return nil, st, fmt.Errorf("wordlist is required")
	}
	if len(opts.Wordlist) > maxWords {
		opts.Wordlist = opts.Wordlist[:maxWords]
	}
	if opts.Method == "" {
		opts.Method = "GET"
	}
	if opts.Workers < 1 {
		opts.Workers = 10
	}
	if opts.Timeout <= 0 {
		opts.Timeout = 10
	}
	st.Total = len(opts.Wordlist)

	tmp, err := os.CreateTemp("", "aura-ffuf-words-*.txt")
	if err != nil {
		return nil, st, err
	}
	defer os.Remove(tmp.Name())
	for _, w := range opts.Wordlist {
		if _, err := tmp.WriteString(w + "\n"); err != nil {
			_ = tmp.Close()
			return nil, st, err
		}
	}
	if err := tmp.Close(); err != nil {
		return nil, st, err
	}

	cmd := exec.CommandContext(ctx, LookPath(), buildArgs(opts, tmp.Name())...)

	// Progress and error counters go to stderr; capture into a temp file and
	// parse it after the run (a pipe risks SIGPIPE against ffuf's writer).
	seTmp, err := os.CreateTemp("", "aura-ffuf-stderr-*.log")
	if err != nil {
		return nil, st, err
	}
	defer os.Remove(seTmp.Name())
	cmd.Stderr = seTmp

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		seTmp.Close()
		return nil, st, err
	}
	if err := cmd.Start(); err != nil {
		seTmp.Close()
		return nil, st, fmt.Errorf("ffuf start: %w", err)
	}

	var out []Result
	sc := bufio.NewScanner(stdout)
	sc.Buffer(make([]byte, 64<<10), 4<<20)
	for sc.Scan() {
		if r, ok := parseResultLine(sc.Text(), opts.Keyword); ok && len(out) < maxResults {
			out = append(out, r)
			if opts.OnResult != nil {
				opts.OnResult(r)
			}
		}
	}
	waitErr := cmd.Wait()
	_ = seTmp.Close()

	// Parse progress lines once stderr is complete.
	if seData, rerr := os.ReadFile(seTmp.Name()); rerr == nil {
		for _, line := range strings.Split(stripANSI(string(seData)), "\n") {
			if d, tot, ok := parseProgressLine(line); ok {
				st.Tried = d
				if tot > 0 {
					st.Total = tot
				}
				if opts.OnProgress != nil {
					opts.OnProgress(d, st.Total)
				}
			}
		}
		st.Errors = lastErrorCount(string(seData))
	}
	// Progress ticks are time-based; fast runs may emit none at all. On a
	// clean exit ffuf consumed the whole wordlist, so report it fully.
	if waitErr == nil && ctx.Err() == nil && st.Tried < st.Total {
		st.Tried = st.Total
	}
	st.Results = len(out)
	if ctx.Err() != nil {
		return out, st, ctx.Err()
	}
	if waitErr != nil {
		tail := ""
		if seData, rerr := os.ReadFile(seTmp.Name()); rerr == nil {
			tail = string(seData)
		}
		return out, st, fmt.Errorf("ffuf exited: %v: %s", waitErr, tailBytes(tail, 500))
	}
	if err := sc.Err(); err != nil {
		return out, st, fmt.Errorf("ffuf output: %w", err)
	}
	return out, st, nil
}

// buildArgs assembles ffuf CLI flags. Exposed for tests.
func buildArgs(o Options, wordlistPath string) []string {
	kw := o.Keyword
	if kw == "" {
		kw = "FUZZ"
	}
	method := o.Method
	if method == "" {
		method = "GET"
	}
	workers := o.Workers
	if workers < 1 {
		workers = 10
	}
	timeout := o.Timeout
	if timeout <= 0 {
		timeout = 10
	}
	args := []string{
		"-json",
		"-noninteractive",
		"-u", o.URL,
		"-w", wordlistPath + ":" + kw,
		"-X", method,
		"-t", strconv.Itoa(workers),
		"-timeout", strconv.Itoa(timeout),
		"-mc", "all",
	}
	if o.RPS > 0 {
		args = append(args, "-rate", strconv.Itoa(o.RPS))
	}
	for _, h := range o.Headers {
		if strings.TrimSpace(h) != "" {
			args = append(args, "-H", h)
		}
	}
	if o.Cookies != "" {
		args = append(args, "-b", o.Cookies)
	}
	if o.Body != "" {
		args = append(args, "-d", o.Body)
	}
	return args
}

// parseResultLine decodes one ffuf JSONL result line. ffuf v2 base64-encodes
// keyword inputs and reports duration in nanoseconds.
func parseResultLine(line, keyword string) (Result, bool) {
	line = strings.TrimSpace(line)
	if !strings.HasPrefix(line, "{") {
		return Result{}, false
	}
	var m map[string]any
	if err := json.Unmarshal([]byte(line), &m); err != nil {
		return Result{}, false
	}
	num := func(k string) int {
		switch v := m[k].(type) {
		case float64:
			return int(v)
		case string:
			n, _ := strconv.Atoi(v)
			return n
		}
		return 0
	}
	str := func(k string) string {
		s, _ := m[k].(string)
		return s
	}
	r := Result{
		Position:    num("position"),
		Status:      num("status"),
		Length:      num("length"),
		Words:       num("words"),
		Lines:       num("lines"),
		URL:         str("url"),
		ContentType: str("content-type"),
		Redirect:    str("redirectlocation"),
	}
	if d := num("duration"); d > 0 {
		r.DurationMs = d / 1e6
	}
	if in, ok := m["input"].(map[string]any); ok {
		if v, ok := in[keyword].(string); ok {
			r.Payload = decodeInput(v)
		} else if len(in) == 1 {
			for _, v := range in {
				if s, ok := v.(string); ok {
					r.Payload = decodeInput(s)
				}
			}
		}
	}
	return r, true
}

// decodeInput reverses ffuf's base64 keyword encoding, falling back to raw.
func decodeInput(v string) string {
	if b, err := base64.StdEncoding.DecodeString(v); err == nil {
		return string(b)
	}
	return v
}

var ansiRe = regexp.MustCompile("\x1b\\[[0-9;]*[A-Za-z]")

func stripANSI(s string) string { return ansiRe.ReplaceAllString(s, "") }

var progressRe = regexp.MustCompile(`Progress:\s*\[(\d+)/(\d+)\]`)

// parseProgressLine extracts done/total from an ffuf progress line.
func parseProgressLine(line string) (done, total int, ok bool) {
	m := progressRe.FindStringSubmatch(line)
	if m == nil {
		return 0, 0, false
	}
	done, _ = strconv.Atoi(m[1])
	total, _ = strconv.Atoi(m[2])
	return done, total, true
}

var errorsRe = regexp.MustCompile(`Errors:\s*(\d+)`)

// lastErrorCount finds the last "Errors: N" figure ffuf printed.
func lastErrorCount(tail string) int {
	found := 0
	for _, m := range errorsRe.FindAllStringSubmatch(tail, -1) {
		if n, err := strconv.Atoi(m[1]); err == nil {
			found = n
		}
	}
	return found
}

// tailBytes returns the last n bytes of s.
func tailBytes(s string, n int) string {
	if len(s) > n {
		return s[len(s)-n:]
	}
	return s
}
