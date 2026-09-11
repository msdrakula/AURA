package intel

import (
	"encoding/json"
	"net"
	"net/url"
	"strconv"
	"strings"
)

// NormalizeInput accepts a domain or URL and returns hostname (no port) plus a base URL.
// A non-default port stays in the base URL. Missing ports use 80 for http and 443 for https
// and are omitted from the printed URL.
func NormalizeInput(raw string) (domain, baseURL string, err error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return "", "", errEmpty
	}
	if !strings.Contains(raw, "://") {
		raw = inferScheme(raw) + "://" + raw
	}
	u, perr := url.Parse(raw)
	if perr != nil || u.Hostname() == "" {
		return "", "", errBadInput
	}
	scheme := strings.ToLower(u.Scheme)
	if scheme != "http" && scheme != "https" {
		return "", "", errBadInput
	}
	host := strings.ToLower(u.Hostname())
	port := u.Port()
	path := u.Path
	if path == "" {
		path = "/"
	}
	return host, joinBase(scheme, host, port, path, u.RawQuery), nil
}

func inferScheme(raw string) string {
	head := raw
	if i := strings.Index(raw, "/"); i >= 0 {
		head = raw[:i]
	}
	if strings.HasPrefix(head, "[") {
		return "https"
	}
	_, port, err := net.SplitHostPort(head)
	if err == nil && (port == "80" || port == "8080" || port == "8000" || port == "8008") {
		return "http"
	}
	return "https"
}

func defaultPort(scheme string) string {
	if strings.ToLower(scheme) == "http" {
		return "80"
	}
	return "443"
}

func joinBase(scheme, host, port, path, query string) string {
	authority := host
	if port != "" && port != defaultPort(scheme) {
		authority = net.JoinHostPort(host, port)
	}
	if path == "" {
		path = "/"
	}
	out := scheme + "://" + authority + path
	if query != "" {
		out += "?" + query
	}
	return out
}

// OriginURL is scheme://host with an explicit non-default port.
func OriginURL(raw string) string {
	u, err := url.Parse(strings.TrimSpace(raw))
	if err != nil || u.Hostname() == "" {
		return strings.TrimSpace(raw)
	}
	scheme := u.Scheme
	if scheme == "" {
		scheme = "https"
	}
	return joinBase(scheme, strings.ToLower(u.Hostname()), u.Port(), "/", "")
}

// extraPortFromBase returns a non-default port from a target URL (e.g. 16126).
func extraPortFromBase(raw string) string {
	u, err := url.Parse(strings.TrimSpace(raw))
	if err != nil || u == nil {
		return ""
	}
	p := u.Port()
	if p == "" || p == defaultPort(u.Scheme) {
		return ""
	}
	return p
}

func defaultWebPorts() []string {
	return []string{
		"80", "443", "3000", "3001", "4000", "4443", "5000", "5001",
		"7001", "8000", "8008", "8080", "8081", "8443", "8888",
		"9000", "9090", "9443", "10443",
	}
}

func parsePortList(raw string) []string {
	raw = strings.ReplaceAll(raw, ",", " ")
	raw = strings.ReplaceAll(raw, ";", " ")
	raw = strings.ReplaceAll(raw, "\n", " ")
	var out []string
	seen := map[string]struct{}{}
	for _, p := range strings.Fields(raw) {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		n, err := strconv.Atoi(p)
		if err != nil || n < 1 || n > 65535 {
			continue
		}
		s := strconv.Itoa(n)
		if _, ok := seen[s]; ok {
			continue
		}
		seen[s] = struct{}{}
		out = append(out, s)
	}
	return out
}

func webPortList(base string, custom ...string) []string {
	ports := defaultWebPorts()
	if len(custom) > 0 && strings.TrimSpace(custom[0]) != "" {
		if parsed := parsePortList(custom[0]); len(parsed) > 0 {
			ports = parsed
		}
	}
	extra := extraPortFromBase(base)
	if extra == "" {
		return ports
	}
	out := []string{extra}
	for _, p := range ports {
		if p != extra {
			out = append(out, p)
		}
	}
	return out
}

func webProbeURLs(host, port, base string) []string {
	host = strings.ToLower(strings.TrimSpace(host))
	var out []string
	add := func(raw string) {
		raw = strings.TrimSpace(raw)
		if raw == "" {
			return
		}
		for _, x := range out {
			if x == raw {
				return
			}
		}
		out = append(out, raw)
	}
	if u, err := url.Parse(strings.TrimSpace(base)); err == nil && u != nil {
		if strings.EqualFold(u.Hostname(), host) {
			bp := u.Port()
			if bp == "" {
				bp = defaultPort(u.Scheme)
			}
			if bp == port {
				if origin := OriginURL(base); origin != "" {
					add(strings.TrimRight(origin, "/") + "/")
				}
			}
		}
	}
	httpsURL := "https://" + net.JoinHostPort(host, port) + "/"
	httpURL := "http://" + net.JoinHostPort(host, port) + "/"
	switch port {
	case "443", "8443":
		add(httpsURL)
	case "80", "8080", "8000", "8008":
		add(httpURL)
		add(httpsURL)
	default:
		add(httpsURL)
		add(httpURL)
	}
	return out
}

type simpleError string

func (e simpleError) Error() string { return string(e) }

const (
	errEmpty    simpleError = "укажите домен или URL"
	errBadInput simpleError = "не удалось разобрать адрес"
)

type crtRow struct {
	NameValue string `json:"name_value"`
	Common    string `json:"common_name"`
}

func parseCRTNames(body []byte, domain string) []string {
	var rows []crtRow
	if err := json.Unmarshal(body, &rows); err != nil {
		return nil
	}
	seen := map[string]struct{}{}
	var out []string
	add := func(name string) {
		name = strings.ToLower(strings.TrimSpace(name))
		name = strings.TrimPrefix(name, "*.")
		if name == "" || strings.Contains(name, " ") {
			return
		}
		if name != domain && !strings.HasSuffix(name, "."+domain) {
			return
		}
		if _, ok := seen[name]; ok {
			return
		}
		seen[name] = struct{}{}
		out = append(out, name)
	}
	for _, r := range rows {
		for _, part := range strings.Split(r.NameValue, "\n") {
			add(part)
		}
		add(r.Common)
	}
	return out
}

func parseWayback(body []byte) []string {
	var rows [][]string
	if err := json.Unmarshal(body, &rows); err != nil {
		return nil
	}
	seen := map[string]struct{}{}
	var out []string
	for i, row := range rows {
		if i == 0 && len(row) > 0 && strings.EqualFold(row[0], "original") {
			continue
		}
		if len(row) == 0 {
			continue
		}
		raw := strings.TrimSpace(row[0])
		if raw == "" {
			continue
		}
		if _, ok := seen[raw]; ok {
			continue
		}
		seen[raw] = struct{}{}
		out = append(out, raw)
	}
	return out
}
