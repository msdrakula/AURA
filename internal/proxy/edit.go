package proxy

import (
	"fmt"
	"strconv"
	"strings"

	"meb/internal/httpio"
)

func errorPage(status int, message string) []byte {
	body := []byte(fmt.Sprintf(
		"<html><body style='font-family:sans-serif;background:#111;color:#eee;padding:24px'><h1>AURA %d</h1><pre>%s</pre></body></html>",
		status, message,
	))
	head := fmt.Sprintf(
		"HTTP/1.1 %d Error\r\nContent-Type: text/html; charset=utf-8\r\nContent-Length: %d\r\nConnection: close\r\n\r\n",
		status, len(body),
	)
	return append([]byte(head), body...)
}

func applyRequestRaw(req *httpio.Request, editedRaw string) {
	if req == nil {
		return
	}
	raw := []byte(editedRaw)
	head, body := httpio.SplitHeadBody(raw)
	if body == nil {
		body = []byte{}
	}
	start, headers := httpio.ParseHeaders(head)
	applyRequestEdit(req, start, headers, body)
}

func applyResponseRaw(resp *httpio.Response, editedRaw string) {
	if resp == nil {
		return
	}
	raw := []byte(editedRaw)
	head, body := httpio.SplitHeadBody(raw)
	if body == nil {
		body = []byte{}
	}
	start, headers := httpio.ParseHeaders(head)
	applyResponseEdit(resp, start, headers, body)
}

func applyRequestEdit(req *httpio.Request, start string, headers []httpio.Header, body []byte) {
	parts := strings.SplitN(start, " ", 3)
	if len(parts) >= 2 {
		req.Method = strings.ToUpper(parts[0])
		req.Path = parts[1]
		if len(parts) > 2 {
			req.Version = parts[2]
		}
	}
	req.Headers = headers
	req.Body = body
	if hostHdr := req.Header("Host"); hostHdr != "" {
		def := 80
		if req.Scheme == "https" {
			def = 443
		}
		host, port := httpio.ParseAuthority(hostHdr, def)
		req.Host = host
		req.Port = port
	}
	req.DropHeaders("transfer-encoding")
	req.SetHeader("Content-Length", strconv.Itoa(len(req.Body)))
}

func applyResponseEdit(resp *httpio.Response, start string, headers []httpio.Header, body []byte) {
	parts := strings.SplitN(start, " ", 3)
	if len(parts) >= 2 {
		resp.Version = parts[0]
		if st, err := strconv.Atoi(parts[1]); err == nil {
			resp.Status = st
		}
		if len(parts) > 2 {
			resp.Reason = parts[2]
		}
	}
	resp.Headers = headers
	resp.Body = body
	resp.DropHeaders("transfer-encoding", "content-encoding")
	resp.SetHeader("Content-Length", strconv.Itoa(len(resp.Body)))
}
