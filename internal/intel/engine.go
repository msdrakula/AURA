package intel

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"meb/internal/debuglog"
	"meb/internal/discover"
	"meb/internal/extractor"
	"meb/internal/fuzz"
	"meb/internal/wordlist"
)

// Engine runs mindmap stages against a Store.
type Engine struct {
	Store    Store
	HTTP     *http.Client
	Log      *zap.Logger
	Lists    *wordlist.Library
	Lookup   func(ctx context.Context, host string) ([]string, error)
	Dial     func(ctx context.Context, network, addr string) (net.Conn, error)
	Progress func(ProgressEvent)
}

func (e *Engine) http() *http.Client {
	if e.HTTP != nil {
		return e.HTTP
	}
	return &http.Client{
		Timeout: 12 * time.Second,
		CheckRedirect: func(*http.Request, []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
}

func (e *Engine) httpWithTimeout(d time.Duration) *http.Client {
	base := e.http()
	c := *base
	c.Timeout = d
	return &c
}

func (e *Engine) log() *zap.Logger {
	if e.Log != nil {
		return e.Log
	}
	return zap.NewNop()
}

// EnsureTarget creates or returns a workspace for a domain/URL.
func (e *Engine) EnsureTarget(raw string, authorized bool) (Target, error) {
	domain, base, err := NormalizeInput(raw)
	if err != nil {
		return Target{}, err
	}
	existing, err := e.Store.ListTargets()
	if err != nil {
		return Target{}, err
	}
	for _, t := range existing {
		if strings.EqualFold(t.Domain, domain) {
			changed := false
			if authorized && !t.Authorized {
				t.Authorized = true
				changed = true
			}
			if base != "" && t.BaseURL != base {
				t.BaseURL = base
				changed = true
			}
			if changed {
				_ = e.Store.UpdateTarget(t)
			}
			return t, nil
		}
	}
	t := Target{
		ID:         uuid.NewString(),
		Name:       domain,
		Domain:     domain,
		BaseURL:    base,
		Authorized: authorized,
		Created:    time.Now().UTC(),
	}
	if err := e.Store.CreateTarget(t); err != nil {
		return Target{}, err
	}
	_ = e.Store.AddArtifact(Artifact{
		TargetID: t.ID,
		Kind:     KindSubdomain,
		Value:    domain,
		Source:   "seed",
		Extra:    map[string]string{"apex": "1"},
	})
	return t, nil
}

// Snapshot returns target, stages (merged with catalog), artifacts, and map.
func (e *Engine) Snapshot(id string) (Target, []StageState, []Artifact, AppMap, error) {
	t, err := e.Store.GetTarget(id)
	if err != nil {
		return Target{}, nil, nil, AppMap{}, err
	}
	arts, err := e.Store.ListArtifacts(id)
	if err != nil {
		return Target{}, nil, nil, AppMap{}, err
	}
	saved, err := e.Store.ListStages(id)
	if err != nil {
		return Target{}, nil, nil, AppMap{}, err
	}
	byID := map[string]StageState{}
	for _, s := range saved {
		byID[s.ID] = s
	}
	var stages []StageState
	for _, c := range Catalog() {
		if st, ok := byID[c.ID]; ok {
			stages = append(stages, st)
			continue
		}
		stages = append(stages, StageState{ID: c.ID, Status: StatusIdle})
	}
	return t, stages, arts, BuildMap(t, arts), nil
}

// RunStage executes one mindmap step. Optional StageOptions override defaults for that module.
func (e *Engine) RunStage(ctx context.Context, targetID, stageID string, opts ...StageOptions) (RunResult, error) {
	cat, ok := stageByID(stageID)
	if !ok {
		return RunResult{}, fmt.Errorf("unknown stage %q", stageID)
	}
	var opt StageOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	t, err := e.Store.GetTarget(targetID)
	if err != nil {
		return RunResult{}, err
	}
	if cat.Mode == ModeActive && !t.Authorized {
		return RunResult{}, fmt.Errorf("active stage requires confirmation that you are testing your own target")
	}
	st := StageState{ID: stageID, Status: StatusRunning, Updated: time.Now().UTC()}
	_ = e.Store.UpsertStage(targetID, st)
	debuglog.Info("intel", "stage_start", stageID, map[string]any{"target_id": targetID, "domain": t.Domain})

	arts, err := e.Store.ListArtifacts(targetID)
	if err != nil {
		return RunResult{}, err
	}

	e.emit(ProgressEvent{Stage: stageID, Action: "start", Host: t.Domain, URL: t.BaseURL})

	var added []Artifact
	var runErr error
	switch stageID {
	case "subdomains_passive":
		added, runErr = e.crtsh(ctx, t, opt)
	case "dns_brute":
		added, runErr = e.dnsBrute(ctx, t, opt)
	case "live_hosts":
		added, runErr = e.liveHosts(ctx, t, arts, opt)
	case "web_ports":
		added, runErr = e.webPorts(ctx, t, arts, opt)
	case "tech":
		added, runErr = e.tech(ctx, t, arts, opt)
	case "urls_passive":
		added, runErr = e.wayback(ctx, t, opt)
	case "scrape":
		added, runErr = e.scrape(ctx, t, opt)
	case "dirs":
		added, runErr = e.dirs(ctx, t, opt)
	case "params":
		added, runErr = e.params(ctx, t, opt)
	default:
		runErr = fmt.Errorf("stage not implemented")
	}

	n := 0
	for _, a := range added {
		a.TargetID = t.ID
		if a.ID == "" {
			a.ID = uuid.NewString()
		}
		if a.Created.IsZero() {
			a.Created = time.Now().UTC()
		}
		if err := e.Store.AddArtifact(a); err == nil {
			n++
		}
	}
	st.Updated = time.Now().UTC()
	if ctx.Err() != nil {
		st.Status = StatusIdle
		st.Summary = "stopped"
		_ = e.Store.UpsertStage(targetID, st)
		e.emit(ProgressEvent{Stage: stageID, Action: "stop", N: n, Host: t.Domain})
		debuglog.Info("intel", "stage_stop", stageID, map[string]any{"target_id": targetID, "added": n})
		return RunResult{Stage: st, Added: n, Artifacts: added}, ctx.Err()
	}
	if runErr != nil {
		st.Status = StatusError
		st.Summary = runErr.Error()
		_ = e.Store.UpsertStage(targetID, st)
		e.emit(ProgressEvent{Stage: stageID, Action: "err", Err: runErr.Error(), N: n, Host: t.Domain})
		debuglog.Error("intel", "stage_error", stageID, runErr, map[string]any{"target_id": targetID, "added": n})
		return RunResult{Stage: st, Added: n, Artifacts: added}, runErr
	}
	st.Status = StatusDone
	st.Summary = fmt.Sprintf("добавлено фактов: %d", n)
	_ = e.Store.UpsertStage(targetID, st)
	e.emit(ProgressEvent{Stage: stageID, Action: "done", N: n, Host: t.Domain})
	debuglog.Info("intel", "stage_done", stageID, map[string]any{"target_id": targetID, "added": n})
	return RunResult{Stage: st, Added: n, Artifacts: added}, nil
}

// IngestHistory folds proxy history URLs into the map.
func (e *Engine) IngestHistory(targetID string, urls []string) (int, error) {
	t, err := e.Store.GetTarget(targetID)
	if err != nil {
		return 0, err
	}
	n := 0
	for _, raw := range urls {
		u, err := url.Parse(raw)
		if err != nil || u.Host == "" {
			continue
		}
		host := strings.ToLower(u.Hostname())
		if host != t.Domain && !strings.HasSuffix(host, "."+t.Domain) {
			continue
		}
		items := []Artifact{
			{TargetID: t.ID, Kind: KindHost, Value: host, Source: "proxy"},
			{TargetID: t.ID, Kind: KindURL, Value: strings.TrimRight(raw, "#"), Source: "proxy"},
		}
		if u.Path != "" && IsSanePath(u.Path) {
			items = append(items, Artifact{TargetID: t.ID, Kind: KindPath, Value: u.Path, Source: "proxy", Extra: map[string]string{"host": host}})
		}
		for key := range u.Query() {
			if IsTrackingParam(key) {
				continue
			}
			items = append(items, Artifact{TargetID: t.ID, Kind: KindParam, Value: key, Source: "proxy", Extra: map[string]string{"host": host, "path": u.Path}})
		}
		for _, a := range items {
			if err := e.Store.AddArtifact(a); err == nil {
				n++
			}
		}
	}
	return n, nil
}

func (e *Engine) crtsh(ctx context.Context, t Target, opt StageOptions) ([]Artifact, error) {
	q := url.QueryEscape("%." + t.Domain)
	raw := "https://crt.sh/?q=" + q + "&output=json"
	sec := opt.timeout(45)
	e.emit(ProgressEvent{Stage: "subdomains_passive", Action: "get", URL: raw, Host: "crt.sh", Timeout: sec})
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, raw, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "AURA-intel/1.0")
	resp, err := e.httpWithTimeout(time.Duration(sec) * time.Second).Do(req)
	if err != nil {
		e.emit(ProgressEvent{Stage: "subdomains_passive", Action: "err", Host: "crt.sh", URL: raw, Err: err.Error()})
		return nil, fmt.Errorf("crt.sh: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(io.LimitReader(resp.Body, 8<<20))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		e.emit(ProgressEvent{Stage: "subdomains_passive", Action: "err", Host: "crt.sh", Status: resp.StatusCode})
		return nil, fmt.Errorf("crt.sh: HTTP %d", resp.StatusCode)
	}
	names := parseCRTNames(body, t.Domain)
	max := opt.limit(400, 2000)
	if len(names) > max {
		names = names[:max]
	}
	e.emit(ProgressEvent{Stage: "subdomains_passive", Action: "ok", Host: "crt.sh", Status: resp.StatusCode, N: len(names)})
	var out []Artifact
	for _, name := range names {
		out = append(out, Artifact{Kind: KindSubdomain, Value: name, Source: "crtsh"})
	}
	return out, nil
}

func (e *Engine) dnsBrute(ctx context.Context, t Target, opt StageOptions) ([]Artifact, error) {
	lookup := e.Lookup
	if lookup == nil {
		lookup = func(ctx context.Context, host string) ([]string, error) {
			var d net.Resolver
			return d.LookupHost(ctx, host)
		}
	}
	words := e.stageWords(opt, wordlist.HostPrefixes(), wordlist.DefaultDNS...)
	max := opt.limit(len(words), 20000)
	if max < len(words) {
		words = words[:max]
	}
	workers := opt.workers(32)
	if workers > 128 {
		workers = 128
	}
	if workers > len(words) {
		workers = len(words)
	}
	if workers < 1 {
		workers = 1
	}
	sec := opt.timeout(5)
	type hit struct {
		host string
		ip   string
		i    int
	}
	jobs := make(chan int)
	hits := make(chan hit, workers)
	var wg sync.WaitGroup
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range jobs {
				if err := ctx.Err(); err != nil {
					return
				}
				host := words[i] + "." + t.Domain
				lctx, cancel := context.WithTimeout(ctx, time.Duration(sec)*time.Second)
				addrs, err := lookup(lctx, host)
				cancel()
				if err != nil || len(addrs) == 0 {
					e.emitTick("dns_brute", i+1, len(words), 25, ProgressEvent{Host: host})
					continue
				}
				hits <- hit{host: host, ip: addrs[0], i: i}
			}
		}()
	}
	go func() {
		defer close(jobs)
		for i := range words {
			select {
			case <-ctx.Done():
				return
			case jobs <- i:
			}
		}
	}()
	go func() {
		wg.Wait()
		close(hits)
	}()
	var out []Artifact
	for h := range hits {
		out = append(out, Artifact{
			Kind:   KindSubdomain,
			Value:  h.host,
			Source: "dns",
			Extra:  map[string]string{"ip": h.ip},
		})
		e.emit(ProgressEvent{Stage: "dns_brute", Action: "hit", Host: h.host, Value: h.ip, N: h.i + 1, Total: len(words)})
	}
	return out, ctx.Err()
}

