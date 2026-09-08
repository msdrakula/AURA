// Package analyzer inspects captured HTTP exchanges against registered rules.
package analyzer

import "meb/internal/httpio"

// Severity is a coarse finding level for the UI.
type Severity string

const (
	SeverityInfo   Severity = "info"
	SeverityLow    Severity = "low"
	SeverityMedium Severity = "medium"
	SeverityHigh   Severity = "high"
)

// Finding is a single rule hit on a request/response pair.
type Finding struct {
	Rule     string   `json:"rule"`
	Severity Severity `json:"severity"`
	Message  string   `json:"message"`
	Evidence string   `json:"evidence,omitempty"`
}

// Rule checks one captured exchange.
type Rule interface {
	// Name returns a stable rule id used in Finding.Rule.
	Name() string
	// Check inspects the pair and returns zero or more findings.
	Check(req *httpio.Request, resp *httpio.Response) []Finding
}

// Engine runs a list of rules over a single exchange.
type Engine struct {
	rules []Rule
}

// NewEngine returns an analyzer with the given rules, in order.
func NewEngine(rules ...Rule) *Engine {
	out := make([]Rule, len(rules))
	copy(out, rules)
	return &Engine{rules: out}
}

// DefaultEngine returns an engine with the built-in header, MIME, and regex rules.
func DefaultEngine() *Engine {
	return NewEngine(DefaultHeaderRule(), DefaultContentTypeRule(), DefaultRegexRule())
}

// Register appends a rule. It is not safe to call concurrently with Analyze.
func (e *Engine) Register(rule Rule) {
	if rule == nil {
		return
	}
	e.rules = append(e.rules, rule)
}

// Analyze runs every registered rule. A nil response yields no findings.
func (e *Engine) Analyze(req *httpio.Request, resp *httpio.Response) []Finding {
	if e == nil || resp == nil {
		return nil
	}
	var out []Finding
	for _, rule := range e.rules {
		if rule == nil {
			continue
		}
		out = append(out, rule.Check(req, resp)...)
	}
	if out == nil {
		return []Finding{}
	}
	return out
}
