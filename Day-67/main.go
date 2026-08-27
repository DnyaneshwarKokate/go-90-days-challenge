package main

import (
	"context"
	"fmt"
	"net/http/httptest"
	"time"

	"day67/tracing"
)

func main() {
	fmt.Println("=== Day 67: Distributed Tracing & OpenTelemetry Context Propagation ===")

	tracer := tracing.NewTracer()

	// 1. Root Gateway Service receives incoming request
	_, gatewaySpan := tracer.StartSpan(context.Background(), "APIGateway_HandleCheckout")
	gatewaySpan.SetTag("http.method", "POST")
	gatewaySpan.SetTag("http.path", "/checkout")
	time.Sleep(15 * time.Millisecond)

	// 2. Gateway calls Order Microservice via HTTP (Inject W3C traceparent header)
	reqToOrder := httptest.NewRequest("POST", "http://order-service/api/orders", nil)
	tracer.InjectHTTP(gatewaySpan, reqToOrder)
	fmt.Printf("[TRACER INJECTED] W3C traceparent header: '%s'\n", reqToOrder.Header.Get("traceparent"))

	// 3. Order Microservice extracts W3C header context
	ctxOrder, orderSpan := tracer.ExtractHTTP(context.Background(), reqToOrder)
	orderSpan.Operation = "OrderService_ProcessOrder"
	orderSpan.SetTag("db.type", "postgresql")

	// 4. Order Microservice calls Payment Microservice
	_, dbSpan := tracer.StartSpan(ctxOrder, "OrderService_InsertDB")
	time.Sleep(25 * time.Millisecond)
	dbSpan.Finish()

	orderSpan.Finish()
	gatewaySpan.Finish()

	// Output collected Distributed Trace Graph
	fmt.Println("\n--- Distributed Trace Telemetry Spans ---")
	for i, s := range tracer.Spans() {
		duration := s.EndTime.Sub(s.StartTime)
		fmt.Printf("Span #%d [%s]\n", i+1, s.Operation)
		fmt.Printf("  TraceID:      %s\n", s.TraceID)
		fmt.Printf("  SpanID:       %s\n", s.SpanID)
		fmt.Printf("  ParentSpanID: %s\n", s.ParentSpanID)
		fmt.Printf("  Duration:     %v\n", duration)
		fmt.Printf("  Tags:         %v\n\n", s.Tags)
	}
}
