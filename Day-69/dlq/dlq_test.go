package dlq_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"day69/dlq"
)

func TestDLQPoisonMessageRouting(t *testing.T) {
	// Max 2 retries before routing to DLQ
	handler := dlq.NewDLQHandler(2, 5)

	validMsg := &dlq.QueueMessage{ID: "msg-101", Payload: "valid_json", CreatedAt: time.Now()}
	poisonMsg := &dlq.QueueMessage{ID: "msg-666", Payload: "corrupted_poison_payload", CreatedAt: time.Now()}

	handler.Enqueue(validMsg)
	handler.Enqueue(poisonMsg)

	processor := func(msg *dlq.QueueMessage) error {
		if msg.Payload == "corrupted_poison_payload" {
			return errors.New("unparseable JSON body")
		}
		return nil
	}

	ctx := context.Background()

	// Process msg-101 (Success)
	_ = handler.ProcessNext(ctx, processor)

	// Process msg-666 Attempt 1 (Fails, requeued)
	_ = handler.ProcessNext(ctx, processor)

	// Process msg-666 Attempt 2 (Fails, routed to DLQ)
	_ = handler.ProcessNext(ctx, processor)

	processed := handler.ProcessedMessages()
	if len(processed) != 1 || processed[0].ID != "msg-101" {
		t.Fatalf("Expected msg-101 in processed list, got %v", processed)
	}

	deadLetters := handler.DLQMessages()
	if len(deadLetters) != 1 || deadLetters[0].ID != "msg-666" {
		t.Fatalf("Expected msg-666 in DLQ, got %v", deadLetters)
	}
}
