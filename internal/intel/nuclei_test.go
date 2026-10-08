package intel

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseNucleiLine(t *testing.T) {
	line := `{"template-id":"cve-2021-41773","host":"http://example.com","matched-at":"http://example.com/cgi-bin/.%2e/bin/sh","type":"http","info":{"name":"Apache 2.4.49 RCE","severity":"critical","description":"Path traversal in Apache","tags":"apache,rce"}}`
	a, ok := parseNucleiLine(line)
	require.True(t, ok)
	assert.Equal(t, KindVuln, a.Kind)
	assert.Equal(t, "cve-2021-41773", a.Value)
	assert.Equal(t, "vulnscan", a.Source)
	assert.Equal(t, "Apache 2.4.49 RCE", a.Extra["name"])
	assert.Equal(t, "critical", a.Extra["severity"])
	assert.Equal(t, "Path traversal in Apache", a.Extra["description"])
	assert.Equal(t, "apache,rce", a.Extra["tags"])
	assert.Equal(t, "http://example.com/cgi-bin/.%2e/bin/sh", a.Extra["matched"])
	assert.Equal(t, "http://example.com/cgi-bin/.%2e/bin/sh", a.Extra["host"])
}

func TestParseNucleiLineFallbackHost(t *testing.T) {
	line := `{"template-id":"exposed-panels","host":"https://panel.example.com","info":{"name":"Login Panel","severity":"info"}}`
	a, ok := parseNucleiLine(line)
	require.True(t, ok)
	assert.Equal(t, "exposed-panels", a.Value)
	assert.Equal(t, "", a.Extra["matched"])
	assert.Equal(t, "https://panel.example.com", a.Extra["host"])
}

func TestParseNucleiLineSkipsJunk(t *testing.T) {
	for _, line := range []string{
		`{"host":"http://example.com","info":{"name":"no template id"}}`,
		`not json at all`,
		``,
	} {
		_, ok := parseNucleiLine(line)
		assert.False(t, ok, "line: %q", line)
	}
}

func TestBuildNucleiTargets(t *testing.T) {
	tg := Target{BaseURL: "https://example.com/"}
	arts := []Artifact{
		{Kind: KindHost, Value: "example.com"},
		{Kind: KindHost, Value: "api.example.com"},
		{Kind: KindURL, Value: "https://api.example.com:8443/users?page=1"},
		{Kind: KindPort, Value: "example.com:8080", Extra: map[string]string{"http": "true"}},
		{Kind: KindPort, Value: "example.com:22", Extra: map[string]string{"proto": "ssh"}},
		{Kind: KindPort, Value: "example.com:9090", Extra: map[string]string{"proto": "http-alt"}},
		{Kind: KindPort, Value: "example.com:1337"},
	}
	targets := buildNucleiTargets(tg, arts)
	assert.Equal(t, []string{
		"https://example.com/",
		"example.com",
		"api.example.com",
		"https://api.example.com:8443",
		"example.com:8080",
		"example.com:9090",
	}, targets)
}

func TestBuildNucleiTargetsDedupAndOrder(t *testing.T) {
	tg := Target{BaseURL: "https://example.com/"}
	arts := []Artifact{
		{Kind: KindHost, Value: "example.com"},
		{Kind: KindHost, Value: "example.com"},
		{Kind: KindURL, Value: "https://example.com/a"},
		{Kind: KindURL, Value: "https://example.com/b"},
	}
	targets := buildNucleiTargets(tg, arts)
	assert.Equal(t, []string{
		"https://example.com/",
		"example.com",
		"https://example.com",
	}, targets)
}

func TestBuildNucleiTargetsCap(t *testing.T) {
	var arts []Artifact
	for i := 0; i < 150; i++ {
		arts = append(arts, Artifact{Kind: KindHost, Value: fmt.Sprintf("h%d.example.com", i)})
	}
	targets := buildNucleiTargets(Target{}, arts)
	assert.Len(t, targets, 100)
	assert.Equal(t, "h0.example.com", targets[0])
	assert.Equal(t, "h99.example.com", targets[99])
}

func TestBuildNucleiTargetsNoBaseURL(t *testing.T) {
	arts := []Artifact{
		{Kind: KindHost, Value: "  "},
		{Kind: KindURL, Value: "not a url"},
	}
	assert.Nil(t, buildNucleiTargets(Target{}, arts))
}
