package analyzer

import (
	"fmt"
	"strings"

	"meb/internal/httpio"
)

// ExpectedHeader describes a response header the rule looks for.
type ExpectedHeader struct {
	Name       string
	Contains   string
	HTTPSOnly  bool
	Severity   Severity
	MissingMsg string
}

// HeaderCheckRule reports missing or weak security-related response headers.
type HeaderCheckRule struct {
	Want []ExpectedHeader
}

// DefaultHeaderRule checks CSP, HSTS, X-Content-Type-Options, and Referrer-Policy.
func DefaultHeaderRule() *HeaderCheckRule {
	return &HeaderCheckRule{
		Want: []ExpectedHeader{
			{
				Name:       "Strict-Transport-Security",
				Contains:   "max-age=",
				HTTPSOnly:  true,
				Severity:   SeverityMedium,
				MissingMsg: "missing or weak Strict-Transport-Security",
			},
			{
				Name:       "Content-Security-Policy",
				Severity:   SeverityLow,
				MissingMsg: "missing Content-Security-Policy",
			},
			{
				Name:       "X-Content-Type-Options",
				Contains:   "nosniff",
				Severity:   SeverityLow,
				MissingMsg: "missing X-Content-Type-Options: nosniff",
			},
			{
				Name:       "Referrer-Policy",
				Severity:   SeverityInfo,
				MissingMsg: "missing Referrer-Policy",
			},
		},
	}
}

// Name implements Rule.
func (r *HeaderCheckRule) Name() string { return "header_check" }

// Check implements Rule.
func (r *HeaderCheckRule) Check(req *httpio.Request, resp *httpio.Response) []Finding {
	if r == nil || resp == nil {
		return nil
	}
	https := req != nil && strings.EqualFold(req.Scheme, "https")
	var out []Finding
	for _, want := range r.Want {
		if want.HTTPSOnly && !https {
			continue
		}
		raw := resp.Header(want.Name)
		sev := want.Severity
		if sev == "" {
			sev = SeverityLow
		}
		if raw == "" {
			out = append(out, Finding{
				Rule:     r.Name(),
				Severity: sev,
				Message:  want.MissingMsg,
				Evidence: want.Name,
			})
			continue
		}
		if want.Contains != "" && !strings.Contains(strings.ToLower(raw), strings.ToLower(want.Contains)) {
			out = append(out, Finding{
				Rule:     r.Name(),
				Severity: sev,
				Message:  fmt.Sprintf("%s present but does not contain %q", want.Name, want.Contains),
				Evidence: clip(raw, 120),
			})
		}
	}
	return out
}
