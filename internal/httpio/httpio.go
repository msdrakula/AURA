package httpio

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/url"
	"strconv"
	"strings"
	"unicode/utf8"
)

var hopByHop = map[string]struct{}{
	"connection":          {},
	"keep-alive":          {},
	"proxy-authenticate":  {},
	"proxy-authorization": {},
	"proxy-connection":    {},
	"te":                  {},
	"trailer":             {},
	"transfer-encoding":   {},
}

const (
	MaxHeaderBlock = 64 * 1024
	MaxBody        = 10 * 1024 * 1024
)

type Header struct {
	Key, Value string
}

type Request struct {
	Method  string
	Path    string
	Version string
	Headers []Header
	Body    []byte
	Scheme  string
	Host    string
	Port    int
}

func (r *Request) Header(name string) string {
	needle := strings.ToLower(name)
	for _, h := range r.Headers {
		if strings.ToLower(h.Key) == needle {
			return h.Value
		}
	}
	return ""
}

func (r *Request) SetHeader(name, value string) {
	needle := strings.ToLower(name)
	for i, h := range r.Headers {
		if strings.ToLower(h.Key) == needle {
			r.Headers[i].Value = value
			return
		}
	}
	r.Headers = append(r.Headers, Header{Key: name, Value: value})
}

func (r *Request) DropHeaders(names ...string) {
	skip := map[string]struct{}{}
	for _, n := range names {
		skip[strings.ToLower(n)] = struct{}{}
	}
	out := r.Headers[:0]
	for _, h := range r.Headers {
		if _, ok := skip[strings.ToLower(h.Key)]; ok {
			continue
		}
		out = append(out, h)
	}
	r.Headers = out
}

func (r *Request) URL() string {
	def := 80
	if r.Scheme == "https" {
		def = 443
	}
	host := r.Host
	if r.Port != def && r.Port != 0 {
		host = fmt.Sprintf("%s:%d", r.Host, r.Port)
	}
	path := r.Path
	if path != "*" && !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return fmt.Sprintf("%s://%s%s", r.Scheme, host, path)
}

func (r *Request) Raw() []byte {
	path := r.Path
	if path == "" {
		path = "/"
	}
	var lines [][]byte
	lines = append(lines, []byte(fmt.Sprintf("%s %s %s", r.Method, path, r.Version)))
	headers := filterHop(r.Headers)
	hasCL := false
	for _, h := range headers {
		if strings.EqualFold(h.Key, "Content-Length") {
			hasCL = true
			break
		}
	}
	cl := strconv.Itoa(len(r.Body))
	if !hasCL {
		headers = append(headers, Header{Key: "Content-Length", Value: cl})
	} else {
		for i, h := range headers {
			if strings.EqualFold(h.Key, "Content-Length") {
				headers[i].Value = cl
			}
		}
	}
	for _, h := range headers {
		lines = append(lines, []byte(h.Key+": "+h.Value))
	}
	out := bytes.Join(lines, []byte("\r\n"))
	out = append(out, []byte("\r\n\r\n")...)
	out = append(out, r.Body...)
	return out
}

type Response struct {
	Version string
	Status  int
	Reason  string
	Headers []Header
	Body    []byte
}

func (r *Response) Header(name string) string {
	needle := strings.ToLower(name)
	for _, h := range r.Headers {
		if strings.ToLower(h.Key) == needle {
			return h.Value
		}
	}
	return ""
}

func (r *Response) SetHeader(name, value string) {
	needle := strings.ToLower(name)
	for i, h := range r.Headers {
		if strings.ToLower(h.Key) == needle {
			r.Headers[i].Value = value
			return
		}
	}
	r.Headers = append(r.Headers, Header{Key: name, Value: value})
}

func (r *Response) DropHeaders(names ...string) {
	skip := map[string]struct{}{}
	for _, n := range names {
		skip[strings.ToLower(n)] = struct{}{}
	}
	out := r.Headers[:0]
	for _, h := range r.Headers {
		if _, ok := skip[strings.ToLower(h.Key)]; ok {
			continue
		}
		out = append(out, h)
	}
	r.Headers = out
}

func (r *Response) Raw() []byte {
	var lines [][]byte
	lines = append(lines, []byte(fmt.Sprintf("%s %d %s", r.Version, r.Status, r.Reason)))
	headers := filterHop(r.Headers)
	outH := headers[:0]
	for _, h := range headers {
		if strings.EqualFold(h.Key, "Content-Length") {
			continue
		}
		outH = append(outH, h)
	}
	outH = append(outH, Header{Key: "Content-Length", Value: strconv.Itoa(len(r.Body))})
	for _, h := range outH {
		lines = append(lines, []byte(h.Key+": "+h.Value))
	}
	out := bytes.Join(lines, []byte("\r\n"))
	out = append(out, []byte("\r\n\r\n")...)
	out = append(out, r.Body...)
	return out
}

func filterHop(headers []Header) []Header {
	out := make([]Header, 0, len(headers))
	for _, h := range headers {
		if _, ok := hopByHop[strings.ToLower(h.Key)]; ok {
			continue
		}
		out = append(out, h)
	}
	return out
}

