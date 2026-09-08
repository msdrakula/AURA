package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHTTPTransactionClone(t *testing.T) {
	t.Parallel()
	orig := HTTPTransaction{
		ID: "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
		Request: HTTPRequest{
			Method:  "POST",
			URL:     "https://example.com/api",
			Headers: map[string]string{"Content-Type": "application/json"},
			Body:    []byte(`{"n":1}`),
		},
		Response: &HTTPResponse{
			Status:   200,
			Headers:  map[string]string{"Content-Type": "text/plain"},
			Body:     []byte("ok"),
			Duration: 12 * time.Millisecond,
		},
		Timestamp: time.Unix(1_700_000_000, 0).UTC(),
		Comment:   "sample",
		Tags:      []string{"auth"},
	}

	clone := orig.Clone()
	require.Equal(t, orig, clone)

	clone.Request.Headers["Content-Type"] = "text/plain"
	clone.Request.Body[0] = 'X'
	clone.Response.Headers["Content-Type"] = "application/json"
	clone.Response.Body[0] = 'X'
	clone.Tags[0] = "changed"
	clone.Comment = "other"

	assert.Equal(t, "application/json", orig.Request.Headers["Content-Type"])
	assert.Equal(t, byte('{'), orig.Request.Body[0])
	assert.Equal(t, "text/plain", orig.Response.Headers["Content-Type"])
	assert.Equal(t, byte('o'), orig.Response.Body[0])
	assert.Equal(t, []string{"auth"}, orig.Tags)
	assert.Equal(t, "sample", orig.Comment)
}

func TestHTTPTransactionCloneNilResponse(t *testing.T) {
	t.Parallel()
	orig := HTTPTransaction{ID: "id", Request: HTTPRequest{Method: "GET", URL: "http://127.0.0.1/"}}
	clone := orig.Clone()
	assert.Nil(t, clone.Response)
	clone.Request.URL = "http://example.com/"
	assert.Equal(t, "http://127.0.0.1/", orig.Request.URL)
}
