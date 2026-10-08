package intel

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

// nucleiLookPath resolves the nuclei binary.
func nucleiLookPath() string {
	bin, err := exec.LookPath("nuclei")
	if err != nil {
		return ""
	}
	return bin
}

// nucleiTemplatesDir returns the first existing nuclei templates directory.
func nucleiTemplatesDir() string {
	home, err := os.UserHomeDir()
	if err != nil || home == "" {
		return ""
	}
	for _, c := range []string{
		filepath.Join(home, ".local", "nuclei-templates"),
		filepath.Join(home, "nuclei-templates"),
	} {
		if st, err := os.Stat(c); err == nil && st.IsDir() {
			return c
		}
	}
	return ""
}

// buildNucleiTargets assembles deduplicated scan targets from the target and artifacts.
func buildNucleiTargets(t Target, arts []Artifact) []string {
	const maxTargets = 100
	seen := map[string]struct{}{}
	var out []string
	add := func(v string) {
		v = strings.TrimSpace(v)
		if v == "" || len(out) >= maxTargets {
			return
		}
		if _, ok := seen[v]; ok {
			return
		}
		seen[v] = struct{}{}
		out = append(out, v)
	}
	if t.BaseURL != "" {
		add(t.BaseURL)
	}
	for _, a := range arts {
		switch a.Kind {
		case KindHost:
			add(a.Value)
		case KindPort:
			// web_ports stores the probed URL in Extra["url"] and its scheme in
			// Extra["scheme"]; older artifacts may carry http=true / proto=http*.
			if raw := a.Extra["url"]; raw != "" {
				if u, err := url.Parse(raw); err == nil && u.Scheme != "" && u.Host != "" {
					add(u.Scheme + "://" + u.Host)
					continue
				}
			}
			if sch := a.Extra["scheme"]; strings.HasPrefix(sch, "http") {
				add(a.Value)
				continue
			}
			if a.Extra["http"] == "true" || strings.HasPrefix(a.Extra["proto"], "http") {
				add(a.Value)
			}
		case KindURL:
			u, err := url.Parse(a.Value)
			if err != nil || u.Scheme == "" || u.Host == "" {
				continue
			}
			add(u.Scheme + "://" + u.Host)
		}
	}
	return out
}

// parseNucleiLine parses one JSONL result line from nuclei -jsonl output.
func parseNucleiLine(line string) (Artifact, bool) {
	var m map[string]any
	if err := json.Unmarshal([]byte(line), &m); err != nil {
		return Artifact{}, false
	}
	tid, _ := m["template-id"].(string)
	if tid == "" {
		return Artifact{}, false
	}
	str := func(k string) string {
		s, _ := m[k].(string)
		return s
	}
	extra := map[string]string{
		"name":        "",
		"severity":    "",
		"host":        "",
		"matched":     "",
		"description": "",
		"tags":        "",
	}
	if info, ok := m["info"].(map[string]any); ok {
		for _, k := range []string{"name", "severity", "description", "tags"} {
			if v, ok := info[k].(string); ok {
				extra[k] = v
			}
		}
	}
	matched := str("matched-at")
	host := str("host")
	extra["matched"] = matched
	if matched != "" {
		extra["host"] = matched
	} else {
		extra["host"] = host
	}
	return Artifact{
		Kind:   KindVuln,
		Value:  tid,
		Source: "vulnscan",
		Extra:  extra,
	}, true
}

// cappedBuffer collects stderr up to a fixed limit.
type cappedBuffer struct {
	buf bytes.Buffer
	max int
}

func (c *cappedBuffer) Write(p []byte) (int, error) {
	n := len(p)
	if c.buf.Len() < c.max {
		space := c.max - c.buf.Len()
		if len(p) > space {
			p = p[:space]
		}
		c.buf.Write(p)
	}
	return n, nil
}

func (c *cappedBuffer) tail(n int) string {
	b := c.buf.Bytes()
	if len(b) > n {
		b = b[len(b)-n:]
	}
	return string(b)
}

// vulnscan runs nuclei against live targets collected by earlier stages.
func (e *Engine) vulnscan(ctx context.Context, t Target, arts []Artifact, opt StageOptions) ([]Artifact, error) {
	bin := nucleiLookPath()
	if bin == "" {
		return nil, fmt.Errorf("nuclei binary not found — install: sudo apt install nuclei")
	}
	targets := buildNucleiTargets(t, arts)
	if len(targets) == 0 {
		return nil, fmt.Errorf("no live targets — run live_hosts and web_ports first")
	}

	tmp, err := os.CreateTemp("", "aura-nuclei-targets-*.txt")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmp.Name())
	for _, tg := range targets {
		if _, err := tmp.WriteString(tg + "\n"); err != nil {
			_ = tmp.Close()
			return nil, err
		}
	}
	if err := tmp.Close(); err != nil {
		return nil, err
	}

	args := []string{
		"-l", tmp.Name(),
		"-jsonl",
		"-silent",
		"-nc",
		"-rl", strconv.Itoa(opt.rps(150)),
		"-c", strconv.Itoa(opt.workers(25)),
		"-timeout", strconv.Itoa(opt.timeout(600)),
	}
	if opt.Severity != "" {
		args = append(args, "-severity", opt.Severity)
	}
	if opt.Tags != "" {
		args = append(args, "-tags", opt.Tags)
	}
	if dir := nucleiTemplatesDir(); dir != "" {
		args = append(args, "-t", dir)
	}

	e.emit(ProgressEvent{Stage: "vulnscan", Action: "get", N: len(targets), Total: len(targets), Timeout: opt.timeout(600)})

	cmd := exec.CommandContext(ctx, bin, args...)
	stderr := &cappedBuffer{max: 64 << 10}
	cmd.Stderr = stderr
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("nuclei start: %w", err)
	}

	var out []Artifact
	sc := bufio.NewScanner(stdout)
	sc.Buffer(make([]byte, 64<<10), 4<<20)
	for sc.Scan() {
		a, ok := parseNucleiLine(sc.Text())
		if !ok {
			continue
		}
		out = append(out, a)
		e.emit(ProgressEvent{Stage: "vulnscan", Action: "hit", Value: a.Value, URL: a.Extra["matched"], Host: a.Extra["host"]})
	}
	waitErr := cmd.Wait()
	if ctx.Err() != nil {
		return out, ctx.Err()
	}
	if waitErr != nil {
		return out, fmt.Errorf("nuclei exited: %v: %s", waitErr, stderr.tail(500))
	}
	if err := sc.Err(); err != nil {
		return out, fmt.Errorf("nuclei output: %w", err)
	}

	max := opt.limit(500, 5000)
	if len(out) > max {
		out = out[:max]
	}
	return out, nil
}
