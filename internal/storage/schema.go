package storage

const schema = `
CREATE TABLE IF NOT EXISTS transactions (
	id TEXT PRIMARY KEY,
	req_method TEXT NOT NULL,
	req_url TEXT NOT NULL,
	req_headers TEXT NOT NULL DEFAULT '{}',
	req_body BLOB,
	resp_status INTEGER,
	resp_headers TEXT,
	resp_body BLOB,
	resp_duration_ns INTEGER,
	timestamp TEXT NOT NULL,
	comment TEXT NOT NULL DEFAULT '',
	tags TEXT NOT NULL DEFAULT '[]'
);

CREATE INDEX IF NOT EXISTS idx_transactions_timestamp ON transactions(timestamp);

CREATE TABLE IF NOT EXISTS findings (
	id TEXT PRIMARY KEY,
	transaction_id TEXT NOT NULL,
	rule_name TEXT NOT NULL,
	severity TEXT NOT NULL,
	title TEXT NOT NULL,
	description TEXT NOT NULL,
	evidence TEXT NOT NULL DEFAULT '',
	FOREIGN KEY(transaction_id) REFERENCES transactions(id)
);

CREATE INDEX IF NOT EXISTS idx_findings_transaction_id ON findings(transaction_id);

CREATE TABLE IF NOT EXISTS organizer_items (
	id TEXT PRIMARY KEY,
	created TEXT NOT NULL,
	tool TEXT NOT NULL DEFAULT '',
	collection TEXT NOT NULL DEFAULT 'inbox',
	method TEXT NOT NULL DEFAULT '',
	host TEXT NOT NULL DEFAULT '',
	path TEXT NOT NULL DEFAULT '',
	query TEXT NOT NULL DEFAULT '',
	param_count INTEGER NOT NULL DEFAULT 0,
	status_code INTEGER NOT NULL DEFAULT 0,
	length INTEGER NOT NULL DEFAULT 0,
	notes TEXT NOT NULL DEFAULT '',
	status TEXT NOT NULL DEFAULT '',
	highlight TEXT NOT NULL DEFAULT '',
	request_raw TEXT NOT NULL DEFAULT '',
	response_raw TEXT NOT NULL DEFAULT ''
);

CREATE INDEX IF NOT EXISTS idx_organizer_collection ON organizer_items(collection);
CREATE INDEX IF NOT EXISTS idx_organizer_created ON organizer_items(created);

CREATE TABLE IF NOT EXISTS recon_targets (
	id TEXT PRIMARY KEY,
	host TEXT NOT NULL UNIQUE,
	created TEXT NOT NULL,
	updated TEXT NOT NULL,
	notes TEXT NOT NULL DEFAULT '',
	base_url TEXT NOT NULL DEFAULT '',
	authorized INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS recon_items (
	id TEXT PRIMARY KEY,
	target_id TEXT NOT NULL,
	stage TEXT NOT NULL,
	kind TEXT NOT NULL,
	value TEXT NOT NULL,
	meta TEXT NOT NULL DEFAULT '{}',
	created TEXT NOT NULL,
	UNIQUE(target_id, kind, value)
);

CREATE INDEX IF NOT EXISTS idx_recon_items_target ON recon_items(target_id);
CREATE INDEX IF NOT EXISTS idx_recon_items_kind ON recon_items(target_id, kind);

CREATE TABLE IF NOT EXISTS recon_stages (
	target_id TEXT NOT NULL,
	stage_id TEXT NOT NULL,
	status TEXT NOT NULL DEFAULT 'idle',
	summary TEXT NOT NULL DEFAULT '',
	updated TEXT NOT NULL,
	PRIMARY KEY (target_id, stage_id)
);
`

