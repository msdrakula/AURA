// Package models holds shared structs for modules, persistence, and the UI.
package models

import (
	"maps"
	"slices"
	"time"
)

// HTTPRequest is the client-facing view of a captured or edited request.
type HTTPRequest struct {
	Method  string            `json:"method" gorm:"column:method" db:"method"`
	URL     string            `json:"url" gorm:"column:url" db:"url"`
	Headers map[string]string `json:"headers" gorm:"serializer:json" db:"headers"`
	Body    []byte            `json:"body" gorm:"column:body;type:blob" db:"body"`
}

// HTTPResponse is the client-facing view of an upstream response.
type HTTPResponse struct {
	Status   int               `json:"status" gorm:"column:status" db:"status"`
	Headers  map[string]string `json:"headers" gorm:"serializer:json" db:"headers"`
	Body     []byte            `json:"body" gorm:"column:body;type:blob" db:"body"`
	Duration time.Duration     `json:"duration" gorm:"column:duration_ns" db:"duration_ns"`
}

// HTTPTransaction is one request/response pair in history or the editor.
type HTTPTransaction struct {
	ID        string        `json:"id" gorm:"column:id;primaryKey" db:"id"`
	Request   HTTPRequest   `json:"request" gorm:"embedded;embeddedPrefix:req_"`
	Response  *HTTPResponse `json:"response,omitempty" gorm:"embedded;embeddedPrefix:resp_"`
	Timestamp time.Time     `json:"timestamp" gorm:"column:timestamp" db:"timestamp"`
	Comment   string        `json:"comment" gorm:"column:comment" db:"comment"`
	Tags      []string      `json:"tags" gorm:"serializer:json" db:"tags"`
}

// Clone returns a deep copy safe to hand to another goroutine or the UI.
func (t HTTPTransaction) Clone() HTTPTransaction {
	out := t
	out.Request.Headers = cloneMap(t.Request.Headers)
	out.Request.Body = slices.Clone(t.Request.Body)
	out.Tags = slices.Clone(t.Tags)
	if t.Response != nil {
		resp := *t.Response
		resp.Headers = cloneMap(t.Response.Headers)
		resp.Body = slices.Clone(t.Response.Body)
		out.Response = &resp
	}
	return out
}

func cloneMap(in map[string]string) map[string]string {
	if in == nil {
		return nil
	}
	return maps.Clone(in)
}
