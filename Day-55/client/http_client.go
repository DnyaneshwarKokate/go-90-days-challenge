package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"day55/circuitbreaker"
)

type ResilientHTTPClient struct {
	client *http.Client
	cb     *circuitbreaker.CircuitBreaker
	retries int
}

func NewResilientHTTPClient(timeout time.Duration, retries int, cbThreshold int, cbTimeout time.Duration) *ResilientHTTPClient {
	return &ResilientHTTPClient{
		client: &http.Client{Timeout: timeout},
		cb:     circuitbreaker.NewCircuitBreaker(cbThreshold, cbTimeout),
		retries: retries,
	}
}

func (c *ResilientHTTPClient) GetWithRetry(ctx context.Context, targetURL string) ([]byte, error) {
	var body []byte

	err := c.cb.Execute(func() error {
		var lastErr error
		for attempt := 0; attempt <= c.retries; attempt++ {
			if attempt > 0 {
				time.Sleep(time.Duration(attempt*50) * time.Millisecond)
			}

			req, err := http.NewRequestWithContext(ctx, http.MethodGet, targetURL, nil)
			if err != nil {
				return err
			}

			resp, err := c.client.Do(req)
			if err != nil {
				lastErr = err
				continue
			}

			if resp.StatusCode >= 500 {
				_ = resp.Body.Close()
				lastErr = fmt.Errorf("server error status: %d", resp.StatusCode)
				continue
			}

			defer resp.Body.Close()
			body, err = io.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			return nil
		}
		return lastErr
	})

	return body, err
}

func (c *ResilientHTTPClient) CircuitState() circuitbreaker.State {
	return c.cb.GetState()
}