const (
	sqlUpsertTarget = `
INSERT INTO recon_targets (id, host, created, updated, notes, base_url, authorized)
VALUES (?, ?, ?, ?, '', ?, ?)
ON CONFLICT(host) DO UPDATE SET
	updated = excluded.updated,
	base_url = CASE WHEN excluded.base_url = '' THEN recon_targets.base_url ELSE excluded.base_url END,
	authorized = CASE WHEN excluded.authorized = 0 THEN recon_targets.authorized ELSE excluded.authorized END`

	sqlGetTargetByHost = `SELECT id, host, created, updated, notes, base_url, authorized FROM recon_targets WHERE host = ?`

	sqlGetTargetByID = `SELECT id, host, created, updated, notes, base_url, authorized FROM recon_targets WHERE id = ?`

	sqlListTargets = `SELECT id, host, created, updated, notes, base_url, authorized FROM recon_targets ORDER BY updated DESC`

	sqlUpdateTargetNotes = `UPDATE recon_targets SET notes = ?, updated = ? WHERE id = ?`

	sqlUpdateTargetAuth = `UPDATE recon_targets SET authorized = ?, base_url = ?, updated = ? WHERE id = ?`

	sqlTouchTarget = `UPDATE recon_targets SET updated = ? WHERE id = ?`

	sqlInsertReconItem = `
INSERT INTO recon_items (id, target_id, stage, kind, value, meta, created)
VALUES (?, ?, ?, ?, ?, ?, ?)
ON CONFLICT(target_id, kind, value) DO UPDATE SET meta = excluded.meta, stage = excluded.stage`

	sqlListReconItems = `SELECT id, target_id, stage, kind, value, meta, created FROM recon_items WHERE target_id = ? ORDER BY kind, created`

	sqlListReconItemsByKind = `SELECT id, target_id, stage, kind, value, meta, created FROM recon_items WHERE target_id = ? AND kind = ? ORDER BY created`

	sqlDeleteReconTarget = `DELETE FROM recon_items WHERE target_id = ?; DELETE FROM recon_stages WHERE target_id = ?; DELETE FROM recon_targets WHERE id = ?`

	sqlUpsertStage = `
INSERT INTO recon_stages (target_id, stage_id, status, summary, updated)
VALUES (?, ?, ?, ?, ?)
ON CONFLICT(target_id, stage_id) DO UPDATE SET status = excluded.status, summary = excluded.summary, updated = excluded.updated`

	sqlListStages = `SELECT target_id, stage_id, status, summary, updated FROM recon_stages WHERE target_id = ?`
)

const (
	sqlInsertTx = `
INSERT OR REPLACE INTO transactions (
	id, req_method, req_url, req_headers, req_body,
	resp_status, resp_headers, resp_body, resp_duration_ns,
	timestamp, comment, tags
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	sqlListTx = `
SELECT
	id, req_method, req_url, req_headers, req_body,
	resp_status, resp_headers, resp_body, resp_duration_ns,
	timestamp, comment, tags
FROM transactions
ORDER BY timestamp DESC
LIMIT ? OFFSET ?`

	sqlInsertFinding = `
INSERT OR REPLACE INTO findings (
	id, transaction_id, rule_name, severity, title, description, evidence
) VALUES (?, ?, ?, ?, ?, ?, ?)`

	sqlGetTx = `
SELECT
	id, req_method, req_url, req_headers, req_body,
	resp_status, resp_headers, resp_body, resp_duration_ns,
	timestamp, comment, tags
FROM transactions
WHERE id = ?`

	sqlClearFindings = `DELETE FROM findings`
	sqlClearTx       = `DELETE FROM transactions`

	sqlListFindings = `
SELECT id, transaction_id, rule_name, severity, title, description, evidence
FROM findings
ORDER BY rowid DESC`

	sqlFindingsByTx = `
SELECT id, transaction_id, rule_name, severity, title, description, evidence
FROM findings
WHERE transaction_id = ?
ORDER BY rowid DESC`

	sqlInsertOrganizer = `
INSERT OR REPLACE INTO organizer_items (
	id, created, tool, collection, method, host, path, query,
	param_count, status_code, length, notes, status, highlight,
	request_raw, response_raw
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	sqlListOrganizer = `
SELECT id, created, tool, collection, method, host, path, query,
	param_count, status_code, length, notes, status, highlight,
	request_raw, response_raw
FROM organizer_items
WHERE collection = ?
ORDER BY created DESC`

	sqlListOrganizerCollections = `
SELECT DISTINCT collection FROM organizer_items ORDER BY collection`

	sqlDeleteOrganizer = `DELETE FROM organizer_items WHERE id = ?`

	sqlUpdateOrganizer = `
UPDATE organizer_items SET notes = ?, status = ?, highlight = ?, collection = ?
WHERE id = ?`
)
