package notification

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"sync"
	"sync/atomic"
	"time"
)

type WebhookStatus string

const (
	StatusDelivered WebhookStatus = "DELIVERED"
	StatusFailed    WebhookStatus = "FAILED"
)

// WebhookJob represents a webhook dispatch task.
type WebhookJob struct {
	ID          string
	TargetURL   string
	Payload     string
	Secret      string
	MaxAttempts int
	Attempts    int
	Status      WebhookStatus
}

// WebhookEngine manages asynchronous delivery worker pools.
type WebhookEngine struct {
	mu           sync.Mutex
	delivered    int64
	failed       int64
	jobs         chan *WebhookJob
	deliveryFunc func(job *WebhookJob, signature string) error
}

// NewWebhookEngine initializes worker engine.
func NewWebhookEngine(workers int, queueCap int, deliveryFunc func(job *WebhookJob, sig string) error) *WebhookEngine {
	e := &WebhookEngine{
		jobs:         make(chan *WebhookJob, queueCap),
		deliveryFunc: deliveryFunc,
	}

	for i := 0; i < workers; i++ {
		go e.workerLoop()
	}
	return e
}

func GenerateHMACSignature(payload string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(payload))
	return hex.EncodeToString(h.Sum(nil))
}

// Dispatch adds job to the delivery queue.
func (e *WebhookEngine) Dispatch(job *WebhookJob) {
	e.jobs <- job
}

func (e *WebhookEngine) workerLoop() {
	for job := range e.jobs {
		sig := GenerateHMACSignature(job.Payload, job.Secret)
		success := false

		for job.Attempts < job.MaxAttempts {
			job.Attempts++
			err := e.deliveryFunc(job, sig)
			if err == nil {
				job.Status = StatusDelivered
				atomic.AddInt64(&e.delivered, 1)
				success = true
				break
			}
			time.Sleep(10 * time.Millisecond)
		}

		if !success {
			job.Status = StatusFailed
			atomic.AddInt64(&e.failed, 1)
		}
	}
}

// Metrics returns delivery counters.
func (e *WebhookEngine) Metrics() (delivered, failed int64) {
	return atomic.LoadInt64(&e.delivered), atomic.LoadInt64(&e.failed)
}

// Shutdown closes the job queue.
func (e *WebhookEngine) Shutdown() {
	close(e.jobs)
}
