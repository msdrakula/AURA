// Package scanner runs basic active vulnerability checks against a request.
package scanner

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"

	"meb/internal/httpio"
	"meb/internal/repeater"
)

// Finding is a scanner finding.
type Finding struct {
	Name        string `json:"name"`
	Severity    string `json:"severity"`
	Param       string `json:"param"`
	Evidence    string `json:"evidence"`
	Description string `json:"description"`
}

// ScanResult is the outcome of a scan.
type ScanResult struct {
	Findings []Finding     `json:"findings"`
	Count    int           `json:"count"`
	Took     time.Duration `json:"took"`
}

// Options configure a scan.
type Options struct {
	Raw     string
	Scheme  string
	Target  string
	Verify  bool
	Timeout time.Duration
}

// Scan runs active checks against the base request. It parses parameters from
// the query string and body (urlencoded), and for each parameter injects probe
// payloads, observing the response for evidence of a vulnerability.
func Scan(ctx context.Context, opts Options) (*ScanResult, error) {
	start := time.Now()
	if opts.Timeout <= 0 {
		opts.Timeout = 15 * time.Second
	}
	// Baseline request.
	base := send(opts, opts.Raw)
	if base == nil {
		return nil, fmt.Errorf("baseline request failed")
	}
	baseBody := bodyOf(base.ResponseRaw)

	// Parse parameters from request line and body.
	_, _, query, headers, body := parseRaw(opts.Raw)
	params := collectParams(query, body)
	if len(params) == 0 {
		params = []string{""} // single insertion at end of path
	}

	var findings []Finding
	for _, p := range params {
		// SQL injection error-based
		if f := checkSQLi(ctx, opts, p, baseBody); f != nil {
			findings = append(findings, *f)
		}
		// XSS reflection
		if f := checkXSS(ctx, opts, p, baseBody); f != nil {
			findings = append(findings, *f)
		}
		// Path traversal
		if f := checkPathTraversal(ctx, opts, p, baseBody); f != nil {
			findings = append(findings, *f)
		}
		// Command injection
		if f := checkCommandInjection(ctx, opts, p, baseBody); f != nil {
			findings = append(findings, *f)
		}
		// Open redirect
		if f := checkOpenRedirect(ctx, opts, p); f != nil {
			findings = append(findings, *f)
		}
	}
	// Header-based checks
	if f := checkCORS(opts, headers, base); f != nil {
		findings = append(findings, *f)
	}
	if f := checkMissingSecurityHeaders(base.ResponseRaw); f != nil {
		findings = append(findings, *f)
	}
	if f := checkCookieFlags(base.ResponseRaw); f != nil {
		findings = append(findings, *f)
	}

	_ = headers

	return &ScanResult{
		Findings: findings,
		Count:    len(findings),
		Took:     time.Since(start),
	}, nil
}

func send(opts Options, raw string) *repeater.Result {
	res := repeater.Send(context.Background(), repeater.Options{
		Raw: raw, Scheme: opts.Scheme, Target: opts.Target,
		Verify: opts.Verify, Timeout: opts.Timeout,
	})
	if !res.OK {
		return nil
	}
	return &res
}

func bodyOf(raw string) string {
	_, b := httpio.SplitHeadBody([]byte(raw))
	return string(b)
}

// collectParams returns parameter names from query and urlencoded body.
func collectParams(query, body string) []string {
	var out []string
	seen := map[string]bool{}
	add := func(kv string) {
		if i := strings.IndexByte(kv, '='); i >= 0 {
			name := kv[:i]
			if name != "" && !seen[name] {
				seen[name] = true
				out = append(out, name)
			}
		}
	}
	for _, kv := range strings.Split(query, "&") {
		add(kv)
	}
	if ct := contentTypeOf(body); strings.Contains(ct, "application/x-www-form-urlencoded") || strings.Contains(ct, "urlencoded") {
		for _, kv := range strings.Split(body, "&") {
			add(kv)
		}
	}
	return out
}

