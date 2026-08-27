package queue

import (
	"sync"
	"testing"
	"time"
)

func TestMessageBrokerPublishSubscribeAndDLQ(t *testing.T) {
	broker := NewMessageBroker(10)
	var wg sync.WaitGroup
	wg.Add(1)

	var receivedPayload string

	// Subscribe to topic
	broker.Subscribe("orders", func(msg Message) bool {
		if msg.Payload == "FAIL" {
			return false // Trigger NACK -> DLQ
		}
		receivedPayload = msg.Payload
		wg.Done()
		return true
	})

	// Publish success message
	err := broker.Publish("orders", Message{
		ID:        "msg_001",
		Topic:     "orders",
		Payload:   "order_created_1001",
		Timestamp: time.Now(),
	})
	if err != nil {
		t.Fatalf("Failed to publish message: %v", err)
	}

	wg.Wait()

	if receivedPayload != "order_created_1001" {
		t.Errorf("Expected payload order_created_1001, got %s", receivedPayload)
	}

	// Publish failing message to trigger DLQ
	_ = broker.Publish("orders", Message{
		ID:        "msg_002",
		Topic:     "orders",
		Payload:   "FAIL",
		Timestamp: time.Now(),
	})

	time.Sleep(50 * time.Millisecond)

	dlqMsgs := broker.ReadDLQ()
	if len(dlqMsgs) != 1 {
		t.Fatalf("Expected 1 message in DLQ, got %d", len(dlqMsgs))
	}

	if dlqMsgs[0].ID != "msg_002" {
		t.Errorf("Expected DLQ message ID msg_002, got %s", dlqMsgs[0].ID)
	}
}
