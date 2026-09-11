const $ = (sel) => document.querySelector(sel);
const $$ = (sel) => [...document.querySelectorAll(sel)];

const DebugLog = (() => {
  const q = [];
  let timer = null;
  let flushing = false;
  function enqueue(ev) {
    q.push({
      ts: new Date().toISOString(),
      src: "ui",
      level: ev.level || "info",
      module: ev.module || "ui",
      action: ev.action || "",
      msg: ev.msg || "",
      err: ev.err || "",
      fields: ev.fields || undefined,
    });
    if (q.length > 300) q.splice(0, q.length - 300);
    if (!timer) timer = setTimeout(flush, 800);
  }
  function flush() {
    timer = null;
    if (!q.length || flushing) return;
    const batch = q.splice(0, 80);
    flushing = true;
    fetch("/api/debug-log", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ events: batch }),
    }).catch(() => {}).finally(() => { flushing = false; });
  }
  window.addEventListener("error", (e) => {
    enqueue({
      level: "error",
      module: "js",
      action: "uncaught",
      msg: e.message || "error",
      err: e.error && e.error.stack ? String(e.error.stack) : "",
      fields: { file: e.filename, line: e.lineno, col: e.colno },
    });
  });
  window.addEventListener("unhandledrejection", (e) => {
    const reason = e.reason;
    enqueue({
      level: "error",
      module: "js",
      action: "unhandledrejection",
      msg: reason && reason.message ? reason.message : String(reason),
      err: reason && reason.stack ? String(reason.stack) : "",
    });
  });
  enqueue({ level: "info", module: "ui", action: "boot", msg: "ui loaded", fields: { href: location.href, lang: document.documentElement.getAttribute("data-lang") } });
  return { log: enqueue, flush };
})();

function isAbortError(err) {
  return !!(err && (err.name === "AbortError" || /aborted|abort/i.test(String(err.message || ""))));
}

function setWork(msg) {
  const bar = $("#workProgress");
  const tx = $("#workProgressText");
  if (!bar) return;
  if (!msg) {
    bar.classList.add("hidden");
    if (tx) tx.textContent = "";
    return;
  }
  bar.classList.remove("hidden");
  if (tx) tx.textContent = msg;
}

let flashTimer = null;
function uiFlash(msg, kind) {
  const el = $("#uiFlash");
  const text = String(msg || "");
  if (!text) return;
  if (!el) {
    console.warn(text);
    return;
  }
  el.textContent = text;
  el.classList.toggle("ok", kind === "ok");
  el.classList.remove("hidden");
  clearTimeout(flashTimer);
  flashTimer = setTimeout(() => el.classList.add("hidden"), kind === "ok" ? 4000 : 8000);
}

function markBusy(btn, on, label) {
  if (!btn) return;
  if (on) {
    if (!btn.dataset.label) btn.dataset.label = btn.textContent.trim();
    btn.disabled = true;
    btn.classList.add("is-busy");
    btn.innerHTML = `<span class="spin" aria-hidden="true"></span>${label || tr("work.running")}`;
    return;
  }
  btn.disabled = false;
  btn.classList.remove("is-busy");
  btn.textContent = btn.dataset.label || label || btn.textContent;
  delete btn.dataset.label;
}

const Jobs = {
  ctrls: new Map(),
  labels: new Map(),
  start(id, label) {
    this.stop(id);
    const c = new AbortController();
    this.ctrls.set(id, c);
    if (label) this.labels.set(id, label);
    this.syncBar();
    return c.signal;
  },
  stop(id) {
    const c = this.ctrls.get(id);
    if (!c) return false;
    c.abort();
    this.ctrls.delete(id);
    this.labels.delete(id);
    this.syncBar();
    return true;
  },
  stopAll() {
    for (const id of [...this.ctrls.keys()]) this.stop(id);
  },
  finish(id) {
    this.ctrls.delete(id);
    this.labels.delete(id);
    this.syncBar();
  },
  running(id) { return this.ctrls.has(id); },
  prefix(p) { return [...this.ctrls.keys()].filter((k) => k.startsWith(p)); },
  syncBar() {
    const labels = [...this.labels.values()];
    if (!labels.length) { setWork(""); return; }
    const last = (typeof MapLog !== "undefined" && MapLog.latest()) || "";
    const mapN = this.prefix("map:").length;
    if (last && mapN && mapN === labels.length) { setWork(last); return; }
    if (labels.length === 1) { setWork(labels[0]); return; }
    if (mapN && mapN === labels.length) { setWork(tr("map.nRunning", { n: mapN })); return; }
    setWork(tr("work.nJobs", { n: labels.length }));
  },
};

async function withWork(msg, btn, fn) {
  setWork(msg);
  if (btn) markBusy(btn, true, msg);
  try { return await fn(); }
  finally {
    if (btn) markBusy(btn, false);
    Jobs.syncBar();
  }
}

async function runJob(id, { label, runBtn, stopBtn, meta, fn }) {
  if (Jobs.running(id)) return { aborted: true };
  const signal = Jobs.start(id, label);
  if (stopBtn) stopBtn.disabled = false;
  if (runBtn) markBusy(runBtn, true, label);
  if (meta) meta.textContent = label;
  try {
    return await fn(signal);
  } catch (err) {
    if (isAbortError(err)) {
      if (meta) meta.textContent = tr("work.stopped");
      return { aborted: true };
    }
    throw err;
  } finally {
    Jobs.finish(id);
    if (runBtn) markBusy(runBtn, false);
    if (stopBtn) stopBtn.disabled = true;
  }
}

function bindJobStop(btn, id) {
  btn?.addEventListener("click", () => Jobs.stop(id));
}
$("#btnWorkStop")?.addEventListener("click", () => Jobs.stopAll());
bindJobStop($("#btnDiscStop"), "discover");
bindJobStop($("#btnFuzzStop"), "fuzz");
bindJobStop($("#btnScanStop"), "scanner");
bindJobStop($("#btnIntrStop"), "intruder");
bindJobStop($("#btnRepStop"), "repeater");

function applyTheme(name) {
  const t = name === "light" ? "light" : "dark";
  document.documentElement.setAttribute("data-theme", t);
  try { localStorage.setItem("aura_theme", t); } catch (_) {}
  $$("#themeSwitch button, #setThemeSwitch button").forEach((b) => b.classList.toggle("active", b.dataset.theme === t));
  if (UILayout && UILayout.data) {
    UILayout.data.theme = t;
    UILayout.save();
  }
}

function migrateNavOrder(order, defaults) {
  const def = (defaults && defaults.length) ? defaults.slice() : ["map", "scope", "proxy"];
  const seen = new Set();
  const out = [];
  for (let id of order || []) {
    if (id === "extensions") continue;
    if (id === "target") id = "scope";
    if (!id || seen.has(id)) continue;
    seen.add(id);
    out.push(id);
  }
  for (const id of def) {
    if (!seen.has(id)) out.push(id);
  }
  return out;
}

const UILayout = {
  key: "aura_ui_v3",
  defaults: {
    theme: "dark",
    uiScale: 100,
    editorScale: 100,
    showTag: true,
    lastView: "map",
    lastProxySub: "intercept",
    lastMapSub: "sitemap",
    uiLocked: false,
    tabOrder: ["map", "scope", "proxy", "discover", "fuzz", "intruder", "repeater", "scanner", "logger", "decoder", "comparer", "sequencer", "collaborator", "organizer", "other"],
    splits: {
      "proxy-intercept": [22, 78],
      "proxy-history": [46, 54],
      "map-layout": [54, 46],
      "map-app": [58, 42],
      "target-sitemap": [34, 66],
      "rep-layout": [38, 38, 24],
      "codec-grid": [42, 16, 42],
      "comparer-main": [58, 42],
      "comparer-grid": [50, 50],
      "seq-layout": [46, 54],
      "disc-layout": [34, 66],
      "fuzz-layout": [34, 66],
      "scan-layout": [36, 64],
      "org-layout": [22, 78],
      "org-main": [36, 64],
      "org-detail": [50, 50],
      "logger-main": [48, 52],
      "intr-main": [42, 22, 36],
      "issues-layout": [48, 52],
    },
  },
  data: null,
  load() {
    let saved = {};
    try { saved = JSON.parse(localStorage.getItem(this.key) || "{}"); } catch (_) { saved = {}; }
    if (!saved.theme) {
      try { saved.theme = localStorage.getItem("aura_theme") || localStorage.getItem("meb_theme") || "dark"; } catch (_) { saved.theme = "dark"; }
    }
    this.data = {
      ...this.defaults,
      ...saved,
      lastView: saved.lastView === "target" ? "map" : (saved.lastView || this.defaults.lastView),
      lastMapSub: saved.lastMapSub === "target" ? "sitemap" : (saved.lastMapSub || this.defaults.lastMapSub),
      splits: { ...this.defaults.splits, ...(saved.splits || {}) },
      tabOrder: migrateNavOrder(Array.isArray(saved.tabOrder) ? saved.tabOrder : [], this.defaults.tabOrder),
      uiLocked: !!saved.uiLocked,
    };
  },
  save() {
    try { localStorage.setItem(this.key, JSON.stringify(this.data)); } catch (_) {}
  },
  applyChrome() {
    const ui = Number(this.data.uiScale) || 100;
    const ed = Number(this.data.editorScale) || 100;
    document.documentElement.style.setProperty("--ui-scale", String(ui / 100));
    document.documentElement.style.setProperty("--editor-scale", String(ed / 100));
    document.documentElement.setAttribute("data-hide-tag", this.data.showTag === false ? "1" : "0");
    const uiEl = $("#setUiScale"), edEl = $("#setEditorScale"), tagEl = $("#setShowTag");
    if (uiEl) uiEl.value = String(ui);
    if (edEl) edEl.value = String(ed);
    if (tagEl) tagEl.checked = this.data.showTag !== false;
    if ($("#setUiScaleVal")) $("#setUiScaleVal").textContent = ui + "%";
    if ($("#setEditorScaleVal")) $("#setEditorScaleVal").textContent = ed + "%";
    document.documentElement.setAttribute("data-ui-locked", this.data.uiLocked ? "1" : "0");
    this.syncLockButton();
    const nav = $("#mainTabs");
    if (nav) nav.title = this.data.uiLocked ? "" : tr("set.dragTabs");
  },
  syncLockButton() {
    const btn = $("#btnLockUi");
    if (!btn) return;
    const locked = !!this.data.uiLocked;
    btn.dataset.i18n = locked ? "set.unlockUi" : "set.lockUi";
    btn.textContent = tr(locked ? "set.unlockUi" : "set.lockUi");
    btn.classList.toggle("primary", !locked);
    if ($("#setLockMsg")) $("#setLockMsg").textContent = locked ? tr("set.lockOn") : "";
  },
  setLocked(on) {
    this.data.uiLocked = !!on;
    this.save();
    this.applyChrome();
  },
  panes(el) {
    return [...el.children].filter((c) => !c.classList.contains("gutter"));
  },
  applySplit(el) {
    const id = el.dataset.split;
    const dir = el.dataset.splitDir === "v" ? "v" : "h";
    const sizes = this.data.splits[id] || this.defaults.splits[id];
    const panes = this.panes(el);
    if (!panes.length) return;
    const tracks = [];
    panes.forEach((p, i) => {
      const n = Math.max(Number(sizes && sizes[i] != null ? sizes[i] : (100 / panes.length)), 8);
      tracks.push(`minmax(88px, ${n}fr)`);
      if (i < panes.length - 1) tracks.push("10px");
      p.style.flex = "";
      p.style.minWidth = "0";
      p.style.minHeight = "0";
      p.style.removeProperty("overflow");
    });
    el.style.display = "grid";
    el.style.flex = "1 1 auto";
    el.style.minHeight = "0";
    el.style.minWidth = "0";
    if (dir === "v") {
      el.style.gridTemplateRows = tracks.join(" ");
      el.style.gridTemplateColumns = "minmax(0, 1fr)";
      el.style.flexDirection = "";
    } else {
      el.style.gridTemplateColumns = tracks.join(" ");
      el.style.gridTemplateRows = "minmax(0, 1fr)";
      el.style.flexDirection = "";
    }
  },
  ensureGutters(el) {
    if (el.dataset.splitReady === "1") {
      this.applySplit(el);
      return;
    }
    const dir = el.dataset.splitDir || "h";
    const kids = this.panes(el);
    kids.forEach((pane, i) => {
      if (i === 0) return;
      const g = document.createElement("div");
      g.className = "gutter";
      g.dataset.gutter = String(i - 1);
      g.title = tr("set.layoutHint");
      el.insertBefore(g, pane);
      this.bindGutter(el, g, i - 1, dir);
    });
    el.dataset.splitReady = "1";
    this.applySplit(el);
  },
  bindGutter(el, gutter, index, dir) {
    const minPx = 80;
    const onDown = (e) => {
      if (this.data.uiLocked) return;
      if (e.button != null && e.button !== 0) return;
      e.preventDefault();
      if (e.pointerId != null && gutter.setPointerCapture) {
        try { gutter.setPointerCapture(e.pointerId); } catch (_) {}
      }
      const panes = this.panes(el);
      const a = panes[index], b = panes[index + 1];
      if (!a || !b) return;
      const start = dir === "v" ? e.clientY : e.clientX;
      const a0 = dir === "v" ? a.getBoundingClientRect().height : a.getBoundingClientRect().width;
      const b0 = dir === "v" ? b.getBoundingClientRect().height : b.getBoundingClientRect().width;
      const total = a0 + b0;
      const dims0 = panes.map((p) => {
        const r = p.getBoundingClientRect();
        return dir === "v" ? r.height : r.width;
      });
      const pair = a0 + b0;
      gutter.classList.add("is-drag");
      document.body.classList.add(dir === "v" ? "is-resizing-v" : "is-resizing-h");
      const move = (ev) => {
        const now = dir === "v" ? ev.clientY : ev.clientX;
        let na = a0 + (now - start);
        let nb = pair - na;
        if (na < minPx) { na = minPx; nb = pair - minPx; }
        if (nb < minPx) { nb = minPx; na = pair - minPx; }
        const dims = dims0.slice();
        dims[index] = na;
        dims[index + 1] = nb;
        const sum = dims.reduce((x, y) => x + y, 0) || 1;
        this.data.splits[el.dataset.split] = dims.map((s) => Math.round((s / sum) * 1000) / 10);
        this.applySplit(el);
      };
      const up = () => {
        gutter.classList.remove("is-drag");
        document.body.classList.remove("is-resizing-v", "is-resizing-h");
        document.removeEventListener("pointermove", move);
        document.removeEventListener("pointerup", up);
        document.removeEventListener("pointercancel", up);
        const panes2 = this.panes(el);
        const dims = panes2.map((p) => {
          const r = p.getBoundingClientRect();
          return dir === "v" ? r.height : r.width;
        });
        const sum = dims.reduce((x, y) => x + y, 0) || 1;
        this.data.splits[el.dataset.split] = dims.map((s) => Math.round((s / sum) * 1000) / 10);
        this.save();
      };
      document.addEventListener("pointermove", move);
      document.addEventListener("pointerup", up);
      document.addEventListener("pointercancel", up);
    };
    gutter.addEventListener("pointerdown", onDown);
    gutter.addEventListener("dblclick", () => {
      if (this.data.uiLocked) return;
      const def = this.defaults.splits[el.dataset.split];
      if (def) this.data.splits[el.dataset.split] = def.slice();
      this.applySplit(el);
      this.save();
    });
  },
  refresh() {
    $$("[data-split]").forEach((el) => this.ensureGutters(el));
  },
  resetLayout() {
    this.data.splits = { ...this.defaults.splits };
    this.data.tabOrder = this.defaults.tabOrder.slice();
    this.save();
    this.refresh();
    NavTabs.applyOrder();
  },
  resetVisual() {
    this.data.theme = this.defaults.theme;
    this.data.uiScale = this.defaults.uiScale;
    this.data.editorScale = this.defaults.editorScale;
    this.data.showTag = true;
    this.data.uiLocked = false;
    this.data.tabOrder = this.defaults.tabOrder.slice();
    this.data.splits = { ...this.defaults.splits };
    this.save();
    applyTheme("dark");
    this.applyChrome();
    this.refresh();
    NavTabs.applyOrder();
  },
  open(tab) {
    const ov = $("#settingsOverlay");
    if (!ov) return;
    ov.classList.remove("hidden");
    const name = tab || "look";
    $$("#settingsNav button").forEach((b) => b.classList.toggle("active", b.dataset.settab === name));
    $$(".settings-pane").forEach((p) => p.classList.toggle("hidden", p.dataset.setpane !== name));
    this.applyChrome();
    $$("#setLangSwitch button").forEach((b) => b.classList.toggle("active", b.dataset.lang === mebLang));
    $$("#setThemeSwitch button").forEach((b) => b.classList.toggle("active", (document.documentElement.getAttribute("data-theme") || "dark") === b.dataset.theme));
  },
  close() {
    $("#settingsOverlay")?.classList.add("hidden");
  },
  init() {
    this.load();
    this.applyChrome();
    applyTheme(this.data.theme || localStorage.getItem("aura_theme") || "dark");
    this.refresh();
    $("#btnSettings")?.addEventListener("click", () => this.open("look"));
    $("#btnSettingsClose")?.addEventListener("click", () => this.close());
    $("#btnProxyOpenSettings")?.addEventListener("click", () => this.open("proxy"));
    $("#settingsOverlay")?.addEventListener("click", (e) => {
      if (e.target === $("#settingsOverlay")) this.close();
    });
    document.addEventListener("keydown", (e) => {
      if (e.key === "Escape" && !$("#settingsOverlay")?.classList.contains("hidden")) this.close();
    });
    $$("#settingsNav button").forEach((b) => b.addEventListener("click", () => this.open(b.dataset.settab)));
    $$("#setThemeSwitch button").forEach((b) => b.addEventListener("click", () => applyTheme(b.dataset.theme)));
    $$("#setLangSwitch button").forEach((b) => b.addEventListener("click", () => applyLang(b.dataset.lang)));
    $("#setUiScale")?.addEventListener("input", () => {
      this.data.uiScale = Number($("#setUiScale").value);
      this.applyChrome();
    });
    $("#setUiScale")?.addEventListener("change", () => this.save());
    $("#setEditorScale")?.addEventListener("input", () => {
      this.data.editorScale = Number($("#setEditorScale").value);
      this.applyChrome();
    });
    $("#setEditorScale")?.addEventListener("change", () => this.save());
    $("#setShowTag")?.addEventListener("change", () => {
      this.data.showTag = $("#setShowTag").checked;
      this.applyChrome();
      this.save();
    });
    $("#btnResetVisual")?.addEventListener("click", () => {
      this.resetVisual();
      if ($("#setVisualMsg")) $("#setVisualMsg").textContent = tr("set.resetDone");
    });
    $("#btnResetLayout")?.addEventListener("click", () => {
      this.resetLayout();
      if ($("#setLayoutMsg")) $("#setLayoutMsg").textContent = tr("set.layoutReset");
    });
    $("#btnLockUi")?.addEventListener("click", () => {
      this.setLocked(!this.data.uiLocked);
      if ($("#setLockMsg")) $("#setLockMsg").textContent = tr(this.data.uiLocked ? "set.lockOn" : "set.lockOff");
    });
    NavTabs.init();
  },
};

