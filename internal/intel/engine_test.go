package intel

import (
	"context"
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

func TestNormalizeInput(t *testing.T) {
	d, b, err := NormalizeInput("Example.COM")
	if err != nil || d != "example.com" || b != "https://example.com" {
		t.Fatalf("got %q %q %v", d, b, err)
	}
	d, b, err = NormalizeInput("https://app.lab.local/v1")
	if err != nil || d != "app.lab.local" || !strings.Contains(b, "/v1") {
		t.Fatalf("got %q %q %v", d, b, err)
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
