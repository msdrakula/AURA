package ffuf

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseResultLine(t *testing.T) {
	payload := base64.StdEncoding.EncodeToString([]byte("admin"))
	line := fmt.Sprintf(
		`{"input":{"FFUFHASH":"MTMwZDIx","FUZZ":%q},"position":1,"status":301,"length":178,"words":12,"lines":4,"content-type":"text/html","redirectlocation":"/login","url":"https://lab.local/admin","duration":4932922}`,
		payload)
	r, ok := parseResultLine(line, "FUZZ")
	require.True(t, ok)
	assert.Equal(t, "admin", r.Payload)
	assert.Equal(t, 1, r.Position)
	assert.Equal(t, 301, r.Status)
	assert.Equal(t, 178, r.Length)
	assert.Equal(t, 12, r.Words)
	assert.Equal(t, 4, r.Lines)
	assert.Equal(t, "text/html", r.ContentType)
	assert.Equal(t, "/login", r.Redirect)
	assert.Equal(t, "https://lab.local/admin", r.URL)
	assert.Equal(t, 4, r.DurationMs) // 4.9ms of nanoseconds truncated
}

func TestParseResultLineSkipsNoise(t *testing.T) {
	for _, line := range []string{"", "banner text", `[::] Progress: [1/2] ::`, `not-json{`} {
		_, ok := parseResultLine(line, "FUZZ")
		assert.False(t, ok, line)
	}
}

func TestDecodeInputFallback(t *testing.T) {
	assert.Equal(t, "plain", decodeInput("plain")) // not base64 → raw
	assert.Equal(t, "admin", decodeInput(base64.StdEncoding.EncodeToString([]byte("admin"))))
}

func TestParseProgressLine(t *testing.T) {
	d, tot, ok := parseProgressLine(":: Progress: [142/4615] :: Job [1/1] :: 124 req/sec :: Duration: [0:00:01] :: Errors: 3 ::")
	require.True(t, ok)
	assert.Equal(t, 142, d)
	assert.Equal(t, 4615, tot)

	_, _, ok = parseProgressLine("some other line")
	assert.False(t, ok)
}

func TestLastErrorCount(t *testing.T) {
	assert.Equal(t, 0, lastErrorCount("no errors here"))
	assert.Equal(t, 3, lastErrorCount(":: Progress: [1/2] :: Errors: 1 ::\n:: Progress: [2/2] :: Errors: 3 ::"))
}

func TestBuildArgs(t *testing.T) {
	args := buildArgs(Options{
		URL:      "https://lab.local/FUZZ",
		Method:   "POST",
		Headers:  []string{"X-Test: 1", ""},
		Cookies:  "session=abc",
		Body:     "user=FUZZ",
		Workers:  5,
		RPS:      20,
		Timeout:  7,
		Keyword:  "FUZZ",
	}, "/tmp/words.txt")
	joined := strings.Join(args, " ")
	for _, want := range []string{"-json", "-noninteractive", "-u https://lab.local/FUZZ", "-w /tmp/words.txt:FUZZ", "-X POST", "-t 5", "-timeout 7", "-mc all", "-rate 20", "-H", "X-Test: 1", "-b session=abc", "-d user=FUZZ"} {
		assert.Contains(t, joined, want)
	}
	assert.NotContains(t, joined, "\"\"") // empty header dropped
}

// installFakeFfuf puts a shell script named ffuf on PATH that emits the given
// stdout/stderr and exits 0.
func installFakeFfuf(t *testing.T, stdout, stderr string) {
	t.Helper()
	dir := t.TempDir()
	script := "#!/bin/sh\nprintf '%s' " + shellQuote(stdout) + "\nprintf '%s' " + shellQuote(stderr) + " 1>&2\n"
	require.NoError(t, os.WriteFile(filepath.Join(dir, "ffuf"), []byte(script), 0o755))
	t.Setenv("PATH", dir+string(os.PathListSeparator)+os.Getenv("PATH"))
}

func shellQuote(s string) string {
	return "'" + strings.ReplaceAll(s, "'", `'\''`) + "'"
}

func TestRunWithFakeBinary(t *testing.T) {
	payload := base64.StdEncoding.EncodeToString([]byte("admin"))
	stdout := fmt.Sprintf(`{"input":{"FUZZ":%q},"position":1,"status":200,"length":10,"words":2,"lines":1,"url":"https://lab.local/admin","duration":1000000}`, payload) + "\n" +
		fmt.Sprintf(`{"input":{"FUZZ":%q},"position":2,"status":404,"length":9,"words":1,"lines":1,"url":"https://lab.local/nope","duration":2000000}`, base64.StdEncoding.EncodeToString([]byte("nope"))) + "\n"
	stderr := "\n        banner\n\n :: Method : GET\n:: Progress: [1/2] :: Job [1/1] :: Errors: 0 ::\n:: Progress: [2/2] :: Job [1/1] :: Errors: 1 ::\n"
	installFakeFfuf(t, stdout, stderr)

	var progresses [][2]int
	results, st, err := Run(context.Background(), Options{
		URL:      "https://lab.local/FUZZ",
		Wordlist: []string{"admin", "nope"},
		Workers:  2,
		OnProgress: func(done, total int) {
			progresses = append(progresses, [2]int{done, total})
		},
	})
	require.NoError(t, err)
	require.Len(t, results, 2)
	assert.Equal(t, "admin", results[0].Payload)
	assert.Equal(t, 200, results[0].Status)
	assert.Equal(t, 1, results[0].DurationMs)
	assert.Equal(t, 2, st.Results)
	assert.Equal(t, 2, st.Tried)
	assert.Equal(t, 2, st.Total)
	assert.Equal(t, 1, st.Errors) // last reported error count
	assert.NotEmpty(t, progresses)
}

func TestRunRequiresFuzzKeyword(t *testing.T) {
	_, _, err := Run(context.Background(), Options{URL: "https://lab.local/", Wordlist: []string{"a"}})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "FUZZ")
}

func TestRunRequiresWordlist(t *testing.T) {
	_, _, err := Run(context.Background(), Options{URL: "https://lab.local/FUZZ"})
	require.Error(t, err)
}

func TestRunContextCancel(t *testing.T) {
	dir := t.TempDir()
	// Fake ffuf that stays alive until killed; exec replaces the shell so the
	// signal reaches the process holding the pipes.
	script := "#!/bin/sh\nexec sleep 30\n"
	require.NoError(t, os.WriteFile(filepath.Join(dir, "ffuf"), []byte(script), 0o755))
	t.Setenv("PATH", dir+string(os.PathListSeparator)+os.Getenv("PATH"))

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	start := time.Now()
	_, _, err := Run(ctx, Options{URL: "https://lab.local/FUZZ", Wordlist: []string{"a"}})
	require.Error(t, err)
	assert.ErrorIs(t, err, context.DeadlineExceeded)
	assert.Less(t, time.Since(start), 5*time.Second)
}
