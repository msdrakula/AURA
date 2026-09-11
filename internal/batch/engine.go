// Package batch provides a reusable worker pool and batch HTTP execution engine.
package batch

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"golang.org/x/time/rate"

	"meb/internal/httpio"
	"meb/internal/models"
	"meb/internal/repeater"
)

// Engine runs batched HTTP requests with payload variations over a worker pool.
type Engine struct {
	workers     int
	rateLimiter *rate.Limiter
	timeout     time.Duration
	verify      bool
	onResult    func(rawReq, respRaw string)
}

// NewEngine creates a batch engine with the given worker count and rate limit.
// rps is the maximum requests per second; burst is set to rps.
func NewEngine(workers, rps int) *Engine {
	if workers < 1 {
		workers = 1
	}
	if rps < 1 {
		rps = 1
	}
	return &Engine{
		workers:     workers,
		rateLimiter: rate.NewLimiter(rate.Limit(rps), rps),
		timeout:     30 * time.Second,
	}
}

// WithOnResult registers a callback invoked for every sent request with the
// raw request and raw response. Used for logging tool-sent flows.
func (e *Engine) WithOnResult(cb func(rawReq, respRaw string)) *Engine {
	e.onResult = cb
	return e
}

// NewWithLogger builds a fresh engine with the given workers/rps and a per-request
// result logger. Satisfies api.BatchFactory.
func NewWithLogger(workers, rps int, onResult func(rawReq, respRaw string)) *Engine {
	return NewEngine(workers, rps).WithOnResult(onResult)
}

// WithTimeout overrides the per-request timeout. Defaults to 30s.
func (e *Engine) WithTimeout(d time.Duration) *Engine {
	if d > 0 {
		e.timeout = d
	}
	return e
}

// WithVerify enables upstream TLS certificate verification.
func (e *Engine) WithVerify(v bool) *Engine {
	e.verify = v
	return e
}

// Execute generates payload combinations from payloadSets, substitutes each
// into the raw template (replacing § markers in order), and sends every
// variant through the repeater using a rate-limited worker pool.
// template is a raw HTTP/1.1 request with § placeholders.
// scheme is "http" or "https"; target overrides the Host header if non-empty.
// This is the legacy Cluster bomb-style entry point.
func (e *Engine) Execute(
	ctx context.Context,
	template string,
	scheme, target string,
	payloadSets [][]string,
) ([]models.BatchResult, error) {
	return e.ExecuteAttack(ctx, template, scheme, target, "combo", payloadSets)
}

// ExecuteAttack runs a payload attack. attackType is one of:
// "one", "same", "zip", "combo" (legacy aliases still accepted).
// The template contains § markers in pairs; each pair is one payload position.
func (e *Engine) ExecuteAttack(
	ctx context.Context,
	template string,
	scheme, target, attackType string,
	payloadSets [][]string,
) ([]models.BatchResult, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	segments, originals := parsePositions(template)
	positions := len(originals)
	if positions == 0 {
		return []models.BatchResult{}, nil
	}

	combinations, err := attackCombinations(attackType, positions, payloadSets)
	if err != nil {
		return nil, err
	}
	if len(combinations) == 0 {
		return []models.BatchResult{}, nil
	}

	tasks := make(chan Task[[]string], len(combinations))
	results := make(chan Result[[]string, models.BatchResult], len(combinations))

	go func() {
		defer close(tasks)
		for i, combo := range combinations {
			select {
			case <-ctx.Done():
				return
			case tasks <- Task[[]string]{ID: taskID(i), Payload: combo}:
			}
		}
	}()

	workerFn := func(ctx context.Context, payloads []string) (models.BatchResult, error) {
		if err := e.rateLimiter.Wait(ctx); err != nil {
			return models.BatchResult{Payloads: payloads, Error: err.Error()}, nil
		}
		raw := buildRequest(segments, originals, payloads)
		res := repeater.Send(ctx, repeater.Options{
			Raw:     raw,
			Scheme:  scheme,
			Target:  target,
			Verify:  e.verify,
			Timeout: e.timeout,
		})
		if e.onResult != nil {
			e.onResult(raw, res.ResponseRaw)
		}
		out := models.BatchResult{
			Payloads:   payloads,
			StatusCode: res.Status,
			Duration:   time.Duration(res.DurationMS) * time.Millisecond,
		}
		if res.ResponseRaw != "" {
			_, body := httpio.SplitHeadBody([]byte(res.ResponseRaw))
			out.BodyLength = len(body)
		}
		if !res.OK && res.Error != "" {
			out.Error = res.Error
		}
		return out, nil
	}

	go func() {
		_ = RunWorkerPool(ctx, tasks, results, e.workers, workerFn)
	}()

	out := make([]models.BatchResult, 0, len(combinations))
	for i := 0; i < len(combinations); i++ {
		select {
		case res := <-results:
			if res.Err != nil {
				out = append(out, models.BatchResult{Error: res.Err.Error()})
			} else {
				out = append(out, res.Data)
			}
		case <-ctx.Done():
			return out, ctx.Err()
		}
	}
	return out, nil
}

