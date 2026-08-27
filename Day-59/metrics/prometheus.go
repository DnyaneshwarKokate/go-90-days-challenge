package metrics

import (
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

type MetricsCollector struct {
	mu           sync.RWMutex
	requestCount int64
	errorCount   int64
	latencies    []time.Duration
}

func NewMetricsCollector() *MetricsCollector {
	return &MetricsCollector{
		latencies: make([]time.Duration, 0),
	}
}

func (m *MetricsCollector) IncRequest() {
	atomic.AddInt64(&m.requestCount, 1)
}

func (m *MetricsCollector) IncError() {
	atomic.AddInt64(&m.errorCount, 1)
}

func (m *MetricsCollector) ObserveLatency(d time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.latencies = append(m.latencies, d)
}

func (m *MetricsCollector) MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		m.IncRequest()

		rw := &responseWriterInterceptor{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(rw, r)

		duration := time.Since(start)
		m.ObserveLatency(duration)

		if rw.statusCode >= 400 {
			m.IncError()
		}
	})
}

type responseWriterInterceptor struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriterInterceptor) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (m *MetricsCollector) Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m.mu.RLock()
		reqs := atomic.LoadInt64(&m.requestCount)
		errs := atomic.LoadInt64(&m.errorCount)
		count := len(m.latencies)
		var total time.Duration
		for _, l := range m.latencies {
			total += l
		}
		m.mu.RUnlock()

		avgMs := float64(0)
		if count > 0 {
			avgMs = float64(total.Milliseconds()) / float64(count)
		}

		w.Header().Set("Content-Type", "text/plain; version=0.0.4")
		_, _ = fmt.Fprintf(w, "# HELP http_requests_total Total HTTP requests handled.\n")
		_, _ = fmt.Fprintf(w, "# TYPE http_requests_total counter\n")
		_, _ = fmt.Fprintf(w, "http_requests_total %d\n", reqs)

		_, _ = fmt.Fprintf(w, "# HELP http_errors_total Total HTTP errors encountered.\n")
		_, _ = fmt.Fprintf(w, "# TYPE http_errors_total counter\n")
		_, _ = fmt.Fprintf(w, "http_errors_total %d\n", errs)

		_, _ = fmt.Fprintf(w, "# HELP http_request_duration_avg_milliseconds Average HTTP latency.\n")
		_, _ = fmt.Fprintf(w, "# TYPE http_request_duration_avg_milliseconds gauge\n")
		_, _ = fmt.Fprintf(w, "http_request_duration_avg_milliseconds %.2f\n", avgMs)
	}
}
