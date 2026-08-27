package tracing

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

type contextKey string

const traceSpanKey contextKey = "opentelemetry_span"

// Span represents a single operation unit in a distributed trace graph.
type Span struct {
	TraceID      string
	SpanID       string
	ParentSpanID string
	Operation    string
	StartTime    time.Time
	EndTime      time.Time
	Tags         map[string]string
}

// Tracer manages span lifecycle and context propagation across RPC/HTTP boundaries.
type Tracer struct {
	mu          sync.RWMutex
	activeSpans []*Span
}

// NewTracer creates an OpenTelemetry-compatible tracer instance.
func NewTracer() *Tracer {
	return &Tracer{
		activeSpans: make([]*Span, 0),
	}
}

func generateID(bytesLen int) string {
	b := make([]byte, bytesLen)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

// StartSpan creates a new root or child span attached to the context.
func (t *Tracer) StartSpan(ctx context.Context, operationName string) (context.Context, *Span) {
	var traceID, parentSpanID string

	parentSpan, ok := ctx.Value(traceSpanKey).(*Span)
	if ok && parentSpan != nil {
		traceID = parentSpan.TraceID
		parentSpanID = parentSpan.SpanID
	} else {
		traceID = generateID(16) // 128-bit hex trace ID
	}

	spanID := generateID(8) // 64-bit hex span ID

	span := &Span{
		TraceID:      traceID,
		SpanID:       spanID,
		ParentSpanID: parentSpanID,
		Operation:    operationName,
		StartTime:    time.Now(),
		Tags:         make(map[string]string),
	}

	t.mu.Lock()
	t.activeSpans = append(t.activeSpans, span)
	t.mu.Unlock()

	newCtx := context.WithValue(ctx, traceSpanKey, span)
	return newCtx, span
}

// Finish records the span end timestamp.
func (s *Span) Finish() {
	s.EndTime = time.Now()
}

// SetTag adds metadata to the span.
func (s *Span) SetTag(key, value string) {
	s.Tags[key] = value
}

// InjectHTTP writes W3C TraceParent headers to an outgoing HTTP request.
func (t *Tracer) InjectHTTP(span *Span, req *http.Request) {
	if span == nil || req == nil {
		return
	}
	// W3C Trace Context specification: 00-{trace_id}-{span_id}-01
	traceparent := fmt.Sprintf("00-%s-%s-01", span.TraceID, span.SpanID)
	req.Header.Set("traceparent", traceparent)
}

// ExtractHTTP reads W3C TraceParent headers from an incoming HTTP request into a Context.
func (t *Tracer) ExtractHTTP(ctx context.Context, req *http.Request) (context.Context, *Span) {
	headerVal := req.Header.Get("traceparent")
	if headerVal == "" {
		return t.StartSpan(ctx, "http_server_request")
	}

	parts := strings.Split(headerVal, "-")
	if len(parts) < 4 {
		return t.StartSpan(ctx, "http_server_request")
	}

	traceID := parts[1]
	parentSpanID := parts[2]
	spanID := generateID(8)

	span := &Span{
		TraceID:      traceID,
		SpanID:       spanID,
		ParentSpanID: parentSpanID,
		Operation:    "http_server_request",
		StartTime:    time.Now(),
		Tags:         make(map[string]string),
	}

	t.mu.Lock()
	t.activeSpans = append(t.activeSpans, span)
	t.mu.Unlock()

	return context.WithValue(ctx, traceSpanKey, span), span
}

// Spans returns all collected trace spans.
func (t *Tracer) Spans() []*Span {
	t.mu.RLock()
	defer t.mu.RUnlock()
	copied := make([]*Span, len(t.activeSpans))
	copy(copied, t.activeSpans)
	return copied
}
