// Package storage persists debug sessions in SQLite (pure Go, no CGO).
package storage

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"meb/internal/models"

	_ "modernc.org/sqlite"
)

// Store is a SQLite-backed session store.
type Store struct {
	db            *sql.DB
	insertTx      *sql.Stmt
	listTx        *sql.Stmt
	getTx         *sql.Stmt
	insertFinding *sql.Stmt
}

// Init opens filepath, runs migrations, and prepares statements.
func (s *Store) Init(filePath string) error {
	if s == nil {
		return fmt.Errorf("storage: nil store")
	}
	if filePath == "" {
		return fmt.Errorf("storage: empty database path")
	}
	if err := s.Close(); err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(filePath), 0o755); err != nil && filepath.Dir(filePath) != "." {
		return fmt.Errorf("failed to create database directory: %w", err)
	}
	dsn := "file:" + filePath + "?_pragma=busy_timeout(5000)&_pragma=foreign_keys(1)&_pragma=journal_mode(WAL)"
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return fmt.Errorf("failed to open sqlite %s: %w", filePath, err)
	}
	db.SetMaxOpenConns(1)
	if _, err := db.Exec(schema); err != nil {
		_ = db.Close()
		return fmt.Errorf("failed to migrate schema: %w", err)
	}
	for _, extra := range []string{
		`ALTER TABLE recon_targets ADD COLUMN base_url TEXT NOT NULL DEFAULT ''`,
		`ALTER TABLE recon_targets ADD COLUMN authorized INTEGER NOT NULL DEFAULT 0`,
	} {
		_, _ = db.Exec(extra)
	}
	insertTx, err := db.Prepare(sqlInsertTx)
	if err != nil {
		_ = db.Close()
		return fmt.Errorf("failed to prepare insert transaction: %w", err)
	}
	listTx, err := db.Prepare(sqlListTx)
	if err != nil {
		_ = insertTx.Close()
		_ = db.Close()
		return fmt.Errorf("failed to prepare list transactions: %w", err)
	}
	insertFinding, err := db.Prepare(sqlInsertFinding)
	if err != nil {
		_ = insertTx.Close()
		_ = listTx.Close()
		_ = db.Close()
		return fmt.Errorf("failed to prepare insert finding: %w", err)
	}
	getTx, err := db.Prepare(sqlGetTx)
	if err != nil {
		_ = insertTx.Close()
		_ = listTx.Close()
		_ = insertFinding.Close()
		_ = db.Close()
		return fmt.Errorf("failed to prepare get transaction: %w", err)
	}
	s.db = db
	s.insertTx = insertTx
	s.listTx = listTx
	s.getTx = getTx
	s.insertFinding = insertFinding
	return nil
}

// Close releases the database and prepared statements.
func (s *Store) Close() error {
	if s == nil {
		return nil
	}
	var first error
	for _, stmt := range []*sql.Stmt{s.insertTx, s.listTx, s.getTx, s.insertFinding} {
		if stmt == nil {
			continue
		}
		if err := stmt.Close(); err != nil && first == nil {
			first = fmt.Errorf("failed to close statement: %w", err)
		}
	}
	s.insertTx, s.listTx, s.getTx, s.insertFinding = nil, nil, nil, nil
	if s.db != nil {
		if err := s.db.Close(); err != nil && first == nil {
			first = fmt.Errorf("failed to close sqlite: %w", err)
		}
		s.db = nil
	}
	return first
}

