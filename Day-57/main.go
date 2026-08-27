package main

import (
	"log"
	"time"

	"day57/queue"
)

func main() {
	broker := queue.NewMessageBroker(50)
	log.Println("Starting Day 57 Message Queue / Pub-Sub Broker Demonstration...")

	broker.Subscribe("payment.completed", func(msg queue.Message) bool {
		log.Printf("[Consumer ACK] Received payment event: ID=%s Payload=%s", msg.ID, msg.Payload)
		return true
	})

	_ = broker.Publish("payment.completed", queue.Message{
		ID:        "evt_9901",
		Topic:     "payment.completed",
		Payload:   `{"payment_id":"pay_888","amount":250.00}`,
		Timestamp: time.Now(),
	})

	time.Sleep(100 * time.Millisecond)
}
