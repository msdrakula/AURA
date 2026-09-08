package storage

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"meb/internal/intel"
	"meb/internal/models"
)

func TestSaveAndGetTransaction(t *testing.T) {
	path := filepath.Join(t.TempDir(), "session.db")
	st := &Store{}
	require.NoError(t, st.Init(path))
	t.Cleanup(func() { require.NoError(t, st.Close()) })

	tx := &models.HTTPTransaction{
		ID: "11111111-2222-3333-4444-555555555555",
		Request: models.HTTPRequest{
			Method:  "POST",
			URL:     "https://example.com/api",
			Headers: map[string]string{"Content-Type": "application/json"},
			Body:    []byte(`{"ok":true}`),
		},
		Response: &models.HTTPResponse{
			Status:   201,
			Headers:  map[string]string{"Content-Type": "text/plain"},
			Body:     []byte("created"),
			Duration: 15 * time.Millisecond,
		},
		Timestamp: time.Date(2026, 9, 4, 12, 0, 0, 0, time.UTC),
		Comment:   "login",
		Tags:      []string{"auth"},
	}
	require.NoError(t, st.SaveTransaction(tx))
	require.NoError(t, st.SaveFinding(&models.AnalysisFinding{
		ID:            "find-1",
		TransactionID: tx.ID,
		RuleName:      "header_check",
		Severity:      models.SeverityLow,
		Title:         "missing CSP",
		Description:   "Content-Security-Policy is absent",
		Evidence:      "Content-Security-Policy",
	}))

	got, err := st.GetTransactions(10, 0)
	require.NoError(t, err)
	require.Len(t, got, 1)
	assert.Equal(t, tx.ID, got[0].ID)
	assert.Equal(t, "POST", got[0].Request.Method)
	assert.Equal(t, "https://example.com/api", got[0].Request.URL)
	assert.Equal(t, "application/json", got[0].Request.Headers["Content-Type"])
	assert.Equal(t, []byte(`{"ok":true}`), got[0].Request.Body)
	require.NotNil(t, got[0].Response)
	assert.Equal(t, 201, got[0].Response.Status)
	assert.Equal(t, []byte("created"), got[0].Response.Body)
	assert.Equal(t, 15*time.Millisecond, got[0].Response.Duration)
	assert.Equal(t, "login", got[0].Comment)
	assert.Equal(t, []string{"auth"}, got[0].Tags)
	assert.True(t, tx.Timestamp.Equal(got[0].Timestamp))

	one, err := st.GetTransaction(tx.ID)
	require.NoError(t, err)
	require.NotNil(t, one)
	assert.Equal(t, tx.Request.URL, one.Request.URL)
}

func TestGetTransactionsEmpty(t *testing.T) {
	path := filepath.Join(t.TempDir(), "empty.db")
	st := &Store{}
	require.NoError(t, st.Init(path))
	t.Cleanup(func() { require.NoError(t, st.Close()) })
	got, err := st.GetTransactions(5, 0)
	require.NoError(t, err)
	assert.Empty(t, got)
}

func TestIntelTargetAndArtifacts(t *testing.T) {
	path := filepath.Join(t.TempDir(), "intel.db")
	st := &Store{}
	require.NoError(t, st.Init(path))
	t.Cleanup(func() { require.NoError(t, st.Close()) })

	tg := intel.Target{ID: "t1", Domain: "lab.local", BaseURL: "https://lab.local/", Authorized: true, Created: time.Now().UTC()}
	require.NoError(t, st.CreateTarget(tg))
	got, err := st.GetTarget("t1")
	require.NoError(t, err)
	assert.Equal(t, "lab.local", got.Domain)
	assert.True(t, got.Authorized)
	require.NoError(t, st.AddArtifact(intel.Artifact{ID: "a1", TargetID: "t1", Kind: intel.KindPath, Value: "/api", Source: "dirs"}))
	arts, err := st.ListArtifacts("t1")
	require.NoError(t, err)
	require.Len(t, arts, 1)
	assert.Equal(t, "/api", arts[0].Value)
}