func DecodeHeaderBlock(raw []byte) (string, []Header) {
	text := string(raw)
	if strings.HasSuffix(text, "\r\n\r\n") {
		text = text[:len(text)-4]
	} else if strings.HasSuffix(text, "\n\n") {
		text = text[:len(text)-2]
	}
	var lines []string
	if strings.Contains(text, "\r\n") {
		lines = strings.Split(text, "\r\n")
	} else {
		lines = strings.Split(text, "\n")
	}
	start := ""
	if len(lines) > 0 {
		start = lines[0]
	}
	var headers []Header
	for _, line := range lines[1:] {
		if line == "" || line[0] == ' ' || line[0] == '\t' {
			continue
		}
		key, val, ok := strings.Cut(line, ":")
		if !ok {
			continue
		}
		headers = append(headers, Header{Key: strings.TrimSpace(key), Value: strings.TrimSpace(val)})
	}
	return start, headers
}

func ParseAuthority(authority string, defaultPort int) (string, int) {
	authority = strings.TrimSpace(authority)
	if strings.HasPrefix(authority, "[") {
		end := strings.Index(authority, "]")
		if end < 0 {
			return authority, defaultPort
		}
		host := authority[1:end]
		rest := authority[end+1:]
		if strings.HasPrefix(rest, ":") {
			p, err := strconv.Atoi(rest[1:])
			if err != nil {
				return host, defaultPort
			}
			return host, p
		}
		return host, defaultPort
	}
	if strings.Count(authority, ":") == 1 {
		host, portS, _ := strings.Cut(authority, ":")
		p, err := strconv.Atoi(portS)
		if err != nil {
			return host, defaultPort
		}
		return host, p
	}
	return authority, defaultPort
}

func ApplyAbsoluteURL(req *Request) {
	if strings.HasPrefix(req.Path, "http://") || strings.HasPrefix(req.Path, "https://") {
		parts, err := url.Parse(req.Path)
		if err != nil {
			return
		}
		if parts.Scheme != "" {
			req.Scheme = parts.Scheme
		}
		if parts.Hostname() != "" {
			req.Host = parts.Hostname()
		}
		if parts.Port() != "" {
			p, _ := strconv.Atoi(parts.Port())
			req.Port = p
		} else if req.Scheme == "https" {
			req.Port = 443
		} else {
			req.Port = 80
		}
		path := parts.Path
		if path == "" {
			path = "/"
		}
		if parts.RawQuery != "" {
			path += "?" + parts.RawQuery
		}
		req.Path = path
		return
	}
	hostHdr := req.Header("Host")
	if hostHdr != "" {
		def := 80
		if req.Scheme == "https" {
			def = 443
		}
		host, port := ParseAuthority(hostHdr, def)
		req.Host = host
		req.Port = port
	}
}

func ReadUntilDoubleCRLF(r *bufio.Reader) ([]byte, error) {
	var buf []byte
	for {
		b, err := r.ReadByte()
		if err != nil {
			return nil, err
		}
		buf = append(buf, b)
		if len(buf) > MaxHeaderBlock {
			return nil, errors.New("header block too large")
		}
		if bytes.HasSuffix(buf, []byte("\r\n\r\n")) {
			return buf, nil
		}
	}
}

func ReadChunkedBody(r *bufio.Reader) ([]byte, error) {
	var chunks []byte
	total := 0
	for {
		line, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF && len(line) == 0 && total == 0 {
				break
			}
			if err != io.EOF {
				return nil, err
			}
		}
		if len(line) == 0 {
			break
		}
		sizeS := strings.TrimSpace(string(bytes.SplitN(line, []byte(";"), 2)[0]))
		size, err := strconv.ParseInt(sizeS, 16, 64)
		if err != nil {
			return nil, err
		}
		if size == 0 {
			for {
				trailer, err := r.ReadBytes('\n')
				if err != nil && err != io.EOF {
					return nil, err
				}
				if bytes.Equal(trailer, []byte("\r\n")) || bytes.Equal(trailer, []byte("\n")) || len(trailer) == 0 {
					break
				}
			}
			break
		}
		total += int(size)
		if total > MaxBody {
			return nil, errors.New("body too large")
		}
		chunk := make([]byte, size)
		if _, err := io.ReadFull(r, chunk); err != nil {
			return nil, err
		}
		if _, err := io.ReadFull(r, make([]byte, 2)); err != nil {
			return nil, err
		}
		chunks = append(chunks, chunk...)
	}
	return chunks, nil
}

func hasChunked(headers []Header) bool {
	for _, h := range headers {
		if strings.EqualFold(h.Key, "Transfer-Encoding") && strings.Contains(strings.ToLower(h.Value), "chunked") {
			return true
		}
	}
	return false
}

func contentLength(headers []Header) (int, bool) {
	for _, h := range headers {
		if strings.EqualFold(h.Key, "Content-Length") {
			n, err := strconv.Atoi(strings.TrimSpace(h.Value))
			if err != nil {
				return 0, false
			}
			return n, true
		}
	}
	return 0, false
}