func contentTypeOf(body string) string {
	// body is the raw body; we can't easily get content-type here, so assume urlencoded if it looks like kv pairs
	if strings.Contains(body, "=") && !strings.Contains(body, "<") {
		return "application/x-www-form-urlencoded"
	}
	return ""
}

// injectParam replaces the value of param in raw with value.
func injectParam(raw, param, value string) string {
	// Replace in query string and urlencoded body.
	// Query: name=old&...  -> name=value
	pat := regexp.MustCompile(regexp.QuoteMeta(param) + "=([^&\\s]*)")
	return pat.ReplaceAllStringFunc(raw, func(s string) string {
		return param + "=" + value
	})
}

var sqlErrors = []string{
	"sql syntax", "mysql_fetch", "ORA-", "SQLSTATE", "ODBC SQL Server Driver",
	"SQLite3::query", "PostgreSQL.*ERROR", "Microsoft SQL Native Client",
	"Unclosed quotation mark", "syntax error at line",
}

func checkSQLi(ctx context.Context, opts Options, param, baseBody string) *Finding {
	probe := "aura'||'test"
	modified := injectParam(opts.Raw, param, probe)
	if modified == opts.Raw {
		// no param matched; append to path
		modified = strings.Replace(opts.Raw, " HTTP/1.1", " HTTP/1.1", 1)
	}
	res := send(opts, modified)
	if res == nil {
		return nil
	}
	b := bodyOf(res.ResponseRaw)
	for _, e := range sqlErrors {
		if strings.Contains(strings.ToLower(b), strings.ToLower(e)) && !strings.Contains(strings.ToLower(baseBody), strings.ToLower(e)) {
			return &Finding{
				Name: "SQL injection (error-based)", Severity: "high", Param: param,
				Evidence: e, Description: "SQL error message reflected in response after injecting a quote.",
			}
		}
	}
	return nil
}

func checkXSS(ctx context.Context, opts Options, param, baseBody string) *Finding {
	probe := "auraxssprobe9x7q"
	modified := injectParam(opts.Raw, param, probe)
	res := send(opts, modified)
	if res == nil {
		return nil
	}
	b := bodyOf(res.ResponseRaw)
	if strings.Contains(b, probe) && !strings.Contains(baseBody, probe) {
		// Check if it's unescaped (no &lt;)
		if !strings.Contains(strings.SplitN(b, probe, 2)[0], "&lt;") {
			return &Finding{
				Name: "Reflected XSS", Severity: "medium", Param: param,
				Evidence: probe, Description: "Injected value reflected unencoded in response.",
			}
		}
		return &Finding{
			Name: "Reflected XSS (encoded)", Severity: "low", Param: param,
			Evidence: probe, Description: "Injected value reflected but HTML-encoded.",
		}
	}
	return nil
}

func checkPathTraversal(ctx context.Context, opts Options, param, baseBody string) *Finding {
	probe := "../../../../../../etc/passwd"
	modified := injectParam(opts.Raw, param, probe)
	res := send(opts, modified)
	if res == nil {
		return nil
	}
	b := bodyOf(res.ResponseRaw)
	if strings.Contains(baseBody, "root:") && strings.Contains(baseBody, ":/bin/") {
		return nil
	}
	if strings.Contains(b, "root:") && strings.Contains(b, ":/bin/") {
		return &Finding{
			Name: "Path traversal", Severity: "high", Param: param,
			Evidence: "root:...:/bin/", Description: "Response contains /etc/passwd content.",
		}
	}
	return nil
}

func checkCommandInjection(ctx context.Context, opts Options, param, baseBody string) *Finding {
	probe := ";echo auracmd9x7q"
	modified := injectParam(opts.Raw, param, probe)
	res := send(opts, modified)
	if res == nil {
		return nil
	}
	b := bodyOf(res.ResponseRaw)
	if strings.Contains(b, "auracmd9x7q") && !strings.Contains(baseBody, "auracmd9x7q") {
		return &Finding{
			Name: "Command injection", Severity: "high", Param: param,
			Evidence: "auracmd9x7q", Description: "Output of injected echo command reflected in response.",
		}
	}
	return nil
}

