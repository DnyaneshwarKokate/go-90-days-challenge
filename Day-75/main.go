package main

import (
	"fmt"

	"day75/snowflake"
)

func main() {
	fmt.Println("=== Day 75: Distributed ID Generation (Twitter Snowflake) ===")

	// Node 1 in datacenter US-East
	node1, _ := snowflake.NewSnowflakeNode(1)
	// Node 2 in datacenter EU-Central
	node2, _ := snowflake.NewSnowflakeNode(2)

	fmt.Println("\n--- Generating High-Throughput Snowflake 64-bit IDs ---")

	for i := 1; i <= 3; i++ {
		id1, _ := node1.NextID()
		id2, _ := node2.NextID()

		fmt.Printf("Worker #1 Generated ID: %d (Binary: %b)\n", id1, id1)
		fmt.Printf("Worker #2 Generated ID: %d (Binary: %b)\n", id2, id2)
	}
}
