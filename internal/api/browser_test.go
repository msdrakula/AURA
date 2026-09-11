package api

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChromiumArgsAreChromiumOnly(t *testing.T) {
	args := chromiumArgs("127.0.0.1:8080", "/tmp/aura-profile")
	joined := strings.Join(args, " ")
	assert.Contains(t, joined, "--proxy-server=http://127.0.0.1:8080")
	assert.Contains(t, joined, "--user-data-dir=/tmp/aura-profile")
	for _, a := range args {
		assert.NotEqual(t, "-profile", a)
	}
}

func TestFirefoxArgsAreNotChromiumFlags(t *testing.T) {
	args := firefoxArgs("/tmp/ff-profile")
	joined := strings.Join(args, " ")
	assert.Contains(t, joined, "-no-remote")
	assert.Contains(t, joined, "-profile")
	assert.NotContains(t, joined, "--proxy-server")
	assert.NotContains(t, joined, "--ignore-certificate-errors")
	assert.NotContains(t, joined, "--user-data-dir")
}

func TestFirefoxPrefsWriteProxy(t *testing.T) {
	dir := t.TempDir()
	require.NoError(t, writeFirefoxProxyPrefs(dir, "127.0.0.1", 8080))
	raw, err := os.ReadFile(filepath.Join(dir, "user.js"))
	require.NoError(t, err)
	s := string(raw)
	assert.Contains(t, s, `user_pref("network.proxy.http", "127.0.0.1")`)
	assert.Contains(t, s, `user_pref("network.proxy.http_port", 8080)`)
	assert.Contains(t, s, `user_pref("network.proxy.ssl_port", 8080)`)
}

func TestSplitProxyAddr(t *testing.T) {
	host, port, err := splitProxyAddr("127.0.0.1:8082")
	require.NoError(t, err)
	assert.Equal(t, "127.0.0.1", host)
	assert.Equal(t, 8082, port)
}
