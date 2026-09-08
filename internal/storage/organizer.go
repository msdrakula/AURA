package storage

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

// OrganizerItem is a stored HTTP message in an Organizer collection.
type OrganizerItem struct {
	ID          string `json:"id"`
	Created     string `json:"created"`
	Tool        string `json:"tool"`
	Collection  string `json:"collection"`
	Method      string `json:"method"`
	Host        string `json:"host"`
	Path        string `json:"path"`
	Query       string `json:"query"`
	ParamCount  int    `json:"param_count"`
	StatusCode  int    `json:"status_code"`
	Length      int    `json:"length"`
	Notes       string `json:"notes"`
	Status      string `json:"status"`
	Highlight   string `json:"highlight"`
	RequestRaw  string `json:"request_raw"`
	ResponseRaw string `json:"response_raw"`
}

// AddOrganizerInput is the payload for adding an item.
type AddOrganizerInput struct {
	ID          string
	Tool        string
	Collection  string
	Method      string
	Host        string
	Path        string
	Query       string
	ParamCount  int
	StatusCode  int
	Length      int
	RequestRaw  string
	ResponseRaw string
}

// AddOrganizer stores a raw HTTP message copy into an Organizer collection.
func (s *Store) AddOrganizer(in AddOrganizerInput) (OrganizerItem, error) {
	if err := s.requireOpen(); err != nil {
		return OrganizerItem{}, err
	}
	if in.ID == "" {
		return OrganizerItem{}, fmt.Errorf("storage: organizer id required")
	}
	if in.Collection == "" {
		in.Collection = "inbox"
	}
	item := OrganizerItem{
		ID:          in.ID,
		Created:     time.Now().UTC().Format(time.RFC3339Nano),
		Tool:        in.Tool,
		Collection:  in.Collection,
		Method:      in.Method,
		Host:        in.Host,
		Path:        in.Path,
		Query:       in.Query,
		ParamCount:  in.ParamCount,
		StatusCode:  in.StatusCode,
		Length:      in.Length,
		RequestRaw:  in.RequestRaw,
		ResponseRaw: in.ResponseRaw,
	}
	_, err := s.db.Exec(sqlInsertOrganizer,
		item.ID, item.Created, item.Tool, item.Collection,
		item.Method, item.Host, item.Path, item.Query,
		item.ParamCount, item.StatusCode, item.Length,
		item.Notes, item.Status, item.Highlight,
		item.RequestRaw, item.ResponseRaw,
	)
	if err != nil {
		return OrganizerItem{}, fmt.Errorf("failed to insert organizer item: %w", err)
	}
	return item, nil
}

// ListOrganizer returns items in the given collection (empty = inbox).
func (s *Store) ListOrganizer(collection string) ([]OrganizerItem, error) {
	if err := s.requireOpen(); err != nil {
		return nil, err
	}
	if collection == "" {
		collection = "inbox"
	}
	rows, err := s.db.Query(sqlListOrganizer, collection)
	if err != nil {
		return nil, fmt.Errorf("failed to list organizer: %w", err)
	}
	defer rows.Close()
	return scanOrganizerRows(rows)
}

// ListOrganizerCollections returns distinct collection names plus the default "inbox".
func (s *Store) ListOrganizerCollections() ([]string, error) {
	if err := s.requireOpen(); err != nil {
		return nil, err
	}
	rows, err := s.db.Query(sqlListOrganizerCollections)
	if err != nil {
		return nil, fmt.Errorf("failed to list collections: %w", err)
	}
	defer rows.Close()
	out := []string{"inbox"}
	seen := map[string]bool{"inbox": true}
	for rows.Next() {
		var c string
		if err := rows.Scan(&c); err != nil {
			return nil, err
		}
		if !seen[c] {
			seen[c] = true
			out = append(out, c)
		}
	}
	return out, nil
}

// UpdateOrganizer patches notes/status/highlight/collection for an item.
func (s *Store) UpdateOrganizer(id, notes, status, highlight, collection string) error {
	if err := s.requireOpen(); err != nil {
		return err
	}
	if collection == "" {
		collection = "inbox"
	}
	res, err := s.db.Exec(sqlUpdateOrganizer, notes, status, highlight, collection, id)
	if err != nil {
		return fmt.Errorf("failed to update organizer item: %w", err)
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		return fmt.Errorf("organizer item not found: %s", id)
	}
	return nil
}

// DeleteOrganizer removes an item.
func (s *Store) DeleteOrganizer(id string) error {
	if err := s.requireOpen(); err != nil {
		return err
	}
	_, err := s.db.Exec(sqlDeleteOrganizer, id)
	return err
}

func scanOrganizerRows(rows *sql.Rows) ([]OrganizerItem, error) {
	out := make([]OrganizerItem, 0)
	for rows.Next() {
		var it OrganizerItem
		if err := rows.Scan(
			&it.ID, &it.Created, &it.Tool, &it.Collection,
			&it.Method, &it.Host, &it.Path, &it.Query,
			&it.ParamCount, &it.StatusCode, &it.Length,
			&it.Notes, &it.Status, &it.Highlight,
			&it.RequestRaw, &it.ResponseRaw,
		); err != nil {
			return nil, err
		}
		out = append(out, it)
	}
	return out, rows.Err()
}

// requireOpen is defined in store.go.

// splitPathQuery splits a URL-ish path into path and query.
func splitPathQuery(u string) (string, string) {
	if i := strings.IndexByte(u, '?'); i >= 0 {
		return u[:i], u[i+1:]
	}
	return u, ""
}
