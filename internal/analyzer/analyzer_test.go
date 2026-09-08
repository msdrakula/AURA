package analyzer

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"meb/internal/httpio"
)

func reqHTTPS() *httpio.Request {
	return &httpio.Request{Method: "GET", Path: "/", Version: "HTTP/1.1", Scheme: "https", Host: "example.com", Port: 443}
}

func respWith(status int, headers []httpio.Header, body string) *httpio.Response {
	return &httpio.Response{
		Version: "HTTP/1.1",
		Status:  status,
		Reason:  "OK",
		Headers: headers,
		Body:    []byte(body),
	}
}

func TestHeaderCheckRule(t *testing.T) {
	t.Parallel()
	rule := DefaultHeaderRule()
	tests := []struct {
		name    string
		req     *httpio.Request
		resp    *httpio.Response
		wantMsg []string
		noMsg   []string
	}{
		{
			name: "https missing security headers",
			req:  reqHTTPS(),
			resp: respWith(200, []httpio.Header{{Key: "Content-Type", Value: "text/plain"}}, "ok"),
			wantMsg: []string{
				"missing or weak Strict-Transport-Security",
				"missing Content-Security-Policy",
				"missing X-Content-Type-Options: nosniff",
			},
		},
		{
			name: "hsts present but weak",
			req:  reqHTTPS(),
			resp: respWith(200, []httpio.Header{
				{Key: "Strict-Transport-Security", Value: "includeSubDomains"},
				{Key: "Content-Security-Policy", Value: "default-src 'self'"},
				{Key: "X-Content-Type-Options", Value: "nosniff"},
				{Key: "Referrer-Policy", Value: "no-referrer"},
			}, "ok"),
			wantMsg: []string{`Strict-Transport-Security present but does not contain "max-age="`},
		},
		{
			name: "http skips hsts",
			req:  &httpio.Request{Method: "GET", Path: "/", Scheme: "http", Host: "example.com", Port: 80},
			resp: respWith(200, []httpio.Header{
				{Key: "Content-Security-Policy", Value: "default-src 'self'"},
				{Key: "X-Content-Type-Options", Value: "nosniff"},
				{Key: "Referrer-Policy", Value: "no-referrer"},
			}, "ok"),
			noMsg: []string{"Strict-Transport-Security"},
		},
		{
			name: "all headers ok",
			req:  reqHTTPS(),
			resp: respWith(200, []httpio.Header{
				{Key: "Strict-Transport-Security", Value: "max-age=63072000"},
				{Key: "Content-Security-Policy", Value: "default-src 'self'"},
				{Key: "X-Content-Type-Options", Value: "nosniff"},
				{Key: "Referrer-Policy", Value: "no-referrer"},
			}, "ok"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := rule.Check(tt.req, tt.resp)
			joined := messages(got)
			for _, msg := range tt.wantMsg {
				assert.Contains(t, joined, msg)
			}
			for _, msg := range tt.noMsg {
				for _, f := range got {
					assert.NotContains(t, f.Message, msg)
					assert.NotEqual(t, msg, f.Evidence)
				}
			}
			if len(tt.wantMsg) == 0 && len(tt.noMsg) == 0 {
				assert.Empty(t, got)
			}
		})
	}
}

func TestContentTypeCheckRule(t *testing.T) {
	t.Parallel()
	rule := DefaultContentTypeRule()
	tests := []struct {
		name      string
		resp      *httpio.Response
		want      string
		wantEmpty bool
	}{
		{
			name: "missing content-type",
			resp: respWith(200, nil, `{"ok":true}`),
			want: "no Content-Type",
		},
		{
			name: "json body as html",
			resp: respWith(200, []httpio.Header{{Key: "Content-Type", Value: "text/html; charset=utf-8"}}, `{"ok":true}`),
			want: "does not match body shape",
		},
		{
			name:      "json ok",
			resp:      respWith(200, []httpio.Header{{Key: "Content-Type", Value: "application/json; charset=utf-8"}}, `{"ok":true}`),
			wantEmpty: true,
		},
		{
			name:      "empty body skipped",
			resp:      respWith(204, nil, ""),
			wantEmpty: true,
		},
		{
			name: "text without charset",
			resp: respWith(200, []httpio.Header{{Key: "Content-Type", Value: "text/plain"}}, "hello"),
			want: "without charset",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := rule.Check(reqHTTPS(), tt.resp)
			if tt.wantEmpty {
				assert.Empty(t, got)
				return
			}
			require.NotEmpty(t, got)
			assert.Contains(t, messages(got), tt.want)
		})
	}
}

func TestRegexMatchRule(t *testing.T) {
	t.Parallel()
	rule := DefaultRegexRule()
	tests := []struct {
		name string
		body string
		want string
	}{
		{
			name: "python traceback",
			body: "Traceback (most recent call last):\n  File \"app.py\", line 1",
			want: "Python traceback",
		},
		{
			name: "java frame",
			body: "Exception\n        at com.example.App.main(App.java:42)\n",
			want: "Java stack frame",
		},
		{
			name: "jwt-like",
			body: `{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIn0.signaturehereok"}`,
			want: "JWT-shaped",
		},
		{
			name: "pem key",
			body: "oops\n-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAK\n",
			want: "PEM private key",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := rule.Check(reqHTTPS(), respWith(500, []httpio.Header{{Key: "Content-Type", Value: "text/plain"}}, tt.body))
			require.NotEmpty(t, got)
			assert.Contains(t, messages(got), tt.want)
		})
	}
	t.Run("clean body", func(t *testing.T) {
		t.Parallel()
		got := rule.Check(reqHTTPS(), respWith(200, []httpio.Header{{Key: "Content-Type", Value: "text/plain"}}, "hello world"))
		assert.Empty(t, got)
	})
}

func TestEngineAnalyze(t *testing.T) {
	t.Parallel()
	eng := DefaultEngine()
	got := eng.Analyze(reqHTTPS(), respWith(200, nil, `{"ok":true}`))
	require.NotEmpty(t, got)
	rules := map[string]int{}
	for _, f := range got {
		rules[f.Rule]++
	}
	assert.Greater(t, rules["header_check"], 0)
	assert.Greater(t, rules["content_type"], 0)
	assert.Empty(t, eng.Analyze(reqHTTPS(), nil))
}

func messages(fs []Finding) string {
	var b []byte
	for _, f := range fs {
		b = append(b, f.Message...)
		b = append(b, '\n')
	}
	return string(b)
}