// SaveTransaction upserts a captured request/response pair.
func (s *Store) SaveTransaction(tx *models.HTTPTransaction) error {
	if err := s.requireOpen(); err != nil {
		return err
	}
	if tx == nil {
		return fmt.Errorf("storage: nil transaction")
	}
	if tx.ID == "" {
		return fmt.Errorf("storage: transaction id required")
	}
	cp := tx.Clone()
	reqHeaders, err := marshalJSON(cp.Request.Headers)
	if err != nil {
		return fmt.Errorf("failed to encode request headers: %w", err)
	}
	tags, err := marshalJSON(cp.Tags)
	if err != nil {
		return fmt.Errorf("failed to encode tags: %w", err)
	}
	var (
		respStatus any
		respHdrs   any
		respBody   any
		respDur    any
	)
	if cp.Response != nil {
		respStatus = cp.Response.Status
		respHdrs, err = marshalJSON(cp.Response.Headers)
		if err != nil {
			return fmt.Errorf("failed to encode response headers: %w", err)
		}
		respBody = cp.Response.Body
		respDur = int64(cp.Response.Duration)
	}
	ts := cp.Timestamp.UTC()
	if ts.IsZero() {
		ts = time.Now().UTC()
	}
	_, err = s.insertTx.Exec(
		cp.ID,
		cp.Request.Method,
		cp.Request.URL,
		reqHeaders,
		cp.Request.Body,
		respStatus,
		respHdrs,
		respBody,
		respDur,
		ts.Format(time.RFC3339Nano),
		cp.Comment,
		tags,
	)
	if err != nil {
		return fmt.Errorf("failed to save transaction %s: %w", cp.ID, err)
	}
	return nil
}

// GetTransactions returns the newest sessions first.
func (s *Store) GetTransactions(limit, offset int) ([]models.HTTPTransaction, error) {
	if err := s.requireOpen(); err != nil {
		return nil, err
	}
	if limit <= 0 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}
	rows, err := s.listTx.Query(limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list transactions: %w", err)
	}
	defer rows.Close()

	out := make([]models.HTTPTransaction, 0)
	for rows.Next() {
		tx, err := scanTransaction(rows)
		if err != nil {
			return nil, err
		}
		out = append(out, tx)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate transactions: %w", err)
	}
	return out, nil
}

// SaveFinding upserts one analysis result.
func (s *Store) SaveFinding(f *models.AnalysisFinding) error {
	if err := s.requireOpen(); err != nil {
		return err
	}
	if f == nil {
		return fmt.Errorf("storage: nil finding")
	}
	if f.ID == "" {
		return fmt.Errorf("storage: finding id required")
	}
	if f.TransactionID == "" {
		return fmt.Errorf("storage: finding transaction_id required")
	}
	_, err := s.insertFinding.Exec(
		f.ID,
		f.TransactionID,
		f.RuleName,
		string(f.Severity),
		f.Title,
		f.Description,
		f.Evidence,
	)
	if err != nil {
		return fmt.Errorf("failed to save finding %s: %w", f.ID, err)
	}
	return nil
}

