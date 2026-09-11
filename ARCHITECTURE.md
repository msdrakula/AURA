# AURA architecture

Independent local intercepting proxy plus a guided application map (Advanced URL and Request Analyzer). Not affiliated with PortSwigger — see `LEGAL.md`. Source-available under PolyForm Noncommercial 1.0.0 (`LICENSE`).

## Layout

- `cmd/server/main.go` — process wiring (proxy, API, callback, intel).
- `internal/proxy/` — HTTP/1.1 MITM and intercept.
- `internal/intel/` — mindmap stages, artifacts, application map.
- `internal/fuzz/` — FUZZ-keyword HTTP fuzzer.
- `internal/discover/` — directory discovery.
- `internal/wordlist/` — small built-in dictionaries.
- `internal/extractor/` — HTML/JS URL harvest (used by intel scrape).
- `internal/repeater/`, `internal/batch/` — single and bulk request senders.
- `internal/analyzer/` — passive findings on captured traffic.
- `internal/callback/` — localhost HTTP interaction log.
- `internal/storage/` — SQLite: history, findings, organizer, recon map.
- `internal/store/` — in-memory proxy runtime.
- `web/` — single-page UI (Карта is the home screen).

## Intel flow

User confirms a target → stages write artifacts (`subdomain`, `host`, `path`, `param`, `js`, `tech`, `url`) → `BuildMap` groups them into host → path tree. Proxy history can be ingested into the same map.
