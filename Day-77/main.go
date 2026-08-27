package main

import (
	"fmt"

	"day77/stream"
)

func main() {
	fmt.Println("=== Day 77: Event Streaming & Log Compaction Engine ===")

	eventLog := stream.NewEventLog()

	fmt.Println("\n--- 1. Appending Stream Events ---")
	eventLog.Append("k8s.pod.web-1", "STATUS: PENDING")
	eventLog.Append("k8s.pod.web-1", "STATUS: RUNNING")
	eventLog.Append("k8s.pod.db-1", "STATUS: RUNNING")
	eventLog.Append("k8s.pod.web-1", "STATUS: TERMINATED")

	records, _ := eventLog.ReadFromOffset(0, 10)
	for _, r := range records {
		fmt.Printf("  [OFFSET %d] Key: %-15s | Value: %s\n", r.Offset, r.Key, r.Value)
	}

	fmt.Println("\n--- 2. Executing Key-Based Log Compaction ---")
	removed := eventLog.Compact()
	fmt.Printf("Log Compaction Complete. Removed %d obsolete records.\n", removed)

	compacted, _ := eventLog.ReadFromOffset(0, 10)
	fmt.Println("\n--- Compacted Stream State ---")
	for _, r := range compacted {
		fmt.Printf("  [OFFSET %d] Key: %-15s | Value: %s\n", r.Offset, r.Key, r.Value)
	}
}
