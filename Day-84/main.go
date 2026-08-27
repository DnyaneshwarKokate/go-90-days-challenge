package main

import (
	"fmt"
	"time"

	"day84/notification"
)

func main() {
	fmt.Println("=== Day 84: High-Scale Notification & Webhook Delivery Engine ===")

	mockHTTPClient := func(job *notification.WebhookJob, sig string) error {
		fmt.Printf("  [HTTP CLIENT] Sending Webhook -> URL: %s | HMAC: %s\n", job.TargetURL, sig[:16]+"...")
		return nil
	}

	engine := notification.NewWebhookEngine(3, 10, mockHTTPClient)

	fmt.Println("\n--- Dispatching Asynchronous Merchant Webhooks ---")
	job1 := &notification.WebhookJob{
		ID:          "wh-1",
		TargetURL:   "https://merchant.com/api/webhooks",
		Payload:     `{"event":"charge.succeeded","amount":99.00}`,
		Secret:      "whsec_secret_key_88",
		MaxAttempts: 3,
	}

	engine.Dispatch(job1)
	time.Sleep(100 * time.Millisecond)

	delivered, failed := engine.Metrics()
	fmt.Printf("\n--- Webhook Delivery Metrics ---\nDelivered: %d | Failed: %d\n", delivered, failed)
	engine.Shutdown()
}
