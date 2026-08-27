package main

import (
	"fmt"

	"day83/queue"
)

func main() {
	fmt.Println("=== Day 83: Distributed Message Broker Engine (Kafka Architecture) ===")

	broker := queue.NewMessageBroker()
	topic := "payment-transactions"
	partition := 0

	fmt.Println("\n--- 1. Producer Publishing Messages ---")
	off1 := broker.Publish(topic, partition, `{"tx":"TX_1001","amount":150.00}`)
	off2 := broker.Publish(topic, partition, `{"tx":"TX_1002","amount":320.50}`)
	fmt.Printf("[PRODUCER] Published 2 messages to '%s' Partition %d. Offsets: %d, %d\n", topic, partition, off1, off2)

	fmt.Println("\n--- 2. Consumer Group 'billing-service' Consuming Batch 1 ---")
	consumed1 := broker.ConsumeGroup("billing-service", topic, partition, 10)
	for _, m := range consumed1 {
		fmt.Printf("  [CONSUMED] Offset: %d | Payload: %s\n", m.Offset, m.Payload)
	}

	fmt.Println("\n--- 3. Consumer Group 'billing-service' Consuming Batch 2 (Committed Offset) ---")
	consumed2 := broker.ConsumeGroup("billing-service", topic, partition, 10)
	fmt.Printf("New Messages Available After Offset Commit: %d\n", len(consumed2))
}