func (e *Engine) liveHosts(ctx context.Context, t Target, arts []Artifact, opt StageOptions) ([]Artifact, error) {
	seen := map[string]struct{}{}
	var out []Artifact
	add := func(host string, extra map[string]string) {
		host = strings.ToLower(strings.TrimSpace(host))
		if host == "" || !SameScope(host, t.Domain) {
			return
		}
		if _, ok := seen[host]; ok {
			return
		}
		seen[host] = struct{}{}
		out = append(out, Artifact{Kind: KindHost, Value: host, Source: "live", Extra: extra})
	}
	sec := opt.timeout(8)
	schemes := opt.Schemes
	if len(schemes) == 0 {
		schemes = []string{"https", "http"}
	}
	if t.BaseURL != "" {
		e.emit(ProgressEvent{Stage: "live_hosts", Action: "get", URL: t.BaseURL, Timeout: sec})
		if live, extra := e.probeURLFast(ctx, t.BaseURL, time.Duration(sec)*time.Second); live {
			if u, err := url.Parse(t.BaseURL); err == nil {
				add(u.Hostname(), extra)
				e.emit(ProgressEvent{Stage: "live_hosts", Action: "hit", URL: t.BaseURL, Host: u.Hostname(), Status: atoi(extra["status"])})
			}
		}
	}
	hosts := []string{t.Domain}
	for _, a := range arts {
		if a.Kind == KindSubdomain || a.Kind == KindHost {
			if SameScope(a.Value, t.Domain) {
				hosts = append(hosts, a.Value)
			}
		}
	}
	for _, h := range hosts {
		if err := ctx.Err(); err != nil {
			return out, err
		}
		if _, ok := seen[strings.ToLower(h)]; ok {
			continue
		}
		e.emit(ProgressEvent{Stage: "live_hosts", Action: "get", Host: h, Timeout: sec})
		live, extra := e.probeHostSchemes(ctx, h, t.BaseURL, schemes, time.Duration(sec)*time.Second)
		if live {
			add(h, extra)
			e.emit(ProgressEvent{Stage: "live_hosts", Action: "hit", Host: h, Status: atoi(extra["status"])})
		}
	}
	return out, nil
}

