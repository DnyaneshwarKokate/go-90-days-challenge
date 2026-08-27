package queue_test

import (
	"testing"

	"day83/queue"
)

func TestKafkaStyleMessageBroker(t *testing.T) {
	broker := queue.NewMessageBroker()

	offset0 := broker.Publish("order-events", 0, `{"id":"ord-1"}`)
	offset1 := broker.Publish("order-events", 0, `{"id":"ord-2"}`)

	if offset0 != 0 || offset1 != 1 {
		t.Fatalf("Expected offsets 0 and 1, got %d and %d", offset0, offset1)
	}

	msgs := broker.ConsumeGroup("analytics-group", "order-events", 0, 10)
	if len(msgs) != 2 {
		t.Fatalf("Expected 2 consumed messages, got %d", len(msgs))
	}

	// Second consume should return 0 (Offset committed)
	msgsSecond := broker.ConsumeGroup("analytics-group", "order-events", 0, 10)
	if len(msgsSecond) != 0 {
		t.Fatalf("Expected 0 messages after offset commit, got %d", len(msgsSecond))
	}
}
