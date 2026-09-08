package batch

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRunWorkerPoolSquares(t *testing.T) {
	t.Parallel()

	const n = 100
	tasks := make(chan Task[int], n)
	results := make(chan Result[int, int], n)
	for i := 1; i <= n; i++ {
		tasks <- Task[int]{ID: strconv.Itoa(i), Payload: i}
	}
	close(tasks)

	err := RunWorkerPool(context.Background(), tasks, results, 10, func(_ context.Context, v int) (int, error) {
		return v * v, nil
	})
	require.NoError(t, err)

	got := make(map[string]int, n)
	for res := range results {
		require.NoError(t, res.Err)
		got[res.TaskID] = res.Data
	}
	require.Len(t, got, n)
	for i := 1; i <= n; i++ {
		assert.Equal(t, i*i, got[strconv.Itoa(i)], "task %d", i)
	}
}

func TestRunWorkerPoolCancel(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	tasks := make(chan Task[int])
	results := make(chan Result[int, int])

	done := make(chan error, 1)
	go func() {
		done <- RunWorkerPool(ctx, tasks, results, 10, func(ctx context.Context, v int) (int, error) {
			select {
			case <-ctx.Done():
				return 0, ctx.Err()
			case <-time.After(time.Second):
				return v * v, nil
			}
		})
	}()
	go func() {
		for range results {
		}
	}()

	for i := 1; i <= 10; i++ {
		select {
		case tasks <- Task[int]{ID: strconv.Itoa(i), Payload: i}:
		case <-time.After(time.Second):
			t.Fatal("timed out sending tasks")
		}
	}

	time.Sleep(10 * time.Millisecond)
	cancel()

	select {
	case err := <-done:
		require.Error(t, err)
		assert.ErrorIs(t, err, context.Canceled)
	case <-time.After(2 * time.Second):
		t.Fatal("RunWorkerPool did not return after cancel")
	}
}