func (e *Engine) probeURL(ctx context.Context, raw string) (bool, map[string]string) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, raw, nil)
	if err != nil {
		return false, nil
	}
	req.Header.Set("User-Agent", "AURA-intel/1.0")
	resp, err := e.http().Do(req)
	if err != nil {
		return false, nil
	}
	_, _ = io.Copy(io.Discard, io.LimitReader(resp.Body, 64*1024))
	_ = resp.Body.Close()
	if resp.StatusCode == 0 {
		return false, nil
	}
	u, _ := url.Parse(raw)
	scheme := "https"
	if u != nil {
		scheme = u.Scheme
	}
	return true, map[string]string{
		"scheme": scheme,
		"status": fmt.Sprintf("%d", resp.StatusCode),
		"server": resp.Header.Get("Server"),
	}
}

func (e *Engine) probeHost(ctx context.Context, host string, bases ...string) (bool, map[string]string) {
	return e.probeHostSchemes(ctx, host, firstString(bases), []string{"https", "http"}, 8*time.Second)
}

func firstString(ss []string) string {
	if len(ss) == 0 {
		return ""
	}
	return ss[0]
}

func (e *Engine) probeHostSchemes(ctx context.Context, host, base string, schemes []string, d time.Duration) (bool, map[string]string) {
	host = strings.ToLower(strings.TrimSpace(host))
	var candidates []string
	if origin := OriginURL(base); origin != "" {
		u, err := url.Parse(origin)
		if err == nil && strings.EqualFold(u.Hostname(), host) {
			candidates = append(candidates, strings.TrimRight(origin, "/")+"/")
		}
	}
	want := map[string]bool{}
	for _, s := range schemes {
		want[strings.ToLower(strings.TrimSpace(s))] = true
	}
	if want["https"] {
		candidates = append(candidates, "https://"+host+"/")
	}
	if want["http"] {
		candidates = append(candidates, "http://"+host+"/")
	}
	if len(candidates) == 0 {
		candidates = []string{"https://" + host + "/", "http://" + host + "/"}
	}
	seen := map[string]struct{}{}
	for _, raw := range candidates {
		if _, ok := seen[raw]; ok {
			continue
		}
		seen[raw] = struct{}{}
		if ok, extra := e.probeURLFast(ctx, raw, d); ok {
			return true, extra
		}
	}
	return false, nil
}

