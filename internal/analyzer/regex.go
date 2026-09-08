package analyzer

import (
	"regexp"
	"unicode/utf8"

	"meb/internal/httpio"
)

// Pattern is a named regular expression applied to the response body.
type Pattern struct {
	ID       string
	Message  string
	Severity Severity
	RE       *regexp.Regexp
}

// RegexMatchRule reports body substrings that match configured patterns.
type RegexMatchRule struct {
	Patterns []Pattern
}

// DefaultRegexRule looks for stack traces and token-shaped strings in bodies.
func DefaultRegexRule() *RegexMatchRule {
	return &RegexMatchRule{
		Patterns: []Pattern{
			{
				ID:       "stack_python",
				Message:  "Python traceback in response body",
				Severity: SeverityMedium,
				RE:       regexp.MustCompile(`Traceback \(most recent call last\)`),
			},
			{
				ID:       "stack_java",
				Message:  "Java stack frame in response body",
				Severity: SeverityMedium,
				RE:       regexp.MustCompile(`\bat [\w$.]+\([^)]+\.java:\d+\)`),
			},
			{
				ID:       "stack_js",
				Message:  "JavaScript stack in response body",
				Severity: SeverityLow,
				RE:       regexp.MustCompile(`(?i)at\s+\S+\s+\([^)]+:\d+:\d+\)`),
			},
			{
				ID:       "jwt_like",
				Message:  "JWT-shaped string in response body",
				Severity: SeverityMedium,
				RE:       regexp.MustCompile(`\beyJ[A-Za-z0-9_-]{10,}\.[A-Za-z0-9_-]{10,}\.[A-Za-z0-9_-]{10,}\b`),
			},
			{
				ID:       "pem_key",
				Message:  "PEM private key block in response body",
				Severity: SeverityHigh,
				RE:       regexp.MustCompile(`-----BEGIN (?:RSA |EC |OPENSSH )?PRIVATE KEY-----`),
			},
		},
	}
}

// Name implements Rule.
func (r *RegexMatchRule) Name() string { return "regex_match" }

// Check implements Rule.
func (r *RegexMatchRule) Check(req *httpio.Request, resp *httpio.Response) []Finding {
	_ = req
	if r == nil || resp == nil || len(resp.Body) == 0 {
		return nil
	}
	if !utf8.Valid(resp.Body) {
		return nil
	}
	body := string(resp.Body)
	var out []Finding
	seen := map[string]struct{}{}
	for _, p := range r.Patterns {
		if p.RE == nil {
			continue
		}
		loc := p.RE.FindStringIndex(body)
		if loc == nil {
			continue
		}
		if _, dup := seen[p.ID]; dup {
			continue
		}
		seen[p.ID] = struct{}{}
		sev := p.Severity
		if sev == "" {
			sev = SeverityLow
		}
		out = append(out, Finding{
			Rule:     r.Name(),
			Severity: sev,
			Message:  p.Message,
			Evidence: clip(body[loc[0]:loc[1]], 80),
		})
	}
	return out
}

func clip(s string, n int) string {
	if n <= 0 || len(s) <= n {
		return s
	}
	return s[:n] + "…"
}
