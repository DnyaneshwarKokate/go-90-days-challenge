package notification_test

import (
	"errors"
	"testing"
	"time"

	"day84/notification"
)

func TestWebhookEngineDeliveryAndHMAC(t *testing.T) {
	var receivedSig string

	mockDelivery := func(job *notification.WebhookJob, sig string) error {
		receivedSig = sig
		if job.TargetURL == "http://fail.com" {
			return errors.New("500 Internal Error")
		}
		return nil
	}

	engine := notification.NewWebhookEngine(2, 10, mockDelivery)

	job1 := &notification.WebhookJob{
		ID:          "wh-100",
		TargetURL:   "http://api.merchant.com/webhook",
		Payload:     `{"event":"payment_success"}`,
		Secret:      "whsec_test_secret_123",
		MaxAttempts: 2,
	}

	engine.Dispatch(job1)
	time.Sleep(50 * time.Millisecond)

	expectedSig := notification.GenerateHMACSignature(job1.Payload, job1.Secret)
	if receivedSig != expectedSig {
		t.Fatalf("HMAC signature mismatch! Expected %s, got %s", expectedSig, receivedSig)
	}

	delivered, _ := engine.Metrics()
	if delivered != 1 {
		t.Fatalf("Expected 1 delivered webhook, got %d", delivered)
	}

	engine.Shutdown()
}