func (e *Engine) webPorts(ctx context.Context, t Target, arts []Artifact, opt StageOptions) ([]Artifact, error) {
	dial := e.Dial
	if dial == nil {
		d := net.Dialer{Timeout: 2 * time.Second}
		dial = d.DialContext
	}
	hosts := liveOrSeed(t, arts)
	ports := webPortList(t.BaseURL, opt.Ports)
	total := len(hosts) * len(ports)
	sec := opt.timeout(3)
	var out []Artifact
	n := 0
	for _, h := range hosts {
		if !SameScope(h, t.Domain) {
			continue
		}
		for _, p := range ports {
			if err := ctx.Err(); err != nil {
				return out, err
			}
			n++
			addr := net.JoinHostPort(h, p)
			c, err := dial(ctx, "tcp", addr)
			if err != nil {
				e.emitTick("web_ports", n, total, 16, ProgressEvent{Host: addr})
				continue
			}
			_ = c.Close()
			live, extra := e.probeWebPort(ctx, h, p, t.BaseURL, time.Duration(sec)*time.Second)
			if !live {
				e.emitTick("web_ports", n, total, 16, ProgressEvent{Host: addr, Msg: "tcp"})
				continue
			}
			if extra == nil {
				extra = map[string]string{}
			}
			extra["host"] = h
			extra["port"] = p
			out = append(out, Artifact{
				Kind:   KindPort,
				Value:  h + ":" + p,
				Source: "ports",
				Extra:  extra,
			})
			e.emit(ProgressEvent{Stage: "web_ports", Action: "hit", Host: addr, Value: addr, URL: extra["url"], Status: atoi(extra["status"]), N: n, Total: total})
		}
	}
	return out, nil
}

