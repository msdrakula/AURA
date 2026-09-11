package intel

import (
	"net/url"
	"strings"
	"unicode"
)

var trackingParams = map[string]struct{}{
	"fbclid": {}, "gclid": {}, "gclsrc": {}, "dclid": {}, "msclkid": {}, "ttclid": {}, "yclid": {},
	"ref": {}, "referrer": {}, "referer": {}, "source": {}, "spm": {}, "trk": {},
	"mc_cid": {}, "mc_eid": {}, "_ga": {}, "_gl": {}, "_hsenc": {}, "_hsmi": {},
	"igshid": {}, "mkt_tok": {}, "oly_anon_id": {}, "oly_enc_id": {},
}

// SameScope reports whether host belongs to the investigated domain (apex or subdomain).
func SameScope(host, domain string) bool {
	host = strings.ToLower(strings.TrimSpace(host))
	domain = strings.ToLower(strings.TrimSpace(domain))
	if i := strings.IndexByte(host, ':'); i >= 0 {
		host = host[:i]
	}
	if host == "" || domain == "" {
		return false
	}
	return host == domain || strings.HasSuffix(host, "."+domain)
}

// IsTrackingParam is a marketing/analytics query name, not an application parameter.
func IsTrackingParam(name string) bool {
	n := strings.ToLower(strings.TrimSpace(name))
	if n == "" {
		return false
	}
	if _, ok := trackingParams[n]; ok {
		return true
	}
	return strings.HasPrefix(n, "utm_")
}

// IsSanePath rejects HTML/JS leftovers, control characters, and foreign URLs stuffed into a path.
func IsSanePath(p string) bool {
	p = strings.TrimSpace(p)
	if p == "" {
		return false
	}
	if p == "/" {
		return true
	}
	if !strings.HasPrefix(p, "/") {
		return false
	}
	if len(p) > 256 {
		return false
	}
	if strings.ContainsAny(p, "\n\r\t") || strings.ContainsRune(p, 0) {
		return false
	}
	if strings.Contains(p, " ") {
		return false
	}
	low := strings.ToLower(p)
	if strings.Contains(low, "http://") || strings.Contains(low, "https://") {
		return false
	}
	if strings.Contains(low, "&quot;") || strings.Contains(low, "&amp;") || strings.Contains(low, "&lt;") ||
		strings.Contains(low, "\\u003c") || strings.Contains(low, "\\u003e") {
		return false
	}
	if strings.ContainsAny(p, `<>"'`+"`") {
		return false
	}
	stripped := strings.Trim(p, "/")
	if stripped == "" {
		return true
	}
	if strings.Contains(stripped, "====") || strings.Contains(stripped, "~~~~") {
		return false
	}
	letters := 0
	other := 0
	for _, r := range stripped {
		switch {
		case unicode.IsLetter(r) || unicode.IsDigit(r):
			letters++
		case r == '-' || r == '_' || r == '.' || r == '/' || r == '~' || r == '+' || r == '=' || r == ':' || r == '@' || r == '%':
			// path punctuation
		default:
			other++
		}
	}
	if other > 0 && other >= letters {
		return false
	}
	return true
}

func artifactInScope(a Artifact, domain string) bool {
	switch a.Kind {
	case KindSubdomain, KindHost:
		return SameScope(a.Value, domain)
	case KindPort:
		h := a.Extra["host"]
		if h == "" {
			h, _, _ = strings.Cut(a.Value, ":")
		}
		return SameScope(h, domain)
	case KindPath:
		return IsSanePath(a.Value) && (a.Extra["host"] == "" || SameScope(a.Extra["host"], domain))
	case KindJS, KindURL:
		u, err := url.Parse(a.Value)
		if err != nil || u.Hostname() == "" {
			return false
		}
		if !SameScope(u.Hostname(), domain) {
			return false
		}
		p := u.Path
		if p == "" {
			p = "/"
		}
		return IsSanePath(p)
	case KindTech:
		h := a.Extra["host"]
		return h == "" || SameScope(h, domain)
	case KindParam:
		return a.Value != "" && !IsTrackingParam(a.Value)
	default:
		return true
	}
}
