package intel

import (
	"encoding/json"
	"net/url"
	"strings"
)

// NormalizeInput accepts a domain or URL and returns apex-ish domain + base URL.
func NormalizeInput(raw string) (domain, baseURL string, err error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return "", "", errEmpty
	}
	if !strings.Contains(raw, "://") {
		if strings.Contains(raw, "/") {
			raw = "https://" + raw
		} else {
			return strings.ToLower(strings.TrimSuffix(raw, ".")), "https://" + strings.ToLower(raw), nil
		}
	}
	u, perr := url.Parse(raw)
	if perr != nil || u.Hostname() == "" {
		return "", "", errBadInput
	}
	host := strings.ToLower(u.Hostname())
	u.Fragment = ""
	if u.Path == "" {
		u.Path = "/"
	}
	return host, u.String(), nil
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
	if len(out) > 200 {
		out = out[:200]
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
