package intel

import (
	"bytes"
	"net/http"
	"regexp"
	"strings"
)

type fingerprint struct {
	Name string
	Test func(h http.Header, body []byte) bool
}

var fingerprints = []fingerprint{
	{Name: "nginx", Test: headerHas("Server", "nginx")},
	{Name: "Apache", Test: headerHas("Server", "apache")},
	{Name: "IIS", Test: headerHas("Server", "microsoft-iis")},
	{Name: "cloudflare", Test: headerOr("Server", "cloudflare", "CF-Ray", "")},
	{Name: "Fastly", Test: headerHas("X-Served-By", "cache-")},
	{Name: "Vercel", Test: headerHas("Server", "vercel")},
	{Name: "Akamai", Test: headerHas("X-Akamai-Transformed", "")},
	{Name: "AWS", Test: headerHas("Server", "awselb")},
	{Name: "Express", Test: headerHas("X-Powered-By", "express")},
	{Name: "PHP", Test: headerHas("X-Powered-By", "php")},
	{Name: "ASP.NET", Test: headerHas("X-Powered-By", "asp.net")},
	{Name: "Django", Test: cookieHas("csrftoken")},
	{Name: "Laravel", Test: cookieHas("laravel_session")},
	{Name: "Clerk", Test: bodyHas("clerk.", "@clerk", "__clerk")},
	{Name: "WordPress", Test: bodyHas("wp-content", "wp-includes")},
	{Name: "Drupal", Test: headerHas("X-Generator", "drupal")},
	{Name: "jQuery", Test: bodyHas("jquery")},
	{Name: "React", Test: bodyHas("data-reactroot", "__NEXT_DATA__")},
	{Name: "Next.js", Test: bodyHas("__NEXT_DATA__")},
	{Name: "Vue", Test: bodyHas("data-v-", "__vue__")},
	{Name: "Angular", Test: bodyHas("ng-version")},
	{Name: "GraphQL", Test: bodyHas("graphql")},
	{Name: "Swagger", Test: bodyHas("swagger-ui", "swagger.json")},
	{Name: "OpenResty", Test: headerHas("Server", "openresty")},
}

func headerHas(name, needle string) func(http.Header, []byte) bool {
	return func(h http.Header, _ []byte) bool {
		v := h.Get(name)
		if needle == "" {
			return strings.TrimSpace(v) != ""
		}
		return strings.Contains(strings.ToLower(v), needle)
	}
}

func headerOr(pairs ...string) func(http.Header, []byte) bool {
	return func(h http.Header, body []byte) bool {
		for i := 0; i+1 < len(pairs); i += 2 {
			if headerHas(pairs[i], pairs[i+1])(h, body) {
				return true
			}
		}
		return false
	}
}

func cookieHas(name string) func(http.Header, []byte) bool {
	lname := strings.ToLower(name)
	return func(h http.Header, _ []byte) bool {
		for _, c := range h.Values("Set-Cookie") {
			if strings.HasPrefix(strings.ToLower(c), lname+"=") {
				return true
			}
		}
		return false
	}
}

func bodyHas(needles ...string) func(http.Header, []byte) bool {
	low := make([][]byte, len(needles))
	for i, n := range needles {
		low[i] = bytes.ToLower([]byte(n))
	}
	return func(_ http.Header, body []byte) bool {
		b := bytes.ToLower(body)
		for _, n := range low {
			if bytes.Contains(b, n) {
				return true
			}
		}
		return false
	}
}

var genRE = regexp.MustCompile(`(?i)<meta[^>]+name=["']generator["'][^>]+content=["']([^"']+)`)

// DetectTech returns technology names from headers and HTML.
func DetectTech(h http.Header, body []byte) []string {
	var out []string
	seen := map[string]struct{}{}
	add := func(n string) {
		n = strings.TrimSpace(n)
		if n == "" {
			return
		}
		if _, ok := seen[n]; ok {
			return
		}
		seen[n] = struct{}{}
		out = append(out, n)
	}
	if srv := strings.TrimSpace(h.Get("Server")); srv != "" {
		add(srv)
	}
	if m := genRE.FindSubmatch(body); len(m) > 1 {
		add(string(m[1]))
	}
	for _, fp := range fingerprints {
		if fp.Test(h, body) {
			add(fp.Name)
		}
	}
	return out
}
