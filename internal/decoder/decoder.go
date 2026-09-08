// Package decoder applies a named text transform (encode or decode).
package decoder

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"html"
	"net/url"
	"strings"
)

// Kind is a supported transform.
type Kind string

const (
	Base64Encode Kind = "base64_encode"
	Base64Decode Kind = "base64_decode"
	URLEncode    Kind = "url_encode"
	URLDecode    Kind = "url_decode"
	HexEncode    Kind = "hex_encode"
	HexDecode    Kind = "hex_decode"
	HTMLUnescape Kind = "html_unescape"
)

var aliases = map[string]Kind{
	"base64_encode": Base64Encode,
	"b64_encode":    Base64Encode,
	"base64_decode": Base64Decode,
	"b64_decode":    Base64Decode,
	"url_encode":    URLEncode,
	"url_decode":    URLDecode,
	"hex_encode":    HexEncode,
	"hex_decode":    HexDecode,
	"html_unescape": HTMLUnescape,
	"html_decode":   HTMLUnescape,
}

// ParseKind maps a UI/API action name onto a Kind.
func ParseKind(s string) (Kind, error) {
	k, ok := aliases[strings.ToLower(strings.TrimSpace(s))]
	if !ok {
		return "", fmt.Errorf("unknown transform %q", s)
	}
	return k, nil
}

// Transform applies kind to input and returns the result.
func Transform(input string, kind Kind) (string, error) {
	switch kind {
	case Base64Encode:
		return base64.StdEncoding.EncodeToString([]byte(input)), nil
	case Base64Decode:
		out, err := decodeBase64(input)
		if err != nil {
			return "", fmt.Errorf("base64 decode: %w", err)
		}
		return string(out), nil
	case URLEncode:
		return encodeURL(input), nil
	case URLDecode:
		out, err := url.PathUnescape(input)
		if err != nil {
			return "", fmt.Errorf("url decode: %w", err)
		}
		return out, nil
	case HexEncode:
		return hex.EncodeToString([]byte(input)), nil
	case HexDecode:
		out, err := decodeHex(input)
		if err != nil {
			return "", fmt.Errorf("hex decode: %w", err)
		}
		return string(out), nil
	case HTMLUnescape:
		return html.UnescapeString(input), nil
	default:
		return "", fmt.Errorf("unknown transform %q", kind)
	}
}

func decodeBase64(data string) ([]byte, error) {
	raw := stripSpace(data)
	if raw == "" {
		return []byte{}, nil
	}
	if m := len(raw) % 4; m != 0 {
		raw += strings.Repeat("=", 4-m)
	}
	out, err := base64.StdEncoding.DecodeString(raw)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func decodeHex(data string) ([]byte, error) {
	cleaned := stripSpace(data)
	if cleaned == "" {
		return []byte{}, nil
	}
	if len(cleaned)%2 != 0 {
		return nil, fmt.Errorf("odd length %d", len(cleaned))
	}
	return hex.DecodeString(cleaned)
}

func stripSpace(s string) string {
	return strings.Map(func(r rune) rune {
		if r == ' ' || r == '\n' || r == '\r' || r == '\t' {
			return -1
		}
		return r
	}, s)
}

func encodeURL(s string) string {
	var b strings.Builder
	for i := 0; i < len(s); i++ {
		c := s[i]
		if unreserved(c) {
			b.WriteByte(c)
			continue
		}
		b.WriteByte('%')
		b.WriteString(strings.ToUpper(hex.EncodeToString([]byte{c})))
	}
	return b.String()
}

func unreserved(c byte) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') ||
		c == '-' || c == '_' || c == '.' || c == '~'
}
