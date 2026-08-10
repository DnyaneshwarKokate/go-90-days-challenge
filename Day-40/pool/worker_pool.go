package pool

import (
	"context"
	"fmt"
	"sync"
	"time"

	"day-40/domain"
)

type WorkerPool struct {
	numWorkers  int
	tasksChan   chan domain.Task
	resultsChan chan domain.TaskResult
	wg          sync.WaitGroup
	logger      domain.Logger
}

func NewWorkerPool(numWorkers int, queueSize int, logger domain.Logger) *WorkerPool {
	return &WorkerPool{
		numWorkers:  numWorkers,
		tasksChan:   make(chan domain.Task, queueSize),
		resultsChan: make(chan domain.TaskResult, queueSize),
		logger:      logger,
	}
}

func (p *WorkerPool) Start(ctx context.Context) {
	p.logger.Info(ctx, "Starting Worker Pool", "worker_count", p.numWorkers)

	for i := 1; i <= p.numWorkers; i++ {
		p.wg.Add(1)
		go p.worker(ctx, i)
	}

	// Monitor worker completion and close results channel
	go func() {
		p.wg.Wait()
		close(p.resultsChan)
		p.logger.Info(ctx, "All workers completed, results channel closed")
	}()
}

func (p *WorkerPool) Submit(t domain.Task) {
	p.tasksChan <- t
}

func (p *WorkerPool) CloseTasks() {
	close(p.tasksChan)
}

func (p *WorkerPool) Results() <-chan domain.TaskResult {
	return p.resultsChan
}

func (p *WorkerPool) worker(ctx context.Context, workerID int) {
	defer p.wg.Done()
	p.logger.Debug(ctx, "Worker goroutine spawned", "worker_id", workerID)

	for task := range p.tasksChan {
		start := time.Now()
		p.logger.Info(ctx, "Worker processing task", "worker_id", workerID, "task_id", task.ID, "job_type", task.JobType)

		// Simulate CPU/IO processing
		time.Sleep(50 * time.Millisecond)

		res := domain.TaskResult{
			TaskID:      task.ID,
			WorkerID:    workerID,
			Result:      fmt.Sprintf("SUCCESS: Processed [%s] for %s", task.JobType, task.Payload),
			ExecutionMs: time.Since(start),
		}

		p.resultsChan <- res
	}
}
