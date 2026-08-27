package system

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type SystemOrder struct {
	ID       string
	Customer string
	Amount   float64
	Status   string
}

type EnterpriseSystem struct {
	mu           sync.RWMutex
	orders       map[string]*SystemOrder
	outboxEvents []string
	readViews    map[string]string
	dlq          []string
	activeSpans  []string
}

func NewEnterpriseSystem() *EnterpriseSystem {
	return &EnterpriseSystem{
		orders:       make(map[string]*SystemOrder),
		outboxEvents: make([]string, 0),
		readViews:    make(map[string]string),
		dlq:          make([]string, 0),
		activeSpans:  make([]string, 0),
	}
}

func (sys *EnterpriseSystem) ProcessFulfillmentPipeline(ctx context.Context, order SystemOrder) error {
	sys.mu.Lock()
	defer sys.mu.Unlock()

	traceID := fmt.Sprintf("trace-%s-%d", order.ID, time.Now().UnixNano())
	sys.activeSpans = append(sys.activeSpans, fmt.Sprintf("Span:StartFulfillment|TraceID:%s", traceID))

	// 1. Write Order entity + Outbox event atomically
	order.Status = "CONFIRMED"
	sys.orders[order.ID] = &order
	sys.outboxEvents = append(sys.outboxEvents, fmt.Sprintf("evt_outbox_%s", order.ID))

	// 2. Project Read Model (CQRS)
	sys.readViews[order.ID] = fmt.Sprintf("READ_VIEW [Customer: %s | Amount: $%.2f | Status: %s]",
		order.Customer, order.Amount, order.Status)

	sys.activeSpans = append(sys.activeSpans, fmt.Sprintf("Span:CompleteFulfillment|TraceID:%s", traceID))
	return nil
}

func (sys *EnterpriseSystem) RoutePoisonToDLQ(msgID string, reason string) {
	sys.mu.Lock()
	defer sys.mu.Unlock()
	sys.dlq = append(sys.dlq, fmt.Sprintf("DLQ_RECORD [ID: %s | Reason: %s]", msgID, reason))
}

func (sys *EnterpriseSystem) Summary() (ordersCount, outboxCount, readViewsCount, dlqCount int) {
	sys.mu.RLock()
	defer sys.mu.RUnlock()
	return len(sys.orders), len(sys.outboxEvents), len(sys.readViews), len(sys.dlq)
}
