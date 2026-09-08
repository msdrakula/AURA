package models

import "time"

// BatchResult represents the outcome of a single batch request.
type BatchResult struct {
	Payloads   []string      `json:"payloads"`
	StatusCode int           `json:"status_code"`
	BodyLength int           `json:"body_length"`
	Duration   time.Duration `json:"duration"`
	Error      string        `json:"error,omitempty"`
}