func checkOpenRedirect(ctx context.Context, opts Options, param string) *Finding {
	probe := "https://evil.example.com/"
	modified := injectParam(opts.Raw, param, probe)
	res := send(opts, modified)
	if res == nil {
		return nil
	}
	loc := ""
	for _, line := range strings.Split(res.ResponseRaw, "\n") {
		l := strings.TrimSpace(line)
		if strings.HasPrefix(strings.ToLower(l), "location:") {
			loc = strings.TrimSpace(l[len("location:"):])
			break
		}
	}
	if strings.HasPrefix(loc, "https://evil.example.com") || strings.HasPrefix(loc, "http://evil.example.com") {
		return &Finding{
			Name: "Open redirect", Severity: "medium", Param: param,
			Evidence: loc, Description: "Location header redirects to attacker-controlled URL.",
		}
	}
	return nil
}

func checkCORS(opts Options, headers string, base *repeater.Result) *Finding {
	// Check if response has Access-Control-Allow-Origin: * with ACAC: true
	h := strings.ToLower(base.ResponseRaw)
	if strings.Contains(h, "access-control-allow-origin: *") {
		// check if credentials allowed
		if strings.Contains(h, "access-control-allow-credentials: true") {
			return &Finding{
				Name: "CORS misconfiguration", Severity: "high",
				Evidence:    "ACAO: * with ACAC: true",
				Description: "CORS reflects wildcard origin with credentials, allowing cross-site data theft.",
			}
		}
		return &Finding{
			Name: "CORS wildcard", Severity: "low",
			Evidence:    "Access-Control-Allow-Origin: *",
			Description: "CORS allows any origin. Review if sensitive data is exposed.",
		}
	}
	return nil
}

var securityHeaders = []string{
	"strict-transport-security", "content-security-policy",
	"x-frame-options", "x-content-type-options",
}

func checkMissingSecurityHeaders(resp string) *Finding {
	l := strings.ToLower(resp)
	var missing []string
	for _, h := range securityHeaders {
		if !strings.Contains(l, h) {
			missing = append(missing, h)
		}
	}
	if len(missing) > 0 {
		return &Finding{
			Name: "Missing security headers", Severity: "low",
			Evidence:    strings.Join(missing, ", "),
			Description: "Response is missing security headers that harden the application.",
		}
	}
	return nil
}

func checkCookieFlags(resp string) *Finding {
	for _, line := range strings.Split(resp, "\n") {
		ll := strings.TrimSpace(line)
		if strings.HasPrefix(ll, "Set-Cookie:") {
			if !strings.Contains(strings.ToLower(ll), "httponly") || !strings.Contains(strings.ToLower(ll), "secure") {
				return &Finding{
					Name: "Cookie without security flags", Severity: "low",
					Evidence: ll, Description: "Set-Cookie missing HttpOnly and/or Secure flags.",
				}
			}
		}
	}
	return nil
}

// guard against unused http import
var _ = http.MethodGet

// parseRaw extracts method, path, query, headers and body from a raw HTTP request.
func parseRaw(raw string) (method, path, query string, headers string, body string) {
	head, b := httpio.SplitHeadBody([]byte(raw))
	body = string(b)
	start, hdrs := httpio.ParseHeaders(head)
	parts := strings.Fields(start)
	if len(parts) >= 1 {
		method = parts[0]
	}
	if len(parts) >= 2 {
		path, query = splitPQ(parts[1])
	}
	for _, h := range hdrs {
		headers += h.Key + ": " + h.Value + "\n"
	}
	return method, path, query, headers, body
}

func splitPQ(u string) (string, string) {
	if i := strings.IndexByte(u, '?'); i >= 0 {
		return u[:i], u[i+1:]
	}
	return u, ""
}
