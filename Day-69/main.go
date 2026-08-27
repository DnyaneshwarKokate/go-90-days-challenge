package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"day69/dlq"
)

func main() {
	fmt.Println("=== Day 69: Dead Letter Queue (DLQ) & Poison Message Handler ===")

	dlqPipeline := dlq.NewDLQHandler(3, 10)

	// Enqueue normal and malformed messages
	dlqPipeline.Enqueue(&dlq.QueueMessage{ID: "EVT-101", Payload: "ORDER_CREATED", CreatedAt: time.Now()})
	dlqPipeline.Enqueue(&dlq.QueueMessage{ID: "EVT-102", Payload: "POISON_PAYLOAD", CreatedAt: time.Now()})
	dlqPipeline.Enqueue(&dlq.QueueMessage{ID: "EVT-103", Payload: "PAYMENT_REFUNDED", CreatedAt: time.Now()})

	// Consumer processor
	messageProcessor := func(msg *dlq.QueueMessage) error {
		if msg.Payload == "POISON_PAYLOAD" {
			return errors.New("SQL injection payload detected in message body")
		}
		fmt.Printf("  ==> [CONSUMER PROCESSED] Event %s: %s\n", msg.ID, msg.Payload)
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	fmt.Println("\n--- Starting Asynchronous Queue Processing ---")
	for i := 0; i < 5; i++ {
		_ = dlqPipeline.ProcessNext(ctx, messageProcessor)
	}

	fmt.Println("\n--- Pipeline Metric Outcomes ---")
	fmt.Printf("Successfully Processed Messages: %d\n", len(dlqPipeline.ProcessedMessages()))

	dlqMsgs := dlqPipeline.DLQMessages()
	fmt.Printf("Dead Letter Queue (DLQ) Messages: %d\n", len(dlqMsgs))

	for _, msg := range dlqMsgs {
		fmt.Printf("  [DLQ RECORD] ID: %s | Retries: %d | Failure Reason: %s\n",
			msg.ID, msg.RetryCount, msg.ErrorReason)
	}
}
