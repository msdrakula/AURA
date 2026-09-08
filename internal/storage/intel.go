package storage

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"

	"meb/internal/intel"
)

func (s *Store) CreateTarget(t intel.Target) error {
	if err := s.requireOpen(); err != nil {
		return err
	}
	if t.ID == "" {
		t.ID = uuid.NewString()
	}
	if t.Created.IsZero() {
		t.Created = time.Now().UTC()
	}
	auth := 0
	if t.Authorized {
		auth = 1
	}
	now := t.Created.UTC().Format(time.RFC3339Nano)
	_, err := s.db.Exec(sqlUpsertTarget, t.ID, t.Domain, now, now, t.BaseURL, auth)
	if err != nil {
		return fmt.Errorf("intel: save target: %w", err)
	}
	return nil
}

func (s *Store) GetTarget(id string) (intel.Target, error) {
	if err := s.requireOpen(); err != nil {
		return intel.Target{}, err
	}
	row := s.db.QueryRow(sqlGetTargetByID, id)
	t, err := scanTarget(row)
	if err == sql.ErrNoRows {
		return intel.Target{}, fmt.Errorf("intel: target not found")
	}
	return t, err
}

func (s *Store) ListTargets() ([]intel.Target, error) {
	if err := s.requireOpen(); err != nil {
		return nil, err
	}
	rows, err := s.db.Query(sqlListTargets)
	if err != nil {
		return nil, fmt.Errorf("intel: list targets: %w", err)
	}
	defer rows.Close()
	var out []intel.Target
	for rows.Next() {
		t, err := scanTarget(rows)
		if err != nil {
			return nil, err
		}
		out = append(out, t)
	}
	if out == nil {
		out = []intel.Target{}
	}
	return out, rows.Err()
}

func (s *Store) UpdateTarget(t intel.Target) error {
	if err := s.requireOpen(); err != nil {
		return err
	}
	auth := 0
	if t.Authorized {
		auth = 1
	}
	now := time.Now().UTC().Format(time.RFC3339Nano)
	_, err := s.db.Exec(sqlUpdateTargetAuth, auth, t.BaseURL, now, t.ID)
	if err != nil {
		return fmt.Errorf("intel: update target: %w", err)
	}
	return nil
}

func (s *Store) UpsertStage(targetID string, st intel.StageState) error {
	if err := s.requireOpen(); err != nil {
		return err
	}
	if st.Updated.IsZero() {
		st.Updated = time.Now().UTC()
	}
	_, err := s.db.Exec(sqlUpsertStage, targetID, st.ID, st.Status, st.Summary, st.Updated.UTC().Format(time.RFC3339Nano))
	if err != nil {
		return fmt.Errorf("intel: save stage: %w", err)
	}
	return nil
}

func (s *Store) ListStages(targetID string) ([]intel.StageState, error) {
	if err := s.requireOpen(); err != nil {
		return nil, err
	}
	rows, err := s.db.Query(sqlListStages, targetID)
	if err != nil {
		return nil, fmt.Errorf("intel: list stages: %w", err)
	}
	defer rows.Close()
	var out []intel.StageState
	for rows.Next() {
		var tid, sid, status, summary, updated string
		if err := rows.Scan(&tid, &sid, &status, &summary, &updated); err != nil {
			return nil, err
		}
		ts, _ := time.Parse(time.RFC3339Nano, updated)
		out = append(out, intel.StageState{ID: sid, Status: status, Summary: summary, Updated: ts})
	}
	if out == nil {
		out = []intel.StageState{}
	}
	return out, rows.Err()
}

func (s *Store) AddArtifact(a intel.Artifact) error {
	if err := s.requireOpen(); err != nil {
		return err
	}
	if a.ID == "" {
		a.ID = uuid.NewString()
	}
	if a.Created.IsZero() {
		a.Created = time.Now().UTC()
	}
	meta := "{}"
	if a.Extra != nil {
		b, err := json.Marshal(a.Extra)
		if err != nil {
			return err
		}
		meta = string(b)
	}
	_, err := s.db.Exec(sqlInsertReconItem, a.ID, a.TargetID, a.Source, a.Kind, a.Value, meta, a.Created.UTC().Format(time.RFC3339Nano))
	if err != nil {
		return fmt.Errorf("intel: save artifact: %w", err)
	}
	_, _ = s.db.Exec(sqlTouchTarget, time.Now().UTC().Format(time.RFC3339Nano), a.TargetID)
	return nil
}

func (s *Store) ListArtifacts(targetID string) ([]intel.Artifact, error) {
	if err := s.requireOpen(); err != nil {
		return nil, err
	}
	rows, err := s.db.Query(sqlListReconItems, targetID)
	if err != nil {
		return nil, fmt.Errorf("intel: list artifacts: %w", err)
	}
	defer rows.Close()
	var out []intel.Artifact
	for rows.Next() {
		var a intel.Artifact
		var meta, created, stage string
		if err := rows.Scan(&a.ID, &a.TargetID, &stage, &a.Kind, &a.Value, &meta, &created); err != nil {
			return nil, err
		}
		a.Source = stage
		a.Extra = map[string]string{}
		if meta != "" && meta != "{}" {
			_ = json.Unmarshal([]byte(meta), &a.Extra)
		}
		a.Created, _ = time.Parse(time.RFC3339Nano, created)
		out = append(out, a)
	}
	if out == nil {
		out = []intel.Artifact{}
	}
	return out, rows.Err()
}

func scanTarget(row rowScanner) (intel.Target, error) {
	var t intel.Target
	var created, updated, notes string
	var auth int
	if err := row.Scan(&t.ID, &t.Domain, &created, &updated, &notes, &t.BaseURL, &auth); err != nil {
		return intel.Target{}, err
	}
	t.Name = t.Domain
	t.Authorized = auth != 0
	t.Created, _ = time.Parse(time.RFC3339Nano, created)
	return t, nil
}
