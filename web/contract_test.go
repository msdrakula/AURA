package web

import (
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"testing"
)

// Contract tests lock UI wiring so later agents run `go test ./web` instead of
// clicking through the GTK window after every comment.

func readWeb(t *testing.T, rel string) string {
	t.Helper()
	_, this, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("caller")
	}
	b, err := os.ReadFile(filepath.Join(filepath.Dir(this), rel))
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

func readRepo(t *testing.T, rel string) string {
	t.Helper()
	_, this, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("caller")
	}
	b, err := os.ReadFile(filepath.Join(filepath.Dir(this), "..", rel))
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

func TestCacheBustAligned(t *testing.T) {
	index := readWeb(t, "index.html")
	main := readRepo(t, "cmd/server/main.go")
	css := regexp.MustCompile(`app\.css\?v=(\d+)`).FindStringSubmatch(index)
	js := regexp.MustCompile(`app\.js\?v=(\d+)`).FindStringSubmatch(index)
	i18n := regexp.MustCompile(`i18n\.js\?v=(\d+)`).FindStringSubmatch(index)
	if len(css) < 2 || len(js) < 2 || len(i18n) < 2 {
		t.Fatal("missing ?v= on css/js/i18n")
	}
	if css[1] != js[1] || css[1] != i18n[1] {
		t.Fatalf("index cache bust mismatch css=%s js=%s i18n=%s", css[1], js[1], i18n[1])
	}
	if !strings.Contains(main, "?v="+css[1]) {
		t.Fatalf("cmd/server/main.go uiURL must use ?v=%s (GTK loads this URL)", css[1])
	}
}

func TestToolbarFieldsNotSquares(t *testing.T) {
	css := readWeb(t, "static/app.css")
	bad := regexp.MustCompile(`\.toolbar input,\s*\.toolbar select\s*\{[^}]*flex:\s*1`)
	if bad.MatchString(css) {
		t.Fatal("`.toolbar input, .toolbar select { flex: 1 … }` stretches Mode/Scheme/Target into squares; keep the rule on `.toolbar > input` only and reset nested label controls")
	}
	if !strings.Contains(css, ".toolbar > label input") {
		t.Fatal("missing `.toolbar > label input` height lock")
	}
	if !strings.Contains(css, "flex: 0 0 auto") {
		t.Fatal("toolbar nested inputs must be flex: 0 0 auto")
	}
}

func TestScannerSynthesizesGETFromTarget(t *testing.T) {
	js := readWeb(t, "static/app.js")
	html := readWeb(t, "index.html")
	for _, needle := range []string{"rawGetFromTarget", `tr("scan.need")`, `tr("scan.usingGet"`} {
		if !strings.Contains(js, needle) {
			t.Errorf("scanner start path missing %s", needle)
		}
	}
	if !strings.Contains(html, `id="scanRaw"`) || !strings.Contains(html, `id="scanTarget"`) {
		t.Fatal("scanner form missing scanRaw/scanTarget")
	}
}

func TestMapActivityLogWired(t *testing.T) {
	html := readWeb(t, "index.html")
	js := readWeb(t, "static/app.js")
	css := readWeb(t, "static/app.css")
	if !strings.Contains(html, `id="mapLog"`) {
		t.Fatal("map activity pane #mapLog missing")
	}
	if !strings.Contains(js, "intel_progress") || !strings.Contains(js, "MapLog.append") {
		t.Fatal("websocket intel_progress must append to MapLog")
	}
	if !strings.Contains(css, ".map-log") {
		t.Fatal("missing .map-log styles")
	}
}

func TestPayloadsWorkersSentToAPI(t *testing.T) {
	js := readWeb(t, "static/app.js")
	html := readWeb(t, "index.html")
	if !strings.Contains(html, `id="intrWorkers"`) || !strings.Contains(html, `id="intrRps"`) {
		t.Fatal("Payloads toolbar missing Workers/RPS")
	}
	if !strings.Contains(js, `$("#intrWorkers")`) || !strings.Contains(js, `$("#intrRps")`) {
		t.Fatal("app.js must read #intrWorkers and #intrRps and send them on /api/batch/execute")
	}
	if !strings.Contains(js, "workers:") || !strings.Contains(js, "rps:") {
		t.Fatal("batch execute JSON must include workers and rps")
	}
}

func TestI18nENRUParity(t *testing.T) {
	src := readWeb(t, "static/i18n.js")
	en := i18nKeys(t, src, "en")
	ru := i18nKeys(t, src, "ru")
	for k := range en {
		if _, ok := ru[k]; !ok {
			t.Errorf("RU missing key %q", k)
		}
	}
	for k := range ru {
		if _, ok := en[k]; !ok {
			t.Errorf("EN missing key %q", k)
		}
	}
}

func TestHTMLI18nKeysExist(t *testing.T) {
	html := readWeb(t, "index.html")
	en := i18nKeys(t, readWeb(t, "static/i18n.js"), "en")
	re := regexp.MustCompile(`data-i18n(?:-html|-placeholder|-title|-aria)?="([^"]+)"`)
	for _, m := range re.FindAllStringSubmatch(html, -1) {
		if _, ok := en[m[1]]; !ok {
			t.Errorf("HTML references missing i18n key %q", m[1])
		}
	}
}

func TestScanNeedKeyPresent(t *testing.T) {
	en := i18nKeys(t, readWeb(t, "static/i18n.js"), "en")
	if _, ok := en["scan.need"]; !ok {
		t.Fatal("scan.need missing")
	}
	if _, ok := en["scan.usingGet"]; !ok {
		t.Fatal("scan.usingGet missing")
	}
}

func i18nKeys(t *testing.T, src, lang string) map[string]struct{} {
	t.Helper()
	marker := lang + ": {"
	i := strings.Index(src, marker)
	if i < 0 {
		t.Fatalf("no I18N.%s block", lang)
	}
	block := src[i:]
	end := strings.Index(block, "\n  },")
	if end < 0 {
		t.Fatalf("unterminated I18N.%s", lang)
	}
	block = block[:end]
	out := map[string]struct{}{}
	dup := map[string]int{}
	re := regexp.MustCompile(`(?m)^\s+"([^"]+)":`)
	for _, m := range re.FindAllStringSubmatch(block, -1) {
		dup[m[1]]++
		out[m[1]] = struct{}{}
	}
	for k, n := range dup {
		if n > 1 {
			t.Errorf("duplicate I18N.%s key %q (%d times)", lang, k, n)
		}
	}
	return out
}
