// Package extractor collects URLs from HTML documents and JavaScript sources.
package extractor

import (
	"bytes"
	"net/url"
	"sort"
	"strings"

	"golang.org/x/net/html"
)

// Result is a de-duplicated, sorted list of absolute http(s) URLs.
type Result struct {
	URLs []string `json:"urls"`
}

// ExtractHTML walks a document for href/src/action and scans inline scripts.
// External script files are not fetched; only the src attribute is recorded.
func ExtractHTML(base *url.URL, doc []byte) (Result, error) {
	root, err := html.Parse(bytes.NewReader(doc))
	if err != nil {
		return Result{}, err
	}
	var raw []string
	var walk func(*html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, key := range []string{"href", "src", "action"} {
				if v := attr(n, key); v != "" {
					raw = append(raw, v)
				}
			}
			if n.Data == "script" && attr(n, "src") == "" {
				raw = append(raw, findJSURLs(textOf(n))...)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(root)
	return Result{URLs: normalize(base, raw)}, nil
}

// ExtractJS finds URL-like strings in a JavaScript source buffer.
func ExtractJS(base *url.URL, src []byte) Result {
	return Result{URLs: normalize(base, findJSURLs(string(src)))}
}

func attr(n *html.Node, name string) string {
	for _, a := range n.Attr {
		if strings.EqualFold(a.Key, name) {
			return strings.TrimSpace(html.UnescapeString(a.Val))
		}
	}
	return ""
}

func textOf(n *html.Node) string {
	var b strings.Builder
	var walk func(*html.Node)
	walk = func(c *html.Node) {
		if c.Type == html.TextNode {
			b.WriteString(c.Data)
		}
		for k := c.FirstChild; k != nil; k = k.NextSibling {
			walk(k)
		}
	}
	walk(n)
	return b.String()
}

func normalize(base *url.URL, refs []string) []string {
	seen := make(map[string]struct{}, len(refs))
	out := make([]string, 0, len(refs))
	for _, ref := range refs {
		abs, ok := resolve(base, ref)
		if !ok {
			continue
		}
		if _, dup := seen[abs]; dup {
			continue
		}
		seen[abs] = struct{}{}
		out = append(out, abs)
	}
	sort.Strings(out)
	return out
}

func resolve(base *url.URL, ref string) (string, bool) {
	ref = strings.TrimSpace(html.UnescapeString(ref))
	if ref == "" || strings.HasPrefix(ref, "#") {
		return "", false
	}
	lower := strings.ToLower(ref)
	for _, scheme := range []string{"javascript:", "data:", "mailto:", "tel:", "blob:", "about:"} {
		if strings.HasPrefix(lower, scheme) {
			return "", false
		}
	}
	u, err := url.Parse(ref)
	if err != nil || u == nil {
		return "", false
	}
	if base != nil {
		u = base.ResolveReference(u)
	}
	u.Fragment = ""
	if u.Scheme != "http" && u.Scheme != "https" {
		return "", false
	}
	if u.Host == "" {
		return "", false
	}
	return u.String(), true
}
