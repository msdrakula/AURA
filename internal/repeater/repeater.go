package repeater

import (
	"bufio"
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"meb/internal/codec"
	"meb/internal/httpio"
)

const defaultTimeout = 30 * time.Second

// Options describes a single replayed HTTP/1.1 request.
type Options struct {
	Raw     string
	Scheme  string
	Target  string
	Verify  bool
	Timeout time.Duration
}

// Result is the JSON-facing outcome of a replayed request.
type Result struct {
	OK          bool   `json:"ok"`
	Error       string `json:"error,omitempty"`
	Status      int    `json:"status,omitempty"`
	Reason      string `json:"reason,omitempty"`
	DurationMS  int    `json:"duration_ms"`
	RequestRaw  string `json:"request_raw,omitempty"`
	ResponseRaw string `json:"response_raw,omitempty"`
}

// ParseRawRequest turns an editor buffer into an origin-form HTTP/1.1 request.
func ParseRawRequest(rawText, scheme, target string) (*httpio.Request, error) {
	if scheme == "" {
		scheme = "https"
	}
	if scheme != "http" && scheme != "https" {
		return nil, fmt.Errorf("scheme must be http or https")
	}
	normalized := strings.ReplaceAll(rawText, "\r\n", "\n")
	normalized = strings.ReplaceAll(normalized, "\n", "\r\n")
	raw := []byte(normalized)
	head, body := httpio.SplitHeadBody(raw)
	if body == nil && !strings.Contains(normalized, "\r\n\r\n") {
		return nil, fmt.Errorf("malformed request: missing header/body separator")
	}
	if body == nil {
		body = []byte{}
	}
	start, headers := httpio.ParseHeaders(head)
	parts := strings.SplitN(start, " ", 3)
	if len(parts) < 2 {
		return nil, fmt.Errorf("malformed request line")
	}
	version := "HTTP/1.1"
	if len(parts) > 2 {
		version = parts[2]
	}
	req := &httpio.Request{
		Method:  strings.ToUpper(parts[0]),
		Path:    parts[1],
		Version: version,
		Headers: headers,
		Body:    body,
		Scheme:  scheme,
	}
	defPort := 80
	if scheme == "https" {
		defPort = 443
	}
	if target != "" {
		host, port := httpio.ParseAuthority(target, defPort)
		req.Host = host
		req.Port = port
		hostHdr := target
		if !strings.Contains(target, ":") {
			if port != 80 && port != 443 {
				hostHdr = fmt.Sprintf("%s:%d", host, port)
			} else {
				hostHdr = host
			}
		}
		req.SetHeader("Host", hostHdr)
	} else {
		hostHdr := req.Header("Host")
		if hostHdr == "" {
			return nil, fmt.Errorf("Host header required")
		}
		host, port := httpio.ParseAuthority(hostHdr, defPort)
		req.Host = host
		req.Port = port
	}
	req.DropHeaders("transfer-encoding", "content-length", "expect")
	req.SetHeader("Content-Length", strconv.Itoa(len(req.Body)))
	return req, nil
}

// Send replays a raw HTTP/1.1 request and returns the upstream response.
func Send(ctx context.Context, opts Options) Result {
	started := time.Now()
	if ctx == nil {
		ctx = context.Background()
	}
	timeout := opts.Timeout
	if timeout <= 0 {
		timeout = defaultTimeout
	}
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	req, err := ParseRawRequest(opts.Raw, opts.Scheme, opts.Target)
	if err != nil {
		return Result{OK: false, Error: err.Error(), DurationMS: elapsedMS(started)}
	}

	conn, err := dialUpstream(ctx, req, opts.Verify, timeout)
	if err != nil {
		return fail(started, req, err)
	}
	defer conn.Close()
	go func() {
		<-ctx.Done()
		_ = conn.Close()
	}()

	deadline, ok := ctx.Deadline()
	if !ok {
		deadline = time.Now().Add(timeout)
	}
	if err := conn.SetDeadline(deadline); err != nil {
		return fail(started, req, fmt.Errorf("failed to set deadline: %w", err))
	}
	if _, err := conn.Write(req.Raw()); err != nil {
		return fail(started, req, fmt.Errorf("failed to write request: %w", err))
	}
	br := bufio.NewReader(conn)
	resp, err := httpio.ReadResponse(br, req.Method)
	if err != nil {
		return fail(started, req, fmt.Errorf("failed to read response: %w", err))
	}
	if resp == nil {
		return fail(started, req, fmt.Errorf("empty response"))
	}
	codec.UnwrapBody(resp)
	return Result{
		OK:          true,
		Status:      resp.Status,
		Reason:      resp.Reason,
		DurationMS:  elapsedMS(started),
		RequestRaw:  httpio.DecodeRaw(req.Raw()),
		ResponseRaw: httpio.DecodeRaw(resp.Raw()),
	}
}

func dialUpstream(ctx context.Context, req *httpio.Request, verify bool, timeout time.Duration) (net.Conn, error) {
	addr := net.JoinHostPort(req.Host, strconv.Itoa(req.Port))
	d := net.Dialer{Timeout: timeout}
	conn, err := d.DialContext(ctx, "tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("failed to dial %s: %w", addr, err)
	}
	if req.Scheme != "https" {
		return conn, nil
	}
	cfg := &tls.Config{
		ServerName:         req.Host,
		InsecureSkipVerify: !verify,
		NextProtos:         []string{"http/1.1"},
		MinVersion:         tls.VersionTLS12,
	}
	tc := tls.Client(conn, cfg)
	if err := tc.HandshakeContext(ctx); err != nil {
		_ = conn.Close()
		return nil, fmt.Errorf("tls handshake %s: %w", req.Host, err)
	}
	return tc, nil
}

func fail(started time.Time, req *httpio.Request, err error) Result {
	out := Result{OK: false, Error: err.Error(), DurationMS: elapsedMS(started)}
	if req != nil {
		out.RequestRaw = httpio.DecodeRaw(req.Raw())
	}
	return out
}

func elapsedMS(started time.Time) int {
	return int(time.Since(started).Milliseconds())
}