const NavTabs = {
  suppressClick: false,
  init() {
    this.applyOrder();
    this.bind();
  },
  buttons() {
    return $$("#mainTabs button[data-view]");
  },
  applyOrder() {
    const nav = $("#mainTabs");
    if (!nav || !UILayout.data) return;
    const map = new Map(this.buttons().map((b) => [b.dataset.view, b]));
    const order = UILayout.data.tabOrder || UILayout.defaults.tabOrder;
    const seen = new Set();
    for (const id of order) {
      if (id === "extensions") continue;
      const b = map.get(id);
      if (!b) continue;
      nav.appendChild(b);
      seen.add(id);
    }
    for (const [id, b] of map) {
      if (id === "extensions" || seen.has(id)) continue;
      nav.appendChild(b);
    }
    const ext = map.get("extensions");
    if (ext) nav.appendChild(ext);
    nav.title = UILayout.data.uiLocked ? "" : tr("set.dragTabs");
  },
  saveOrder() {
    if (!UILayout.data) return;
    UILayout.data.tabOrder = this.buttons()
      .map((b) => b.dataset.view)
      .filter((id) => id && id !== "extensions");
    UILayout.save();
  },
  bind() {
    const nav = $("#mainTabs");
    if (!nav || nav.dataset.navReady === "1") return;
    nav.dataset.navReady = "1";
    nav.addEventListener("click", (e) => {
      if (this.suppressClick) {
        e.preventDefault();
        e.stopImmediatePropagation();
        return;
      }
      const b = e.target.closest("button[data-view]");
      if (!b || b.hidden || b.classList.contains("hidden")) return;
      showView(b.dataset.view);
    });
    nav.addEventListener("pointerdown", (e) => {
      if (UILayout.data && UILayout.data.uiLocked) return;
      if (e.button != null && e.button !== 0) return;
      const btn = e.target.closest("#mainTabs button[data-view]");
      if (!btn || btn.hidden || btn.dataset.view === "extensions") return;
      this.startDrag(btn, e);
    });
  },
  startDrag(btn, e) {
    const nav = $("#mainTabs");
    const startX = e.clientX;
    const startY = e.clientY;
    let dragging = false;
    const onMove = (ev) => {
      if (!dragging) {
        if (Math.abs(ev.clientX - startX) < 8 && Math.abs(ev.clientY - startY) < 8) return;
        dragging = true;
        btn.classList.add("is-dragging");
        document.body.classList.add("is-dragging-tabs");
        if (ev.pointerId != null && btn.setPointerCapture) {
          try { btn.setPointerCapture(ev.pointerId); } catch (_) {}
        }
      }
      btn.style.pointerEvents = "none";
      const under = document.elementFromPoint(ev.clientX, ev.clientY);
      btn.style.pointerEvents = "";
      const over = under && under.closest("#mainTabs button[data-view]");
      if (!over || over === btn || over.dataset.view === "extensions") return;
      const rect = over.getBoundingClientRect();
      if (ev.clientX > rect.left + rect.width / 2) {
        nav.insertBefore(btn, over.nextSibling);
      } else {
        nav.insertBefore(btn, over);
      }
      const ext = nav.querySelector('[data-view="extensions"]');
      if (ext) nav.appendChild(ext);
    };
    const onUp = () => {
      document.removeEventListener("pointermove", onMove);
      document.removeEventListener("pointerup", onUp);
      document.removeEventListener("pointercancel", onUp);
      btn.classList.remove("is-dragging");
      document.body.classList.remove("is-dragging-tabs");
      if (!dragging) return;
      this.suppressClick = true;
      setTimeout(() => { this.suppressClick = false; }, 0);
      this.saveOrder();
    };
    document.addEventListener("pointermove", onMove);
    document.addEventListener("pointerup", onUp);
    document.addEventListener("pointercancel", onUp);
  },
};

function applyLang(name) {
  setLang(name);
  applyI18nDom();
  $$("#langSwitch button, #setLangSwitch button").forEach((b) => b.classList.toggle("active", b.dataset.lang === mebLang));
  refreshTranslatedUI();
}

const state = { pending: [], selectedFlow: null, histReq: "", histResp: "", histScheme: "https", intrResults: [] };

async function api(path, opts = {}) {
  const method = (opts.method || "GET").toUpperCase();
  const t0 = performance.now();
  let logged = false;
  try {
    const res = await fetch(path, { headers: { "Content-Type": "application/json", ...(opts.headers || {}) }, ...opts });
    const ms = Math.round(performance.now() - t0);
    if (!res.ok) {
      let msg = res.statusText;
      try { const d = await res.json(); msg = d.detail || JSON.stringify(d); } catch (_) {}
      DebugLog.log({ level: "error", module: "api", action: method + " " + path, msg, fields: { status: res.status, ms } });
      logged = true;
      throw new Error(msg);
    }
    if (method !== "GET" && !path.startsWith("/api/intercept") && path !== "/api/debug-log") {
      DebugLog.log({ level: "info", module: "api", action: method + " " + path, msg: String(res.status), fields: { status: res.status, ms } });
    }
    const ct = res.headers.get("content-type") || "";
    return ct.includes("application/json") ? res.json() : res.text();
  } catch (err) {
    if (isAbortError(err)) {
      DebugLog.log({ level: "info", module: "api", action: method + " " + path, msg: "stopped" });
      throw err;
    }
    if (!logged) {
      DebugLog.log({ level: "error", module: "api", action: method + " " + path, msg: err && err.message ? err.message : String(err) });
    }
    throw err;
  }
}

function showView(name) {
  if (name === "extensions" || name === "target") name = "map";
  DebugLog.log({ level: "info", module: "nav", action: "view", msg: name, fields: { view: name } });
  $$(".view").forEach((el) => el.classList.add("hidden"));
  const pane = $(`#view-${name}`);
  if (pane) pane.classList.remove("hidden");
  $$("#mainTabs button").forEach((b) => b.classList.toggle("active", b.dataset.view === name));
  if (UILayout && UILayout.data) {
    UILayout.data.lastView = name;
    UILayout.save();
    requestAnimationFrame(() => UILayout.refresh());
  }
  if (name === "intruder" && typeof intruder !== "undefined" && intruder.tabs.length === 0) newIntruderTab("");
  if (name === "discover" || name === "fuzz") loadSecLists();
  if (name === "map") {
    loadSiteMap();
    loadMap();
    let sub = UILayout.data && UILayout.data.lastMapSub;
    if (sub === "target") sub = "sitemap";
    showSub("#view-map", sub || "sitemap");
  }
  if (name === "dashboard") refreshDashboard();
  if (name === "scope") { renderScope(); }
  if (name === "other") { loadIssues(); renderIssueDefs(); }
  if (name === "proxy") { renderMr(); }
  if (name === "organizer") { loadOrganizer(); }
  if (name === "logger") { loadLogs(); }
  if (name === "repeater" && repeater.tabs.length === 0) { newRepeaterTab(""); }
  if (name === "scanner" && state.histReq) { if ($("#scanFromHistory").checked) $("#scanRaw").value = state.histReq; }
}
function showSub(group, name) {
  $$(`${group} .subpane`).forEach((el) => el.classList.toggle("hidden", el.dataset.subpane !== name));
  $$(`${group} .subtabs button`).forEach((b) => b.classList.toggle("active", b.dataset.sub === name));
  if (group === "#view-proxy" && UILayout && UILayout.data) {
    UILayout.data.lastProxySub = name;
    UILayout.save();
  }
  if (group === "#view-map" && UILayout && UILayout.data) {
    UILayout.data.lastMapSub = name;
    UILayout.save();
  }
  requestAnimationFrame(() => UILayout && UILayout.refresh());
}

const fmtBytes = (n) => n < 1024 ? `${n}B` : n < 1048576 ? `${(n/1024).toFixed(1)}k` : `${(n/1048576).toFixed(1)}M`;
const esc = (s) => String(s ?? "").replaceAll("&","&amp;").replaceAll("<","&lt;").replaceAll(">","&gt;");
const statusClass = (c) => !c ? "" : c >= 400 ? "status-red" : c >= 300 ? "status-3" : "status-ok";

let lastStatus = null;
function renderStatus(s) {
  lastStatus = s;
  const badge = $("#proxyBadge");
  badge.textContent = s.proxy_running ? `proxy ${s.settings.listen_host}:${s.settings.listen_port}` : "proxy off";
  badge.className = `badge ${s.proxy_running ? "on" : "off"}`;
  $("#interceptReq").checked = s.settings.intercept_requests;
  $("#interceptResp").checked = s.settings.intercept_responses;
  $("#listenHost").value = s.settings.listen_host;
  $("#listenPort").value = s.settings.listen_port;
  $("#filter").value = s.settings.intercept_filter;
  $("#verifyUp").checked = s.settings.verify_upstream;
  $("#recordHist").checked = s.settings.record_history;
  refreshDashboard();
}

function refreshDashboard() {
  if (!lastStatus) return;
  const s = lastStatus, st = s.stats || {};
  $("#dStatProxy").textContent = s.proxy_running ? "on" : "off";
  $("#dStatProxy").style.color = s.proxy_running ? "var(--ok)" : "var(--muted)";
  $("#dStatFlows").textContent = st.flows || 0;
  $("#dStatIntercepted").textContent = st.intercepted || 0;
  $("#dStatErrors").textContent = st.errors || 0;
  $("#dStatsTraffic").innerHTML = `<span>flows ${st.flows||0}</span><span>intercepted ${st.intercepted||0}</span><span>in ${fmtBytes(st.bytes_in||0)}</span><span>out ${fmtBytes(st.bytes_out||0)}</span><span>errors ${st.errors||0}</span>`;
  loadHistory().then((items) => {
    $("#dRecent").innerHTML = items.slice(0,12).map((f) => `<tr data-id="${f.id}">
      <td class="method ${f.method}">${esc(f.method)}</td>
      <td>${esc(f.host)}</td>
      <td title="${esc(f.path)}">${esc(f.path).slice(0,60)}</td>
      <td class="${statusClass(f.status)}">${f.status||"—"}</td>
      <td><button class="mini" data-act="rep" data-id="${f.id}">${tr("dash.toReplay")}</button> <button class="mini" data-act="intr" data-id="${f.id}">${tr("dash.toPayloads")}</button></td>
    </tr>`).join("") || `<tr><td colspan="5" class="muted">${tr("empty.none")}</td></tr>`;
  });
}

function renderPending(force = false) {
  const n = state.pending.length;
  $("#pendingBadge").classList.toggle("hidden", n === 0);
  $("#pendingBadge").textContent = tr("proxy.inQueue", { n });
  $("#proxyDot").classList.toggle("hidden", n === 0);
  $("#interceptDot").classList.toggle("hidden", n === 0);
  const list = $("#interceptList");
  list.innerHTML = n ? `<table><thead><tr><th>${tr("proxy.col.time")}</th><th>${tr("proxy.col.type")}</th><th>${tr("proxy.col.dir")}</th><th>${tr("proxy.col.method")}</th><th>${tr("proxy.col.url")}</th></tr></thead><tbody>` +
    state.pending.map((p, i) => `<tr data-i="${i}" class="${i === 0 ? "sel" : ""}">
      <td>${new Date((p.flow_id || 0) * 1000 || Date.now()).toLocaleTimeString()}</td>
      <td>HTTP</td>
      <td>${p.phase === "response" ? tr("proxy.dir.resp") : tr("proxy.dir.req")}</td>
      <td class="method ${p.method}">${esc(p.method)}</td>
      <td>${esc(p.url)}</td>
    </tr>`).join("") + `</tbody></table>` : `<div class="muted" style="padding:12px">${tr("empty.queue")}</div>`;
  const cur = state.pending[0];
  if (!cur) { $("#interceptEditor").value = ""; $("#interceptMeta").textContent = ""; return; }
  const phase = cur.phase === "response" ? tr("proxy.phase.resp") : tr("proxy.phase.req");
  $("#interceptMeta").textContent = `${phase}  ${cur.method}  ${cur.url}`;
  const raw = cur.phase === "response" ? cur.response_raw : cur.request_raw;
  if (force || document.activeElement !== $("#interceptEditor")) $("#interceptEditor").value = raw || "";
}

function renderHistory(items) {
  const body = $("#histBody");
  const filtered = typeof applyHistFilter === "function" ? items.filter(applyHistFilter) : items;
  body.innerHTML = filtered.map((f,i) => `<tr data-id="${f.id}"><td>${filtered.length-i}</td><td class="method ${f.method}">${esc(f.method)}</td><td>${esc(f.host)}</td><td title="${esc(f.path)}">${esc(f.path).slice(0,80)}</td><td class="${statusClass(f.status)}">${f.status||f.error||"—"}</td><td>${fmtBytes(f.resp_bytes||0)}</td><td>${f.duration_ms??"—"}</td></tr>`).join("");
  if (state.selectedFlow) { const r = body.querySelector(`[data-id="${state.selectedFlow}"]`); if (r) r.classList.add("sel"); }
}

function renderCallback(items) {
  const el = $("#callbackList");
  if (!items.length) { el.innerHTML = `<p class="muted">${tr("callback.none", { id: esc($("#callbackId").value) })}</p>`; return; }
  el.innerHTML = items.map((it) => `<div class="callback-item"><div class="ci-meta">${new Date(it.timestamp).toLocaleString()} — ${esc(it.method)} ${esc(it.url)} from ${esc(it.remote_addr)}</div><div class="ci-meta">Headers: ${Object.entries(it.headers||{}).map(([k,v])=>`${esc(k)}: ${esc(v)}`).join(", ")}</div><div class="ci-body">${esc(it.body)}</div></div>`).join("");
}

