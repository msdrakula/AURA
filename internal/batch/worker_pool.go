// Package batch provides a reusable worker pool for parallel task processing.
package batch

import (
	"context"
	"fmt"
	"sync"
)

// Task is one unit of work identified by ID.
type Task[T any] struct {
	ID      string
	Payload T
}

// Result is the outcome of processing one Task.
type Result[T any, R any] struct {
	TaskID string
	Data   R
	Err    error
}

// WorkerFunc processes a single task payload.
type WorkerFunc[T any, R any] func(ctx context.Context, task T) (R, error)

// RunWorkerPool starts workers goroutines that read from tasks, call workerFn,
// and write to results. It closes results after every worker returns.
// If ctx is cancelled, it returns ctx.Err() after in-flight workers stop.
func RunWorkerPool[T any, R any](
	ctx context.Context,
	tasks <-chan Task[T],
	results chan<- Result[T, R],
	workers int,
	workerFn WorkerFunc[T, R],
) error {
	if ctx == nil {
		ctx = context.Background()
	}
	if results == nil {
		return fmt.Errorf("batch: nil results channel")
	}
	if workerFn == nil {
		return fmt.Errorf("batch: nil worker function")
	}
	if workers < 1 {
		workers = 1
	}

	defer close(results)

	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			runWorker(ctx, tasks, results, workerFn)
		}()
	}
	wg.Wait()
	return ctx.Err()
}

func runWorker[T any, R any](
	ctx context.Context,
	tasks <-chan Task[T],
	results chan<- Result[T, R],
	workerFn WorkerFunc[T, R],
) {
	for {
		select {
		case <-ctx.Done():
			return
		case task, ok := <-tasks:
			if !ok {
				return
			}
			if ctx.Err() != nil {
				return
			}
			data, err := workerFn(ctx, task.Payload)
			res := Result[T, R]{TaskID: task.ID, Data: data, Err: err}
			select {
			case <-ctx.Done():
				return
			case results <- res:
			}
		}
	}
}