// parsePositions splits a template on § markers into alternating text
// segments and the original text enclosed by each marker pair.
// For "GET /?a=§x§&b=§y§" it returns segments=["GET /?a=", "&b=", ""] and
// originals=["x", "y"]. Unpaired § markers are treated as literal text.
func parsePositions(template string) (segments, originals []string) {
	parts := strings.Split(template, "§")
	// A well-formed template has an odd number of parts (pairs of markers).
	// positions = number of complete pairs.
	positions := (len(parts) - 1) / 2
	segments = make([]string, positions+1)
	originals = make([]string, positions)
	for i := 0; i < positions; i++ {
		segments[i] = parts[2*i]
		originals[i] = parts[2*i+1]
	}
	// Last segment: parts[2*positions]. If there's a trailing unpaired §,
	// append the remainder as literal text.
	last := parts[2*positions]
	if len(parts)%2 == 0 {
		// Odd number of § markers: a trailing unpaired marker. Append "§" + remainder.
		last = last + "§" + parts[len(parts)-1]
	}
	segments[positions] = last
	return segments, originals
}

// buildRequest reassembles the template substituting each position with its
// payload. If a payload is the empty string and the position had original
// text, the original is kept (used when a position is not filled).
func buildRequest(segments, originals, payloads []string) string {
	var b strings.Builder
	for i, seg := range segments {
		b.WriteString(seg)
		if i < len(payloads) {
			if payloads[i] != "" {
				b.WriteString(payloads[i])
			} else if i < len(originals) {
				b.WriteString(originals[i])
			}
		}
	}
	return b.String()
}

func normalizeAttackType(attackType string) string {
	switch strings.ToLower(strings.TrimSpace(attackType)) {
	case "one", "one_position", "sniper":
		return "one"
	case "same", "same_value", "battering_ram":
		return "same"
	case "zip", "paired", "pitchfork":
		return "zip"
	case "combo", "combinations", "cluster_bomb", "":
		return "combo"
	default:
		return strings.ToLower(strings.TrimSpace(attackType))
	}
}

// attackCombinations returns the list of payload arrays per request for the
// given attack type and number of positions.
func attackCombinations(attackType string, positions int, sets [][]string) ([][]string, error) {
	switch normalizeAttackType(attackType) {
	case "one":
		if len(sets) != 1 {
			return nil, fmt.Errorf("one-position mode requires exactly 1 payload set, got %d", len(sets))
		}
		set := sets[0]
		n := positions * len(set)
		if positions > 0 && n/positions != len(set) {
			return nil, fmt.Errorf("too many combinations, max %d", MaxCombinations)
		}
		if n > MaxCombinations {
			return nil, fmt.Errorf("too many combinations (%d), max %d", n, MaxCombinations)
		}
		out := make([][]string, 0, n)
		for pos := 0; pos < positions; pos++ {
			for _, p := range set {
				combo := make([]string, positions)
				combo[pos] = p
				out = append(out, combo)
			}
		}
		return out, nil
	case "same":
		if len(sets) != 1 {
			return nil, fmt.Errorf("same-value mode requires exactly 1 payload set, got %d", len(sets))
		}
		set := sets[0]
		if len(set) > MaxCombinations {
			return nil, fmt.Errorf("too many combinations (%d), max %d", len(set), MaxCombinations)
		}
		out := make([][]string, 0, len(set))
		for _, p := range set {
			combo := make([]string, positions)
			for i := range combo {
				combo[i] = p
			}
			out = append(out, combo)
		}
		return out, nil
	case "zip":
		if len(sets) != positions {
			return nil, fmt.Errorf("paired mode requires %d payload sets (one per position), got %d", positions, len(sets))
		}
		minLen := len(sets[0])
		for _, s := range sets[1:] {
			if len(s) < minLen {
				minLen = len(s)
			}
		}
		if minLen > MaxCombinations {
			return nil, fmt.Errorf("too many combinations (%d), max %d", minLen, MaxCombinations)
		}
		out := make([][]string, 0, minLen)
		for i := 0; i < minLen; i++ {
			combo := make([]string, positions)
			for j := 0; j < positions; j++ {
				combo[j] = sets[j][i]
			}
			out = append(out, combo)
		}
		return out, nil
	case "combo":
		if len(sets) != positions {
			return nil, fmt.Errorf("combinations mode requires %d payload sets (one per position), got %d", positions, len(sets))
		}
		n, overflow := combinationCount(sets)
		if overflow || n > MaxCombinations {
			shown := n
			if overflow {
				shown = -1
			}
			return nil, fmt.Errorf("too many combinations (%d), max %d — use paired mode or smaller sets", shown, MaxCombinations)
		}
		return GenerateCombinations(sets), nil
	default:
		return nil, fmt.Errorf("unknown attack type: %q", attackType)
	}
}

func taskID(i int) string {
	var b [4]byte
	_, _ = rand.Read(b[:])
	return fmt.Sprintf("batch-%d-%s", i, hex.EncodeToString(b[:]))
}