async function loadStatus() { renderStatus(await api("/api/status")); }
async function loadHistory() {
  const q = encodeURIComponent($("#histQ")?.value || "");
  const items = (await api(`/api/history?q=${q}`)).items;
  renderHistory(items); return items;
}
async function loadIntercept() { state.pending = (await api("/api/intercept")).items; renderPending(); }
async function loadLogs() {
  const data = await api("/api/history?limit=1000");
  const items = data.items || [];
  const src = $("#logSource")?.value || "";
  const q = ($("#logFilter")?.value || "").toLowerCase();
  const filtered = items.filter((f) => {
    if (src && f.source !== src) return false;
    if (q && !(`${f.source||""} ${f.host||""} ${f.path||""} ${f.method||""} ${f.status||""}`.toLowerCase().includes(q))) return false;
    return true;
  });
  $("#logFlows").innerHTML = filtered.map((f, i) => `<tr data-id="${f.id}">
    <td>${filtered.length - i}</td>
    <td class="src-${f.source||"proxy"}">${esc(f.source||"proxy")}</td>
    <td class="muted">${fmtTime(f.created)}</td>
    <td class="method ${f.method}">${esc(f.method)}</td>
    <td>${esc(f.host)}</td>
    <td title="${esc(f.path)}">${esc((f.path||"").slice(0,70))}</td>
    <td class="${statusClass(f.status)}">${f.status||"—"}</td>
    <td>${fmtBytes(f.resp_bytes||0)}</td>
    <td>${f.duration_ms??"—"}</td>
  </tr>`).join("") || `<tr><td colspan="9" class="muted">${tr("empty.requests")}</td></tr>`;
  $("#logMeta").textContent = `${filtered.length} / ${items.length}`;
}
$("#logFilter").addEventListener("input", loadLogs);
$("#logSource").addEventListener("change", loadLogs);
$("#btnLogRefresh").addEventListener("click", loadLogs);
$("#logFlows").addEventListener("click", async (e) => {
  const tr = e.target.closest("tr[data-id]"); if (!tr) return;
  $$("#logFlows tr").forEach((r) => r.classList.remove("sel")); tr.classList.add("sel");
  const f = await api(`/api/history/${tr.dataset.id}`);
  $("#logReq").textContent = f.request_raw || "";
  $("#logResp").textContent = f.response_raw || f.error || "";
});
$$("#view-logger .tabs.small button").forEach((b) => b.addEventListener("click", () => {
  $$("#view-logger .tabs.small button").forEach((x) => x.classList.remove("active")); b.classList.add("active");
  $("#logReq").classList.toggle("hidden", b.dataset.pane !== "req");
  $("#logResp").classList.toggle("hidden", b.dataset.pane === "req");
}));

function connectWs() {
  const proto = location.protocol === "https:" ? "wss" : "ws";
  const ws = new WebSocket(`${proto}://${location.host}/api/ws`);
  ws.onmessage = (ev) => {
    const m = JSON.parse(ev.data);
    if (m.type === "intel_progress") MapLog.append(m);
    if (m.type === "hello") renderStatus({ ...m.status, stats: m.status.stats });
    if (m.type === "flow") { loadHistory(); refreshDashboard(); loadSiteMap(); scheduleMapIngest(); if (currentView() === "logger") loadLogs(); }
    if (m.type === "intercept") { state.pending.push(m.item); renderPending(); }
    if (m.type === "intercept_done") { state.pending = state.pending.filter((p) => p.id !== m.id); renderPending(true); }
    if (m.type === "log") loadLogs();
    if (m.type === "history_cleared") loadHistory();
    if (m.type === "settings") loadStatus();
  };
  ws.onerror = () => DebugLog.log({ level: "error", module: "ws", action: "error", msg: "websocket error" });
  ws.onclose = () => {
    DebugLog.log({ level: "warn", module: "ws", action: "close", msg: "websocket closed" });
    setTimeout(connectWs, 1500);
  };
}

$$("#proxySubtabs button").forEach((b) => b.addEventListener("click", () => {
  showSub("#view-proxy", b.dataset.sub);
  if (b.dataset.sub === "settings") UILayout.open("proxy");
}));
$$("#mapSubtabs button").forEach((b) => b.addEventListener("click", () => showSub("#view-map", b.dataset.sub)));

$("#interceptReq").addEventListener("change", saveSettings);
$("#interceptResp").addEventListener("change", saveSettings);
$("#verifyUp").addEventListener("change", saveSettings);
$("#recordHist").addEventListener("change", saveSettings);
$("#filter").addEventListener("change", saveSettings);

async function saveSettings(extra = {}) {
  const body = { listen_host: $("#listenHost").value.trim(), listen_port: Number($("#listenPort").value), intercept_requests: $("#interceptReq").checked, intercept_responses: $("#interceptResp").checked, intercept_filter: $("#filter").value, verify_upstream: $("#verifyUp").checked, record_history: $("#recordHist").checked, ...extra };
  await api("/api/settings", { method: "PUT", body: JSON.stringify(body) });
  await loadStatus();
}

async function startProxy() { await saveSettings(); await api("/api/proxy/start", { method: "POST" }); await loadStatus(); }
async function stopProxy() { await api("/api/proxy/stop", { method: "POST" }); await loadStatus(); }
async function launchBrowser() {
  try { const r = await api("/api/proxy/launch-browser", { method: "POST" }); uiFlash(tr("browser.launched", { proxy: r.proxy }), "ok"); }
  catch (e) { uiFlash(e.message); }
}
$("#btnStart").addEventListener("click", startProxy);
$("#btnStop").addEventListener("click", stopProxy);
$("#btnLaunchBrowser").addEventListener("click", launchBrowser);

$("#btnForward").addEventListener("click", async () => {
  const c = state.pending[0]; if (!c) return;
  const id = c.id;
  if (!id) { uiFlash(tr("proxy.forwardErr", { msg: "no id" })); return; }
  let raw = $("#interceptEditor").value;
  try { raw = applyMatchReplace(raw, c.phase); } catch (e) { DebugLog.log({ level: "error", module: "proxy", action: "matchreplace", msg: e.message }); }
  try {
    await api(`/api/intercept/${encodeURIComponent(id)}`, { method: "POST", body: JSON.stringify({ action: "forward", raw }) });
  } catch (e) {
    DebugLog.log({ level: "error", module: "proxy", action: "forward", msg: e.message, fields: { id } });
    uiFlash(tr("proxy.forwardErr", { msg: e.message }));
  }
});
$("#btnForwardAll").addEventListener("click", async () => {
  try { await api("/api/intercept/forward-all", { method: "POST" }); }
  catch (e) { uiFlash(tr("proxy.forwardErr", { msg: e.message })); }
});
$("#btnDrop").addEventListener("click", async () => {
  const c = state.pending[0]; if (!c) return;
  try { await api(`/api/intercept/${encodeURIComponent(c.id)}`, { method: "POST", body: JSON.stringify({ action: "drop" }) }); }
  catch (e) { uiFlash(tr("proxy.forwardErr", { msg: e.message })); }
});
$("#btnDropAll").addEventListener("click", async () => {
  try { await api("/api/intercept/drop-all", { method: "POST" }); }
  catch (e) { uiFlash(tr("proxy.forwardErr", { msg: e.message })); }
});
$("#interceptList").addEventListener("click", (e) => {
  const tr = e.target.closest("tr"); if (!tr) return;
  const i = +tr.dataset.i;
  if (state.pending[i]) {
    state.pending = [state.pending[i], ...state.pending.filter((_, idx) => idx !== i)];
    renderPending(true);
  }
});
$("#btnWsOpenBrowser").addEventListener("click", launchBrowser);

/* Match and replace */
const mrRules = JSON.parse(localStorage.getItem("aura_mr_rules") || localStorage.getItem("meb_mr_rules") || "[]");
function saveMr() { localStorage.setItem("aura_mr_rules", JSON.stringify(mrRules)); }
function applyMatchReplace(raw, phase) {
  for (const r of mrRules) {
    if (!r.enabled) continue;
    if (r.item === "request_header" && phase !== "request") continue;
    if (r.item === "request_body" && phase !== "request") continue;
    if (r.item === "response_header" && phase !== "response") continue;
    if (r.item === "response_body" && phase !== "response") continue;
    try {
      if (r.type === "regex") {
        const re = new RegExp(r.match, "g");
        raw = raw.replace(re, r.replace || "");
      } else {
        raw = raw.split(r.match).join(r.replace || "");
      }
    } catch (_) {}
  }
  return raw;
}
function renderMr() {
  $("#mrBody").innerHTML = mrRules.map((r, i) => `<tr>
    <td><input type="checkbox" data-i="${i}" data-f="enabled" ${r.enabled ? "checked" : ""}/></td>
    <td><select data-i="${i}" data-f="item">
      <option value="request_header" ${r.item==="request_header"?"selected":""}>Request header</option>
      <option value="request_body" ${r.item==="request_body"?"selected":""}>Request body</option>
      <option value="response_header" ${r.item==="response_header"?"selected":""}>Response header</option>
      <option value="response_body" ${r.item==="response_body"?"selected":""}>Response body</option>
    </select></td>
    <td><input type="text" data-i="${i}" data-f="name" value="${esc(r.name)}" placeholder="name"/></td>
    <td><input type="text" data-i="${i}" data-f="match" value="${esc(r.match)}" placeholder="^User-Agent:.*$"/></td>
    <td><input type="text" data-i="${i}" data-f="replace" value="${esc(r.replace)}" placeholder="Mozilla/5.0 ..."/></td>
    <td><select data-i="${i}" data-f="type">
      <option value="literal" ${r.type==="literal"?"selected":""}>Literal</option>
      <option value="regex" ${r.type==="regex"?"selected":""}>Regex</option>
    </select></td>
    <td><input type="text" data-i="${i}" data-f="comment" value="${esc(r.comment)}" placeholder="comment"/></td>
  </tr>`).join("") || `<tr><td colspan="7" class="muted">${tr("empty.rules")}</td></tr>`;
}
$("#mrAdd").addEventListener("click", () => { mrRules.push({ enabled: true, item: "request_header", name: "", match: "", replace: "", type: "literal", comment: "" }); saveMr(); renderMr(); });
$("#mrRemove").addEventListener("click", () => { const s = selectedMrRow(); if (s >= 0) { mrRules.splice(s, 1); saveMr(); renderMr(); } });
$("#mrUp").addEventListener("click", () => { const s = selectedMrRow(); if (s > 0) { [mrRules[s-1], mrRules[s]] = [mrRules[s], mrRules[s-1]]; saveMr(); renderMr(); } });
$("#mrDown").addEventListener("click", () => { const s = selectedMrRow(); if (s >= 0 && s < mrRules.length-1) { [mrRules[s+1], mrRules[s]] = [mrRules[s], mrRules[s+1]]; saveMr(); renderMr(); } });
function selectedMrRow() { const s = document.querySelector("#mrBody tr.sel"); return s ? +s.dataset.i : -1; }
document.addEventListener("change", (e) => {
  const t = e.target;
  if (t.dataset && t.dataset.f && t.closest("#mrBody")) {
    const r = mrRules[+t.dataset.i]; if (!r) return;
    if (t.type === "checkbox") r[t.dataset.f] = t.checked; else r[t.dataset.f] = t.value;
    saveMr();
  }
});
document.addEventListener("click", (e) => {
  const tr = e.target.closest("#mrBody tr"); if (!tr) return;
  $$("#mrBody tr").forEach((r) => r.classList.remove("sel")); tr.classList.add("sel");
});

$("#btnReloadHist").addEventListener("click", loadHistory);
$("#histQ").addEventListener("input", () => { clearTimeout(window._histT); window._histT = setTimeout(loadHistory, 200); });
$("#btnClearHist").addEventListener("click", async () => { await api("/api/history", { method: "DELETE" }); $("#histReq").textContent=""; $("#histResp").textContent=""; state.selectedFlow=null; await loadHistory(); });
$("#histBody").addEventListener("click", async (e) => {
  const tr = e.target.closest("tr"); if (!tr) return;
  $$("#histBody tr").forEach((r) => r.classList.remove("sel")); tr.classList.add("sel");
  state.selectedFlow = tr.dataset.id;
  const data = await api(`/api/history/${tr.dataset.id}`);
  state.histReq = data.request_raw || ""; state.histResp = data.response_raw || data.error || ""; state.histScheme = data.scheme || "https";
  $("#histReq").textContent = state.histReq; $("#histResp").textContent = state.histResp;
});
$("#btnToRepeater").addEventListener("click", () => { if (!state.histReq) return; showView("repeater"); newRepeaterTab(state.histReq); if (state.histScheme) { const t = curRep(); if (t) t.scheme = state.histScheme; } });
$("#btnToIntruder").addEventListener("click", () => {
  if (!state.histReq) return;
  showView("intruder");
  newIntruderTab(state.histReq.replace(/\r\n/g, "\n"));
});
$("#btnToComparer").addEventListener("click", () => { if (!state.histReq) return; $("#cmpA").value = state.histReq; $("#cmpB").value = state.histResp; showView("comparer"); });
$("#btnToOrganizer").addEventListener("click", () => { if (!state.histReq) return; sendToOrganizer(state.histReq, state.histResp, "proxy"); });

$$("#view-proxy .tabs.small button").forEach((b) => b.addEventListener("click", () => {
  $$("#view-proxy .tabs.small button").forEach((x) => x.classList.remove("active")); b.classList.add("active");
  const req = b.dataset.pane === "req";
  $("#histReq").classList.toggle("hidden", !req); $("#histResp").classList.toggle("hidden", req);
}));

/* Repeater — multi-tab */
const repeater = { tabs: [], active: 0, seq: 0 };
function newRepeaterTab(raw) {
  repeater.seq++;
  const def = "GET / HTTP/1.1\r\nHost: example.com\r\nUser-Agent: AURA/0.1\r\nAccept: */*\r\n\r\n";
  const tab = { id: repeater.seq, name: String(repeater.seq), raw: raw || def, scheme: "https", target: "", response: "" };
  repeater.tabs.push(tab);
  repeater.active = repeater.tabs.length - 1;
  renderRepTabs(); renderRepTab();
}
function closeRepeaterTab(i) {
  repeater.tabs.splice(i, 1);
  if (repeater.active >= repeater.tabs.length) repeater.active = repeater.tabs.length - 1;
  if (repeater.tabs.length === 0) newRepeaterTab("");
  renderRepTabs(); renderRepTab();
}
function renderRepTabs() {
  $("#repTabList").innerHTML = repeater.tabs.map((t, i) => `<div class="rep-tab ${i === repeater.active ? "active" : ""}" data-i="${i}">
    <span>${esc(t.name)}</span><span class="close" data-close="${i}">✕</span>
  </div>`).join("");
}
function curRep() { return repeater.tabs[repeater.active]; }
function renderRepTab() {
  const t = curRep(); if (!t) return;
  $("#repReq").value = t.raw;
  $("#repScheme").value = t.scheme;
  $("#repTarget").value = t.target;
  $("#repResp").textContent = t.response;
  renderInspector($("#repInspector"), t.raw);
}
$("#repTabList").addEventListener("click", (e) => {
  const c = e.target.closest("[data-close]");
  if (c) { closeRepeaterTab(+c.dataset.close); return; }
  const t = e.target.closest(".rep-tab"); if (!t) return;
  repeater.active = +t.dataset.i;
  renderRepTabs(); renderRepTab();
});
$("#btnRepNewTab").addEventListener("click", () => newRepeaterTab(""));
$("#repReq").addEventListener("input", () => { const t = curRep(); if (t) { t.raw = $("#repReq").value; renderInspector($("#repInspector"), t.raw); } });
$("#repScheme").addEventListener("change", () => { const t = curRep(); if (t) t.scheme = $("#repScheme").value; });
$("#repTarget").addEventListener("input", () => { const t = curRep(); if (t) t.target = $("#repTarget").value; });

