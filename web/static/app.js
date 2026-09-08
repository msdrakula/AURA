const $ = (sel) => document.querySelector(sel);
const $$ = (sel) => [...document.querySelectorAll(sel)];

function applyLang(name) {
  setLang(name);
  applyI18nDom();
  $$("#langSwitch button").forEach((b) => b.classList.toggle("active", b.dataset.lang === mebLang));
  refreshTranslatedUI();
}

function applyTheme(name) {
  const t = name === "light" ? "light" : "dark";
  document.documentElement.setAttribute("data-theme", t);
  try { localStorage.setItem("meb_theme", t); } catch (_) {}
  $$("#themeSwitch button").forEach((b) => b.classList.toggle("active", b.dataset.theme === t));
}

const state = { pending: [], selectedFlow: null, histReq: "", histResp: "", histScheme: "https", intrResults: [] };

async function api(path, opts = {}) {
  const res = await fetch(path, { headers: { "Content-Type": "application/json", ...(opts.headers || {}) }, ...opts });
  if (!res.ok) { let msg = res.statusText; try { const d = await res.json(); msg = d.detail || JSON.stringify(d); } catch (_) {} throw new Error(msg); }
  const ct = res.headers.get("content-type") || "";
  return ct.includes("application/json") ? res.json() : res.text();
}

