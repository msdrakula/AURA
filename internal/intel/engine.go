package intel

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"meb/internal/discover"
	"meb/internal/extractor"
	"meb/internal/fuzz"
	"meb/internal/wordlist"
)

// Engine runs mindmap stages against a Store.
type Engine struct {
	Store  Store
	HTTP   *http.Client
	Log    *zap.Logger
	Lists  *wordlist.Library
	Lookup func(ctx context.Context, host string) ([]string, error)
	Dial   func(ctx context.Context, network, addr string) (net.Conn, error)
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
			if authorized && !t.Authorized {
				t.Authorized = true
				t.BaseURL = base
				_ = e.Store.UpdateTarget(t)
			}
			if t.BaseURL == "" && base != "" {
				t.BaseURL = base
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

// RunStage executes one mindmap step.
func (e *Engine) RunStage(ctx context.Context, targetID, stageID string) (RunResult, error) {
	cat, ok := stageByID(stageID)
	if !ok {
		return RunResult{}, fmt.Errorf("unknown stage %q", stageID)
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

	arts, err := e.Store.ListArtifacts(targetID)
	if err != nil {
		return RunResult{}, err
	}

	var added []Artifact
	var runErr error
	switch stageID {
	case "subdomains_passive":
		added, runErr = e.crtsh(ctx, t)
	case "dns_brute":
		added, runErr = e.dnsBrute(ctx, t)
	case "live_hosts":
		added, runErr = e.liveHosts(ctx, t, arts)
	case "web_ports":
		added, runErr = e.webPorts(ctx, t, arts)
	case "tech":
		added, runErr = e.tech(ctx, t, arts)
	case "urls_passive":
		added, runErr = e.wayback(ctx, t)
	case "scrape":
		added, runErr = e.scrape(ctx, t)
	case "dirs":
		added, runErr = e.dirs(ctx, t)
	case "params":
		added, runErr = e.params(ctx, t)
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
	if runErr != nil {
		st.Status = StatusError
		st.Summary = runErr.Error()
		_ = e.Store.UpsertStage(targetID, st)
		return RunResult{Stage: st, Added: n, Artifacts: added}, runErr
	}
	st.Status = StatusDone
	st.Summary = fmt.Sprintf("добавлено фактов: %d", n)
	_ = e.Store.UpsertStage(targetID, st)
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
		if u.Path != "" {
			items = append(items, Artifact{TargetID: t.ID, Kind: KindPath, Value: u.Path, Source: "proxy", Extra: map[string]string{"host": host}})
		}
		for key := range u.Query() {
			items = append(items, Artifact{TargetID: t.ID, Kind: KindParam, Value: key, Source: "proxy"})
		}
		for _, a := range items {
			if err := e.Store.AddArtifact(a); err == nil {
				n++
			}
		}
	}
	return n, nil
}

func (e *Engine) crtsh(ctx context.Context, t Target) ([]Artifact, error) {
	q := url.QueryEscape("%." + t.Domain)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://crt.sh/?q="+q+"&output=json", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "MEB-intel/1.0")
	resp, err := e.http().Do(req)
	if err != nil {
		return nil, fmt.Errorf("crt.sh: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(io.LimitReader(resp.Body, 8<<20))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("crt.sh: HTTP %d", resp.StatusCode)
	}
	names := parseCRTNames(body, t.Domain)
	var out []Artifact
	for _, name := range names {
		out = append(out, Artifact{Kind: KindSubdomain, Value: name, Source: "crtsh"})
	}
	return out, nil
}

func (e *Engine) dnsBrute(ctx context.Context, t Target) ([]Artifact, error) {
	lookup := e.Lookup
	if lookup == nil {
		lookup = func(ctx context.Context, host string) ([]string, error) {
			var d net.Resolver
			return d.LookupHost(ctx, host)
		}
	}
	var out []Artifact
	for _, p := range e.Lists.LoadFirst(wordlist.HostPrefixes(), wordlist.DefaultDNS...) {
		host := p + "." + t.Domain
		addrs, err := lookup(ctx, host)
		if err != nil || len(addrs) == 0 {
			continue
		}
		out = append(out, Artifact{
			Kind:   KindSubdomain,
			Value:  host,
			Source: "dns",
			Extra:  map[string]string{"ip": addrs[0]},
		})
	}
	return out, nil
}

func (e *Engine) liveHosts(ctx context.Context, t Target, arts []Artifact) ([]Artifact, error) {
	seen := map[string]struct{}{}
	var out []Artifact
	add := func(host string, extra map[string]string) {
		host = strings.ToLower(strings.TrimSpace(host))
		if host == "" {
			return
		}
		if _, ok := seen[host]; ok {
			return
		}
		seen[host] = struct{}{}
		out = append(out, Artifact{Kind: KindHost, Value: host, Source: "live", Extra: extra})
	}
	if t.BaseURL != "" {
		if live, extra := e.probeURL(ctx, t.BaseURL); live {
			if u, err := url.Parse(t.BaseURL); err == nil {
				add(u.Hostname(), extra)
			}
		}
	}
	hosts := []string{t.Domain}
	for _, a := range arts {
		if a.Kind == KindSubdomain || a.Kind == KindHost {
			hosts = append(hosts, a.Value)
		}
	}
	for _, h := range hosts {
		if _, ok := seen[strings.ToLower(h)]; ok {
			continue
		}
		live, extra := e.probeHost(ctx, h)
		if live {
			add(h, extra)
		}
	}
	return out, nil
}

func (e *Engine) probeURL(ctx context.Context, raw string) (bool, map[string]string) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, raw, nil)
	if err != nil {
		return false, nil
	}
	req.Header.Set("User-Agent", "MEB-intel/1.0")
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

func (e *Engine) probeHost(ctx context.Context, host string) (bool, map[string]string) {
	for _, scheme := range []string{"https", "http"} {
		u := scheme + "://" + host + "/"
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
		if err != nil {
			continue
		}
		req.Header.Set("User-Agent", "MEB-intel/1.0")
		resp, err := e.http().Do(req)
		if err != nil {
			continue
		}
		_, _ = io.Copy(io.Discard, io.LimitReader(resp.Body, 64*1024))
		_ = resp.Body.Close()
		if resp.StatusCode > 0 {
			return true, map[string]string{
				"scheme": scheme,
				"status": fmt.Sprintf("%d", resp.StatusCode),
				"server": resp.Header.Get("Server"),
			}
		}
	}
	return false, nil
}

func (e *Engine) webPorts(ctx context.Context, t Target, arts []Artifact) ([]Artifact, error) {
	dial := e.Dial
	if dial == nil {
		d := net.Dialer{Timeout: 2 * time.Second}
		dial = d.DialContext
	}
	hosts := liveOrSeed(t, arts)
	ports := []string{"80", "443", "8080", "8443", "3000", "8000", "8008", "8888"}
	var out []Artifact
	for _, h := range hosts {
		for _, p := range ports {
			addr := net.JoinHostPort(h, p)
			c, err := dial(ctx, "tcp", addr)
			if err != nil {
				continue
			}
			_ = c.Close()
			out = append(out, Artifact{
				Kind:   KindPort,
				Value:  h + ":" + p,
				Source: "ports",
				Extra:  map[string]string{"host": h, "port": p},
			})
		}
	}
	return out, nil
}

func (e *Engine) tech(ctx context.Context, t Target, arts []Artifact) ([]Artifact, error) {
	var out []Artifact
	seen := map[string]struct{}{}
	fetch := func(raw, host string) {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, raw, nil)
		if err != nil {
			return
		}
		req.Header.Set("User-Agent", "MEB-intel/1.0")
		resp, err := e.http().Do(req)
		if err != nil {
			return
		}
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 256*1024))
		_ = resp.Body.Close()
		for _, name := range DetectTech(resp.Header, body) {
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

func (e *Engine) wayback(ctx context.Context, t Target) ([]Artifact, error) {
	u := "https://web.archive.org/cdx/search/cdx?url=*." + url.QueryEscape(t.Domain) + "/*&output=json&fl=original&collapse=urlkey&limit=150"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "MEB-intel/1.0")
	resp, err := e.http().Do(req)
	if err != nil {
		return nil, fmt.Errorf("wayback: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(io.LimitReader(resp.Body, 4<<20))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("wayback: HTTP %d", resp.StatusCode)
	}
	urls := parseWayback(body)
	var out []Artifact
	for _, raw := range urls {
		out = append(out, Artifact{Kind: KindURL, Value: raw, Source: "wayback"})
		if pu, err := url.Parse(raw); err == nil && pu.Path != "" && pu.Path != "/" {
			out = append(out, Artifact{Kind: KindPath, Value: pu.Path, Source: "wayback", Extra: map[string]string{"host": pu.Hostname()}})
		}
	}
	return out, nil
}

func (e *Engine) scrape(ctx context.Context, t Target) ([]Artifact, error) {
	base := t.BaseURL
	if base == "" {
		base = "https://" + t.Domain + "/"
	}
	bu, err := url.Parse(base)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, base, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "MEB-intel/1.0")
	resp, err := e.http().Do(req)
	if err != nil {
		return nil, err
	}
	body, _ := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	_ = resp.Body.Close()
	htmlRes, err := extractor.ExtractHTML(bu, body)
	if err != nil {
		htmlRes = extractor.ExtractJS(bu, body)
	}
	var out []Artifact
	jsSeen := map[string]struct{}{}
	for _, raw := range htmlRes.URLs {
		out = append(out, Artifact{Kind: KindURL, Value: raw, Source: "scrape"})
		u, err := url.Parse(raw)
		if err != nil {
			continue
		}
		if strings.HasSuffix(strings.ToLower(u.Path), ".js") {
			out = append(out, Artifact{Kind: KindJS, Value: raw, Source: "scrape"})
			jsSeen[raw] = struct{}{}
		}
		if u.Path != "" {
			out = append(out, Artifact{Kind: KindPath, Value: u.Path, Source: "scrape", Extra: map[string]string{"host": u.Hostname()}})
		}
	}
	n := 0
	for raw := range jsSeen {
		if n >= 8 {
			break
		}
		n++
		ju, err := url.Parse(raw)
		if err != nil {
			continue
		}
		jreq, err := http.NewRequestWithContext(ctx, http.MethodGet, raw, nil)
		if err != nil {
			continue
		}
		jresp, err := e.http().Do(jreq)
		if err != nil {
			continue
		}
		jsBody, _ := io.ReadAll(io.LimitReader(jresp.Body, 512*1024))
		_ = jresp.Body.Close()
		for _, found := range extractor.ExtractJS(ju, jsBody).URLs {
			out = append(out, Artifact{Kind: KindURL, Value: found, Source: "js"})
		}
	}
	return out, nil
}

func (e *Engine) dirs(ctx context.Context, t Target) ([]Artifact, error) {
	base := t.BaseURL
	if base == "" {
		base = "https://" + t.Domain
	}
	res, err := discover.Run(ctx, discover.Options{
		BaseURL:  base,
		Wordlist: e.Lists.LoadFirst(wordlist.Dirs(), wordlist.DefaultDirs...),
		Workers:  8,
		RPS:      12,
		Timeout:  8 * time.Second,
	}, e.log())
	if err != nil && ctx.Err() == nil && len(res) == 0 {
		return nil, err
	}
	var out []Artifact
	for _, r := range res {
		if r.Error != "" || r.StatusCode == 0 || r.StatusCode == 404 {
			continue
		}
		path := "/" + strings.TrimLeft(r.Path, "/")
		out = append(out, Artifact{
			Kind:   KindPath,
			Value:  path,
			Source: "dirs",
			Extra:  map[string]string{"status": fmt.Sprintf("%d", r.StatusCode), "length": fmt.Sprintf("%d", r.Length)},
		})
	}
	return out, nil
}

func (e *Engine) params(ctx context.Context, t Target) ([]Artifact, error) {
	base := t.BaseURL
	if base == "" {
		base = "https://" + t.Domain + "/"
	}
	sep := "?"
	if strings.Contains(base, "?") {
		sep = "&"
	}
	template := strings.TrimRight(base, "&") + sep + "FUZZ=1"
	hits, err := fuzz.Run(ctx, fuzz.Options{
		URL:      template,
		Wordlist: e.Lists.LoadFirst(wordlist.Params(), wordlist.DefaultParams...),
		Workers:  6,
		RPS:      10,
		Timeout:  8 * time.Second,
		Hide:     []int{404},
	})
	if err != nil && len(hits) == 0 {
		return nil, err
	}
	var out []Artifact
	baseLen := 0
	counts := map[int]int{}
	for _, h := range hits {
		counts[h.Length]++
	}
	for lenVal, n := range counts {
		if n > baseLen {
			baseLen = n
			_ = lenVal
		}
	}
	// Keep responses whose length is uncommon among the batch — likely a real param.
	commonLen := 0
	commonN := 0
	for l, n := range counts {
		if n > commonN {
			commonN = n
			commonLen = l
		}
	}
	for _, h := range hits {
		if h.StatusCode == 404 || h.Error != "" {
			continue
		}
		if h.Length == commonLen && commonN > 3 {
			continue
		}
		out = append(out, Artifact{
			Kind:   KindParam,
			Value:  h.Payload,
			Source: "params",
			Extra:  map[string]string{"status": fmt.Sprintf("%d", h.StatusCode)},
		})
	}
	return out, nil
}

func liveOrSeed(t Target, arts []Artifact) []string {
	var hosts []string
	seen := map[string]struct{}{}
	add := func(h string) {
		h = strings.ToLower(strings.TrimSpace(h))
		if h == "" {
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