func (e *Engine) probeWebPort(ctx context.Context, host, port, base string, d ...time.Duration) (bool, map[string]string) {
	wait := 3 * time.Second
	if len(d) > 0 && d[0] > 0 {
		wait = d[0]
	}
	for _, raw := range webProbeURLs(host, port, base) {
		ok, extra := e.probeURLFast(ctx, raw, wait)
		if ok {
			if extra == nil {
				extra = map[string]string{}
			}
			extra["url"] = raw
			return true, extra
		}
	}
	return false, nil
}

func (e *Engine) probeURLFast(ctx context.Context, raw string, d time.Duration) (bool, map[string]string) {
	ctx, cancel := context.WithTimeout(ctx, d)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, raw, nil)
	if err != nil {
		return false, nil
	}
	req.Header.Set("User-Agent", "AURA-intel/1.0")
	resp, err := e.httpWithTimeout(d).Do(req)
	if err != nil {
		return false, nil
	}
	_, _ = io.Copy(io.Discard, io.LimitReader(resp.Body, 64*1024))
	_ = resp.Body.Close()
	if resp.StatusCode == 0 {
		return false, nil
	}
	u, _ := url.Parse(raw)
	scheme := "https"
	if u != nil && u.Scheme != "" {
		scheme = u.Scheme
	}
	return true, map[string]string{
		"scheme": scheme,
		"status": fmt.Sprintf("%d", resp.StatusCode),
		"server": resp.Header.Get("Server"),
	}
}

func (e *Engine) tech(ctx context.Context, t Target, arts []Artifact, opt StageOptions) ([]Artifact, error) {
	var out []Artifact
	seen := map[string]struct{}{}
	sec := opt.timeout(10)
	fetch := func(raw, host string) {
		if !SameScope(host, t.Domain) {
			return
		}
		e.emit(ProgressEvent{Stage: "tech", Action: "get", URL: raw, Host: host, Timeout: sec})
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, raw, nil)
		if err != nil {
			return
		}
		req.Header.Set("User-Agent", "AURA-intel/1.0")
		resp, err := e.httpWithTimeout(time.Duration(sec) * time.Second).Do(req)
		if err != nil {
			e.emit(ProgressEvent{Stage: "tech", Action: "err", URL: raw, Host: host, Err: err.Error()})
			return
		}
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 256*1024))
		_ = resp.Body.Close()
		names := DetectTech(resp.Header, body)
		e.emit(ProgressEvent{Stage: "tech", Action: "ok", URL: raw, Host: host, Status: resp.StatusCode, N: len(names)})
		for _, name := range names {
			key := host + "|" + name
			if _, ok := seen[key]; ok {
				continue
			}
			seen[key] = struct{}{}
			out = append(out, Artifact{
				Kind:   KindTech,
				Value:  name,
				Source: "tech",
				Extra:  map[string]string{"host": host},
			})
			e.emit(ProgressEvent{Stage: "tech", Action: "hit", Host: host, Value: name, Status: resp.StatusCode})
		}
	}
	if t.BaseURL != "" {
		host := t.Domain
		if u, err := url.Parse(t.BaseURL); err == nil && u.Hostname() != "" {
			host = u.Hostname()
		}
		fetch(t.BaseURL, host)
	}
	for _, h := range liveOrSeed(t, arts) {
		if !SameScope(h, t.Domain) {
			continue
		}
		scheme := "https"
		for _, a := range arts {
			if a.Kind == KindHost && a.Value == h && a.Extra["scheme"] == "http" {
				scheme = "http"
			}
		}
		if t.BaseURL != "" {
			if u, err := url.Parse(t.BaseURL); err == nil && strings.EqualFold(u.Hostname(), h) {
				continue
			}
		}
		fetch(scheme+"://"+h+"/", h)
	}
	return out, nil
}