$("#btnSend").addEventListener("click", async () => {
  const t = curRep(); if (!t) return;
  try {
    const data = await runJob("repeater", {
      label: tr("work.repeater"),
      runBtn: $("#btnSend"),
      stopBtn: $("#btnRepStop"),
      meta: $("#repMeta"),
      fn: (signal) => api("/api/repeater", {
        method: "POST",
        signal,
        body: JSON.stringify({ raw: $("#repReq").value, scheme: $("#repScheme").value, target: $("#repTarget").value.trim() || null }),
      }),
    });
    if (!data || data.aborted) return;
    if (!data.ok) { $("#repMeta").textContent = data.error || "error"; $("#repResp").textContent = data.error || ""; t.response = data.error||""; return; }
    $("#repMeta").textContent = `${data.status} ${data.reason}  ${data.duration_ms}ms`;
    $("#repResp").textContent = data.response_raw || "";
    t.response = data.response_raw || "";
  } catch (e) { $("#repMeta").textContent = e.message; }
});
$("#btnRepToComparer").addEventListener("click", () => { const t = curRep(); if (!t) return; $("#cmpA").value = t.raw; $("#cmpB").value = t.response; showView("comparer"); });
$("#btnRepToOrganizer").addEventListener("click", () => { const t = curRep(); if (!t) return; sendToOrganizer(t.raw, t.response, "repeater"); });
$("#btnRepToIntruder").addEventListener("click", () => { const t = curRep(); if (!t) return; showView("intruder"); newIntruderTab(t.raw); });
$("#btnRepToScanner").addEventListener("click", () => { const t = curRep(); if (!t) return; showView("scanner"); $("#scanRaw").value = t.raw; $("#scanScheme").value = t.scheme; $("#scanTarget").value = t.target; });

$$(".codec-btns button[data-act]").forEach((b) => b.addEventListener("click", async () => {
  $("#codecErr").textContent = "";
  try { const d = await api("/api/codec", { method: "POST", body: JSON.stringify({ action: b.dataset.act, data: $("#codecIn").value }) }); if (!d.ok) { $("#codecErr").textContent = d.error; return; } $("#codecOut").value = d.result; }
  catch (e) { $("#codecErr").textContent = e.message; }
}));
$("#codecSwap").addEventListener("click", () => { const a = $("#codecIn").value; $("#codecIn").value = $("#codecOut").value; $("#codecOut").value = a; });

