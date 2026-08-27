package dlq

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// QueueMessage represents an asynchronous payload in a queue pipeline.
type QueueMessage struct {
	ID          string
	Payload     string
	RetryCount  int
	ErrorReason string
	CreatedAt   time.Time
}

// DLQHandler isolates failed / unprocessable messages.
type DLQHandler struct {
	mu           sync.RWMutex
	maxRetries   int
	mainQueue    chan *QueueMessage
	deadLetterQueue []*QueueMessage
	processed    []*QueueMessage
}

// NewDLQHandler initializes main queue worker with dead-letter storage.
func NewDLQHandler(maxRetries int, queueCapacity int) *DLQHandler {
	return &DLQHandler{
		maxRetries:      maxRetries,
		mainQueue:       make(chan *QueueMessage, queueCapacity),
		deadLetterQueue: make([]*QueueMessage, 0),
		processed:       make([]*QueueMessage, 0),
	}
}

// Enqueue adds a message to the main consumer pipeline.
func (h *DLQHandler) Enqueue(msg *QueueMessage) {
	h.mainQueue <- msg
}

// ProcessNext attempts message processing; routes to DLQ if max retries exceeded.
func (h *DLQHandler) ProcessNext(ctx context.Context, processor func(msg *QueueMessage) error) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case msg := <-h.mainQueue:
		err := processor(msg)
		if err != nil {
			msg.RetryCount++
			msg.ErrorReason = err.Error()

			if msg.RetryCount >= h.maxRetries {
				fmt.Printf("[DLQ ROUTER] Poison message %s exceeded max retries (%d/%d). Moving to Dead Letter Queue.\n",
					msg.ID, msg.RetryCount, h.maxRetries)

				h.mu.Lock()
				h.deadLetterQueue = append(h.deadLetterQueue, msg)
				h.mu.Unlock()
				return fmt.Errorf("message %s routed to DLQ: %w", msg.ID, err)
			}

			fmt.Printf("[QUEUE WORKER] Message %s processing failed (Attempt %d/%d). Re-queueing...\n",
				msg.ID, msg.RetryCount, h.maxRetries)
			h.mainQueue <- msg
			return err
		}

		h.mu.Lock()
		h.processed = append(h.processed, msg)
		h.mu.Unlock()
		return nil
	}
}

// DLQMessages returns all messages currently residing in the Dead Letter Queue.
func (h *DLQHandler) DLQMessages() []*QueueMessage {
	h.mu.RLock()
	defer h.mu.RUnlock()
	copied := make([]*QueueMessage, len(h.deadLetterQueue))
	copy(copied, h.deadLetterQueue)
	return copied
}

// ProcessedMessages returns all successfully handled messages.
func (h *DLQHandler) ProcessedMessages() []*QueueMessage {
	h.mu.RLock()
	defer h.mu.RUnlock()
	copied := make([]*QueueMessage, len(h.processed))
	copy(copied, h.processed)
	return copied
}