func (e *Engine) wayback(ctx context.Context, t Target, opt StageOptions) ([]Artifact, error) {
	limit := opt.limit(400, 2000)
	match := "*." + t.Domain + "/*"
	u := "https://web.archive.org/cdx/search/cdx?url=" + url.QueryEscape(match) + "&output=json&fl=original&collapse=urlkey&limit=" + strconv.Itoa(limit)
	sec := opt.timeout(45)
	e.emit(ProgressEvent{Stage: "urls_passive", Action: "get", URL: u, Host: "web.archive.org", Timeout: sec})
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "AURA-intel/1.0")
	resp, err := e.httpWithTimeout(time.Duration(sec) * time.Second).Do(req)
	if err != nil {
		e.emit(ProgressEvent{Stage: "urls_passive", Action: "err", Host: "web.archive.org", URL: u, Err: err.Error()})
		return nil, fmt.Errorf("wayback: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(io.LimitReader(resp.Body, 4<<20))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		e.emit(ProgressEvent{Stage: "urls_passive", Action: "err", Host: "web.archive.org", Status: resp.StatusCode})
		return nil, fmt.Errorf("wayback: HTTP %d", resp.StatusCode)
	}
	urls := parseWayback(body)
	e.emit(ProgressEvent{Stage: "urls_passive", Action: "ok", Host: "web.archive.org", Status: resp.StatusCode, N: len(urls)})
	var out []Artifact
	for _, raw := range urls {
		pu, err := url.Parse(raw)
		if err != nil || pu.Hostname() == "" {
			continue
		}
		if opt.sameHost() && !SameScope(pu.Hostname(), t.Domain) {
			continue
		}
		path := pu.Path
		if path == "" {
			path = "/"
		}
		if !IsSanePath(path) {
			continue
		}
		out = append(out, Artifact{Kind: KindURL, Value: raw, Source: "wayback"})
		if path != "/" {
			out = append(out, Artifact{Kind: KindPath, Value: path, Source: "wayback", Extra: map[string]string{"host": pu.Hostname()}})
		}
	}
	return out, nil
}