/* Intruder — multi-tab, attack types, positions */
const intruder = {
  tabs: [],
  active: 0,
  seq: 0,
};
function newIntruderTab(template) {
  intruder.seq++;
  const tab = { id: intruder.seq, name: String(intruder.seq), template: template || "", scheme: "https", target: "", attackType: "combo", payloadSets: [""], results: [] };
  intruder.tabs.push(tab);
  intruder.active = intruder.tabs.length - 1;
  renderIntruderTabs();
  renderIntruderTab();
}
function closeIntruderTab(i) {
  intruder.tabs.splice(i, 1);
  if (intruder.active >= intruder.tabs.length) intruder.active = intruder.tabs.length - 1;
  if (intruder.tabs.length === 0) newIntruderTab("");
  renderIntruderTabs();
  renderIntruderTab();
}
function renderIntruderTabs() {
  $("#intrTabList").innerHTML = intruder.tabs.map((t, i) => `<div class="intr-tab ${i === intruder.active ? "active" : ""}" data-i="${i}">
    <span>${esc(t.name)}</span><span class="close" data-close="${i}">✕</span>
  </div>`).join("");
}
$("#intrTabList").addEventListener("click", (e) => {
  const c = e.target.closest("[data-close]");
  if (c) { closeIntruderTab(+c.dataset.close); return; }
  const t = e.target.closest(".intr-tab"); if (!t) return;
  intruder.active = +t.dataset.i;
  renderIntruderTabs(); renderIntruderTab();
});
$("#btnIntrNewTab").addEventListener("click", () => newIntruderTab(""));
function curIntr() { return intruder.tabs[intruder.active]; }
function renderIntruderTab() {
  const t = curIntr(); if (!t) return;
  $("#intrTemplate").value = t.template;
  $("#intrScheme").value = t.scheme;
  $("#intrTarget").value = t.target;
  $("#intrAttackType").value = t.attackType;
  renderIntruderPayloads();
  renderIntruderResults();
  updatePositionsCount();
}
function renderIntruderPayloads() {
  const t = curIntr(); if (!t) return;
  $("#intrPayloadSets").innerHTML = t.payloadSets.map((p, i) => `<div class="payload-set">
    <textarea class="payload-area" data-i="${i}" placeholder="${esc(tr("intr.payloadPh", { n: i + 1 }))}">${esc(p)}</textarea>
  </div>`).join("");
}
function renderIntruderResults() {
  const t = curIntr(); if (!t) return;
  $("#intrBody").innerHTML = (t.results || []).map((r, i) => `<tr>
    <td>${i+1}</td><td>${esc((r.payloads||[]).join(" | "))}</td>
    <td class="${statusClass(r.status_code)}">${r.status_code || "—"}</td>
    <td>${r.body_length || 0}</td>
    <td>${r.duration ? (r.duration/1e6).toFixed(0) : "—"}</td>
    <td class="${r.error ? "status-red" : ""}">${esc(r.error || "")}</td>
  </tr>`).join("") || `<tr><td colspan="6" class="muted">${tr("empty.results")}</td></tr>`;
}
function updatePositionsCount() {
  const t = curIntr(); if (!t) return;
  const n = (t.template.match(/§/g) || []).length / 2;
  $("#intrPositionsCount").textContent = tr("intr.positions", { n: Math.floor(n) });
}
$("#intrTemplate").addEventListener("input", () => { const t = curIntr(); if (t) { t.template = $("#intrTemplate").value; updatePositionsCount(); } });
$("#intrScheme").addEventListener("change", () => { const t = curIntr(); if (t) t.scheme = $("#intrScheme").value; });
$("#intrTarget").addEventListener("input", () => { const t = curIntr(); if (t) t.target = $("#intrTarget").value; });
$("#intrAttackType").addEventListener("change", () => { const t = curIntr(); if (t) t.attackType = $("#intrAttackType").value; });
$("#intrPayloadSets").addEventListener("input", (e) => {
  const ta = e.target.closest(".payload-area"); if (!ta) return;
  const t = curIntr(); if (!t) return;
  t.payloadSets[+ta.dataset.i] = ta.value;
});
$("#btnAddPayloadSet").addEventListener("click", () => { const t = curIntr(); if (!t) return; t.payloadSets.push(""); renderIntruderPayloads(); });
$("#btnIntrAddMark").addEventListener("click", () => {
  const ta = $("#intrTemplate"); const t = curIntr(); if (!t) return;
  const s = ta.selectionStart, e = ta.selectionEnd;
  const v = ta.value;
  ta.value = v.slice(0, s) + "§" + v.slice(s, e) + "§" + v.slice(e);
  t.template = ta.value; updatePositionsCount();
});
$("#btnIntrClearMarks").addEventListener("click", () => {
  const t = curIntr(); if (!t) return;
  t.template = t.template.replace(/§/g, "");
  $("#intrTemplate").value = t.template; updatePositionsCount();
});
$("#btnIntrAutoMark").addEventListener("click", () => {
  const t = curIntr(); if (!t) return;
  // Auto-mark query parameter values and cookie values: text after = up to & or " or end of line
  let s = t.template;
  s = s.replace(/(=)([^&\s§"]*)/g, (m, eq, val) => val ? eq + "§" + val + "§" : m);
  t.template = s; $("#intrTemplate").value = s; updatePositionsCount();
});
$("#btnIntrRun").addEventListener("click", async () => {
  const t = curIntr(); if (!t) return;
  if (!t.template) { uiFlash(tr("intr.emptyTpl")); return; }
  const sets = t.payloadSets.map((p) => p.split("\n").map((x) => x.trim()).filter(Boolean));
  const positions = Math.floor((t.template.match(/§/g) || []).length / 2);
  if (positions === 0) { uiFlash(tr("intr.noMarks")); return; }
  const mode = ({ sniper: "one", battering_ram: "same", pitchfork: "zip", cluster_bomb: "combo" }[t.attackType] || t.attackType);
  if ((mode === "one" || mode === "same") && sets.length !== 1) { uiFlash(tr("intr.needOneSet", { type: tr("intr.mode." + mode), n: sets.length })); return; }
  if ((mode === "zip" || mode === "combo") && sets.length !== positions) { uiFlash(tr("intr.needNSets", { type: tr("intr.mode." + mode), need: positions, n: sets.length })); return; }
  if (sets.every((s) => s.length === 0)) { uiFlash(tr("intr.emptySets")); return; }
  $("#btnIntrRun").disabled = true;
  t.results = [];
  try {
    const data = await runJob("intruder", {
      label: tr("work.intruder"),
      runBtn: $("#btnIntrRun"),
      stopBtn: $("#btnIntrStop"),
      fn: (signal) => api("/api/batch/execute", {
        method: "POST",
        signal,
        body: JSON.stringify({
          template_raw: t.template,
          scheme: t.scheme,
          target: t.target,
          attack_type: t.attackType,
          payload_sets: sets,
          workers: Number($("#intrWorkers")?.value) || 10,
          rps: Number($("#intrRps")?.value) || 50,
        }),
      }),
    });
    if (!data || data.aborted) return;
    t.results = data || []; renderIntruderResults();
  } catch (e) { uiFlash(e.message); }
});
$("#btnIntrToIntruder").addEventListener("click", () => {
  const t = curIntr(); if (!t) return;
  newIntruderTab(t.template);
});

async function loadCallback() {
  const id = $("#callbackId").value.trim();
  if (!id) { $("#callbackList").innerHTML = `<p class="muted">${tr("callback.needId")}</p>`; return; }
  try { renderCallback(await api(`/api/callback/${encodeURIComponent(id)}`) || []); } catch (e) { $("#callbackList").innerHTML = `<p class="error">${esc(e.message)}</p>`; }
}
$("#btnCallbackRefresh").addEventListener("click", loadCallback);
$("#callbackId").addEventListener("keydown", (e) => { if (e.key === "Enter") loadCallback(); });
$("#btnCallbackCopy").addEventListener("click", () => { const id = $("#callbackId").value.trim() || "<request-id>"; navigator.clipboard.writeText(`http://127.0.0.1:8082/${id}/`); });

/* TARGET: Site map (on Map), Scope, Issues */
$$("#otherSubtabs button").forEach((b) => b.addEventListener("click", () => showSub("#view-other", b.dataset.sub)));

const scope = {
  inc: JSON.parse(localStorage.getItem("aura_scope_inc") || localStorage.getItem("meb_scope_inc") || "[]"),
  exc: JSON.parse(localStorage.getItem("aura_scope_exc") || localStorage.getItem("meb_scope_exc") || "[]"),
};
function saveScope() {
  localStorage.setItem("aura_scope_inc", JSON.stringify(scope.inc));
  localStorage.setItem("aura_scope_exc", JSON.stringify(scope.exc));
}
function scopePrefixMatch(prefix, subdomains, url) {
  if (!prefix) return false;
  try {
    const u = new URL(url);
    const host = u.host;
    const p = prefix.replace(/^https?:\/\//, "").replace(/\/.*$/, "");
    if (subdomains) return host === p || host.endsWith("." + p);
    return host === p || url.startsWith(prefix);
  } catch (_) { return url.startsWith(prefix); }
}
function scopeMatches(url) {
  if (!scope.inc.length) return true;
  const inScope = scope.inc.some((r) => r.enabled && scopePrefixMatch(r.prefix, r.subdomains, url));
  if (!inScope) return false;
  return !scope.exc.some((r) => r.enabled && scopePrefixMatch(r.prefix, r.subdomains, url));
}
function renderScope() {
  for (const [key, body] of [["inc", "#scopeIncBody"], ["exc", "#scopeExcBody"]]) {
    $(body).innerHTML = scope[key].map((r, i) => `<tr>
      <td><input type="checkbox" data-k="${key}" data-i="${i}" data-f="enabled" ${r.enabled ? "checked" : ""}/></td>
      <td><input type="text" data-k="${key}" data-i="${i}" data-f="prefix" value="${esc(r.prefix)}" placeholder="example.com or https://example.com/api"/></td>
      <td><input type="checkbox" data-k="${key}" data-i="${i}" data-f="subdomains" ${r.subdomains ? "checked" : ""}/></td>
      <td><button class="del-btn" data-k="${key}" data-i="${i}" data-act="del">✕</button></td>
    </tr>`).join("") || `<tr><td colspan="4" class="muted">${tr("empty.rules")}</td></tr>`;
  }
}
function addScopeRow(key) { scope[key].push({ enabled: true, prefix: "", subdomains: false }); saveScope(); renderScope(); }
$("#scopeIncAdd").addEventListener("click", () => addScopeRow("inc"));
$("#scopeExcAdd").addEventListener("click", () => addScopeRow("exc"));
$("#scopeIncPaste").addEventListener("click", () => navigator.clipboard.readText().then((t) => { scope.inc.push({ enabled: true, prefix: t.trim(), subdomains: true }); saveScope(); renderScope(); }).catch(() => addScopeRow("inc")));
$("#scopeExcPaste").addEventListener("click", () => navigator.clipboard.readText().then((t) => { scope.exc.push({ enabled: true, prefix: t.trim(), subdomains: true }); saveScope(); renderScope(); }).catch(() => addScopeRow("exc")));
document.addEventListener("change", (e) => {
  const t = e.target;
  if (t.dataset && t.dataset.k && t.dataset.f) {
    const r = scope[t.dataset.k][+t.dataset.i];
    if (!r) return;
    if (t.type === "checkbox") r[t.dataset.f] = t.checked; else r[t.dataset.f] = t.value;
    saveScope();
    if (t.dataset.f === "prefix") renderSiteMap();
  }
});
document.addEventListener("click", (e) => {
  const b = e.target.closest("[data-act='del']");
  if (b && b.dataset.k) { scope[b.dataset.k].splice(+b.dataset.i, 1); saveScope(); renderScope(); }
});

let siteItems = [];
let siteSelectedId = null;
async function loadSiteMap() {
  try {
    siteItems = (await api("/api/history?q=")).items || [];
  } catch (_) {
    siteItems = siteItems || [];
  }
  renderSiteMap();
}
function emptySiteNode() {
  return { _leaf: null, _intelHost: null, _intelPath: null, _children: {} };
}
function walkSiteNode(tree, host, parts) {
  let node = tree[host] = tree[host] || emptySiteNode();
  for (const p of parts) node = node._children[p] = node._children[p] || emptySiteNode();
  return node;
}
function intelHostUrl(host, path) {
  const p = path || "/";
  return "https://" + host + (p.startsWith("/") ? p : "/" + p);
}
function mergeIntelIntoSiteTree(tree, filter, applyScope) {
  for (const h of (mapState.map && mapState.map.hosts) || []) {
    const host = h.host || "?";
    if (filter && !host.toLowerCase().includes(filter)) {
      const anyPath = (h.paths || []).some((p) => `${host} ${p.path || "/"}`.toLowerCase().includes(filter));
      if (!anyPath) continue;
    }
    if (applyScope && !scopeMatches(intelHostUrl(host, "/")) && !(h.paths || []).some((p) => scopeMatches(intelHostUrl(host, p.path || "/")))) {
      continue;
    }
    const hostNode = walkSiteNode(tree, host, []);
    hostNode._intelHost = h;
    for (const p of h.paths || []) {
      let path = p.path || "/";
      if (path !== "/" && !path.startsWith("/")) path = "/" + path;
      if (filter && !(`${host} ${path} ${(p.params || []).join(" ")}`.toLowerCase().includes(filter))) continue;
      if (applyScope && !scopeMatches(intelHostUrl(host, path))) continue;
      const parts = path.split("/").filter(Boolean);
      const node = walkSiteNode(tree, host, parts);
      node._intelPath = p;
    }
  }
}
function hasSiteMapData() {
  if ((siteItems || []).length) return true;
  if (mapState.map && mapState.map.hosts && mapState.map.hosts.length) return true;
  return false;
}
function syncMapSitePanes() {
  const showWork = !!mapState.target || hasSiteMapData();
  $("#mapEmpty")?.classList.toggle("hidden", showWork);
  $("#mapWork")?.classList.toggle("hidden", !showWork);
  if ($("#btnMapHistory")) $("#btnMapHistory").disabled = !mapState.target;
  if ($("#mapStats")) $("#mapStats").classList.toggle("hidden", !mapState.target);
}
function renderSiteMap() {
  const treeEl = $("#siteTree");
  if (!treeEl) return;
  const filter = ($("#siteFilter")?.value || "").toLowerCase();
  const applyScope = $("#siteApplyScope")?.checked;
  const items = (siteItems || []).filter((f) => {
    if (filter && !(`${f.method} ${f.host} ${f.path}`.toLowerCase().includes(filter))) return false;
    if (applyScope && !scopeMatches(`${f.scheme || "https"}://${f.host}${f.path}`)) return false;
    return true;
  });
  const tree = {};
  for (const f of items) {
    const host = f.host || "?";
    const parts = (f.path || "/").split("/").filter(Boolean);
    walkSiteNode(tree, host, parts)._leaf = f;
  }
  mergeIntelIntoSiteTree(tree, filter, applyScope);
  treeEl.innerHTML = Object.keys(tree).sort().map((host) => renderTreeNode(host, tree[host], true, host, "/")).join("") || `<div class="muted">${tr("empty.sitemap")}</div>`;
  syncMapSitePanes();
}
function renderTreeNode(name, node, isHost, host, path) {
  const childKeys = Object.keys(node._children).sort();
  const hasChildren = childKeys.length > 0;
  const leaf = node._leaf;
  const intelP = node._intelPath;
  const intelH = isHost ? node._intelHost : null;
  const nid = "n" + Math.random().toString(36).slice(2, 9);
  const icon = isHost ? "🔒" : (leaf || intelP ? "📄" : "📁");
  const status = leaf ? leaf.status : (intelP && intelP.status);
  const extra = [];
  if ((intelP?.params || []).length) extra.push((intelP.params || []).slice(0, 4).join(", "));
  if (isHost && (intelH?.tech || []).length) extra.push((intelH.tech || []).slice(0, 3).join(", "));
  if (isHost && (intelH?.ports || []).length) extra.push(":" + (intelH.ports || []).slice(0, 4).join(","));
  if (!leaf && (intelP || intelH)) extra.push(tr("map.srcRecon"));
  return `<div class="tree-node">
    <div class="tree-row" data-id="${leaf ? leaf.id : ""}" data-host="${esc(host)}" data-path="${esc(path || "/")}">
      <span class="tree-tw" data-toggle="${nid}">${hasChildren ? "▶" : ""}</span>
      <span class="tree-icon">${icon}</span>
      <span>${esc(name)}</span>
      ${leaf ? `<span class="tree-method ${leaf.method}">${esc(leaf.method)}</span>` : ""}
      ${status ? `<span class="tree-status ${statusClass(status)}">${esc(status)}</span>` : (intelH && intelH.live ? `<span class="tree-status status-ok">live</span>` : "")}
      ${extra.length ? `<span class="tree-status">${esc(extra.filter(Boolean).join(" · "))}</span>` : ""}
    </div>
    <div id="${nid}" class="tree-children hidden">${childKeys.map((k) => renderTreeNode(k, node._children[k], false, host, (isHost ? "" : path) + "/" + k)).join("")}</div>
  </div>`;
}
function showMapDetailPane(name) {
  $$("#mapDetailTabs button").forEach((b) => b.classList.toggle("active", b.dataset.pane === name));
  $("#siteReq")?.classList.toggle("hidden", name !== "req");
  $("#siteResp")?.classList.toggle("hidden", name !== "resp");
  $("#mapDetail")?.classList.toggle("hidden", name !== "info");
  $("#mapLog")?.classList.toggle("hidden", name !== "log");
}
function fillMapInfo(host, path) {
  const detail = $("#mapDetail");
  if (!detail) return;
  const h = ((mapState.map && mapState.map.hosts) || []).find((x) => x.host === host);
  if (!h) {
    detail.innerHTML = `<strong>${esc(host)}${path && path !== "/" ? esc(path) : ""}</strong><div class="muted">${esc(tr("map.pickNode"))}</div>`;
    return;
  }
  if (path && path !== "/") {
    const p = (h.paths || []).find((x) => x.path === path || x.path === path.replace(/^\//, ""));
    detail.innerHTML = `<strong>${esc(host)}${esc(path)}</strong>
      <div>status ${esc(p?.status || "—")}</div>
      <div>params: ${esc((p?.params || []).join(", ") || "—")}</div>
      <div>js: ${esc((p?.js || []).slice(0, 8).join(", ") || "—")}</div>`;
    return;
  }
  detail.innerHTML = `<strong>${esc(displayMapHost(h, mapState.map && mapState.map.base_url))}</strong> ${h.live ? "live" : ""}
    <div>tech: ${esc((h.tech || []).join(", ") || "—")}</div>
    <div>ports: ${esc((h.ports || []).join(", ") || "—")}</div>`;
}
$("#siteTree").addEventListener("click", async (e) => {
  const tw = e.target.closest(".tree-tw");
  if (tw) {
    const c = document.getElementById(tw.dataset.toggle);
    if (c) { c.classList.toggle("hidden"); tw.textContent = c.classList.contains("hidden") ? "▶" : "▼"; }
    return;
  }
  const row = e.target.closest(".tree-row");
  if (!row) return;
  $$("#siteTree .tree-row").forEach((r) => r.classList.remove("sel"));
  row.classList.add("sel");
  const host = row.dataset.host || "";
  const path = row.dataset.path || "/";
  fillMapInfo(host, path);
  const id = row.dataset.id;
  if (!id) {
    siteSelectedId = null;
    $("#siteReq").textContent = "";
    $("#siteResp").textContent = "";
    showMapDetailPane("info");
    return;
  }
  const data = await api(`/api/history/${id}`);
  siteSelectedId = id;
  $("#siteReq").textContent = data.request_raw || "";
  $("#siteResp").textContent = data.response_raw || data.error || "";
  showMapDetailPane("req");
});
$("#btnSiteToRepeater").addEventListener("click", () => { if (siteSelectedId) flowToTool(siteSelectedId, "repeater"); });
$("#btnSiteToIntruder").addEventListener("click", () => { if (siteSelectedId) flowToTool(siteSelectedId, "intruder"); });
$("#btnSiteToScanner").addEventListener("click", () => { if (!siteSelectedId) return; api(`/api/history/${siteSelectedId}`).then((f) => { showView("scanner"); $("#scanRaw").value = f.request_raw || ""; $("#scanScheme").value = f.scheme || "https"; }); });
$$("#mapDetailTabs button").forEach((b) => b.addEventListener("click", () => showMapDetailPane(b.dataset.pane)));
$("#siteFilter").addEventListener("input", () => { clearTimeout(window._siteT); window._siteT = setTimeout(renderSiteMap, 200); });
$("#siteApplyScope").addEventListener("change", renderSiteMap);

/* Issues */
const ISSUE_DEFS = [
  { nameKey: "issue.headers.name", sev: "low", descKey: "issue.headers.desc", fixKey: "issue.headers.fix", refs: ["https://owasp.org/www-project-secure-headers/"] },
  { nameKey: "issue.verbose.name", sev: "info", descKey: "issue.verbose.desc", fixKey: "issue.verbose.fix", refs: [] },
  { nameKey: "issue.token.name", sev: "high", descKey: "issue.token.desc", fixKey: "issue.token.fix", refs: ["https://owasp.org/www-community/vulnerabilities/Information_exposure_through_query_strings_in_url"] },
  { nameKey: "issue.cookie.name", sev: "medium", descKey: "issue.cookie.desc", fixKey: "issue.cookie.fix", refs: ["https://owasp.org/www-community/controls/SecureCookieAttribute"] },
  { nameKey: "issue.redos.name", sev: "info", descKey: "issue.redos.desc", fixKey: "issue.redos.fix", refs: ["https://owasp.org/www-community/attacks/Regular_expression_Denial_of_Service_-_ReDoS"] },
];
function renderIssueDefs() {
  $("#issueDefsBody").innerHTML = ISSUE_DEFS.map((d, i) => `<tr data-i="${i}"><td>${esc(tr(d.nameKey))}</td><td class="sev-${d.sev}">${d.sev}</td></tr>`).join("");
}
$("#issueDefsBody").addEventListener("click", (e) => {
  const row = e.target.closest("tr"); if (!row) return;
  const d = ISSUE_DEFS[+row.dataset.i];
  $("#issueTitle").textContent = tr(d.nameKey);
  $("#issueMeta").innerHTML = `${esc(tr("issue.severity"))}: <span class="sev-${d.sev}">${d.sev}</span>`;
  $("#issueDesc").textContent = tr(d.descKey);
  $("#issueFix").textContent = tr(d.fixKey);
  $("#issueRefs").innerHTML = (d.refs || []).map((r) => `<div><a href="${r}" target="_blank">${esc(r)}</a></div>`).join("") || `<span class="muted">—</span>`;
});
async function loadIssues() {
  try {
    const items = (await api("/api/findings")).items || [];
    $("#issuesBody").innerHTML = items.map((f) => `<tr data-tx="${f.transaction_id}">
      <td class="sev-${f.severity}">${esc(f.severity)}</td>
      <td>${esc(f.rule_name)}</td>
      <td>${esc(f.title)}</td>
      <td><code>${esc(f.transaction_id)}</code></td>
    </tr>`).join("") || `<tr><td colspan="4" class="muted">${tr("empty.findings")}</td></tr>`;
  } catch (e) { $("#issuesBody").innerHTML = `<tr><td colspan="4" class="error">${esc(e.message)}</td></tr>`; }
}
$("#issuesBody").addEventListener("click", async (e) => {
  const tr = e.target.closest("tr"); if (!tr || !tr.dataset.tx) return;
  const tx = await api(`/api/history/${tr.dataset.tx}`);
  $("#issueTitle").textContent = tr.querySelector("td:nth-child(3)").textContent;
  $("#issueMeta").innerHTML = `Transaction: <code>${esc(tr.dataset.tx)}</code>`;
  $("#issueDesc").textContent = (tx.request_raw || "").slice(0, 800);
  $("#issueFix").textContent = tr("issue.openHist");
  $("#issueRefs").innerHTML = "";
});

/* Comparer — client-side LCS line diff */
function lineDiff(a, b, ignoreWs) {
  const norm = (s) => ignoreWs ? s.replace(/\s+/g, " ").trim() : s;
  const A = a.split("\n"), B = b.split("\n");
  const m = A.length, n = B.length;
  const dp = Array.from({ length: m+1 }, () => new Int32Array(n+1));
  for (let i = m-1; i >= 0; i--) for (let j = n-1; j >= 0; j--) dp[i][j] = norm(A[i]) === norm(B[j]) ? dp[i+1][j+1]+1 : Math.max(dp[i+1][j], dp[i][j+1]);
  const out = []; let i = 0, j = 0;
  while (i < m && j < n) {
    if (norm(A[i]) === norm(B[j])) { out.push({ t: "eq", s: A[i] }); i++; j++; }
    else if (dp[i+1][j] >= dp[i][j+1]) { out.push({ t: "del", s: A[i] }); i++; }
    else { out.push({ t: "add", s: B[j] }); j++; }
  }
  while (i < m) { out.push({ t: "del", s: A[i++] }); }
  while (j < n) { out.push({ t: "add", s: B[j++] }); }
  return out;
}
$("#btnCompare").addEventListener("click", () => {
  const a = $("#cmpA").value, b = $("#cmpB").value;
  const ignoreWs = $("#cmpIgnoreWs").checked;
  const diff = lineDiff(a, b, ignoreWs);
  const eq = diff.filter((d) => d.t === "eq").length;
  $("#cmpMeta").textContent = tr("cmp.meta", { n: diff.length, eq, diff: diff.length-eq });
  $("#cmpDiff").innerHTML = diff.map((d) => {
    const cls = d.t === "add" ? "diff-add" : d.t === "del" ? "diff-del" : "diff-eq";
    const sign = d.t === "add" ? "+" : d.t === "del" ? "-" : " ";
    return `<div class="${cls}">${sign} ${esc(d.s)}</div>`;
  }).join("");
});

/* Sequencer */
$("#btnSeqAnalyze").addEventListener("click", async () => {
  const tokens = $("#seqTokens").value;
  await withWork(tr("work.running"), $("#btnSeqAnalyze"), async () => {
    const rep = await api("/api/sequencer/analyze", { method: "POST", body: JSON.stringify({ tokens }) });
    renderSequencer(rep);
  }).catch((e) => uiFlash(e.message));
});
$("#btnSeqClear").addEventListener("click", () => { $("#seqTokens").value = ""; $("#seqResults").innerHTML = `<p class="muted">${tr("seq.results")}</p>`; $("#seqCount").textContent = ""; });
$("#seqTokens").addEventListener("input", () => {
  const n = $("#seqTokens").value.split("\n").filter((x) => x.trim()).length;
  $("#seqCount").textContent = tr("seq.tokens", { n });
});
function renderSequencer(r) {
  const ratingCls = r.overall.rating === "good" ? "good" : r.overall.rating === "fair" ? "fair" : "poor";
  const fips = (r.bit_level.fips_tests || []).map((t) => `<div class="seq-fips-row">
    <span class="name">${esc(t.name)}</span>
    <span class="${t.pass ? "pass" : "fail"}">${t.pass ? "PASS" : "FAIL"}</span>
    <span class="muted">p=${(t.p_value||0).toExponential(2)}</span>
    <span class="muted">${esc(t.detail||"")}</span>
  </div>`).join("");
  const entropy = (r.effective_entropy.levels || []).map((l) => `<div class="seq-fips-row">
    <span class="name">${esc(l.level)}</span>
    <div class="seq-bar"><span style="width:${Math.min(100, l.bits_passing*4)}%"></span></div>
    <span>${l.bits_passing} bits</span>
  </div>`).join("");
  const pos = (r.char_level.positions || []).map((p) => `<tr>
    <td>${p.index}</td><td>${p.charset_size}</td><td>${p.max_entropy_bits.toFixed(2)}</td>
    <td class="${p.pass ? "status-green" : "status-red"}">${p.pass ? "pass" : "fail"}</td>
    <td class="muted">p=${(p.p_value||0).toExponential(2)}</td>
  </tr>`).join("");
  $("#seqResults").innerHTML = `
    <div class="seq-rating ${ratingCls}">Overall: ${esc(r.overall.rating)} (score ${(r.overall.score*100).toFixed(0)}%)</div>
    <div class="muted">Sample: ${r.sample_size} tokens, length ${r.token_length.min}..${r.token_length.max}, padded: ${r.padded}. Reliability: ${esc(r.reliability)}</div>
    <div class="seq-section"><h3>Effective entropy</h3>${entropy || `<p class="muted">${tr("empty.seq")}</p>`}</div>
    <div class="seq-section"><h3>FIPS bit-level tests (${r.bit_level.total_bits} bits)</h3>${fips}</div>
    <div class="seq-section"><h3>Character-level analysis</h3>
      <table class="seq-pos-table"><thead><tr><th>Pos</th><th>Charset</th><th>Max bits</th><th>Result</th><th>p</th></tr></thead><tbody>${pos}</tbody></table>
    </div>`;
}

/* Organizer */
const org = { items: [], collections: ["inbox"], active: "inbox", selectedId: null };
async function loadOrganizer() {
  try {
    const d = await api(`/api/organizer?collection=${encodeURIComponent(org.active)}`);
    org.items = d.items || [];
    org.collections = d.collections || ["inbox"];
    renderOrganizer();
  } catch (e) { uiFlash(e.message); }
}
function renderOrganizer() {
  $("#orgCollections").innerHTML = org.collections.map((c) => `<div class="org-coll-item ${c === org.active ? "active" : ""}" data-c="${esc(c)}">${esc(c)}</div>`).join("");
  $("#orgMoveColl").innerHTML = org.collections.map((c) => `<option value="${esc(c)}">${esc(c)}</option>`).join("");
  $("#orgBody").innerHTML = org.items.map((it, i) => `<tr data-id="${esc(it.id)}">
    <td>${i+1}</td><td class="muted">${fmtTime(it.created)}</td>
    <td>${esc(it.status||"")}</td><td>${esc(it.method)}</td><td>${esc(it.host)}</td>
    <td>${esc(it.path)}</td><td class="${statusClass(it.status_code)}">${it.status_code||"—"}</td>
    <td>${it.length||0}</td><td>${esc(it.notes||"")}</td>
  </tr>`).join("") || `<tr><td colspan="9" class="muted">${tr("empty.org")}</td></tr>`;
  renderOrganizerDetail();
}
function renderOrganizerDetail() {
  const it = org.items.find((x) => x.id === org.selectedId);
  if (!it) {
    $("#orgReq").value = ""; $("#orgResp").textContent = "";
    $("#orgNotes").value = ""; $("#orgStatus").value = "";
    $("#orgHighlight").value = ""; return;
  }
  $("#orgReq").value = it.request_raw || "";
  $("#orgResp").textContent = it.response_raw || "";
  $("#orgNotes").value = it.notes || "";
  $("#orgStatus").value = it.status || "";
  $("#orgHighlight").value = it.highlight || "";
  $("#orgMoveColl").value = it.collection || "inbox";
}
$("#orgCollections").addEventListener("click", (e) => {
  const c = e.target.closest("[data-c]"); if (!c) return;
  org.active = c.dataset.c; loadOrganizer();
});
$("#btnOrgAddColl").addEventListener("click", () => {
  const v = $("#orgNewColl").value.trim(); if (!v) return;
  if (!org.collections.includes(v)) org.collections.push(v);
  org.active = v; $("#orgNewColl").value = ""; loadOrganizer();
});
$("#orgBody").addEventListener("click", (e) => {
  const r = e.target.closest("tr[data-id]"); if (!r) return;
  org.selectedId = r.dataset.id; renderOrganizerDetail();
});
$("#btnOrgSave").addEventListener("click", async () => {
  if (!org.selectedId) return;
  try {
    await api(`/api/organizer/${encodeURIComponent(org.selectedId)}`, { method: "PATCH", body: JSON.stringify({
      notes: $("#orgNotes").value, status: $("#orgStatus").value,
      highlight: $("#orgHighlight").value, collection: $("#orgMoveColl").value,
    }) });
    loadOrganizer();
  } catch (e) { uiFlash(e.message); }
});
$("#btnOrgDelete").addEventListener("click", async () => {
  if (!org.selectedId) return;
  if (!confirm(tr("org.confirmDel"))) return;
  try {
    await api(`/api/organizer/${encodeURIComponent(org.selectedId)}`, { method: "DELETE" });
    org.selectedId = null; loadOrganizer();
  } catch (e) { uiFlash(e.message); }
});
$("#btnOrgToRepeater").addEventListener("click", () => {
  const it = org.items.find((x) => x.id === org.selectedId); if (!it) return;
  showView("repeater"); newRepeaterTab(it.request_raw || "");
});
$("#btnOrgToComparer").addEventListener("click", () => {
  const it = org.items.find((x) => x.id === org.selectedId); if (!it) return;
  $("#cmpA").value = it.request_raw || ""; $("#cmpB").value = it.response_raw || "";
  showView("comparer");
});
async function sendToOrganizer(requestRaw, responseRaw, tool) {
  try {
    await api("/api/organizer", { method: "POST", body: JSON.stringify({
      request_raw: requestRaw, response_raw: responseRaw || "", tool: tool || "proxy", collection: org.active || "inbox",
    }) });
    if (currentView() === "organizer") loadOrganizer();
  } catch (e) { uiFlash(e.message); }
}
/* Send-to handlers for Dashboard recent flows and Target site map */
async function flowToTool(flowId, tool) {
  try {
    const f = await api(`/api/history/${flowId}`);
    const raw = f.request_raw || "";
    if (!raw) { uiFlash(tr("flow.noReq")); return; }
    if (tool === "repeater") { showView("repeater"); newRepeaterTab(raw); }
    else if (tool === "intruder") { showView("intruder"); newIntruderTab(raw); }
  } catch (e) { uiFlash(e.message); }
}
$("#dRecent").addEventListener("click", (e) => {
  const b = e.target.closest("button[data-act]"); if (!b) return;
  flowToTool(b.dataset.id, b.dataset.act === "intr" ? "intruder" : "repeater");
});

function currentView() {
  const vis = $$("main > .view").find((el) => !el.classList.contains("hidden"));
  return vis?.id?.replace(/^view-/, "") || $("#mainTabs button.active")?.dataset.view || "map";
}
function fmtTime(iso) { try { return new Date(iso).toLocaleTimeString(); } catch { return iso; } }

/* Send to Sequencer / Decoder */
$("#btnToSequencer").addEventListener("click", () => {
  if (!state.histReq) return;
  showView("sequencer");
  $("#seqTokens").value = "";
  // For sequencer we need tokens, not a request — but allow sending response body or a header value.
  // Use the response body lines as a fallback token source.
  $("#seqTokens").value = (state.histResp || "").split("\n").filter(Boolean).join("\n");
  $("#seqCount").textContent = tr("seq.tokens", { n: $("#seqTokens").value.split("\n").length });
});
$("#btnToDecoder").addEventListener("click", () => {
  if (!state.histReq) return;
  showView("decoder");
  $("#codecIn").value = state.histReq;
});
$("#btnToScanner").addEventListener("click", () => {
  if (!state.histReq) return;
  showView("scanner");
  $("#scanRaw").value = state.histReq;
  $("#scanScheme").value = state.histScheme || "https";
});

/* History advanced filter */
let histFilter = { scope: false, params: false, status: "", mime: "", ext: "", search: "", regex: false, neg: false };
$("#btnHistFilter").addEventListener("click", () => $("#histFilterPanel").classList.toggle("hidden"));
$("#btnHfApply").addEventListener("click", () => {
  histFilter = {
    scope: $("#hfScope").checked, params: $("#hfParams").checked,
    status: $("#hfStatus").value.trim(), mime: $("#hfMime").value.trim(),
    ext: $("#hfExt").value.trim(), search: $("#hfSearch").value,
    regex: $("#hfRegex").checked, neg: $("#hfNeg").checked,
  };
  $("#histFilterPanel").classList.add("hidden");
  loadHistory();
});
$("#btnHfReset").addEventListener("click", () => {
  histFilter = { scope: false, params: false, status: "", mime: "", ext: "", search: "", regex: false, neg: false };
  ["hfScope","hfParams","hfRegex","hfNeg"].forEach((id) => $("#"+id).checked = false);
  ["hfStatus","hfMime","hfExt","hfSearch"].forEach((id) => $("#"+id).value = "");
  loadHistory();
});
function applyHistFilter(f) {
  if (histFilter.scope && typeof scopeMatches === "function" && !scopeMatches(`${f.scheme}://${f.host}${f.path}`)) return false;
  if (histFilter.params && !hasParams(f)) return false;
  if (histFilter.status) {
    const codes = histFilter.status.split(",").map((s) => s.trim()).filter(Boolean);
    if (!codes.includes(String(f.status))) return false;
  }
  if (histFilter.ext) {
    const exts = histFilter.ext.split(",").map((s) => s.trim().toLowerCase()).filter(Boolean);
    const e = (f.path.split(".").pop() || "").toLowerCase();
    if (!exts.includes(e)) return false;
  }
  if (histFilter.search) {
    const hay = `${f.method} ${f.host} ${f.path} ${f.status} ${(f.response_body||"").slice(0,4096)}`;
    let m;
    try { m = histFilter.regex ? new RegExp(histFilter.search).test(hay) : hay.includes(histFilter.search); }
    catch { m = hay.includes(histFilter.search); }
    if (histFilter.neg && m) return false;
    if (!histFilter.neg && !m) return false;
  }
  return true;
}
function hasParams(f) { return (f.path||"").includes("?") || !!(f.request_body && f.request_body.includes("=")); }

/* Discover */
let discAbort = null;
let seclists = { items: [], count: 0 };

async function loadSecLists() {
  try {
    seclists = await api("/api/wordlists");
  } catch (_) {
    seclists = { items: [], count: 0 };
  }
  const n = seclists.count || 0;
  const msg = n ? tr("wl.meta", { n }) : tr("wl.missing");
  if ($("#discWlMeta")) $("#discWlMeta").textContent = msg;
  if ($("#fuzzWlMeta")) $("#fuzzWlMeta").textContent = msg;
  fillWlSelect($("#discWlSelect"), $("#discWlSearch")?.value, "Discovery/Web-Content/common.txt");
  fillWlSelect($("#fuzzWlSelect"), $("#fuzzWlSearch")?.value, "Discovery/Web-Content/burp-parameter-names.txt");
}

function fillWlSelect(sel, q, prefer) {
  if (!sel) return;
  q = (q || "").toLowerCase();
  const items = (seclists.items || []).filter((e) => !q || e.path.toLowerCase().includes(q));
  const shown = items.slice(0, 400);
  sel.innerHTML = shown.map((e) => `<option value="${esc(e.path)}">${esc(e.path)} · ${fmtBytes(e.size)}</option>`).join("")
    || `<option value="">${tr("wl.noMatch")}</option>`;
  if (prefer && !q) {
    const opt = [...sel.options].find((o) => o.value === prefer);
    if (opt) sel.value = prefer;
  }
  if (items.length > shown.length) {
    const o = document.createElement("option");
    o.disabled = true;
    o.textContent = tr("wl.more", { n: items.length - shown.length });
    sel.appendChild(o);
  }
}

$("#btnDiscRun").addEventListener("click", async () => {
  const base = $("#discBase").value.trim();
  const custom = $("#discWords").value.split("\n").map((s) => s.trim()).filter(Boolean);
  const path = $("#discWlSelect")?.value || "";
  if (!base || (custom.length === 0 && !path)) { uiFlash(tr("disc.need")); return; }
  if (!requireLabAuth($("#discAuth"))) return;
  $("#discBody").innerHTML = "";
  try {
    const data = await runJob("discover", {
      label: tr("work.discover"),
      runBtn: $("#btnDiscRun"),
      stopBtn: $("#btnDiscStop"),
      meta: $("#discMeta"),
      fn: (signal) => api("/api/discover/run", { method: "POST", signal, body: JSON.stringify({
        base_url: base, wordlist: custom, wordlist_path: custom.length ? "" : path,
        workers: +$("#discWorkers").value || 10,
        rps: +$("#discRps").value || 20, cookies: $("#discCookies").value,
        authorized: true,
      }) }),
    });
    if (!data || data.aborted) return;
    discCache = data || [];
    renderDisc(discCache);
    $("#discMeta").textContent = tr("disc.nResults", { n: discCache.length });
  } catch (e) { uiFlash(e.message); $("#discMeta").textContent = ""; }
});
function renderDisc(rows) {
  const q = ($("#discFilter").value || "").toLowerCase();
  const hide2 = $("#discHide2xx").checked, hide404 = $("#discHide404").checked;
  const out = rows.filter((r) => {
    if (hide2 && r.status_code >= 200 && r.status_code < 300) return false;
    if (hide404 && r.status_code === 404) return false;
    if (q && !(`${r.path} ${r.status_code}`.toLowerCase().includes(q))) return false;
    return true;
  });
  $("#discBody").innerHTML = out.map((r) => `<tr>
    <td>${esc(r.path)}</td>
    <td class="${statusClass(r.status_code)}">${r.status_code||"—"}</td>
    <td>${r.length||0}</td>
    <td>${r.duration ? (r.duration/1e6).toFixed(0) : "—"}</td>
    <td class="${r.error?"status-red":""}">${esc(r.error||"")}</td>
  </tr>`).join("") || `<tr><td colspan="5" class="muted">${tr("empty.results")}</td></tr>`;
}
$("#discFilter").addEventListener("input", () => { if (discCache.length) renderDisc(discCache); });
$("#discHide2xx").addEventListener("change", () => { if (discCache.length) renderDisc(discCache); });
$("#discHide404").addEventListener("change", () => { if (discCache.length) renderDisc(discCache); });
let discCache = [];

/* Scanner */
function rawGetFromTarget(target) {
  let host = String(target || "").trim();
  if (!host) return "";
  let path = "/";
  try {
    const u = new URL(host.includes("://") ? host : "https://" + host);
    host = u.host;
    path = u.pathname || "/";
    if (u.search) path += u.search;
  } catch (_) {}
  if (!host) return "";
  return `GET ${path} HTTP/1.1\r\nHost: ${host}\r\nUser-Agent: AURA/0.1\r\nAccept: */*\r\n\r\n`;
}
$("#btnScanRun").addEventListener("click", async () => {
  let raw = $("#scanRaw").value;
  if (!raw.trim() && $("#scanFromHistory")?.checked && state.histReq) {
    raw = state.histReq;
    $("#scanRaw").value = raw;
  }
  if (!raw.trim()) {
    const target = $("#scanTarget").value.trim();
    raw = rawGetFromTarget(target);
    if (raw) {
      $("#scanRaw").value = raw;
      let host = target, path = "/";
      try {
        const u = new URL(target.includes("://") ? target : "https://" + target);
        host = u.host;
        path = (u.pathname || "/") + (u.search || "");
      } catch (_) {}
      $("#scanMeta").textContent = tr("scan.usingGet", { host, path });
    }
  }
  if (!raw.trim()) {
    $("#scanMeta").textContent = tr("scan.need");
    $("#scanRaw")?.focus();
    return;
  }
  if (!requireLabAuth($("#scanAuth"))) return;
  $("#scanBody").innerHTML = "";
  try {
    const data = await runJob("scanner", {
      label: tr("work.scan"),
      runBtn: $("#btnScanRun"),
      stopBtn: $("#btnScanStop"),
      meta: $("#scanMeta"),
      fn: (signal) => api("/api/scanner/scan", { method: "POST", signal, body: JSON.stringify({
        raw, scheme: $("#scanScheme").value, target: $("#scanTarget").value.trim(),
        authorized: true,
      }) }),
    });
    if (!data || data.aborted) return;
    renderScan(data || { findings: [] });
    $("#scanMeta").textContent = `${(data.findings||[]).length} findings, ${(data.took/1e9).toFixed(1)}s`;
  } catch (e) { $("#scanMeta").textContent = e.message; }
});
function renderScan(res) {
  const sev = { high: "status-red", medium: "status-orange", low: "status-yellow" };
  $("#scanBody").innerHTML = (res.findings||[]).map((f) => `<tr>
    <td class="${sev[f.severity]||""}">${esc(f.severity||"")}</td>
    <td>${esc(f.name||"")}</td>
    <td>${esc(f.param||"")}</td>
    <td>${esc((f.evidence||"").slice(0,80))}</td>
    <td>${esc(f.description||"")}</td>
  </tr>`).join("") || `<tr><td colspan="5" class="muted">${tr("empty.vulns")}</td></tr>`;
}
$("#scanFromHistory").addEventListener("change", () => {
  if ($("#scanFromHistory").checked && state.histReq) $("#scanRaw").value = state.histReq;
});

/* Inspector — parses a raw HTTP request into structured fields */
function parseRequestForInspector(raw) {
  const out = { attrs: {}, requestHeaders: [], params: [], cookies: [] };
  if (!raw) return out;
  const lines = raw.split(/\r?\n/);
  const first = lines[0] ? lines[0].split(" ") : [];
  out.attrs = { method: first[0]||"", path: first[1]||"", version: first[2]||"" };
  let inBody = false; const bodyLines = [];
  for (let i = 1; i < lines.length; i++) {
    const l = lines[i];
    if (l === "") { inBody = true; continue; }
    if (inBody) { bodyLines.push(l); continue; }
    const idx = l.indexOf(":");
    if (idx < 0) continue;
    const k = l.slice(0, idx).trim(), v = l.slice(idx+1).trim();
    out.requestHeaders.push([k, v]);
    if (k.toLowerCase() === "cookie") {
      v.split(";").forEach((c) => { const eq = c.indexOf("="); if (eq>=0) out.cookies.push([c.slice(0,eq).trim(), c.slice(eq+1).trim()]); });
    }
  }
  // params from query
  const qIdx = out.attrs.path.indexOf("?");
  if (qIdx >= 0) {
    const qs = out.attrs.path.slice(qIdx+1).split("&");
    qs.forEach((kv) => { const eq = kv.indexOf("="); out.params.push(["query", eq>=0?kv.slice(0,eq):kv, eq>=0?decodeURIComponent(kv.slice(eq+1)): ""]); });
  }
  const body = bodyLines.join("\n");
  if (body.includes("=") && !body.includes("<")) {
    body.split("&").forEach((kv) => { const eq = kv.indexOf("="); if (eq>=0) out.params.push(["body", kv.slice(0,eq), decodeURIComponent(kv.slice(eq+1))]); });
  }
  out.body = body;
  return out;
}
function renderInspector(target, raw) {
  const p = parseRequestForInspector(raw);
  if (!raw) { target.innerHTML = `<p class="muted">${tr("empty.inspector")}</p>`; return; }
  const row = (k, v) => `<div class="insp-row"><span class="k">${esc(k)}</span><span class="v">${esc(String(v))}</span></div>`;
  const rows = (arr) => arr.map((r) => row(r[0], r[1])).join("");
  target.innerHTML = `
    <div class="insp-group"><h4>Request attributes (1)</h4>${row("Method", p.attrs.method)}${row("Path", p.attrs.path)}${row("Protocol", p.attrs.version)}</div>
    ${p.requestHeaders.length?`<div class="insp-group"><h4>Headers (${p.requestHeaders.length})</h4>${rows(p.requestHeaders)}</div>`:""}
    ${p.params.length?`<div class="insp-group"><h4>Parameters (${p.params.length})</h4>${p.params.map((r)=>row(r[0]+"."+r[1], r[2])).join("")}</div>`:""}
    ${p.cookies.length?`<div class="insp-group"><h4>Cookies (${p.cookies.length})</h4>${rows(p.cookies)}</div>`:""}
  `;
}

/* Map */
let mapState = { target: null, catalog: [], stages: [], artifacts: [], map: null, opts: loadStoredMapOpts() };

function loadStoredMapOpts() {
  try {
    const raw = localStorage.getItem("aura_map_opts");
    if (raw) return JSON.parse(raw) || {};
  } catch (_) {}
  return {};
}

const MAP_STAGE_DEFAULTS = {
  subdomains_passive: { timeout_sec: 45, limit: 400 },
  dns_brute: { timeout_sec: 5, workers: 32, limit: 5000, wordlist: "" },
  live_hosts: { timeout_sec: 8, https: true, http: true },
  web_ports: { timeout_sec: 3, ports: "80 443 3000 3001 4000 4443 5000 5001 7001 8000 8008 8080 8081 8443 8888 9000 9090 9443 10443" },
  tech: { timeout_sec: 10 },
  urls_passive: { timeout_sec: 45, limit: 400, same_host: true },
  scrape: { timeout_sec: 12, follow_js: 8, same_host: true },
  dirs: { timeout_sec: 8, workers: 16, rps: 25, hide: "404", wordlist: "" },
  params: { timeout_sec: 8, workers: 10, rps: 15, hide: "404", wordlist: "" },
};

function mapStageOpts(id) {
  const d = MAP_STAGE_DEFAULTS[id] || {};
  const s = (mapState.opts && mapState.opts[id]) || {};
  const out = { ...d, ...s };
  if (Array.isArray(out.wordlist)) out.wordlist = out.wordlist.join("\n");
  if (Array.isArray(out.hide)) out.hide = out.hide.join(", ");
  if (Array.isArray(out.schemes)) {
    out.https = out.schemes.includes("https");
    out.http = out.schemes.includes("http");
  }
  return out;
}

function stageField(label, inner) {
  const wide = String(inner).includes("<textarea") ? " stage-field-wide" : "";
  return `<label class="stage-field${wide}"><span>${esc(label)}</span>${inner}</label>`;
}

function mapStageFields(id, busy) {
  const o = mapStageOpts(id);
  const dis = busy ? " disabled" : "";
  const num = (name, val, extra = "") => `<input type="number" name="${name}" value="${esc(val)}"${dis} ${extra}>`;
  const text = (name, val, extra = "") => `<input type="text" name="${name}" value="${esc(val || "")}"${dis} ${extra}>`;
  const area = (name, val, ph) => `<textarea name="${name}" rows="4" placeholder="${esc(ph || "")}"${dis}>${esc(val || "")}</textarea>`;
  const chk = (name, on, lab) => `<label class="check"><input type="checkbox" name="${name}"${on ? " checked" : ""}${dis}/> <span>${esc(lab)}</span></label>`;
  let fields = "";
  switch (id) {
    case "subdomains_passive":
      fields = stageField(tr("map.opt.timeout"), num("timeout_sec", o.timeout_sec, "min=\"5\" max=\"120\""))
        + stageField(tr("map.opt.limit"), num("limit", o.limit, "min=\"10\" max=\"2000\""));
      break;
    case "dns_brute":
      fields = stageField(tr("map.opt.timeout"), num("timeout_sec", o.timeout_sec, "min=\"1\" max=\"30\""))
        + stageField(tr("map.opt.workers"), num("workers", o.workers, "min=\"1\" max=\"128\""))
        + stageField(tr("map.opt.limit"), num("limit", o.limit, "min=\"0\" max=\"20000\""))
        + stageField(tr("map.opt.seclists"), text("wordlist_path", o.wordlist_path || "", `placeholder="Discovery/DNS/subdomains-top1million-5000.txt"`))
        + stageField(tr("map.opt.wordlist"), area("wordlist", o.wordlist, tr("map.opt.wordlistDnsPh")));
      break;
    case "live_hosts":
      fields = stageField(tr("map.opt.timeout"), num("timeout_sec", o.timeout_sec, "min=\"1\" max=\"30\""))
        + `<div class="stage-checks">${chk("https", o.https !== false, "HTTPS")}${chk("http", o.http !== false, "HTTP")}</div>`;
      break;
    case "web_ports":
      fields = stageField(tr("map.opt.timeout"), num("timeout_sec", o.timeout_sec, "min=\"1\" max=\"15\""))
        + stageField(tr("map.opt.ports"), area("ports", o.ports, "80 443 8080 8443"));
      break;
    case "tech":
      fields = stageField(tr("map.opt.timeout"), num("timeout_sec", o.timeout_sec, "min=\"1\" max=\"30\""));
      break;
    case "urls_passive":
      fields = stageField(tr("map.opt.timeout"), num("timeout_sec", o.timeout_sec, "min=\"10\" max=\"90\""))
        + stageField(tr("map.opt.limit"), num("limit", o.limit, "min=\"20\" max=\"2000\""))
        + `<div class="stage-checks">${chk("same_host", o.same_host !== false, tr("map.opt.sameHost"))}</div>`;
      break;
    case "scrape":
      fields = stageField(tr("map.opt.timeout"), num("timeout_sec", o.timeout_sec, "min=\"3\" max=\"30\""))
        + stageField(tr("map.opt.followJs"), num("follow_js", o.follow_js, "min=\"0\" max=\"40\""))
        + `<div class="stage-checks">${chk("same_host", o.same_host !== false, tr("map.opt.sameHost"))}</div>`;
      break;
    case "dirs":
      fields = stageField(tr("map.opt.timeout"), num("timeout_sec", o.timeout_sec, "min=\"2\" max=\"30\""))
        + stageField(tr("map.opt.workers"), num("workers", o.workers, "min=\"1\" max=\"64\""))
        + stageField(tr("map.opt.rps"), num("rps", o.rps, "min=\"1\" max=\"200\""))
        + stageField(tr("map.opt.hide"), text("hide", o.hide || "404"))
        + stageField(tr("map.opt.seclists"), text("wordlist_path", o.wordlist_path || "", `placeholder="Discovery/Web-Content/common.txt"`))
        + stageField(tr("map.opt.wordlist"), area("wordlist", o.wordlist, tr("map.opt.wordlistDirPh")));
      break;
    case "params":
      fields = stageField(tr("map.opt.timeout"), num("timeout_sec", o.timeout_sec, "min=\"2\" max=\"30\""))
        + stageField(tr("map.opt.workers"), num("workers", o.workers, "min=\"1\" max=\"64\""))
        + stageField(tr("map.opt.rps"), num("rps", o.rps, "min=\"1\" max=\"200\""))
        + stageField(tr("map.opt.hide"), text("hide", o.hide || "404"))
        + stageField(tr("map.opt.seclists"), text("wordlist_path", o.wordlist_path || "", `placeholder="Discovery/Web-Content/burp-parameter-names.txt"`))
        + stageField(tr("map.opt.wordlist"), area("wordlist", o.wordlist, tr("map.opt.wordlistParamPh")));
      break;
    default:
      return "";
  }
  return `<fieldset class="stage-fields"${busy ? " disabled" : ""}>${fields}</fieldset>`;
}

function readStageForm(card) {
  const el = (n) => card.querySelector(`[name="${n}"]`);
  const num = (n) => {
    const node = el(n);
    if (!node) return undefined;
    const x = Number(node.value);
    return Number.isFinite(x) ? x : undefined;
  };
  const text = (n) => (el(n)?.value || "").trim();
  const lines = (n) => text(n).split(/\n/).map((s) => s.trim()).filter(Boolean);
  const chk = (n) => !!el(n)?.checked;
  const id = card.dataset.stage;
  const o = {};
  switch (id) {
    case "subdomains_passive":
      o.timeout_sec = num("timeout_sec");
      o.limit = num("limit");
      break;
    case "dns_brute":
      o.timeout_sec = num("timeout_sec");
      o.workers = num("workers");
      o.limit = num("limit");
      o.wordlist = lines("wordlist");
      o.wordlist_path = text("wordlist_path");
      break;
    case "live_hosts":
      o.timeout_sec = num("timeout_sec");
      o.schemes = [];
      if (chk("https")) o.schemes.push("https");
      if (chk("http")) o.schemes.push("http");
      if (!o.schemes.length) o.schemes = ["https", "http"];
      break;
    case "web_ports":
      o.timeout_sec = num("timeout_sec");
      o.ports = text("ports");
      break;
    case "tech":
      o.timeout_sec = num("timeout_sec");
      break;
    case "urls_passive":
      o.timeout_sec = num("timeout_sec");
      o.limit = num("limit");
      o.same_host = chk("same_host");
      break;
    case "scrape":
      o.timeout_sec = num("timeout_sec");
      o.follow_js = num("follow_js");
      o.same_host = chk("same_host");
      break;
    case "dirs":
    case "params":
      o.timeout_sec = num("timeout_sec");
      o.workers = num("workers");
      o.rps = num("rps");
      o.wordlist = lines("wordlist");
      o.wordlist_path = text("wordlist_path");
      o.hide = text("hide").split(/[,\s]+/).map(Number).filter((n) => n > 0);
      break;
  }
  return o;
}

function collectMapOpts() {
  const out = { ...(mapState.opts || {}) };
  $$("#view-map .map-stage-page[data-stage]").forEach((card) => {
    if (!card.dataset.stage) return;
    out[card.dataset.stage] = { ...(out[card.dataset.stage] || {}), ...readStageForm(card) };
  });
  mapState.opts = out;
  try { localStorage.setItem("aura_map_opts", JSON.stringify(out)); } catch (_) {}
  return out;
}

async function loadMap() {
  const skipIngest = anyMapRunning();
  if (!mapState.target) {
    const saved = localStorage.getItem("aura_map_id") || localStorage.getItem("meb_map_id");
    if (saved) {
      try {
        const path = skipIngest ? `/api/intel/targets/${saved}` : `/api/intel/targets/${saved}/ingest-history`;
        const data = skipIngest ? await api(path) : await api(path, { method: "POST" });
        renderMapSnapshot(data);
        return;
      } catch (_) {
        try {
          renderMapSnapshot(await api(`/api/intel/targets/${saved}`));
          return;
        } catch (__) {
          localStorage.removeItem("aura_map_id");
          localStorage.removeItem("meb_map_id");
        }
      }
    }
    await suggestMapFromProxy();
    return;
  }
  try {
    if (skipIngest) {
      renderMapSnapshot(await api(`/api/intel/targets/${mapState.target.id}`));
    } else {
      renderMapSnapshot(await api(`/api/intel/targets/${mapState.target.id}/ingest-history`, { method: "POST" }));
    }
  } catch (_) {
    renderMapSnapshot(await api(`/api/intel/targets/${mapState.target.id}`));
  }
}

async function ingestMapHistory() {
  if (!mapState.target || anyMapRunning()) return;
  renderMapSnapshot(await api(`/api/intel/targets/${mapState.target.id}/ingest-history`, { method: "POST" }));
}

function mapJobId(stageId) { return "map:" + stageId; }
function anyMapRunning() { return Jobs.prefix("map:").length > 0; }
let mapIngestTimer = null;
function scheduleMapIngest() {
  if (anyMapRunning()) return;
  if (!mapState.target) {
    if (!$("#mapWork") || $("#mapWork").classList.contains("hidden")) {
      loadSiteMap();
      suggestMapFromProxy();
    }
    return;
  }
  clearTimeout(mapIngestTimer);
  mapIngestTimer = setTimeout(() => { ingestMapHistory().catch(() => {}); }, 700);
}

async function suggestMapFromProxy() {
  const box = $("#mapProxyHosts");
  if (!box) return;
  try {
    const data = await api("/api/history?limit=400");
    const counts = new Map();
    for (const f of data.items || []) {
      let h = String(f.host || "").toLowerCase();
      if (!h || h === "127.0.0.1" || h === "localhost" || h.startsWith("127.0.0.1:")) continue;
      counts.set(h, (counts.get(h) || 0) + 1);
    }
    const hosts = [...counts.entries()].sort((a, b) => b[1] - a[1]).slice(0, 10);
    if (!hosts.length) {
      box.innerHTML = `<p class="muted">${tr("map.noProxyYet")}</p>`;
      return;
    }
    box.innerHTML = `<p class="muted">${tr("map.fromProxy")}</p><div class="map-host-list">${
      hosts.map(([h, n]) => `<button type="button" class="mini" data-map-host="${esc(h)}">${esc(h)} · ${n}</button>`).join("")
    }</div>`;
  } catch (_) {
    box.innerHTML = `<p class="muted">${tr("map.noProxyYet")}</p>`;
  }
}

function clipText(s, n) {
  s = String(s || "");
  return s.length > n ? s.slice(0, n) + "…" : s;
}

function intelProgressLine(m) {
  if (!m) return "";
  const vars = {
    stage: m.stage ? tr("stage." + m.stage + ".title") : "",
    n: m.n ?? 0,
    total: m.total ?? 0,
    url: m.url || "",
    host: m.host || "",
    status: m.status || "",
    err: m.err || "",
    value: m.value || m.url || m.host || "",
    timeout: m.timeout || "",
  };
  let action = m.action || "";
  if (action === "get" && m.timeout) action = "getWait";
  if (action === "hit" && m.status) action = "hitStatus";
  const key = "map.prog." + action;
  let line = tr(key, vars);
  if (!line || line === key) line = m.msg || "";
  if (vars.stage && line && m.action !== "start") line = vars.stage + " · " + line;
  return line;
}

const MapLog = {
  items: [],
  last: {},
  max: 400,
  latest() {
    return this.items.length ? this.items[this.items.length - 1].line : "";
  },
  lastFor(stage) {
    return this.last[stage] || "";
  },
  append(ev) {
    const line = intelProgressLine(ev);
    if (!line) return;
    const prev = this.items[this.items.length - 1];
    if (prev && prev.line === line) return;
    this.items.push({ ...ev, line });
    if (this.items.length > this.max) this.items.splice(0, this.items.length - this.max);
    if (ev && ev.stage) this.last[ev.stage] = line;
    this.render();
    const live = ev && ev.stage ? document.querySelector(`.map-stage-page[data-stage="${ev.stage}"] .stage-live`) : null;
    if (live) live.textContent = line;
    if (Jobs.prefix("map:").length) setWork(line);
  },
  clear() {
    this.items = [];
    this.last = {};
    this.render();
  },
  render() {
    const el = $("#mapLog");
    if (!el) return;
    el.dataset.empty = tr("map.logIdle");
    el.textContent = this.items.map((x) => x.line).join("\n");
    el.scrollTop = el.scrollHeight;
    $$(".map-stage-log").forEach((pre) => {
      const id = pre.closest("[data-stage]")?.dataset.stage;
      pre.textContent = this.items.filter((x) => x.stage === id).map((x) => x.line).join("\n");
      pre.scrollTop = pre.scrollHeight;
    });
  },
};

function clearMapWorkspace() {
  const opts = mapState.opts || loadStoredMapOpts();
  mapState = { target: null, catalog: mapState.catalog || [], stages: [], artifacts: [], map: null, opts };
  try {
    localStorage.removeItem("aura_map_id");
    localStorage.removeItem("meb_map_id");
  } catch (_) {}
  if ($("#mapStats")) $("#mapStats").innerHTML = "";
  if ($("#mapMeta")) $("#mapMeta").textContent = "";
  const detail = $("#mapDetail");
  if (detail) detail.textContent = tr("map.pickNode");
  MapLog.clear();
  renderMapStagePanes();
  renderSiteMap();
  suggestMapFromProxy();
}

function displayMapHost(h, base) {
  const host = String(h.host || "");
  try {
    const u = new URL(base);
    if (u.hostname.toLowerCase() === host.toLowerCase() && u.port) return host + ":" + u.port;
  } catch (_) {}
  return host;
}

function mapCatalog() {
  if (mapState.catalog && mapState.catalog.length) return mapState.catalog;
  return ["subdomains_passive", "dns_brute", "live_hosts", "web_ports", "tech", "urls_passive", "scrape", "dirs", "params"]
    .map((id) => ({ id, mode: id === "subdomains_passive" || id === "urls_passive" ? "passive" : "active" }));
}

function syncMapSubtabBusy() {
  $$("#mapSubtabs button[data-sub]").forEach((b) => {
    const id = b.dataset.sub;
    if (id === "sitemap" || id === "target") return;
    const running = Jobs.running(mapJobId(id));
    b.classList.toggle("is-busy", running);
  });
}

function renderMapStagePanes() {
  collectMapOpts();
  const byId = Object.fromEntries((mapState.stages || []).map((s) => [s.id, s]));
  const catalog = mapCatalog();
  catalog.forEach((c, i) => {
    const host = document.querySelector(`#view-map .map-stage-page[data-stage="${c.id}"]`);
    if (!host) return;
    const s = byId[c.id] || { status: "idle" };
    const busy = s.status === "running" || Jobs.running(mapJobId(c.id));
    const status = busy ? "running" : (s.status || "idle");
    const stKey = "map.status." + status;
    let summary = s.summary || "";
    if (summary === "stopped" || summary === "остановлено") summary = tr("work.stopped");
    if (summary && !busy) summary = clipText(summary, 220);
    const live = busy ? (MapLog.lastFor(c.id) || tr("map.stageWork." + c.id)) : "";
    const ready = !!mapState.target;
    host.className = `map-stage-page ${status}`;
    host.innerHTML = `
      <div class="stage-head">
        <h2>${i + 1}. ${esc(tr("stage." + c.id + ".title"))}</h2>
        <span class="stage-mode ${c.mode}">${c.mode === "active" ? tr("map.modeActive") : tr("map.modePassive")} · ${esc(tr(stKey))}</span>
      </div>
      <p class="map-stage-hint">${esc(tr("stage." + c.id + ".hint"))}</p>
      ${!ready ? `<p class="muted">${esc(tr("map.needStart"))}</p>` : ""}
      ${mapStageFields(c.id, busy || !ready)}
      ${summary && !busy ? `<p>${esc(summary)}</p>` : ""}
      ${busy ? `<p class="stage-live">${esc(live)}</p>` : ""}
      ${busy ? `<div class="stage-progress" aria-hidden="true"><span></span></div>` : ""}
      <div class="toolbar job-actions">
        <button type="button" class="primary${busy ? " is-busy" : ""}" data-run="${c.id}" ${busy || !ready ? "disabled" : ""}>${busy ? `<span class="spin"></span>${tr("map.running")}` : tr("map.run")}</button>
        <button type="button" class="danger" data-stop="${c.id}" ${busy ? "" : "disabled"}>${tr("map.stop")}</button>
        ${c.id === "dirs" ? `<button type="button" data-jump="discover">${tr("map.openPaths")}</button>` : ""}
        ${c.id === "params" ? `<button type="button" data-jump="fuzz">${tr("map.openFuzz")}</button>` : ""}
      </div>
      <pre class="map-log map-stage-log" data-empty="${esc(tr("map.logIdle"))}"></pre>`;
  });
  syncMapSubtabBusy();
  MapLog.render();
}

function renderMapSnapshot(data) {
  mapState = { ...mapState, ...data, target: data.target, opts: mapState.opts };
  if (data.target?.id) localStorage.setItem("aura_map_id", data.target.id);
  if (!data.target) {
    clearMapWorkspace();
    return;
  }
  const field = $("#mapTarget");
  if (field && document.activeElement !== field) {
    field.value = data.target.base_url || data.target.domain || "";
  }
  $("#mapAuth").checked = !!data.target.authorized;
  $("#mapMeta").textContent = (data.target.base_url || data.target.domain) + (data.target.authorized ? tr("map.ownTarget") : tr("map.passiveOnly"));
  const st = data.map?.stats || {};
  $("#mapStats").innerHTML = [
    [tr("map.statHosts"), st.live_hosts ?? 0],
    [tr("map.statNames"), st.subdomains ?? 0],
    [tr("map.statPaths"), st.paths ?? 0],
    [tr("map.statParams"), st.params ?? 0],
    ["JS", st.js ?? 0],
    ["URL", st.urls ?? 0],
  ].map(([k, v]) => `<div class="map-stat"><b>${v}</b><span>${k}</span></div>`).join("");
  renderMapStagePanes();
  renderSiteMap();
  requestAnimationFrame(() => {
    if (UILayout && UILayout.refresh) UILayout.refresh();
    MapLog.render();
  });
}

function renderAppMap(m) {
  if (m) mapState.map = m;
  renderSiteMap();
}

$("#mapTarget")?.addEventListener("input", () => {
  if (!$("#mapTarget").value.trim() && mapState.target) clearMapWorkspace();
});
$("#btnMapClear")?.addEventListener("click", () => {
  if ($("#mapTarget")) $("#mapTarget").value = "";
  if ($("#mapAuth")) $("#mapAuth").checked = false;
  clearMapWorkspace();
});
$("#mapForm").addEventListener("submit", async (e) => {
  e.preventDefault();
  if (!$("#mapTarget").value.trim()) {
    clearMapWorkspace();
    return;
  }
  const startBtn = $("#btnMapStart");
  markBusy(startBtn, true, tr("map.starting"));
  setWork(tr("map.starting"));
  try {
    const data = await api("/api/intel/targets", {
      method: "POST",
      body: JSON.stringify({ target: $("#mapTarget").value.trim(), authorized: $("#mapAuth").checked }),
    });
    renderMapSnapshot(data);
    setWork(tr("map.ingesting"));
    MapLog.append({ msg: tr("map.ingesting") });
    let n = 0;
    try {
      const ing = await api(`/api/intel/targets/${data.target.id}/ingest-history`, { method: "POST" });
      n = ing.ingested || 0;
      renderMapSnapshot(ing);
    } catch (_) {}
    $("#mapMeta").textContent = tr("map.startedOk", { host: data.target.base_url || data.target.domain || "", n });
    showSub("#view-map", "sitemap");
  } catch (err) { uiFlash(err.message); }
  finally {
    markBusy(startBtn, false);
    Jobs.syncBar();
  }
});
$("#mapProxyHosts")?.addEventListener("click", (e) => {
  const b = e.target.closest("[data-map-host]");
  if (!b) return;
  $("#mapTarget").value = b.dataset.mapHost;
  $("#mapTarget").focus();
});
$("#view-map").addEventListener("change", (e) => {
  if (e.target.closest(".map-stage-page")) collectMapOpts();
});
$("#view-map").addEventListener("input", (e) => {
  if (e.target.closest(".map-stage-page")) collectMapOpts();
});
$("#view-map").addEventListener("click", async (e) => {
  const jump = e.target.closest("[data-jump]");
  if (jump) { showView(jump.dataset.jump); return; }
  const stop = e.target.closest("[data-stop]");
  if (stop) {
    const stageId = stop.dataset.stop;
    Jobs.stop(mapJobId(stageId));
    MapLog.append({ stage: stageId, action: "stop" });
    mapState.stages = (mapState.stages || []).map((s) => s.id === stageId ? { ...s, status: "idle", summary: "stopped" } : s);
    renderMapSnapshot(mapState);
    $("#mapMeta").textContent = tr("work.stopped");
    return;
  }
  const btn = e.target.closest("[data-run]");
  if (!btn || !mapState.target || btn.disabled) return;
  const stageId = btn.dataset.run;
  const jobId = mapJobId(stageId);
  if (Jobs.running(jobId)) return;
  collectMapOpts();
  const opts = { ...(mapState.opts && mapState.opts[stageId]) || {} };
  const hint = tr("map.stageWork." + stageId);
  const label = hint && hint !== "map.stageWork." + stageId ? hint : tr("map.stageWorking");
  const signal = Jobs.start(jobId, label);
  MapLog.append({ stage: stageId, action: "start", host: mapState.target.domain || mapState.target.base_url || "" });
  mapState.stages = [...(mapState.stages || []).filter((s) => s.id !== stageId), { id: stageId, status: "running", summary: "" }];
  renderMapSnapshot(mapState);
  try {
    const data = await api(`/api/intel/targets/${mapState.target.id}/stages/${stageId}`, {
      method: "POST",
      signal,
      body: JSON.stringify(opts),
    });
    Jobs.finish(jobId);
    renderMapSnapshot(data);
    if (data.stopped) $("#mapMeta").textContent = tr("work.stopped");
    else if (data.error) {
      MapLog.append({ stage: stageId, action: "err", err: data.error, host: mapState.target.domain || "" });
      $("#mapMeta").textContent = clipText(data.error, 180);
    }
    else $("#mapMeta").textContent = data.run?.stage?.summary || tr("map.startedOk", { host: mapState.target.domain, n: data.run?.added || 0 });
  } catch (err) {
    Jobs.finish(jobId);
    if (isAbortError(err)) {
      $("#mapMeta").textContent = tr("work.stopped");
      try {
        const snap = await api(`/api/intel/targets/${mapState.target.id}`);
        const stages = (snap.stages || []).map((s) => s.id === stageId && s.status === "running"
          ? { ...s, status: "idle", summary: "stopped" } : s);
        renderMapSnapshot({ ...snap, stages });
      } catch (_) {
        mapState.stages = (mapState.stages || []).map((s) => s.id === stageId ? { ...s, status: "idle", summary: "stopped" } : s);
        renderMapSnapshot(mapState);
      }
      return;
    }
    uiFlash(err.message);
    MapLog.append({ stage: stageId, action: "err", err: err.message });
    mapState.stages = (mapState.stages || []).map((s) => s.id === stageId ? { ...s, status: "error", summary: err.message } : s);
    renderMapSnapshot(mapState);
  }
});
$("#btnMapHistory").addEventListener("click", async () => {
  if (!mapState.target) return;
  await withWork(tr("work.history"), $("#btnMapHistory"), async () => {
    renderMapSnapshot(await api(`/api/intel/targets/${mapState.target.id}/ingest-history`, { method: "POST" }));
  }).catch((err) => uiFlash(err.message));
});
$("#btnMapRefresh").addEventListener("click", async () => {
  await withWork(tr("work.history"), $("#btnMapRefresh"), async () => {
    await loadSiteMap();
    await loadMap();
  }).catch((err) => uiFlash(err.message));
});
$("#btnMapLogClear")?.addEventListener("click", () => MapLog.clear());

$("#btnDiscBuiltin")?.addEventListener("click", async () => {
  const w = await api("/api/wordlists");
  $("#discWords").value = (w.dirs || []).join("\n");
});
$("#discWlSearch")?.addEventListener("input", () => fillWlSelect($("#discWlSelect"), $("#discWlSearch").value));
$("#fuzzWlSearch")?.addEventListener("input", () => fillWlSelect($("#fuzzWlSelect"), $("#fuzzWlSearch").value));

$("#btnFuzzDirs")?.addEventListener("click", async () => {
  const w = await api("/api/wordlists");
  $("#fuzzWords").value = (w.dirs || []).join("\n");
  if (!$("#fuzzUrl").value) $("#fuzzUrl").value = "https://example.com/FUZZ";
});
$("#btnFuzzParams")?.addEventListener("click", async () => {
  const w = await api("/api/wordlists");
  $("#fuzzWords").value = (w.params || []).join("\n");
  if (!$("#fuzzUrl").value.includes("FUZZ")) $("#fuzzUrl").value = "https://example.com/?FUZZ=1";
});
function requireLabAuth(box) {
  if (box && box.checked) return true;
  uiFlash(tr("lab.needAuth"));
  box?.focus();
  return false;
}

$("#btnFuzzRun")?.addEventListener("click", async () => {
  const hide = ($("#fuzzHide").value || "").split(",").map((s) => Number(s.trim())).filter(Boolean);
  const custom = $("#fuzzWords").value.split("\n").map((s) => s.trim()).filter(Boolean);
  const path = $("#fuzzWlSelect")?.value || "";
  if (!requireLabAuth($("#fuzzAuth"))) return;
  try {
    const data = await runJob("fuzz", {
      label: tr("work.fuzz"),
      runBtn: $("#btnFuzzRun"),
      stopBtn: $("#btnFuzzStop"),
      meta: $("#fuzzMeta"),
      fn: (signal) => api("/api/fuzz/run", {
        method: "POST",
        signal,
        body: JSON.stringify({
          url: $("#fuzzUrl").value.trim(),
          method: $("#fuzzMethod").value,
          body: $("#fuzzBody").value,
          wordlist: custom,
          wordlist_path: custom.length ? "" : path,
          workers: Number($("#fuzzWorkers").value),
          rps: Number($("#fuzzRps").value),
          hide,
          authorized: true,
        }),
      }),
    });
    if (!data || data.aborted) return;
    const items = data.items || [];
    $("#fuzzBodyRows").innerHTML = items.map((h) => `<tr>
      <td>${esc(h.payload)}</td>
      <td class="${statusClass(h.status_code)}">${h.status_code || "—"}</td>
      <td>${h.length ?? "—"}</td>
      <td>${Math.round((h.duration || 0) / 1e6)}</td>
      <td title="${esc(h.url)}">${esc((h.url || "").slice(0, 80))}</td>
    </tr>`).join("") || `<tr><td colspan="5" class="muted">${tr("empty.hits")}</td></tr>`;
    $("#fuzzMeta").textContent = `${items.length} hits / ${data.tried || items.length} tried` + (data.truncated ? tr("fuzz.truncated") : "");
  } catch (err) { $("#fuzzMeta").textContent = err.message; }
});

function refreshTranslatedUI() {
  try { renderPending(); } catch (_) {}
  try { if (typeof mapState !== "undefined" && mapState.target) renderMapSnapshot(mapState); else if (typeof renderMapStagePanes === "function") renderMapStagePanes(); } catch (_) {}
  try { if (typeof renderSiteMap === "function" && !(mapState && mapState.target)) renderSiteMap(); } catch (_) {}
  try { if (typeof MapLog !== "undefined") MapLog.render(); } catch (_) {}
  try { renderMr(); } catch (_) {}
  try { renderScope(); } catch (_) {}
  try { renderExt(); } catch (_) {}
  try { renderIssueDefs(); } catch (_) {}
  try { if (typeof discCache !== "undefined") renderDisc(discCache); } catch (_) {}
  try { if (typeof loadSecLists === "function") loadSecLists(); } catch (_) {}
  const detail = $("#mapDetail");
  if (detail && !detail.querySelector("strong")) detail.textContent = tr("map.pickNode");
  const insp = $("#repInspector");
  const req = $("#repReq");
  if (insp && req) renderInspector(insp, req.value);
  try { UILayout.syncLockButton(); } catch (_) {}
  try { if ($("#mainTabs")) $("#mainTabs").title = (UILayout.data && UILayout.data.uiLocked) ? "" : tr("set.dragTabs"); } catch (_) {}
}

/* Init */
try { UILayout.init(); } catch (e) { DebugLog.log({ level: "error", module: "ui", action: "layout", msg: String(e) }); }
try { applyLang(detectLang()); } catch (_) { applyLang("en"); }
try { MapLog.render(); renderMapStagePanes(); } catch (_) {}
$$("#themeSwitch button").forEach((b) => b.addEventListener("click", () => applyTheme(b.dataset.theme)));
$$("#langSwitch button").forEach((b) => b.addEventListener("click", () => applyLang(b.dataset.lang)));
(function bootView() {
  const allowed = new Set($$("#mainTabs button").filter((b) => !b.hidden && !b.classList.contains("hidden")).map((b) => b.dataset.view));
  const last = UILayout.data && UILayout.data.lastView;
  showView(allowed.has(last) ? last : "map");
  const sub = UILayout.data && UILayout.data.lastProxySub;
  if (sub && sub !== "settings") showSub("#view-proxy", sub);
})();
loadStatus();
loadHistory();
loadIntercept();
loadLogs();
connectWs();
