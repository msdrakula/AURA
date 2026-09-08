package analyzer

import (
	"bytes"
	"mime"
	"strings"

	"meb/internal/httpio"
)

// ContentTypeCheckRule reports empty or mismatched Content-Type values.
type ContentTypeCheckRule struct{}

// DefaultContentTypeRule returns the built-in MIME checker.
func DefaultContentTypeRule() *ContentTypeCheckRule {
	return &ContentTypeCheckRule{}
}

// Name implements Rule.
func (r *ContentTypeCheckRule) Name() string { return "content_type" }

// Check implements Rule.
func (r *ContentTypeCheckRule) Check(req *httpio.Request, resp *httpio.Response) []Finding {
	_ = req
	if resp == nil {
		return nil
	}
	if len(resp.Body) == 0 || resp.Status == 204 || resp.Status == 304 {
		return nil
	}
	raw := strings.TrimSpace(resp.Header("Content-Type"))
	if raw == "" {
		return []Finding{{
			Rule:     r.Name(),
			Severity: SeverityMedium,
			Message:  "response has a body but no Content-Type",
		}}
	}
	media, params, err := mime.ParseMediaType(raw)
	if err != nil {
		return []Finding{{
			Rule:     r.Name(),
			Severity: SeverityLow,
			Message:  "malformed Content-Type",
			Evidence: clip(raw, 80),
		}}
	}
	var out []Finding
	if isTextual(media) && params["charset"] == "" {
		out = append(out, Finding{
			Rule:     r.Name(),
			Severity: SeverityInfo,
			Message:  "textual Content-Type without charset",
			Evidence: media,
		})
	}
	kind := sniffBody(resp.Body)
	if kind != "" && !mimeMatches(media, kind) {
		out = append(out, Finding{
			Rule:     r.Name(),
			Severity: SeverityMedium,
			Message:  "Content-Type does not match body shape",
			Evidence: media + " vs " + kind,
		})
	}
	return out
}

func isTextual(media string) bool {
	media = strings.ToLower(media)
	return strings.HasPrefix(media, "text/") ||
		media == "application/json" ||
		media == "application/javascript" ||
		media == "application/xml" ||
		strings.HasSuffix(media, "+json") ||
		strings.HasSuffix(media, "+xml")
}

func sniffBody(body []byte) string {
	s := bytes.TrimSpace(body)
	if len(s) == 0 {
		return ""
	}
	if s[0] == '{' || s[0] == '[' {
		return "json"
	}
	low := bytes.ToLower(s[:min(64, len(s))])
	if bytes.Contains(low, []byte("<html")) || bytes.HasPrefix(low, []byte("<!doctype html")) {
		return "html"
	}
	if bytes.HasPrefix(low, []byte("<?xml")) {
		return "xml"
	}
	return ""
}

func mimeMatches(media, kind string) bool {
	media = strings.ToLower(media)
	switch kind {
	case "json":
		return strings.Contains(media, "json")
	case "html":
		return strings.Contains(media, "html")
	case "xml":
		return strings.Contains(media, "xml")
	default:
		return true
	}
}