func (e *Engine) scrape(ctx context.Context, t Target, opt StageOptions) ([]Artifact, error) {
	base := t.BaseURL
	if base == "" {
		base = "https://" + t.Domain + "/"
	}
	bu, err := url.Parse(base)
	if err != nil {
		return nil, err
	}
	sec := opt.timeout(12)
	e.emit(ProgressEvent{Stage: "scrape", Action: "get", URL: base, Timeout: sec})
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, base, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "AURA-intel/1.0")
	resp, err := e.httpWithTimeout(time.Duration(sec) * time.Second).Do(req)
	if err != nil {
		e.emit(ProgressEvent{Stage: "scrape", Action: "err", URL: base, Err: err.Error()})
		return nil, err
	}
	body, _ := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	_ = resp.Body.Close()
	htmlRes, err := extractor.ExtractHTML(bu, body)
	if err != nil {
		htmlRes = extractor.ExtractJS(bu, body)
	}
	e.emit(ProgressEvent{Stage: "scrape", Action: "ok", URL: base, Status: resp.StatusCode, N: len(htmlRes.URLs)})
	keep := func(raw string) (*url.URL, bool) {
		u, err := url.Parse(raw)
		if err != nil || u.Hostname() == "" {
			return nil, false
		}
		if opt.sameHost() && !SameScope(u.Hostname(), t.Domain) {
			return nil, false
		}
		path := u.Path
		if path == "" {
			path = "/"
		}
		if !IsSanePath(path) {
			return nil, false
		}
		return u, true
	}
	var out []Artifact
	jsSeen := map[string]struct{}{}
	for _, raw := range htmlRes.URLs {
		u, ok := keep(raw)
		if !ok {
			continue
		}
		out = append(out, Artifact{Kind: KindURL, Value: raw, Source: "scrape"})
		if strings.HasSuffix(strings.ToLower(u.Path), ".js") {
			out = append(out, Artifact{Kind: KindJS, Value: raw, Source: "scrape"})
			jsSeen[raw] = struct{}{}
		}
		if u.Path != "" && IsSanePath(u.Path) {
			out = append(out, Artifact{Kind: KindPath, Value: u.Path, Source: "scrape", Extra: map[string]string{"host": u.Hostname()}})
		}
	}
	follow := opt.followJS(8)
	if follow > 40 {
		follow = 40
	}
	n := 0
	jsTotal := len(jsSeen)
	if jsTotal > follow {
		jsTotal = follow
	}
	for raw := range jsSeen {
		if err := ctx.Err(); err != nil {
			return out, err
		}
		if n >= follow {
			break
		}
		n++
		e.emit(ProgressEvent{Stage: "scrape", Action: "get", URL: raw, N: n, Total: jsTotal, Timeout: sec})
		ju, err := url.Parse(raw)
		if err != nil {
			continue
		}
		jreq, err := http.NewRequestWithContext(ctx, http.MethodGet, raw, nil)
		if err != nil {
			continue
		}
		jresp, err := e.httpWithTimeout(time.Duration(sec) * time.Second).Do(jreq)
		if err != nil {
			e.emit(ProgressEvent{Stage: "scrape", Action: "err", URL: raw, Err: err.Error()})
			continue
		}
		jsBody, _ := io.ReadAll(io.LimitReader(jresp.Body, 512*1024))
		_ = jresp.Body.Close()
		found := extractor.ExtractJS(ju, jsBody).URLs
		e.emit(ProgressEvent{Stage: "scrape", Action: "ok", URL: raw, Status: jresp.StatusCode, N: len(found)})
		for _, u := range found {
			if _, ok := keep(u); !ok {
				continue
			}
			out = append(out, Artifact{Kind: KindURL, Value: u, Source: "js"})
		}
	}
	return out, nil
}

func (e *Engine) dirs(ctx context.Context, t Target, opt StageOptions) ([]Artifact, error) {
	base := t.BaseURL
	if base == "" {
		base = "https://" + t.Domain
	}
	words := e.stageWords(opt, wordlist.Dirs(), wordlist.DefaultDirs...)
	max := opt.limit(len(words), 20000)
	if max < len(words) {
		words = words[:max]
	}
	host := t.Domain
	if u, err := url.Parse(base); err == nil && u.Hostname() != "" {
		host = u.Hostname()
	}
	hide := opt.Hide
	if len(hide) == 0 {
		hide = []int{404}
	}
	hideSet := map[int]struct{}{}
	for _, c := range hide {
		hideSet[c] = struct{}{}
	}
	e.emit(ProgressEvent{Stage: "dirs", Action: "get", URL: base, Total: len(words)})
	res, err := discover.Run(ctx, discover.Options{
		BaseURL:  base,
		Wordlist: words,
		Workers:  opt.workers(16),
		RPS:      opt.rps(25),
		Timeout:  time.Duration(opt.timeout(8)) * time.Second,
		OnProgress: func(done, total int, r discover.Result) {
			ev := ProgressEvent{URL: r.Path, Status: r.StatusCode, Err: r.Error}
			if r.Error == "" && r.StatusCode != 0 {
				if _, skip := hideSet[r.StatusCode]; !skip {
					ev.Action = "hit"
					ev.Value = r.Path
					ev.Stage = "dirs"
					ev.N = done
					ev.Total = total
					e.emit(ev)
					return
				}
			}
			e.emitTick("dirs", done, total, 25, ev)
		},
	}, e.log())
	if err != nil && ctx.Err() == nil && len(res) == 0 {
		return nil, err
	}
	var out []Artifact
	for _, r := range res {
		if r.Error != "" || r.StatusCode == 0 {
			continue
		}
		if _, skip := hideSet[r.StatusCode]; skip {
			continue
		}
		path := "/" + strings.TrimLeft(r.Path, "/")
		if !IsSanePath(path) {
			continue
		}
		out = append(out, Artifact{
			Kind:   KindPath,
			Value:  path,
			Source: "dirs",
			Extra:  map[string]string{"status": fmt.Sprintf("%d", r.StatusCode), "length": fmt.Sprintf("%d", r.Length), "host": host},
		})
	}
	return out, nil
}

