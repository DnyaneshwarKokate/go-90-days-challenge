package outbox

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

type EventStatus string

const (
	StatusPending   EventStatus = "PENDING"
	StatusProcessed EventStatus = "PROCESSED"
	StatusFailed    EventStatus = "FAILED"
)

// OutboxEvent represents a domain event saved atomically in the database outbox table.
type OutboxEvent struct {
	ID            string
	AggregateType string
	AggregateID   string
	EventType     string
	Payload       string
	Status        EventStatus
	CreatedAt     time.Time
	ProcessedAt   *time.Time
	Attempts      int
}

// OrderRecord represents an application entity in business storage.
type OrderRecord struct {
	ID        string
	Customer  string
	Amount    float64
	Status    string
	CreatedAt time.Time
}

// DatabaseSimulator simulates a relational database with transactional outbox support.
type DatabaseSimulator struct {
	mu     sync.RWMutex
	orders map[string]OrderRecord
	events map[string]*OutboxEvent
}

// NewDatabaseSimulator initializes the mock DB store.
func NewDatabaseSimulator() *DatabaseSimulator {
	return &DatabaseSimulator{
		orders: make(map[string]OrderRecord),
		events: make(map[string]*OutboxEvent),
	}
}

// CreateOrderWithOutbox executes an atomic dual-write (Order record + Outbox event).
func (db *DatabaseSimulator) CreateOrderWithOutbox(order OrderRecord, eventType string, payload string) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if _, exists := db.orders[order.ID]; exists {
		return fmt.Errorf("order ID %s already exists", order.ID)
	}

	// 1. Write Order entity
	db.orders[order.ID] = order

	// 2. Write Outbox Event atomically in same transaction scope
	eventID := fmt.Sprintf("evt_%s_%d", order.ID, time.Now().UnixNano())
	event := &OutboxEvent{
		ID:            eventID,
		AggregateType: "ORDER",
		AggregateID:   order.ID,
		EventType:     eventType,
		Payload:       payload,
		Status:        StatusPending,
		CreatedAt:     time.Now(),
		Attempts:      0,
	}

	db.events[eventID] = event
	return nil
}

// FetchPendingEvents queries outbox events in PENDING status.
func (db *DatabaseSimulator) FetchPendingEvents(limit int) []*OutboxEvent {
	db.mu.RLock()
	defer db.mu.RUnlock()

	pending := make([]*OutboxEvent, 0)
	for _, evt := range db.events {
		if evt.Status == StatusPending {
			pending = append(pending, evt)
			if len(pending) >= limit {
				break
			}
		}
	}
	return pending
}

// MarkEventProcessed updates outbox record status to PROCESSED.
func (db *DatabaseSimulator) MarkEventProcessed(eventID string) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	evt, exists := db.events[eventID]
	if !exists {
		return errors.New("event not found")
	}

	now := time.Now()
	evt.Status = StatusProcessed
	evt.ProcessedAt = &now
	return nil
}

// OutboxProcessor polls the outbox database table and publishes events reliably.
type OutboxProcessor struct {
	db        *DatabaseSimulator
	publisher func(evt *OutboxEvent) error
}

// NewOutboxProcessor creates a processor daemon.
func NewOutboxProcessor(db *DatabaseSimulator, publisher func(evt *OutboxEvent) error) *OutboxProcessor {
	return &OutboxProcessor{
		db:        db,
		publisher: publisher,
	}
}

// ProcessBatch scans and dispatches pending outbox events.
func (p *OutboxProcessor) ProcessBatch(ctx context.Context) (int, error) {
	events := p.db.FetchPendingEvents(10)
	processedCount := 0

	for _, evt := range events {
		select {
		case <-ctx.Done():
			return processedCount, ctx.Err()
		default:
		}

		fmt.Printf("[OUTBOX PROCESSOR] Dispatching event %s (%s)...\n", evt.ID, evt.EventType)
		if err := p.publisher(evt); err != nil {
			fmt.Printf("[OUTBOX PROCESSOR] Failed to publish event %s: %v\n", evt.ID, err)
			continue
		}

		if err := p.db.MarkEventProcessed(evt.ID); err != nil {
			return processedCount, fmt.Errorf("failed to mark processed: %w", err)
		}

		processedCount++
	}

	return processedCount, nil
}
