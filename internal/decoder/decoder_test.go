package decoder

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransform(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		kind    Kind
		in      string
		want    string
		wantErr string
	}{
		{name: "base64 encode", kind: Base64Encode, in: "hello", want: "aGVsbG8="},
		{name: "base64 decode", kind: Base64Decode, in: "aGVsbG8=", want: "hello"},
		{name: "base64 decode padded whitespace", kind: Base64Decode, in: "aGVs bG8=\n", want: "hello"},
		{name: "base64 decode missing pad", kind: Base64Decode, in: "aGVsbG8", want: "hello"},
		{name: "base64 decode invalid", kind: Base64Decode, in: "!!!!", wantErr: "base64 decode"},
		{name: "url encode", kind: URLEncode, in: "a b/c", want: "a%20b%2Fc"},
		{name: "url decode", kind: URLDecode, in: "a%20b%2Fc", want: "a b/c"},
		{name: "url decode bad percent", kind: URLDecode, in: "%zz", wantErr: "url decode"},
		{name: "hex encode", kind: HexEncode, in: "OK", want: "4f4b"},
		{name: "hex decode", kind: HexDecode, in: "4f4b", want: "OK"},
		{name: "hex decode spaced", kind: HexDecode, in: "4f 4b", want: "OK"},
		{name: "hex decode odd length", kind: HexDecode, in: "abc", wantErr: "hex decode"},
		{name: "hex decode invalid nibble", kind: HexDecode, in: "zz", wantErr: "hex decode"},
		{name: "html unescape entities", kind: HTMLUnescape, in: "&lt;div&gt;&amp;", want: "<div>&"},
		{name: "html unescape numeric", kind: HTMLUnescape, in: "&#39;x&#x27;", want: "'x'"},
		{name: "unknown kind", kind: Kind("rot13"), in: "x", wantErr: "unknown transform"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := Transform(tt.in, tt.kind)
			if tt.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestParseKind(t *testing.T) {
	t.Parallel()
	k, err := ParseKind("b64_encode")
	require.NoError(t, err)
	assert.Equal(t, Base64Encode, k)

	_, err = ParseKind("nope")
	require.Error(t, err)
}
