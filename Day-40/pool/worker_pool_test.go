package pool_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"day-40/domain"
	"day-40/logger"
	"day-40/pool"
)

func TestWorkerPool_Execution(t *testing.T) {
	zapLog, _ := logger.NewZapLogger("test", "")
	numWorkers := 3
	numTasks := 6

	wp := pool.NewWorkerPool(numWorkers, numTasks, zapLog)
	ctx := context.Background()

	wp.Start(ctx)

	for i := 1; i <= numTasks; i++ {
		wp.Submit(domain.Task{
			ID:        i,
			JobType:   "TEST_JOB",
			Payload:   fmt.Sprintf("data_%d", i),
			CreatedAt: time.Now(),
		})
	}
	wp.CloseTasks()

	count := 0
	for res := range wp.Results() {
		count++
		if res.TaskID <= 0 {
			t.Errorf("invalid task id in result: %v", res)
		}
	}

	if count != numTasks {
		t.Errorf("expected %d completed tasks, got %d", numTasks, count)
	}
}
