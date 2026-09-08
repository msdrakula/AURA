package codec

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"compress/zlib"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"html"
	"io"
	"net"
	"net/url"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"strconv"
	"strings"

	"meb/internal/decoder"
	"meb/internal/httpio"
)

func unreserved(c byte) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') ||
		c == '-' || c == '_' || c == '.' || c == '~'
}

func encodeURL(s string, plus bool) string {
	var b strings.Builder
	for i := 0; i < len(s); i++ {
		c := s[i]
		if unreserved(c) {
			b.WriteByte(c)
			continue
		}
		if plus && c == ' ' {
			b.WriteByte('+')
			continue
		}
		b.WriteByte('%')
		b.WriteString(strings.ToUpper(hex.EncodeToString([]byte{c})))
	}
	return b.String()
}

func decodeURL(s string, plus bool) (string, error) {
	if plus {
		s = strings.ReplaceAll(s, "+", "%20")
	}
	return url.PathUnescape(s)
}

func padB64(raw string) string {
	if m := len(raw) % 4; m != 0 {
		raw += strings.Repeat("=", 4-m)
	}
	return raw
}

func asText(data []byte) string {
	return string(data)
}

func decodeJWT(token string) (string, error) {
	parts := strings.Split(strings.TrimSpace(token), ".")
	if len(parts) < 2 {
		return "", fmt.Errorf("not a JWT")
	}
	labels := []string{"header", "payload", "signature"}
	var out []string
	n := len(parts)
	if n > 3 {
		n = 3
	}
	for i := 0; i < n; i++ {
		label := labels[i]
		part := parts[i]
		if i < 2 {
			std := strings.ReplaceAll(strings.ReplaceAll(part, "-", "+"), "_", "/")
			blob, err := base64.StdEncoding.DecodeString(padB64(std))
			if err != nil {
				blob, err = base64.RawStdEncoding.DecodeString(std)
				if err != nil {
					return "", err
				}
			}
			var parsed any
			if json.Unmarshal(blob, &parsed) == nil {
				pretty, _ := json.MarshalIndent(parsed, "", "  ")
				out = append(out, label+":\n"+string(pretty))
			} else {
				out = append(out, label+":\n"+asText(blob))
			}
		} else {
			out = append(out, label+":\n"+part)
		}
	}
	return strings.Join(out, "\n\n"), nil
}

func Run(action, data string) (string, error) {
	if k, err := decoder.ParseKind(action); err == nil {
		return decoder.Transform(data, k)
	}
	switch action {
	case "url_encode_plus":
		return encodeURL(data, true), nil
	case "url_decode_plus":
		return decodeURL(data, true)
	case "b64url_encode":
		return strings.TrimRight(base64.URLEncoding.EncodeToString([]byte(data)), "="), nil
	case "b64url_decode":
		raw := strings.ReplaceAll(data, " ", "")
		out, err := base64.URLEncoding.DecodeString(padB64(raw))
		if err != nil {
			out, err = base64.RawURLEncoding.DecodeString(raw)
			if err != nil {
				return "", err
			}
		}
		return asText(out), nil
	case "html_encode":
		return html.EscapeString(data), nil
	case "gzip_encode":
		var buf bytes.Buffer
		zw := gzip.NewWriter(&buf)
		if _, err := zw.Write([]byte(data)); err != nil {
			return "", err
		}
		if err := zw.Close(); err != nil {
			return "", err
		}
		return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
	case "gzip_decode":
		blob := []byte(data)
		cleaned := strings.Map(func(r rune) rune {
			if r == ' ' || r == '\n' || r == '\r' || r == '\t' {
				return -1
			}
			return r
		}, data)
		if decoded, err := base64.StdEncoding.DecodeString(padB64(cleaned)); err == nil {
			blob = decoded
		}
		zr, err := gzip.NewReader(bytes.NewReader(blob))
		if err != nil {
			zr, err = gzip.NewReader(bytes.NewReader([]byte(data)))
			if err != nil {
				return "", err
			}
		}
		defer zr.Close()
		out, err := io.ReadAll(zr)
		if err != nil {
			return "", err
		}
		return string(out), nil
	case "jwt_decode":
		return decodeJWT(data)
	case "unicode_escape":
		quoted := `"` + strings.ReplaceAll(data, `"`, `\"`) + `"`
		s, err := strconv.Unquote(quoted)
		if err != nil {
			return "", err
		}
		return s, nil
	case "hash_md5":
		return hashHex(md5.Sum([]byte(data))), nil
	case "hash_sha1":
		return hashHex(sha1.Sum([]byte(data))), nil
	case "hash_sha256":
		return hashHex(sha256.Sum256([]byte(data))), nil
	case "hash_sha512":
		return hashHex(sha512.Sum512([]byte(data))), nil
	case "smart_decode":
		return smartDecode(data)
	case "ascii_hex_encode":
		return hex.EncodeToString([]byte(data)), nil
	case "ascii_hex_decode":
		out, err := hex.DecodeString(strings.ReplaceAll(data, " ", ""))
		if err != nil {
			return "", err
		}
		return asText(out), nil
	case "octal_encode":
		var b strings.Builder
		for _, c := range []byte(data) {
			fmt.Fprintf(&b, "\\%03o", c)
		}
		return b.String(), nil
	case "binary_encode":
		var b strings.Builder
		for _, c := range []byte(data) {
			fmt.Fprintf(&b, "%08b", c)
		}
		return b.String(), nil
	default:
		return "", fmt.Errorf("unknown action: %s", action)
	}
}

