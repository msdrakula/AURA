package intel

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestIsSanePath(t *testing.T) {
	ok := []string{"/", "/api", "/api/v1/users", "/.well-known/security.txt", "/about", "/admin/"}
	for _, p := range ok {
		if !IsSanePath(p) {
			t.Errorf("want sane %q", p)
		}
	}
	bad := []string{
		"/ Testez",
		"/ https://gorest.co.in",
		"/http://httpbin.org/",
		"/api&quot;;\nString",
		"/\nREST",
		"/[][2]",
		"/\"",
		"/'",
		"/)",
		"<a>",
	}
	for _, p := range bad {
		if IsSanePath(p) {
			t.Errorf("want junk %q", p)
		}
	}
}

func TestSameScope(t *testing.T) {
	if !SameScope("reqres.in", "reqres.in") || !SameScope("api.reqres.in", "reqres.in") {
		t.Fatal("in-scope")
	}
	if SameScope("gorest.co.in", "reqres.in") || SameScope("httpbin.org", "reqres.in") {
		t.Fatal("foreign host")
	}
}

func TestBuildMapDropsJunkAndTracking(t *testing.T) {
	t0 := Target{Domain: "reqres.in", BaseURL: "https://reqres.in/"}
	m := BuildMap(t0, []Artifact{
		{Kind: KindURL, Value: "https://reqres.in/?utm_source=x&fbclid=1&id=2"},
		{Kind: KindURL, Value: "https://gorest.co.in/users"},
		{Kind: KindPath, Value: "/ api", Extra: map[string]string{"host": "reqres.in"}},
		{Kind: KindPath, Value: "/api/users", Extra: map[string]string{"host": "reqres.in"}},
		{Kind: KindParam, Value: "page", Extra: map[string]string{"host": "reqres.in", "path": "/api/users"}},
		{Kind: KindParam, Value: "utm_medium"},
	})
	if len(m.Hosts) != 1 || m.Hosts[0].Host != "reqres.in" {
		t.Fatalf("hosts=%+v", m.Hosts)
	}
	var paths []string
	var params []string
	for _, p := range m.Hosts[0].Paths {
		paths = append(paths, p.Path)
		params = append(params, p.Params...)
	}
	if contains(paths, "/ api") {
		t.Fatalf("junk path survived: %v", paths)
	}
	if !contains(paths, "/api/users") {
		t.Fatalf("missing /api/users in %v", paths)
	}
	if contains(params, "utm_source") || contains(params, "fbclid") || contains(params, "utm_medium") {
		t.Fatalf("tracking params on tree: %v", params)
	}
	if !contains(params, "id") || !contains(params, "page") {
		t.Fatalf("real params missing: %v", params)
	}
}

func TestWaybackKeepsInScopeDropsJunk(t *testing.T) {
	st := newMem()
	eng := &Engine{
		Store: st,
		HTTP: &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			body := `[["original"],["https://reqres.in/api/users"],["https://gorest.co.in/x"],["https://reqres.in/ https://httpbin.org"]]`
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(body)),
				Header:     make(http.Header),
				Request:    r,
			}, nil
		})},
	}
	tg, err := eng.EnsureTarget("https://reqres.in", true)
	if err != nil {
		t.Fatal(err)
	}
	same := true
	res, err := eng.RunStage(context.Background(), tg.ID, "urls_passive", StageOptions{SameHost: &same, Limit: 50})
	if err != nil {
		t.Fatal(err)
	}
	if res.Added < 1 {
		t.Fatalf("added=%d", res.Added)
	}
	for _, a := range res.Artifacts {
		if a.Kind == KindURL {
			u, err := url.Parse(a.Value)
			if err != nil || !SameScope(u.Hostname(), "reqres.in") {
				t.Fatalf("off-host url %s", a.Value)
			}
		}
		if a.Kind == KindPath && !IsSanePath(a.Value) {
			t.Fatalf("junk path %s", a.Value)
		}
	}
}

func TestWebPortsCustomList(t *testing.T) {
	ports := webPortList("https://lab.local:16126/", "8080 8443")
	if len(ports) == 0 || ports[0] != "16126" {
		t.Fatalf("url port first: %v", ports)
	}
	if !contains(ports, "8080") || !contains(ports, "8443") {
		t.Fatalf("custom: %v", ports)
	}
	if contains(ports, "443") {
		t.Fatalf("default ports should not leak into custom list: %v", ports)
	}
}