// GetTransaction loads one session by id. Returns nil, nil if missing.
func (s *Store) GetTransaction(id string) (*models.HTTPTransaction, error) {
	if err := s.requireOpen(); err != nil {
		return nil, err
	}
	if id == "" {
		return nil, fmt.Errorf("storage: transaction id required")
	}
	tx, err := scanTransaction(s.getTx.QueryRow(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &tx, nil
}

// ClearTransactions deletes findings and history.
func (s *Store) ClearTransactions() error {
	if err := s.requireOpen(); err != nil {
		return err
	}
	if _, err := s.db.Exec(sqlClearFindings); err != nil {
		return fmt.Errorf("failed to clear findings: %w", err)
	}
	if _, err := s.db.Exec(sqlClearTx); err != nil {
		return fmt.Errorf("failed to clear transactions: %w", err)
	}
	return nil
}

// GetFindings returns all stored analysis findings, newest first by rowid.
func (s *Store) GetFindings() ([]models.AnalysisFinding, error) {
	if err := s.requireOpen(); err != nil {
		return nil, err
	}
	rows, err := s.db.Query(sqlListFindings)
	if err != nil {
		return nil, fmt.Errorf("failed to list findings: %w", err)
	}
	defer rows.Close()
	out := make([]models.AnalysisFinding, 0)
	for rows.Next() {
		var f models.AnalysisFinding
		if err := rows.Scan(&f.ID, &f.TransactionID, &f.RuleName, &f.Severity, &f.Title, &f.Description, &f.Evidence); err != nil {
			return nil, fmt.Errorf("failed to scan finding: %w", err)
		}
		out = append(out, f)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate findings: %w", err)
	}
	return out, nil
}

// GetFindingsByTransaction returns findings tied to one transaction.
func (s *Store) GetFindingsByTransaction(txID string) ([]models.AnalysisFinding, error) {
	if err := s.requireOpen(); err != nil {
		return nil, err
	}
	if txID == "" {
		return nil, fmt.Errorf("storage: transaction id required")
	}
	rows, err := s.db.Query(sqlFindingsByTx, txID)
	if err != nil {
		return nil, fmt.Errorf("failed to query findings: %w", err)
	}
	defer rows.Close()
	out := make([]models.AnalysisFinding, 0)
	for rows.Next() {
		var f models.AnalysisFinding
		if err := rows.Scan(&f.ID, &f.TransactionID, &f.RuleName, &f.Severity, &f.Title, &f.Description, &f.Evidence); err != nil {
			return nil, fmt.Errorf("failed to scan finding: %w", err)
		}
		out = append(out, f)
	}
	return out, nil
}

func (s *Store) requireOpen() error {
	if s == nil || s.db == nil {
		return fmt.Errorf("storage: not initialized")
	}
	return nil
}

type rowScanner interface {
	Scan(dest ...any) error
}

func scanTransaction(row rowScanner) (models.HTTPTransaction, error) {
	var (
		id, method, rawURL, reqHeaders, ts, comment, tags string
		reqBody                                           []byte
		respStatus                                        sql.NullInt64
		respHeaders                                       sql.NullString
		respBody                                          []byte
		respDur                                           sql.NullInt64
	)
	if err := row.Scan(
		&id, &method, &rawURL, &reqHeaders, &reqBody,
		&respStatus, &respHeaders, &respBody, &respDur,
		&ts, &comment, &tags,
	); err != nil {
		return models.HTTPTransaction{}, fmt.Errorf("failed to scan transaction: %w", err)
	}
	parsedTS, err := time.Parse(time.RFC3339Nano, ts)
	if err != nil {
		parsedTS, err = time.Parse(time.RFC3339, ts)
		if err != nil {
			return models.HTTPTransaction{}, fmt.Errorf("failed to parse timestamp %q: %w", ts, err)
		}
	}
	tx := models.HTTPTransaction{
		ID: id,
		Request: models.HTTPRequest{
			Method:  method,
			URL:     rawURL,
			Headers: map[string]string{},
			Body:    reqBody,
		},
		Timestamp: parsedTS.UTC(),
		Comment:   comment,
		Tags:      []string{},
	}
	if err := unmarshalJSON(reqHeaders, &tx.Request.Headers); err != nil {
		return models.HTTPTransaction{}, fmt.Errorf("failed to decode request headers: %w", err)
	}
	if err := unmarshalJSON(tags, &tx.Tags); err != nil {
		return models.HTTPTransaction{}, fmt.Errorf("failed to decode tags: %w", err)
	}
	if respStatus.Valid {
		resp := &models.HTTPResponse{
			Status:  int(respStatus.Int64),
			Headers: map[string]string{},
			Body:    respBody,
		}
		if respDur.Valid {
			resp.Duration = time.Duration(respDur.Int64)
		}
		if respHeaders.Valid && respHeaders.String != "" {
			if err := unmarshalJSON(respHeaders.String, &resp.Headers); err != nil {
				return models.HTTPTransaction{}, fmt.Errorf("failed to decode response headers: %w", err)
			}
		}
		tx.Response = resp
	}
	return tx, nil
}

func marshalJSON(v any) (string, error) {
	if v == nil {
		switch v.(type) {
		case map[string]string:
			return "{}", nil
		case []string:
			return "[]", nil
		}
	}
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	if string(b) == "null" {
		switch v.(type) {
		case map[string]string, *map[string]string:
			return "{}", nil
		default:
			return "[]", nil
		}
	}
	return string(b), nil
}

func unmarshalJSON(raw string, dest any) error {
	if raw == "" {
		return nil
	}
	return json.Unmarshal([]byte(raw), dest)
}