func ReadBody(r *bufio.Reader, headers []Header, allowUntilEOF bool) ([]byte, error) {
	if hasChunked(headers) {
		return ReadChunkedBody(r)
	}
	length, ok := contentLength(headers)
	if !ok {
		if allowUntilEOF {
			data, err := io.ReadAll(io.LimitReader(r, MaxBody+1))
			if err != nil {
				return nil, err
			}
			if len(data) > MaxBody {
				return nil, errors.New("body too large")
			}
			return data, nil
		}
		return []byte{}, nil
	}
	if length > MaxBody {
		return nil, errors.New("body too large")
	}
	if length == 0 {
		return []byte{}, nil
	}
	buf := make([]byte, length)
	_, err := io.ReadFull(r, buf)
	return buf, err
}

func ResponseMayHaveBody(reqMethod string, status int) bool {
	if reqMethod == "HEAD" {
		return false
	}
	if status == 204 || status == 304 {
		return false
	}
	if status >= 100 && status < 200 {
		return false
	}
	return true
}

func ReadRequest(r *bufio.Reader, defaultScheme string) (*Request, error) {
	raw, err := ReadUntilDoubleCRLF(r)
	if err != nil {
		return nil, err
	}
	start, headers := DecodeHeaderBlock(raw)
	parts := strings.SplitN(start, " ", 3)
	if len(parts) < 2 {
		return nil, errors.New("bad request line")
	}
	method := strings.ToUpper(parts[0])
	path := parts[1]
	version := "HTTP/1.1"
	if len(parts) > 2 {
		version = parts[2]
	}
	req := &Request{
		Method:  method,
		Path:    path,
		Version: version,
		Headers: headers,
		Scheme:  defaultScheme,
		Body:    []byte{},
	}
	if method != "CONNECT" {
		body, err := ReadBody(r, headers, false)
		if err != nil {
			return nil, err
		}
		req.Body = body
	}
	ApplyAbsoluteURL(req)
	if req.Host == "" {
		hostHdr := req.Header("Host")
		if hostHdr != "" {
			def := 80
			if defaultScheme == "https" {
				def = 443
			}
			host, port := ParseAuthority(hostHdr, def)
			req.Host = host
			req.Port = port
		}
	}
	req.DropHeaders("expect")
	return req, nil
}

func ReadResponse(r *bufio.Reader, reqMethod string) (*Response, error) {
	raw, err := ReadUntilDoubleCRLF(r)
	if err != nil {
		return nil, err
	}
	start, headers := DecodeHeaderBlock(raw)
	parts := strings.SplitN(start, " ", 3)
	version := "HTTP/1.1"
	if len(parts) > 0 {
		version = parts[0]
	}
	status := 0
	if len(parts) > 1 {
		status, _ = strconv.Atoi(parts[1])
	}
	reason := ""
	if len(parts) > 2 {
		reason = parts[2]
	}
	body := []byte{}
	if ResponseMayHaveBody(reqMethod, status) {
		conn := ""
		for _, h := range headers {
			if strings.EqualFold(h.Key, "Connection") {
				conn = strings.ToLower(h.Value)
			}
		}
		_, hasCL := contentLength(headers)
		untilEOF := strings.Contains(conn, "close") || (!hasCL && !hasChunked(headers))
		b, err := ReadBody(r, headers, untilEOF)
		if err != nil {
			return nil, err
		}
		body = b
	}
	return &Response{Version: version, Status: status, Reason: reason, Headers: headers, Body: body}, nil
}

func LooksLikeUpgrade(headers []Header) bool {
	conn, upgrade := "", ""
	for _, h := range headers {
		low := strings.ToLower(h.Key)
		if low == "connection" {
			conn = strings.ToLower(h.Value)
		} else if low == "upgrade" {
			upgrade = strings.ToLower(h.Value)
		}
	}
	return strings.Contains(conn, "upgrade") && upgrade != ""
}

func DecodeRaw(raw []byte) string {
	if utf8.Valid(raw) {
		return string(raw)
	}
	runes := make([]rune, len(raw))
	for i, c := range raw {
		runes[i] = rune(c)
	}
	return string(runes)
}

func SplitHeadBody(raw []byte) (head []byte, body []byte) {
	if i := bytes.Index(raw, []byte("\r\n\r\n")); i >= 0 {
		return raw[:i], raw[i+4:]
	}
	if i := bytes.Index(raw, []byte("\n\n")); i >= 0 {
		return raw[:i], raw[i+2:]
	}
	return raw, nil
}

func ParseHeaders(head []byte) (start string, headers []Header) {
	text := string(head)
	var lines []string
	if strings.Contains(text, "\r\n") {
		lines = strings.Split(text, "\r\n")
	} else {
		lines = strings.Split(text, "\n")
	}
	if len(lines) == 0 {
		return "", nil
	}
	start = lines[0]
	for _, line := range lines[1:] {
		if line == "" || !strings.Contains(line, ":") {
			continue
		}
		key, val, _ := strings.Cut(line, ":")
		headers = append(headers, Header{Key: strings.TrimSpace(key), Value: strings.TrimSpace(val)})
	}
	return start, headers
}
