package main

import (
	"fmt"
	"runtime"

	"day86/memory"
)

func main() {
	fmt.Println("=== Day 86: Go Memory Management, GC & Escape Analysis ===")

	fmt.Println("\n--- 1. Testing Stack vs Heap Escaped Allocation ---")
	stackVal := memory.StackAllocation(50)
	heapPtr := memory.HeapEscapedAllocation(1001)

	fmt.Printf("Stack Value Result: %d\n", stackVal)
	fmt.Printf("Heap Escaped Pointer: %p (Struct ID: %d)\n", heapPtr, heapPtr.ID)

	fmt.Println("\n--- 2. Triggering Manual Garbage Collection (runtime.GC()) ---")
	runtime.GC()

	heapAlloc, numGC, pauseNs := memory.GetGCMetrics()
	fmt.Printf("Active Heap Allocations: %d bytes\n", heapAlloc)
	fmt.Printf("Total GC Runs Performed: %d\n", numGC)
	fmt.Printf("Total GC Pause Duration: %d ns\n", pauseNs)
}
