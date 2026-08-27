package event

import (
	"context"
	"sync"
	"time"
)

type Event struct {
	ID        string
	Type      string
	AggregateID string
	Data      interface{}
	Timestamp time.Time
}

type EventHandler func(ctx context.Context, event Event) error

type EventBus struct {
	mu          sync.RWMutex
	handlers    map[string][]EventHandler
	outboxTable []Event
}

func NewEventBus() *EventBus {
	return &EventBus{
		handlers: make(map[string][]EventHandler),
	}
}

func (eb *EventBus) Subscribe(eventType string, handler EventHandler) {
	eb.mu.Lock()
	defer eb.mu.Unlock()

	eb.handlers[eventType] = append(eb.handlers[eventType], handler)
}

func (eb *EventBus) Publish(ctx context.Context, event Event) error {
	eb.mu.Lock()
	eb.outboxTable = append(eb.outboxTable, event) // Transactional Outbox write
	handlers := append([]EventHandler{}, eb.handlers[event.Type]...)
	eb.mu.Unlock()

	var wg sync.WaitGroup
	for _, h := range handlers {
		wg.Add(1)
		handler := h
		go func() {
			defer wg.Done()
			_ = handler(ctx, event)
		}()
	}
	wg.Wait()

	return nil
}

func (eb *EventBus) GetOutboxEvents() []Event {
	eb.mu.RLock()
	defer eb.mu.RUnlock()

	eventsCopy := make([]Event, len(eb.outboxTable))
	copy(eventsCopy, eb.outboxTable)
	return eventsCopy
}