func hashHex(sum interface{}) string {
	switch v := sum.(type) {
	case [16]byte:
		return hex.EncodeToString(v[:])
	case [20]byte:
		return hex.EncodeToString(v[:])
	case [32]byte:
		return hex.EncodeToString(v[:])
	case [64]byte:
		return hex.EncodeToString(v[:])
	}
	return ""
}

// smartDecode repeatedly applies recognizable decodings until no more apply.
func smartDecode(data string) (string, error) {
	cur := data
	for i := 0; i < 10; i++ {
		next, err := smartDecodeStep(cur)
		if err != nil || next == cur {
			return cur, nil
		}
		cur = next
	}
	return cur, nil
}

func smartDecodeStep(data string) (string, error) {
	// URL-encoded
	if strings.Contains(data, "%") {
		if d, err := decodeURL(data, false); err == nil && d != data {
			return d, nil
		}
	}
	// Base64 (length multiple of 4, valid charset)
	trimmed := strings.ReplaceAll(strings.ReplaceAll(data, " ", ""), "\n", "")
	if len(trimmed) >= 8 && len(trimmed)%4 == 0 && isBase64(trimmed) {
		if out, err := base64.StdEncoding.DecodeString(trimmed); err == nil && isPrintable(out) {
			return string(out), nil
		}
	}
	// HTML entities
	if strings.Contains(data, "&") {
		if d := html.UnescapeString(data); d != data {
			return d, nil
		}
	}
	// Hex string
	hx := strings.ReplaceAll(data, " ", "")
	if len(hx) >= 4 && len(hx)%2 == 0 && isHex(hx) {
		if out, err := hex.DecodeString(hx); err == nil && isPrintable(out) {
			return string(out), nil
		}
	}
	return data, nil
}

func isBase64(s string) bool {
	for _, c := range s {
		if !((c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || c == '+' || c == '/' || c == '=') {
			return false
		}
	}
	return true
}

func isHex(s string) bool {
	for _, c := range s {
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')) {
			return false
		}
	}
	return true
}

func isPrintable(b []byte) bool {
	if len(b) == 0 {
		return false
	}
	printable := 0
	for _, c := range b {
		if c == '\n' || c == '\r' || c == '\t' || (c >= 32 && c < 127) {
			printable++
		}
	}
	return printable*100/len(b) > 80
}

// UnwrapBody replaces a gzip/deflate payload with the decoded bytes and
// drops Content-Encoding / Transfer-Encoding so the body can be shown as-is.
func UnwrapBody(resp *httpio.Response) {
	if resp == nil {
		return
	}
	decoded, ok := MaybeDecompress(resp.Body, resp.Header("Content-Encoding"))
	if !ok {
		return
	}
	resp.Body = decoded
	resp.DropHeaders("content-encoding", "transfer-encoding")
	resp.SetHeader("Content-Length", strconv.Itoa(len(resp.Body)))
}

func MaybeDecompress(body []byte, encoding string) ([]byte, bool) {
	if encoding == "" {
		return body, false
	}
	enc := strings.ToLower(encoding)
	if strings.Contains(enc, "gzip") {
		zr, err := gzip.NewReader(bytes.NewReader(body))
		if err != nil {
			return body, false
		}
		defer zr.Close()
		out, err := io.ReadAll(zr)
		if err != nil {
			return body, false
		}
		return out, true
	}
	if strings.Contains(enc, "deflate") {
		if zr, err := zlib.NewReader(bytes.NewReader(body)); err == nil {
			out, err := io.ReadAll(zr)
			zr.Close()
			if err == nil {
				return out, true
			}
		}
		return inflateRaw(body)
	}
	return body, false
}

func inflateRaw(body []byte) ([]byte, bool) {
	r := flate.NewReader(bytes.NewReader(body))
	defer r.Close()
	out, err := io.ReadAll(r)
	if err != nil {
		return body, false
	}
	return out, true
}

func Pipe(a, b net.Conn, extraA, extraB io.Reader) {
	var srcA, srcB io.Reader = a, b
	if extraA != nil {
		srcA = io.MultiReader(extraA, a)
	}
	if extraB != nil {
		srcB = io.MultiReader(extraB, b)
	}
	done := make(chan struct{}, 2)
	go func() {
		_, _ = io.Copy(b, srcA)
		done <- struct{}{}
	}()
	go func() {
		_, _ = io.Copy(a, srcB)
		done <- struct{}{}
	}()
	<-done
	_ = a.Close()
	_ = b.Close()
	<-done
}
