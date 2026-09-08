package extractor

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func mustURL(t *testing.T, raw string) *url.URL {
	t.Helper()
	u, err := url.Parse(raw)
	require.NoError(t, err)
	return u
}

func TestExtractHTML(t *testing.T) {
	t.Parallel()
	base := mustURL(t, "https://app.example.com/app/home")
	doc := []byte(`<!doctype html>
<html>
<head>
  <link href="/static/app.css" rel="stylesheet">
  <script src="../js/vendor.js"></script>
</head>
<body>
  <a href="/api/v1/users">users</a>
  <a href="https://app.example.com/api/v1/users">dup</a>
  <a href="#section">skip hash</a>
  <a href="javascript:void(0)">skip js</a>
  <form action="/login" method="post"></form>
  <img src="/img/logo.png">
  <script>
    fetch("/api/v1/orders");
    axios.get('https://app.example.com/graphql');
    xhr.open("GET", "/health");
    const extra = "/api/v2/items";
  </script>
</body>
</html>`)

	got, err := ExtractHTML(base, doc)
	require.NoError(t, err)
	assert.Equal(t, []string{
		"https://app.example.com/api/v1/orders",
		"https://app.example.com/api/v1/users",
		"https://app.example.com/api/v2/items",
		"https://app.example.com/graphql",
		"https://app.example.com/health",
		"https://app.example.com/img/logo.png",
		"https://app.example.com/js/vendor.js",
		"https://app.example.com/login",
		"https://app.example.com/static/app.css",
	}, got.URLs)
}

func TestExtractJS(t *testing.T) {
	t.Parallel()
	base := mustURL(t, "http://127.0.0.1:8080/debug/")
	src := []byte(`
		fetch("/api/v1/status");
		fetch('http://127.0.0.1:8080/api/v1/status');
		client.open("POST", "/submit");
	`)
	got := ExtractJS(base, src)
	assert.Equal(t, []string{
		"http://127.0.0.1:8080/api/v1/status",
		"http://127.0.0.1:8080/submit",
	}, got.URLs)
}

func TestResolveSkipsNonHTTP(t *testing.T) {
	t.Parallel()
	base := mustURL(t, "https://example.com/")
	got := ExtractJS(base, []byte(`fetch("mailto:a@b.c"); fetch("data:text/plain,hi");`))
	assert.Empty(t, got.URLs)
}
