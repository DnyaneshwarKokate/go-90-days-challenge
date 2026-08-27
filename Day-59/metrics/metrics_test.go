package metrics

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMetricsCollectorAndMiddleware(t *testing.T) {
	col := NewMetricsCollector()

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/err" {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	wrapped := col.MetricsMiddleware(testHandler)

	// Call success endpoint
	req1 := httptest.NewRequest("GET", "/ok", nil)
	rr1 := httptest.NewRecorder()
	wrapped.ServeHTTP(rr1, req1)

	// Call error endpoint
	req2 := httptest.NewRequest("GET", "/err", nil)
	rr2 := httptest.NewRecorder()
	wrapped.ServeHTTP(rr2, req2)

	// Fetch /metrics output
	reqM := httptest.NewRequest("GET", "/metrics", nil)
	rrM := httptest.NewRecorder()
	col.Handler().ServeHTTP(rrM, reqM)

	output := rrM.Body.String()

	if !strings.Contains(output, "http_requests_total 2") {
		t.Errorf("Expected 2 total requests in metrics, got:\n%s", output)
	}

	if !strings.Contains(output, "http_errors_total 1") {
		t.Errorf("Expected 1 total error in metrics, got:\n%s", output)
	}
}
