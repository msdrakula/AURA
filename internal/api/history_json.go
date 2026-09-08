package api

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"meb/internal/httpio"
	"meb/internal/models"
	"meb/internal/store"
)

func flowDetail(flow *store.Flow) map[string]any {
	req := flow.Request
	resp := flow.Response
	out := flow.Summary()
	out["client"] = flow.Client
	out["request_raw"] = httpio.DecodeRaw(req.Raw())
	if resp != nil {
		out["response_raw"] = httpio.DecodeRaw(resp.Raw())
		out["response_headers"] = headerPairsFromList(resp.Headers)
	} else {
		out["response_raw"] = ""
		out["response_headers"] = []any{}
	}
	out["request_headers"] = headerPairsFromList(req.Headers)
	return out
}

func headerPairsFromList(headers []httpio.Header) [][2]string {
	out := make([][2]string, 0, len(headers))
	for _, h := range headers {
		out = append(out, [2]string{h.Key, h.Value})
	}
	return out
}

func transactionSummary(tx models.HTTPTransaction) map[string]any {
	host, path, scheme := splitURL(tx.Request.URL)
	var status any
	var mime any
	respBytes := 0
	var duration any
	if tx.Response != nil {
		status = tx.Response.Status
		respBytes = len(tx.Response.Body)
		mime = tx.Response.Headers["Content-Type"]
		if mime == "" {
			mime = tx.Response.Headers["content-type"]
		}
		duration = int(tx.Response.Duration / time.Millisecond)
	}
	// Source: derive from tags/comment (tool:repeater etc.) or default to "proxy".
	source := "proxy"
	for _, t := range tx.Tags {
		if strings.HasPrefix(t, "tool:") {
			source = strings.TrimPrefix(t, "tool:")
			break
		}
	}
	if strings.HasPrefix(tx.Comment, "tool:") {
		source = strings.TrimPrefix(tx.Comment, "tool:")
	}
	return map[string]any{
		"id":          tx.ID,
		"created":     float64(tx.Timestamp.UnixNano()) / 1e9,
		"method":      tx.Request.Method,
		"url":         tx.Request.URL,
		"host":        host,
		"path":        path,
		"scheme":      scheme,
		"status":      status,
		"reason":      nil,
		"req_bytes":   len(tx.Request.Body),
		"resp_bytes":  respBytes,
		"duration_ms": duration,
		"error":       nil,
		"mime":        mime,
		"source":      source,
	}
}

func transactionDetail(tx models.HTTPTransaction) map[string]any {
	out := transactionSummary(tx)
	out["client"] = ""
	out["request_raw"] = rawHTTP(tx.Request.Method, tx.Request.URL, tx.Request.Headers, tx.Request.Body, true)
	out["request_headers"] = headerPairs(tx.Request.Headers)
	if tx.Response != nil {
		out["response_raw"] = rawHTTP(fmt.Sprintf("HTTP/1.1 %d", tx.Response.Status), "", tx.Response.Headers, tx.Response.Body, false)
		out["response_headers"] = headerPairs(tx.Response.Headers)
	} else {
		out["response_raw"] = ""
		out["response_headers"] = []any{}
	}
	return out
}

func splitURL(raw string) (host, path, scheme string) {
	u, err := url.Parse(raw)
	if err != nil || u.Host == "" {
		return "", raw, "http"
	}
	path = u.RequestURI()
	if path == "" {
		path = "/"
	}
	return u.Host, path, u.Scheme
}

func headerPairs(h map[string]string) [][2]string {
	out := make([][2]string, 0, len(h))
	for k, v := range h {
		out = append(out, [2]string{k, v})
	}
	return out
}

func rawHTTP(start, rawURL string, headers map[string]string, body []byte, isReq bool) string {
	var b strings.Builder
	line := start
	if isReq {
		_, path, _ := splitURL(rawURL)
		line = fmt.Sprintf("%s %s HTTP/1.1", start, path)
	} else if !strings.Contains(start, " ") {
		line = start + " "
	}
	b.WriteString(line)
	if !strings.HasSuffix(line, "\r\n") {
		b.WriteString("\r\n")
	}
	for k, v := range headers {
		b.WriteString(k)
		b.WriteString(": ")
		b.WriteString(v)
		b.WriteString("\r\n")
	}
	b.WriteString("\r\n")
	b.Write(body)
	return b.String()
}
