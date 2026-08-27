package main

import (
	"context"
	"log"
	"time"

	"day58/event"
)

func main() {
	bus := event.NewEventBus()
	log.Println("Starting Day 58 Event-Driven Architecture (EDA) & Outbox Demo...")

	bus.Subscribe("UserRegistered", func(ctx context.Context, evt event.Event) error {
		log.Printf("[Email Service Subscriber] Sending Welcome Email for User %s", evt.AggregateID)
		return nil
	})

	bus.Subscribe("UserRegistered", func(ctx context.Context, evt event.Event) error {
		log.Printf("[Analytics Service Subscriber] Tracking Registration for User %s", evt.AggregateID)
		return nil
	})

	_ = bus.Publish(context.Background(), event.Event{
		ID:          "evt_reg_001",
		Type:        "UserRegistered",
		AggregateID: "usr_777",
		Data:        map[string]string{"email": "user@example.com"},
		Timestamp:   time.Now(),
	})

	log.Printf("Transactional Outbox total events stored: %d", len(bus.GetOutboxEvents()))
}