func (e *Engine) params(ctx context.Context, t Target, opt StageOptions) ([]Artifact, error) {
	base := t.BaseURL
	if base == "" {
		base = "https://" + t.Domain + "/"
	}
	sep := "?"
	if strings.Contains(base, "?") {
		sep = "&"
	}
	template := strings.TrimRight(base, "&") + sep + "FUZZ=1"
	words := e.stageWords(opt, wordlist.Params(), wordlist.DefaultParams...)
	filtered := words[:0]
	for _, w := range words {
		if IsTrackingParam(w) {
			continue
		}
		filtered = append(filtered, w)
	}
	words = filtered
	max := opt.limit(len(words), 20000)
	if max < len(words) {
		words = words[:max]
	}
	path := "/"
	host := t.Domain
	if u, err := url.Parse(base); err == nil {
		if u.Path != "" {
			path = u.Path
		}
		if u.Hostname() != "" {
			host = u.Hostname()
		}
	}
	hide := opt.Hide
	if len(hide) == 0 {
		hide = []int{404}
	}
	e.emit(ProgressEvent{Stage: "params", Action: "get", URL: template, Total: len(words)})
	hits, err := fuzz.Run(ctx, fuzz.Options{
		URL:      template,
		Wordlist: words,
		Workers:  opt.workers(10),
		RPS:      opt.rps(15),
		Timeout:  time.Duration(opt.timeout(8)) * time.Second,
		Hide:     hide,
		OnProgress: func(done, total int, h fuzz.Hit) {
			ev := ProgressEvent{URL: h.URL, Value: h.Payload, Status: h.StatusCode, Err: h.Error}
			if h.Error == "" && h.StatusCode != 0 && h.StatusCode != 404 {
				ev.Action = "hit"
				ev.Stage = "params"
				ev.N = done
				ev.Total = total
				e.emit(ev)
				return
			}
			e.emitTick("params", done, total, 25, ev)
		},
	})
	if err != nil && len(hits) == 0 {
		return nil, err
	}
	var out []Artifact
	counts := map[int]int{}
	for _, h := range hits {
		counts[h.Length]++
	}
	commonLen := 0
	commonN := 0
	for l, n := range counts {
		if n > commonN {
			commonN = n
			commonLen = l
		}
	}
	for _, h := range hits {
		if h.StatusCode == 404 || h.Error != "" || IsTrackingParam(h.Payload) {
			continue
		}
		if h.Length == commonLen && commonN > 3 {
			continue
		}
		out = append(out, Artifact{
			Kind:   KindParam,
			Value:  h.Payload,
			Source: "params",
			Extra:  map[string]string{"status": fmt.Sprintf("%d", h.StatusCode), "path": path, "host": host},
		})
	}
	return out, nil
}

func (e *Engine) stageWords(opt StageOptions, fallback []string, seclists ...string) []string {
	if len(opt.Wordlist) > 0 {
		out := make([]string, 0, len(opt.Wordlist))
		for _, w := range opt.Wordlist {
			w = strings.TrimSpace(w)
			if w != "" {
				out = append(out, w)
			}
		}
		if len(out) > 0 {
			return out
		}
	}
	if opt.WordlistPath != "" && e.Lists != nil {
		if words, err := e.Lists.Load(opt.WordlistPath); err == nil && len(words) > 0 {
			return words
		}
	}
	if e.Lists != nil {
		return e.Lists.LoadFirst(fallback, seclists...)
	}
	return fallback
}

func liveOrSeed(t Target, arts []Artifact) []string {
	var hosts []string
	seen := map[string]struct{}{}
	add := func(h string) {
		h = strings.ToLower(strings.TrimSpace(h))
		if h == "" || (t.Domain != "" && !SameScope(h, t.Domain)) {
			return
		}
		if _, ok := seen[h]; ok {
			return
		}
		seen[h] = struct{}{}
		hosts = append(hosts, h)
	}
	for _, a := range arts {
		if a.Kind == KindHost {
			add(a.Value)
		}
	}
	if len(hosts) == 0 {
		add(t.Domain)
		for _, a := range arts {
			if a.Kind == KindSubdomain {
				add(a.Value)
			}
		}
	}
	return hosts
}

func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
