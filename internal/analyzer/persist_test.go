package analyzer

import (
	"errors"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"

	"meb/internal/models"
)

type memFindings struct {
	mu   sync.Mutex
	all  []*models.AnalysisFinding
	fail error
}

func (m *memFindings) SaveFinding(f *models.AnalysisFinding) error {
	if m.fail != nil {
		return m.fail
	}
	cp := *f
	m.mu.Lock()
	m.all = append(m.all, &cp)
	m.mu.Unlock()
	return nil
}

func TestPersistSavesFindings(t *testing.T) {
	t.Parallel()
	db := &memFindings{}
	eng := DefaultEngine()
	resp := respWith(200, nil, `{"ok":true}`)
	eng.Persist(zap.NewNop(), db, "tx-1", reqHTTPS(), resp)
	require.NotEmpty(t, db.all)
	for _, f := range db.all {
		assert.Equal(t, "tx-1", f.TransactionID)
		assert.NotEmpty(t, f.ID)
		assert.NotEmpty(t, f.RuleName)
		assert.NotEmpty(t, f.Title)
	}
}

func TestPersistSkipsNilStore(t *testing.T) {
	t.Parallel()
	DefaultEngine().Persist(zap.NewNop(), nil, "tx-1", reqHTTPS(), respWith(200, nil, "x"))
}

func TestPersistLogsSaveErrors(t *testing.T) {
	t.Parallel()
	core, logs := observer.New(zap.ErrorLevel)
	log := zap.New(core)
	db := &memFindings{fail: errors.New("disk full")}
	DefaultEngine().Persist(log, db, "tx-err", reqHTTPS(), respWith(200, nil, `{"ok":true}`))
	assert.Empty(t, db.all)
	require.NotEmpty(t, logs.All())
	assert.Contains(t, logs.All()[0].Message, "save finding")
}
