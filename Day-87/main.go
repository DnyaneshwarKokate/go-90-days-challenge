package main

import (
	"fmt"
	"sync"

	"day87/concurrency"
)

func main() {
	fmt.Println("=== Day 87: Advanced Concurrency & Data Race Prevention (`-race`) ===")

	counter := concurrency.NewSafeCounter()
	bufferPool := concurrency.NewBufferPool()
	var wg sync.WaitGroup

	fmt.Println("\n--- 1. Launching 100 Concurrent Routines (Atomic Increment) ---")
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			counter.Increment()

			buf := bufferPool.Get()
			buf.WriteString(fmt.Sprintf("Worker #%d payload", id))
			bufferPool.Put(buf)
		}(i)
	}

	wg.Wait()

	fmt.Printf("[RACE-FREE COUNTER] Total Increments: %d\n", counter.Value())
	fmt.Println("[SYNC POOL] All recycled buffer allocations returned cleanly to memory pool.")
}
