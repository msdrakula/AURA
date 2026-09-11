package intel

import (
	"context"
	"errors"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"
)

type memStore struct {
	mu   sync.Mutex
	tg   map[string]Target
	st   map[string]map[string]StageState
	arts map[string][]Artifact
}

func newMem() *memStore {
	return &memStore{
		tg:   map[string]Target{},
		st:   map[string]map[string]StageState{},
		arts: map[string][]Artifact{},
	}
}

func (m *memStore) CreateTarget(t Target) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.tg[t.ID] = t
	return nil
}
func (m *memStore) GetTarget(id string) (Target, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	t, ok := m.tg[id]
	if !ok {
		return Target{}, simpleError("not found")
	}
	return t, nil
}
func (m *memStore) ListTargets() ([]Target, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	var out []Target
	for _, t := range m.tg {
		out = append(out, t)
	}
	return out, nil
}
func (m *memStore) UpdateTarget(t Target) error { return m.CreateTarget(t) }
func (m *memStore) UpsertStage(targetID string, st StageState) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.st[targetID] == nil {
		m.st[targetID] = map[string]StageState{}
	}
	m.st[targetID][st.ID] = st
	return nil
}
func (m *memStore) ListStages(targetID string) ([]StageState, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	var out []StageState
	for _, st := range m.st[targetID] {
		out = append(out, st)
	}
	return out, nil
}
func (m *memStore) AddArtifact(a Artifact) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, x := range m.arts[a.TargetID] {
		if x.Kind == a.Kind && x.Value == a.Value {
			return nil
		}
	}
	m.arts[a.TargetID] = append(m.arts[a.TargetID], a)
	return nil
}
func (m *memStore) ListArtifacts(targetID string) ([]Artifact, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	return append([]Artifact{}, m.arts[targetID]...), nil
}

func TestCatalogTechIsActive(t *testing.T) {
	cat, ok := stageByID("tech")
	if !ok || cat.Mode != ModeActive {
		t.Fatalf("tech mode=%q ok=%v", cat.Mode, ok)
	}
}

func TestEnsureTargetRevokesAuthorized(t *testing.T) {
	st := newMem()
	eng := &Engine{Store: st}
	tg, err := eng.EnsureTarget("lab.local", true)
	if err != nil {
		t.Fatal(err)
	}
	if !tg.Authorized {
		t.Fatal("want authorized")
	}
	tg, err = eng.EnsureTarget("lab.local", false)
	if err != nil {
		t.Fatal(err)
	}
	if tg.Authorized {
		t.Fatal("checkbox off must revoke lab flag")
	}
}

func TestNormalizeInput(t *testing.T) {
	d, b, err := NormalizeInput("Example.COM")
	if err != nil || d != "example.com" || b != "https://example.com/" {
		t.Fatalf("got %q %q %v", d, b, err)
	}
	d, b, err = NormalizeInput("https://app.lab.local/v1")
	if err != nil || d != "app.lab.local" || !strings.Contains(b, "/v1") {
		t.Fatalf("got %q %q %v", d, b, err)
	}
	d, b, err = NormalizeInput("http://62.173.140.174:16126/")
	if err != nil || d != "62.173.140.174" || b != "http://62.173.140.174:16126/" {
		t.Fatalf("port kept: %q %q %v", d, b, err)
	}
	d, b, err = NormalizeInput("https://app.lab.local:443/api")
	if err != nil || d != "app.lab.local" || b != "https://app.lab.local/api" {
		t.Fatalf("default https port omitted: %q %q %v", d, b, err)
	}
	d, b, err = NormalizeInput("http://app.lab.local")
	if err != nil || b != "http://app.lab.local/" {
		t.Fatalf("http default: %q %q %v", d, b, err)
	}
	d, b, err = NormalizeInput("10.0.0.8:8080")
	if err != nil || d != "10.0.0.8" || b != "http://10.0.0.8:8080/" {
		t.Fatalf("bare host:port: %q %q %v", d, b, err)
	}
}

func TestParseCRTNames(t *testing.T) {
	body := []byte(`[{"name_value":"example.com\ndev.example.com","common_name":"www.example.com"}]`)
	got := parseCRTNames(body, "example.com")
	if len(got) < 3 {
		t.Fatalf("names=%v", got)
	}
}

func TestDetectTech(t *testing.T) {
	h := http.Header{}
	h.Set("Server", "nginx")
	h.Set("X-Powered-By", "Express")
	got := DetectTech(h, []byte(`<html>jquery wp-content</html>`))
	joined := strings.Join(got, ",")
	if !strings.Contains(strings.ToLower(joined), "nginx") || !strings.Contains(joined, "WordPress") {
		t.Fatalf("tech=%v", got)
	}
}

func TestBuildMap(t *testing.T) {
	t0 := Target{Domain: "lab.local", BaseURL: "https://lab.local/"}
	m := BuildMap(t0, []Artifact{
		{Kind: KindHost, Value: "lab.local"},
		{Kind: KindPath, Value: "/api", Extra: map[string]string{"host": "lab.local", "status": "200"}},
		{Kind: KindParam, Value: "id"},
		{Kind: KindTech, Value: "nginx", Extra: map[string]string{"host": "lab.local"}},
	})
	if m.Stats.Paths < 1 || len(m.Hosts) == 0 || !m.Hosts[0].Live {
		t.Fatalf("map=%+v", m)
	}
}

