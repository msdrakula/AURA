package api

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

// installFakeFfufForAPI puts a shell script named ffuf on PATH that emits the
// given stdout/stderr and exits 0.
func installFakeFfufForAPI(t *testing.T, stdout, stderr string) {
	t.Helper()
	dir := t.TempDir()
	script := "#!/bin/sh\nprintf '%s' '" + stdout + "'\nprintf '%s' '" + stderr + "' 1>&2\n"
	require.NoError(t, os.WriteFile(filepath.Join(dir, "ffuf"), []byte(script), 0o755))
	t.Setenv("PATH", dir+string(os.PathListSeparator)+os.Getenv("PATH"))
}

func TestFuzzRunStreamsNDJSON(t *testing.T) {
	payload := base64.StdEncoding.EncodeToString([]byte("admin"))
	stdout := fmt.Sprintf(`{"input":{"FUZZ":%q},"position":1,"status":200,"length":10,"words":2,"lines":1,"url":"https://lab.local/admin","duration":1000000}`, payload) + "\n"
	stderr := ":: Progress: [1/1] :: Job [1/1] :: Errors: 0 ::\n"
	installFakeFfufForAPI(t, stdout, stderr)

	h := New(Options{Log: zap.NewNop()})
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/fuzz/run", strings.NewReader(
		`{"mode":"fuzz","url":"https://lab.local/FUZZ","wordlist":["admin"],"authorized":true}`))
	h.ServeHTTP(rec, req)

	require.Equal(t, 200, rec.Code)
	require.Contains(t, rec.Header().Get("Content-Type"), "application/x-ndjson")

	var types []string
	var sawResult bool
	var done map[string]any
	sc := bufio.NewScanner(rec.Body)
	for sc.Scan() {
		var ev map[string]any
		require.NoError(t, json.Unmarshal(sc.Bytes(), &ev))
		types = append(types, ev["type"].(string))
		if ev["type"] == "result" {
			sawResult = true
			assert.Equal(t, "admin", ev["payload"])
			assert.Equal(t, float64(200), ev["status"])
		}
		if ev["type"] == "done" {
			done = ev
		}
	}
	require.NoError(t, sc.Err())
	require.Equal(t, []string{"start", "result", "done"}, types)
	assert.True(t, sawResult)
	require.NotNil(t, done)
	assert.Equal(t, float64(1), done["count"])
	assert.Equal(t, float64(1), done["tried"])
	assert.Equal(t, "", done["error"])
}

func TestFuzzRunPathsModeAppendsFuzz(t *testing.T) {
	dir := t.TempDir()
	script := "#!/bin/sh\nprintf '%s' '" + `{"input":{"FUZZ":"YQ=="},"position":1,"status":200,"length":1,"url":"https://lab.local/a","duration":1000}` + "'\nprintf '%s' ':: Progress: [1/1] :: Errors: 0 ::' 1>&2\n"
	require.NoError(t, os.WriteFile(filepath.Join(dir, "ffuf"), []byte(script), 0o755))
	t.Setenv("PATH", dir+string(os.PathListSeparator)+os.Getenv("PATH"))

	h := New(Options{Log: zap.NewNop()})
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/fuzz/run", strings.NewReader(
		`{"mode":"paths","url":"https://lab.local/","wordlist":["a"],"authorized":true}`))
	h.ServeHTTP(rec, req)

	require.Equal(t, 200, rec.Code)
	body := rec.Body.String()
	assert.Contains(t, body, `"type":"start"`)
	assert.Contains(t, body, `"url":"https://lab.local/FUZZ"`)
}
