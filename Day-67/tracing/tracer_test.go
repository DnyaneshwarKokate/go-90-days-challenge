package tracing_test

import (
	"context"
	"net/http/httptest"
	"testing"

	"day67/tracing"
)

func TestTracerParentChildSpans(t *testing.T) {
	tracer := tracing.NewTracer()

	ctx, rootSpan := tracer.StartSpan(context.Background(), "RootController")
	rootSpan.SetTag("user_id", "usr_100")

	_, childSpan := tracer.StartSpan(ctx, "DatabaseQuery")
	childSpan.SetTag("db.statement", "SELECT * FROM users")

	rootSpan.Finish()
	childSpan.Finish()

	if childSpan.TraceID != rootSpan.TraceID {
		t.Fatalf("Expected child TraceID %s to match root TraceID %s", childSpan.TraceID, rootSpan.TraceID)
	}

	if childSpan.ParentSpanID != rootSpan.SpanID {
		t.Fatalf("Expected child ParentSpanID %s to equal root SpanID %s", childSpan.ParentSpanID, rootSpan.SpanID)
	}

	spans := tracer.Spans()
	if len(spans) != 2 {
		t.Fatalf("Expected 2 collected spans, got %d", len(spans))
	}
}

func TestW3CTraceContextHTTPPropagation(t *testing.T) {
	tracer := tracing.NewTracer()

	// Service A creates outgoing request
	_, spanA := tracer.StartSpan(context.Background(), "ServiceA_Client")
	req := httptest.NewRequest("GET", "http://service-b/api/data", nil)
	tracer.InjectHTTP(spanA, req)

	traceHeader := req.Header.Get("traceparent")
	if traceHeader == "" {
		t.Fatalf("Expected 'traceparent' header to be injected into request")
	}

	// Service B receives request and extracts context
	ctxB, spanB := tracer.ExtractHTTP(context.Background(), req)
	_ = ctxB

	if spanB.TraceID != spanA.TraceID {
		t.Fatalf("Cross-service TraceID mismatch! Service A: %s | Service B: %s", spanA.TraceID, spanB.TraceID)
	}

	if spanB.ParentSpanID != spanA.SpanID {
		t.Fatalf("Service B ParentSpanID %s did not equal Service A SpanID %s", spanB.ParentSpanID, spanA.SpanID)
	}
}