func TestEngineLiveAndActiveGuard(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Server", "nginx")
		w.Header().Set("X-Powered-By", "PHP")
		_, _ = io.WriteString(w, `<html><a href="/login">x</a><script src="/app.js"></script></html>`)
	}))
	t.Cleanup(srv.Close)

	st := newMem()
	eng := &Engine{Store: st, HTTP: srv.Client()}
	target, err := eng.EnsureTarget(srv.URL, false)
	if err != nil {
		t.Fatal(err)
	}

	_, err = eng.RunStage(context.Background(), target.ID, "live_hosts")
	if err == nil {
		t.Fatal("expected authorization error")
	}
	_, err = eng.RunStage(context.Background(), target.ID, "tech")
	if err == nil {
		t.Fatal("expected tech to require authorization")
	}
	target.Authorized = true
	_ = st.UpdateTarget(target)
	res, err := eng.RunStage(context.Background(), target.ID, "live_hosts")
	if err != nil {
		t.Fatal(err)
	}
	if res.Added < 1 {
		t.Fatalf("added=%d", res.Added)
	}
	res, err = eng.RunStage(context.Background(), target.ID, "tech")
	if err != nil {
		t.Fatal(err)
	}
	if res.Added < 1 {
		t.Fatalf("tech added=%d", res.Added)
	}
}

func TestWebPortsUsesDial(t *testing.T) {
	st := newMem()
	eng := &Engine{
		Store: st,
		Dial: func(ctx context.Context, network, addr string) (net.Conn, error) {
			c1, c2 := net.Pipe()
			_ = c2.Close()
			return c1, nil
		},
		HTTP: &http.Client{Transport: okTransport{}},
	}
	tg, err := eng.EnsureTarget("lab.local", true)
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	res, err := eng.RunStage(ctx, tg.ID, "web_ports")
	if err != nil {
		t.Fatal(err)
	}
	if res.Added < 1 {
		t.Fatalf("ports added=%d", res.Added)
	}
}

func TestWebPortsIncludesURLPort(t *testing.T) {
	st := newMem()
	var seen []string
	eng := &Engine{
		Store: st,
		Dial: func(ctx context.Context, network, addr string) (net.Conn, error) {
			seen = append(seen, addr)
			return nil, errors.New("closed")
		},
	}
	tg, err := eng.EnsureTarget("https://62.173.140.174:16126/", true)
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if _, err := eng.RunStage(ctx, tg.ID, "web_ports"); err != nil {
		t.Fatal(err)
	}
	want := "62.173.140.174:16126"
	if len(seen) == 0 || seen[0] != want {
		t.Fatalf("URL port should be first, got %v", seen)
	}
}

func TestWebPortsTCPOnlyNotCounted(t *testing.T) {
	st := newMem()
	eng := &Engine{
		Store: st,
		Dial: func(ctx context.Context, network, addr string) (net.Conn, error) {
			c1, c2 := net.Pipe()
			_ = c2.Close()
			return c1, nil
		},
		HTTP: &http.Client{Transport: failTransport{}},
	}
	tg, err := eng.EnsureTarget("lab.local", true)
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	res, err := eng.RunStage(ctx, tg.ID, "web_ports")
	if err != nil {
		t.Fatal(err)
	}
	if res.Added != 0 {
		t.Fatalf("TCP-only ports must not be counted as web, added=%d", res.Added)
	}
}

func TestWebPortListPutsURLPortFirst(t *testing.T) {
	ports := webPortList("https://62.173.140.174:16126/")
	if len(ports) == 0 || ports[0] != "16126" {
		t.Fatalf("got %v", ports)
	}
}

type okTransport struct{}

func (okTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader("ok")),
		Header:     make(http.Header),
		Request:    req,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
	}, nil
}

type failTransport struct{}

func (failTransport) RoundTrip(*http.Request) (*http.Response, error) {
	return nil, errors.New("no http")
}

type blockingTransport struct{}

func (blockingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	<-req.Context().Done()
	return nil, req.Context().Err()
}

func TestRunStageCancelResetsIdle(t *testing.T) {
	st := newMem()
	eng := &Engine{
		Store: st,
		HTTP:  &http.Client{Transport: blockingTransport{}},
	}
	tg, err := eng.EnsureTarget("https://lab.local", true)
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	errCh := make(chan error, 1)
	resCh := make(chan RunResult, 1)
	go func() {
		res, err := eng.RunStage(ctx, tg.ID, "live_hosts")
		resCh <- res
		errCh <- err
	}()
	time.Sleep(40 * time.Millisecond)
	cancel()
	select {
	case err := <-errCh:
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("err=%v", err)
		}
		res := <-resCh
		if res.Stage.Status != StatusIdle {
			t.Fatalf("status=%s summary=%s", res.Stage.Status, res.Stage.Summary)
		}
	case <-time.After(2 * time.Second):
		cancel()
		t.Fatal("RunStage did not return after cancel")
	}
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func TestProgressEventsOnWayback(t *testing.T) {
	st := newMem()
	var mu sync.Mutex
	var got []ProgressEvent
	eng := &Engine{
		Store: st,
		HTTP: &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			body := `[["original"],["https://lab.local/old"]]`
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(body)),
				Header:     make(http.Header),
				Request:    r,
			}, nil
		})},
		Progress: func(ev ProgressEvent) {
			mu.Lock()
			got = append(got, ev)
			mu.Unlock()
		},
	}
	tg, err := eng.EnsureTarget("https://lab.local", true)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := eng.RunStage(context.Background(), tg.ID, "urls_passive"); err != nil {
		t.Fatal(err)
	}
	mu.Lock()
	defer mu.Unlock()
	actions := map[string]bool{}
	for _, ev := range got {
		if ev.Stage != "urls_passive" {
			continue
		}
		actions[ev.Action] = true
		if ev.Action == "get" && !strings.Contains(ev.URL, "web.archive.org") {
			t.Fatalf("get url = %q", ev.URL)
		}
	}
	for _, want := range []string{"start", "get", "ok", "done"} {
		if !actions[want] {
			t.Fatalf("missing action %s in %+v", want, got)
		}
	}
}