function showView(name) {
  $$(".view").forEach((el) => el.classList.add("hidden"));
  const pane = $(`#view-${name}`);
  if (pane) pane.classList.remove("hidden");
  $$("#mainTabs button").forEach((b) => b.classList.toggle("active", b.dataset.view === name));
  if (name === "intruder" && typeof intruder !== "undefined" && intruder.tabs.length === 0) newIntruderTab("");
  if (name === "discover" || name === "fuzz") loadSecLists();
  if (name === "map") { loadMap(); }
  if (name === "dashboard") refreshDashboard();
  if (name === "target") { loadSiteMap(); loadIssues(); renderScope(); renderIssueDefs(); }
  if (name === "proxy") { renderMr(); }
  if (name === "organizer") { loadOrganizer(); }
  if (name === "logger") { loadLogs(); }
  if (name === "repeater" && repeater.tabs.length === 0) { newRepeaterTab(""); }
  if (name === "scanner" && state.histReq) { if ($("#scanFromHistory").checked) $("#scanRaw").value = state.histReq; }
}
function showSub(group, name) {
  $$(`${group} .subpane`).forEach((el) => el.classList.toggle("hidden", el.dataset.subpane !== name));
  $$(`${group} .subtabs button`).forEach((b) => b.classList.toggle("active", b.dataset.sub === name));
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
      <td><button class="mini" data-act="rep" data-id="${f.id}">→Repeater</button> <button class="mini" data-act="intr" data-id="${f.id}">→Intruder</button></td>
    </tr>`).join("") || `<tr><td colspan="5" class="muted">${tr("empty.none")}</td></tr>`;
  });
}

function renderPending(force = false) {
  const n = state.pending.length;
  $("#pendingBadge").classList.toggle("hidden", n === 0);
  $("#pendingBadge").textContent = `${n} in queue`;
  $("#proxyDot").classList.toggle("hidden", n === 0);
  $("#interceptDot").classList.toggle("hidden", n === 0);
  const list = $("#interceptList");
  list.innerHTML = n ? `<table><thead><tr><th>Time</th><th>Type</th><th>Dir</th><th>Method</th><th>URL</th></tr></thead><tbody>` +
    state.pending.map((p, i) => `<tr data-i="${i}" class="${i === 0 ? "sel" : ""}">
      <td>${new Date((p.flow_id || 0) * 1000 || Date.now()).toLocaleTimeString()}</td>
      <td>HTTP</td>
      <td>${p.phase === "response" ? "← resp" : "→ req"}</td>
      <td class="method ${p.method}">${esc(p.method)}</td>
      <td>${esc(p.url)}</td>
    </tr>`).join("") + `</tbody></table>` : `<div class="muted" style="padding:12px">${tr("empty.queue")}</div>`;
  const cur = state.pending[0];
  if (!cur) { $("#interceptEditor").value = ""; $("#interceptMeta").textContent = ""; return; }
  $("#interceptMeta").textContent = `${cur.phase.toUpperCase()}  ${cur.method}  ${cur.url}`;
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
    if (m.type === "hello") renderStatus({ ...m.status, stats: m.status.stats });
    if (m.type === "flow") { loadHistory(); refreshDashboard(); loadSiteMap(); if (currentView() === "logger") loadLogs(); }
    if (m.type === "intercept") { state.pending.push(m.item); renderPending(); }
    if (m.type === "intercept_done") { state.pending = state.pending.filter((p) => p.id !== m.id); renderPending(true); }
    if (m.type === "log") loadLogs();
    if (m.type === "history_cleared") loadHistory();
    if (m.type === "settings") loadStatus();
  };
  ws.onclose = () => setTimeout(connectWs, 1500);
}

$$("#mainTabs button").forEach((b) => b.addEventListener("click", () => showView(b.dataset.view)));
$$("#proxySubtabs button").forEach((b) => b.addEventListener("click", () => showSub("#view-proxy", b.dataset.sub)));

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
  try { const r = await api("/api/proxy/launch-browser", { method: "POST" }); alert(tr("browser.launched", { proxy: r.proxy })); }
  catch (e) { alert(e.message); }
}
$("#btnStart").addEventListener("click", startProxy);
$("#btnStop").addEventListener("click", stopProxy);
$("#btnLaunchBrowser").addEventListener("click", launchBrowser);

$("#btnForward").addEventListener("click", async () => {
  const c = state.pending[0]; if (!c) return;
  let raw = $("#interceptEditor").value;
  raw = applyMatchReplace(raw, c.phase);
  await api(`/api/intercept/${c.id}`, { method: "POST", body: JSON.stringify({ action: "forward", raw }) });
});
$("#btnForwardAll").addEventListener("click", () => api("/api/intercept/forward-all", { method: "POST" }));
$("#btnDrop").addEventListener("click", async () => {
  const c = state.pending[0]; if (!c) return;
  await api(`/api/intercept/${c.id}`, { method: "POST", body: JSON.stringify({ action: "drop" }) });
});
$("#btnDropAll").addEventListener("click", () => api("/api/intercept/drop-all", { method: "POST" }));
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
const mrRules = JSON.parse(localStorage.getItem("meb_mr_rules") || "[]");
function saveMr() { localStorage.setItem("meb_mr_rules", JSON.stringify(mrRules)); }
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
  const def = "GET / HTTP/1.1\r\nHost: example.com\r\nUser-Agent: MEB/0.1\r\nAccept: */*\r\n\r\n";
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
  $("#repMeta").textContent = "sending…";
  try {
    const data = await api("/api/repeater", { method: "POST", body: JSON.stringify({ raw: $("#repReq").value, scheme: $("#repScheme").value, target: $("#repTarget").value.trim() || null }) });
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
  const tab = { id: intruder.seq, name: String(intruder.seq), template: template || "", scheme: "https", target: "", attackType: "cluster_bomb", payloadSets: [""], results: [] };
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
  $("#intrPositionsCount").textContent = `${Math.floor(n)} payload positions`;
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
  if (!t.template) { alert(tr("intr.emptyTpl")); return; }
  const sets = t.payloadSets.map((p) => p.split("\n").map((x) => x.trim()).filter(Boolean));
  const positions = Math.floor((t.template.match(/§/g) || []).length / 2);
  if (positions === 0) { alert(tr("intr.noMarks")); return; }
  // validate set count per attack type
  if ((t.attackType === "sniper" || t.attackType === "battering_ram") && sets.length !== 1) { alert(tr("intr.needOneSet", { type: t.attackType, n: sets.length })); return; }
  if ((t.attackType === "pitchfork" || t.attackType === "cluster_bomb") && sets.length !== positions) { alert(tr("intr.needNSets", { type: t.attackType, need: positions, n: sets.length })); return; }
  if (sets.every((s) => s.length === 0)) { alert(tr("intr.emptySets")); return; }
  $("#btnIntrRun").disabled = true;
  t.results = [];
  try {
    const data = await api("/api/batch/execute", {
      method: "POST",
      body: JSON.stringify({ template_raw: t.template, scheme: t.scheme, target: t.target, attack_type: t.attackType, payload_sets: sets }),
    });
    t.results = data || []; renderIntruderResults();
  } catch (e) { alert(e.message); }
  $("#btnIntrRun").disabled = false;
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

/* TARGET: Site map, Scope, Issues */
$$("#targetSubtabs button").forEach((b) => b.addEventListener("click", () => showSub("#view-target", b.dataset.sub)));

const scope = {
  inc: JSON.parse(localStorage.getItem("meb_scope_inc") || "[]"),
  exc: JSON.parse(localStorage.getItem("meb_scope_exc") || "[]"),
};
function saveScope() {
  localStorage.setItem("meb_scope_inc", JSON.stringify(scope.inc));
  localStorage.setItem("meb_scope_exc", JSON.stringify(scope.exc));
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
async function loadSiteMap() { siteItems = (await api("/api/history?q=")).items; renderSiteMap(); }
function renderSiteMap() {
  const filter = ($("#siteFilter")?.value || "").toLowerCase();
  const applyScope = $("#siteApplyScope")?.checked;
  const items = siteItems.filter((f) => {
    if (filter && !(`${f.method} ${f.host} ${f.path}`.toLowerCase().includes(filter))) return false;
    if (applyScope && !scopeMatches(`${f.scheme || "https"}://${f.host}${f.path}`)) return false;
    return true;
  });
  const tree = {};
  for (const f of items) {
    const host = f.host || "?";
    const parts = (f.path || "/").split("/").filter(Boolean);
    let node = tree[host] = tree[host] || { _leaf: null, _children: {} };
    for (const p of parts) node = node._children[p] = node._children[p] || { _leaf: null, _children: {} };
    node._leaf = f;
  }
  $("#siteTree").innerHTML = Object.keys(tree).sort().map((host) => renderTreeNode(host, tree[host], true)).join("") || `<div class="muted">${tr("empty.sitemap")}</div>`;
}
function renderTreeNode(name, node, isHost) {
  const childKeys = Object.keys(node._children).sort();
  const hasChildren = childKeys.length > 0;
  const leaf = node._leaf;
  const id = "n" + Math.random().toString(36).slice(2, 9);
  const icon = isHost ? "🔒" : (leaf ? "📄" : "📁");
  return `<div class="tree-node">
    <div class="tree-row" data-id="${leaf ? leaf.id : ""}">
      <span class="tree-tw" data-toggle="${id}">${hasChildren ? "▶" : ""}</span>
      <span class="tree-icon">${icon}</span>
      <span>${esc(name)}</span>
      ${leaf ? `<span class="tree-method ${leaf.method}">${esc(leaf.method)}</span><span class="tree-status ${statusClass(leaf.status)}">${leaf.status || ""}</span>` : ""}
    </div>
    <div id="${id}" class="tree-children hidden">${childKeys.map((k) => renderTreeNode(k, node._children[k], false)).join("")}</div>
  </div>`;
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
  const id = row.dataset.id;
  if (!id) return;
  const data = await api(`/api/history/${id}`);
  siteSelectedId = id;
  $("#siteReq").textContent = data.request_raw || "";
  $("#siteResp").textContent = data.response_raw || data.error || "";
});
$("#btnSiteToRepeater").addEventListener("click", () => { if (siteSelectedId) flowToTool(siteSelectedId, "repeater"); });
$("#btnSiteToIntruder").addEventListener("click", () => { if (siteSelectedId) flowToTool(siteSelectedId, "intruder"); });
$("#btnSiteToScanner").addEventListener("click", () => { if (!siteSelectedId) return; api(`/api/history/${siteSelectedId}`).then((f) => { showView("scanner"); $("#scanRaw").value = f.request_raw || ""; $("#scanScheme").value = f.scheme || "https"; }); });
$$("#view-target .tabs.small button").forEach((b) => b.addEventListener("click", () => {
  $$("#view-target .tabs.small button").forEach((x) => x.classList.remove("active")); b.classList.add("active");
  const req = b.dataset.pane === "req";
  $("#siteReq").classList.toggle("hidden", !req); $("#siteResp").classList.toggle("hidden", req);
}));
$("#siteFilter").addEventListener("input", () => { clearTimeout(window._siteT); window._siteT = setTimeout(renderSiteMap, 200); });
$("#siteApplyScope").addEventListener("change", renderSiteMap);
$("#btnSiteRefresh").addEventListener("click", loadSiteMap);

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
  $("#btnSeqAnalyze").disabled = true;
  try {
    const rep = await api("/api/sequencer/analyze", { method: "POST", body: JSON.stringify({ tokens }) });
    renderSequencer(rep);
  } catch (e) { alert(e.message); }
  $("#btnSeqAnalyze").disabled = false;
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
  } catch (e) { alert(e.message); }
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
  } catch (e) { alert(e.message); }
});
$("#btnOrgDelete").addEventListener("click", async () => {
  if (!org.selectedId) return;
  if (!confirm(tr("org.confirmDel"))) return;
  try {
    await api(`/api/organizer/${encodeURIComponent(org.selectedId)}`, { method: "DELETE" });
    org.selectedId = null; loadOrganizer();
  } catch (e) { alert(e.message); }
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
  } catch (e) { alert(e.message); }
}
/* Send-to handlers for Dashboard recent flows and Target site map */
async function flowToTool(flowId, tool) {
  try {
    const f = await api(`/api/history/${flowId}`);
    const raw = f.request_raw || "";
    if (!raw) { alert(tr("flow.noReq")); return; }
    if (tool === "repeater") { showView("repeater"); newRepeaterTab(raw); }
    else if (tool === "intruder") { showView("intruder"); newIntruderTab(raw); }
  } catch (e) { alert(e.message); }
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
  if (!base || (custom.length === 0 && !path)) { alert(tr("disc.need")); return; }
  $("#btnDiscRun").disabled = true; $("#discMeta").textContent = tr("disc.scanning");
  $("#discBody").innerHTML = "";
  try {
    const data = await api("/api/discover/run", { method: "POST", body: JSON.stringify({
      base_url: base, wordlist: custom, wordlist_path: custom.length ? "" : path,
      workers: +$("#discWorkers").value || 10,
      rps: +$("#discRps").value || 20, cookies: $("#discCookies").value,
    }) });
    discCache = data || [];
    renderDisc(discCache);
    $("#discMeta").textContent = tr("disc.nResults", { n: discCache.length });
  } catch (e) { alert(e.message); $("#discMeta").textContent = ""; }
  $("#btnDiscRun").disabled = false;
});
$("#btnDiscStop").addEventListener("click", () => { /* best-effort: results already returned */ });
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
$("#btnScanRun").addEventListener("click", async () => {
  let raw = $("#scanRaw").value;
  if (!raw.trim()) { alert(tr("scan.need")); return; }
  $("#btnScanRun").disabled = true; $("#scanMeta").textContent = tr("scan.scanning");
  $("#scanBody").innerHTML = "";
  try {
    const data = await api("/api/scanner/scan", { method: "POST", body: JSON.stringify({
      raw, scheme: $("#scanScheme").value, target: $("#scanTarget").value.trim(),
    }) });
    renderScan(data || { findings: [] });
    $("#scanMeta").textContent = `${(data.findings||[]).length} findings, ${(data.took/1e9).toFixed(1)}s`;
  } catch (e) { alert(e.message); $("#scanMeta").textContent = ""; }
  $("#btnScanRun").disabled = false;
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

/* Extensions (client-side stub) */
const extLoaded = [];
$("#btnExtLoad").addEventListener("click", () => {
  const name = $("#extPath").value.trim();
  if (!name) { alert(tr("ext.needPath")); return; }
  extLoaded.push({ name, version: "1.0", author: "user", status: "loaded", description: tr("ext.stub") });
  renderExt();
});
$("#btnExtUnload").addEventListener("click", () => { extLoaded.pop(); renderExt(); });
function renderExt() {
  $("#extBody").innerHTML = extLoaded.length ? extLoaded.map((e) => `<tr>
    <td>${esc(e.name)}</td><td>${esc(e.version)}</td><td>${esc(e.author)}</td>
    <td>${esc(e.status)}</td><td>${esc(e.description)}</td>
  </tr>`).join("") : `<tr><td colspan="5" class="muted">${tr("ext.empty")}</td></tr>`;
}

/* Map */
let mapState = { target: null, catalog: [], stages: [], artifacts: [], map: null };

async function loadMap() {
  if (!mapState.target) {
    const saved = localStorage.getItem("meb_map_id");
    if (!saved) return;
    try {
      renderMapSnapshot(await api(`/api/intel/targets/${saved}`));
    } catch (_) {
      localStorage.removeItem("meb_map_id");
    }
  } else {
    renderMapSnapshot(await api(`/api/intel/targets/${mapState.target.id}`));
  }
}

function renderMapSnapshot(data) {
  mapState = { ...mapState, ...data, target: data.target };
  if (data.target?.id) localStorage.setItem("meb_map_id", data.target.id);
  $("#mapEmpty").classList.toggle("hidden", !!data.target);
  $("#mapWork").classList.toggle("hidden", !data.target);
  if (!data.target) return;
  $("#mapTarget").value = data.target.base_url || data.target.domain || "";
  $("#mapAuth").checked = !!data.target.authorized;
  $("#mapMeta").textContent = data.target.domain + (data.target.authorized ? tr("map.ownTarget") : tr("map.passiveOnly"));
  const st = data.map?.stats || {};
  $("#mapStats").innerHTML = [
    [tr("map.statHosts"), st.live_hosts ?? 0],
    [tr("map.statNames"), st.subdomains ?? 0],
    [tr("map.statPaths"), st.paths ?? 0],
    [tr("map.statParams"), st.params ?? 0],
    ["JS", st.js ?? 0],
    ["URL", st.urls ?? 0],
  ].map(([k, v]) => `<div class="map-stat"><b>${v}</b><span>${k}</span></div>`).join("");
  const byId = Object.fromEntries((data.stages || []).map((s) => [s.id, s]));
  $("#mapStages").innerHTML = (data.catalog || []).map((c, i) => {
    const s = byId[c.id] || { status: "idle" };
    const busy = s.status === "running";
    return `<article class="stage-card ${s.status}" data-stage="${c.id}">
      <div class="stage-head">
        <h3>${i + 1}. ${esc(tr("stage." + c.id + ".title"))}</h3>
        <span class="stage-mode ${c.mode}">${c.mode === "active" ? tr("map.modeActive") : tr("map.modePassive")} · ${esc(s.status || "idle")}</span>
      </div>
      <p>${esc(tr("stage." + c.id + ".hint"))}</p>
      ${s.summary ? `<p>${esc(s.summary)}</p>` : ""}
      <div class="toolbar">
        <button class="primary mini" data-run="${c.id}" ${busy ? "disabled" : ""}>${tr("map.run")}</button>
        ${c.id === "dirs" ? `<button class="mini" data-jump="discover">${tr("map.openPaths")}</button>` : ""}
        ${c.id === "params" ? `<button class="mini" data-jump="fuzz">${tr("map.openFuzz")}</button>` : ""}
      </div>
    </article>`;
  }).join("");
  renderAppMap(data.map);
}

function renderAppMap(m) {
  const tree = $("#mapTree");
  const hosts = m?.hosts || [];
  if (!hosts.length) {
    tree.innerHTML = `<div class="muted" style="padding:12px">${tr("map.emptyTree")}</div>`;
    return;
  }
  tree.innerHTML = hosts.map((h) => {
    const paths = (h.paths || []).map((p) => {
      const bits = [p.status, (p.params || []).length ? `params ${(p.params || []).join(", ")}` : ""].filter(Boolean).join(" · ");
      return `<div class="tree-row tree-leaf" data-host="${esc(h.host)}" data-path="${esc(p.path)}">
        <span class="tree-icon">/</span><span>${esc(p.path)}</span>
        <span class="tree-status">${esc(bits)}</span>
      </div>`;
    }).join("");
    return `<div class="tree-node">
      <div class="tree-row" data-host="${esc(h.host)}">
        <span class="tree-icon">${h.live ? "●" : "○"}</span>
        <strong>${esc(h.host)}</strong>
        <span class="tree-status">${esc((h.tech || []).join(", "))}</span>
      </div>
      <div class="tree-children">${paths || `<div class="muted" style="padding:4px 12px">${tr("map.noPaths")}</div>`}</div>
    </div>`;
  }).join("");
}

$("#mapForm").addEventListener("submit", async (e) => {
  e.preventDefault();
  try {
    const data = await api("/api/intel/targets", {
      method: "POST",
      body: JSON.stringify({ target: $("#mapTarget").value.trim(), authorized: $("#mapAuth").checked }),
    });
    renderMapSnapshot(data);
  } catch (err) { alert(err.message); }
});
$("#mapStages").addEventListener("click", async (e) => {
  const jump = e.target.closest("[data-jump]");
  if (jump) { showView(jump.dataset.jump); return; }
  const btn = e.target.closest("[data-run]");
  if (!btn || !mapState.target) return;
  btn.disabled = true;
  try {
    const data = await api(`/api/intel/targets/${mapState.target.id}/stages/${btn.dataset.run}`, { method: "POST" });
    renderMapSnapshot(data);
    if (data.error) $("#mapMeta").textContent = data.error;
  } catch (err) { alert(err.message); }
});
$("#btnMapHistory").addEventListener("click", async () => {
  if (!mapState.target) return;
  try { renderMapSnapshot(await api(`/api/intel/targets/${mapState.target.id}/ingest-history`, { method: "POST" })); }
  catch (err) { alert(err.message); }
});
$("#btnMapRefresh").addEventListener("click", loadMap);
$("#mapTree").addEventListener("click", (e) => {
  const row = e.target.closest("[data-host]");
  if (!row) return;
  const host = row.dataset.host;
  const path = row.dataset.path;
  const h = (mapState.map?.hosts || []).find((x) => x.host === host);
  if (!h) return;
  if (path) {
    const p = (h.paths || []).find((x) => x.path === path);
    $("#mapDetail").innerHTML = `<strong>${esc(host)}${esc(path)}</strong><div>status ${esc(p?.status || "—")}</div>
      <div>params: ${esc((p?.params || []).join(", ") || "—")}</div>
      <div>js: ${esc((p?.js || []).slice(0, 8).join(", ") || "—")}</div>`;
    return;
  }
  $("#mapDetail").innerHTML = `<strong>${esc(host)}</strong> ${h.live ? "live" : ""}
    <div>tech: ${esc((h.tech || []).join(", ") || "—")}</div>
    <div>ports: ${esc((h.ports || []).join(", ") || "—")}</div>`;
});

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
$("#btnFuzzRun")?.addEventListener("click", async () => {
  const hide = ($("#fuzzHide").value || "").split(",").map((s) => Number(s.trim())).filter(Boolean);
  const custom = $("#fuzzWords").value.split("\n").map((s) => s.trim()).filter(Boolean);
  const path = $("#fuzzWlSelect")?.value || "";
  $("#fuzzMeta").textContent = "running…";
  try {
    const data = await api("/api/fuzz/run", {
      method: "POST",
      body: JSON.stringify({
        url: $("#fuzzUrl").value.trim(),
        method: $("#fuzzMethod").value,
        body: $("#fuzzBody").value,
        wordlist: custom,
        wordlist_path: custom.length ? "" : path,
        workers: Number($("#fuzzWorkers").value),
        rps: Number($("#fuzzRps").value),
        hide,
      }),
    });
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
  try { if (typeof mapState !== "undefined" && mapState.target) renderMapSnapshot(mapState); } catch (_) {}
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
}

/* Init */
try { applyTheme(localStorage.getItem("meb_theme") || "dark"); } catch (_) { applyTheme("dark"); }
try { applyLang(detectLang()); } catch (_) { applyLang("en"); }
$$("#themeSwitch button").forEach((b) => b.addEventListener("click", () => applyTheme(b.dataset.theme)));
$$("#langSwitch button").forEach((b) => b.addEventListener("click", () => applyLang(b.dataset.lang)));
showView("map");
loadStatus();
loadHistory();
loadIntercept();
loadLogs();
connectWs();
