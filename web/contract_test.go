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
	if !strings.Contains(css, ".toolbar > label input:not([type=\"checkbox\"])") {
		t.Fatal("toolbar nested text inputs must lock height; checkboxes must be excluded")
	}
	if !strings.Contains(css, ".toolbar > .tb-mode") {
		t.Fatal("Payloads Mode must use .tb-mode so RU «Все сочетания» is not clipped")
	}
	if !strings.Contains(css, "select:not([multiple]):not([size])") || !strings.Contains(css, "-webkit-appearance: none") {
		t.Fatal("compact selects need appearance:none — GTK native chevron overlays the value")
	}
	if !strings.Contains(css, "min-width: 22em") {
		t.Fatal("tb-mode must stay wide enough for «По одной позиции» / «Все сочетания»")
	}
	if !strings.Contains(readWeb(t, "index.html"), `class="tb-field tb-mode"`) {
		t.Fatal("intrAttackType label needs tb-mode")
	}
}

func TestNavTabsReorderAndLock(t *testing.T) {
	js := readWeb(t, "static/app.js")
	html := readWeb(t, "index.html")
	i18n := readWeb(t, "static/i18n.js")
	css := readWeb(t, "static/app.css")
	for _, needle := range []string{"NavTabs", "tabOrder", "uiLocked", "startDrag", "set.lockUi"} {
		if !strings.Contains(js, needle) {
			t.Errorf("nav reorder/lock missing %s", needle)
		}
	}
	if !strings.Contains(html, `id="btnLockUi"`) {
		t.Fatal("settings must have Lock interface button")
	}
	if !strings.Contains(css, `html[data-ui-locked="1"]`) {
		t.Fatal("locked interface must disable tab grab and pane gutters")
	}
	if !strings.Contains(i18n, `"set.lockUi"`) || !strings.Contains(i18n, `"set.unlockUi"`) {
		t.Fatal("lock/unlock i18n keys missing")
	}
}

func TestReplayAndPayloadsNewTabSameSide(t *testing.T) {
	css := readWeb(t, "static/app.css")
	html := readWeb(t, "index.html")
	if !strings.Contains(css, ".intr-tabs, .rep-tabs") {
		t.Fatal("Payloads and Replay tab strips must share the same flex rules")
	}
	if !strings.Contains(css, "order: -1") || !strings.Contains(css, "margin-left: 0") {
		t.Fatal("+ new-tab must sit on the left of the tab strip in Payloads and Replay")
	}
	if !strings.Contains(html, `id="btnIntrNewTab"`) || !strings.Contains(html, `id="btnRepNewTab"`) {
		t.Fatal("missing new-tab buttons")
	}
}

func TestCheckboxesNotSquares(t *testing.T) {
	css := readWeb(t, "static/app.css")
	js := readWeb(t, "static/app.js")
	if !strings.Contains(css, `input[type="checkbox"]`) {
		t.Fatal("missing checkbox size lock")
	}
	if !strings.Contains(css, "max-width: 15px") {
		t.Fatal("checkboxes must stay 15px, not stretch to 100% of the label")
	}
	if strings.Contains(js, `p.style.overflow = "hidden"`) {
		t.Fatal("UILayout.applySplit must not force overflow:hidden on panes — Discover/Scanner forms need a vertical scrollbar")
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

func TestMapModulesAndNav(t *testing.T) {
	html := readWeb(t, "index.html")
	js := readWeb(t, "static/app.js")
	if !strings.Contains(html, `data-view="scope"`) || !strings.Contains(html, `data-view="other"`) {
		t.Fatal("Scope and Other must be main tabs")
	}
	if strings.Contains(html, `data-view="target"`) {
		t.Fatal("Target is no longer a main tab — the map lives on Map")
	}
	if !strings.Contains(html, `id="view-scope"`) || !strings.Contains(html, `id="view-other"`) {
		t.Fatal("missing #view-scope / #view-other")
	}
	if !strings.Contains(html, `id="mapSubtabs"`) || !strings.Contains(html, `data-sub="sitemap"`) || !strings.Contains(html, `data-sub="subdomains_passive"`) {
		t.Fatal("Map must nest Site map and modules like Proxy subtabs")
	}
	if strings.Contains(html, `id="otherSubtabs"`) || strings.Contains(html, `data-i18n="other.sitemap"`) {
		t.Fatal("Site map belongs on Map, not Other")
	}
	if !strings.Contains(html, `id="siteTree"`) || !strings.Contains(html, `data-subpane="sitemap"`) {
		t.Fatal("Site map tree must live in the first Map subtab")
	}
	if !strings.Contains(html, `class="map-stage-page" data-stage="dirs"`) {
		t.Fatal("each Map module needs its own subpane")
	}
	for _, needle := range []string{"mapStageFields", "collectMapOpts", "readStageForm", "timeout_sec", "same_host", "mergeIntelIntoSiteTree"} {
		if !strings.Contains(js, needle) {
			t.Errorf("map module settings missing %s", needle)
		}
	}
	if !strings.Contains(js, "JSON.stringify(opts)") {
		t.Fatal("Run must POST per-module options")
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

func TestLabConfirmWiredOnActiveTools(t *testing.T) {
	html := readWeb(t, "index.html")
	js := readWeb(t, "static/app.js")
	for _, id := range []string{`id="discAuth"`, `id="fuzzAuth"`, `id="scanAuth"`} {
		if !strings.Contains(html, id) {
			t.Fatalf("missing %s", id)
		}
	}
	if !strings.Contains(js, "function requireLabAuth") {
		t.Fatal("requireLabAuth helper missing")
	}
	if strings.Contains(js, `$("#mapAuth").checked = true`) {
		t.Fatal("host chip must not auto-authorize the lab checkbox")
	}
	for _, needle := range []string{
		`authorized: true`,
		`$("#discAuth")`,
		`$("#fuzzAuth")`,
		`$("#scanAuth")`,
	} {
		if !strings.Contains(js, needle) {
			t.Fatalf("active tool wiring missing %s", needle)
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

func TestNoBrowserAlertDialogs(t *testing.T) {
	js := readWeb(t, "static/app.js")
	if strings.Contains(js, "alert(") {
		t.Fatal("app.js must not call alert(); GTK shows a blocking JavaScript dialog")
	}
	if !strings.Contains(js, "function uiFlash") {
		t.Fatal("missing uiFlash for inline errors")
	}
}

func TestExtensionsStubHidden(t *testing.T) {
	html := readWeb(t, "index.html")
	js := readWeb(t, "static/app.js")
	if !strings.Contains(html, `data-view="extensions"`) || !strings.Contains(html, "ext.unavailable") {
		t.Fatal("extensions tab should remain in DOM but show unavailable copy")
	}
	if strings.Contains(js, "btnExtLoad") {
		t.Fatal("fake extension loader must not remain")
	}
	if !strings.Contains(js, `name === "extensions"`) {
		t.Fatal("showView must skip the extensions stub")
	}
}

func TestOriginGuardPresent(t *testing.T) {
	src := readRepo(t, "internal/api/origin.go")
	if !strings.Contains(src, "origin not allowed") || !strings.Contains(src, "Sec-Fetch-Site") {
		t.Fatal("API origin guard missing")
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
