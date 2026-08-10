package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"day-40/domain"
	"day-40/logger"
	"day-40/pool"
)

func main() {
	fmt.Println("==========================================================================")
	fmt.Println("🚀 Day 40: Goroutine Worker Pools in Go (Channels, WaitGroup, Rate Control)")
	fmt.Println("==========================================================================")

	logFilePath := "worker_app.log"
	zapLog, err := logger.NewZapLogger("development", logFilePath)
	if err != nil {
		fmt.Printf("Failed to initialize Zap Logger: %v\n", err)
		os.Exit(1)
	}
	defer zapLog.Sync()

	ctx := context.Background()

	numWorkers := 4
	totalTasks := 12
	wp := pool.NewWorkerPool(numWorkers, totalTasks, zapLog)

	startOverall := time.Now()
	wp.Start(ctx)

	fmt.Printf("\n--- 1️⃣ Submitting %d Tasks to %d-Worker Pool ---\n", totalTasks, numWorkers)
	for i := 1; i <= totalTasks; i++ {
		jobType := "IMAGE_RESIZE"
		if i%2 == 0 {
			jobType = "PDF_GENERATE"
		}
		wp.Submit(domain.Task{
			ID:        i,
			JobType:   jobType,
			Payload:   fmt.Sprintf("file_%d.dat", i),
			CreatedAt: time.Now(),
		})
	}
	wp.CloseTasks()

	fmt.Println("\n--- 2️⃣ Collecting Aggregated Worker Results ---")
	completedCount := 0
	for result := range wp.Results() {
		completedCount++
		fmt.Printf("Result #%02d | Task #%02d | Processed by Worker #%d | %s (%v)\n",
			completedCount, result.TaskID, result.WorkerID, result.Result, result.ExecutionMs)
	}

	totalDuration := time.Since(startOverall)
	fmt.Printf("\n🎉 All %d tasks processed in %v using %d Workers! (Sequential would take ~%v)\n",
		completedCount, totalDuration, numWorkers, time.Duration(totalTasks*50)*time.Millisecond)

	fmt.Println("\n✅ Day 40 Goroutine Worker Pools executed successfully! Check worker_app.log for logs.")
}
