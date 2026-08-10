package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"

	"day-35/domain"
)

type ResilientHTTPClient struct {
	client  *http.Client
	config  domain.ClientConfig
	logger  domain.Logger
}

func NewResilientHTTPClient(cfg domain.ClientConfig, logger domain.Logger) *ResilientHTTPClient {
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   10,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	httpClient := &http.Client{
		Transport: transport,
		Timeout:   cfg.Timeout,
	}

	return &ResilientHTTPClient{
		client:  httpClient,
		config:  cfg,
		logger:  logger,
	}
}

func (c *ResilientHTTPClient) GetUser(ctx context.Context, userID int) (*domain.ExternalUser, error) {
	url := fmt.Sprintf("%s/users/%d", c.config.BaseURL, userID)

	var lastErr error
	backoff := c.config.InitialBackoff

	for attempt := 1; attempt <= c.config.MaxRetries; attempt++ {
		c.logger.Info(ctx, "HTTP GET request attempt", "url", url, "attempt", attempt)

		reqCtx, cancel := context.WithTimeout(ctx, c.config.Timeout)
		req, err := http.NewRequestWithContext(reqCtx, "GET", url, nil)
		if err != nil {
			cancel()
			return nil, fmt.Errorf("failed to create HTTP request: %w", err)
		}

		if reqID, ok := ctx.Value(domain.RequestIDKey).(string); ok && reqID != "" {
			req.Header.Set("X-Correlation-ID", reqID)
		}
		req.Header.Set("Accept", "application/json")

		resp, err := c.client.Do(req)
		if err != nil {
			cancel()
			c.logger.Warn(ctx, "HTTP request error, retrying...", "attempt", attempt, "error", err)
			lastErr = err
			time.Sleep(backoff)
			backoff *= 2
			continue
		}

		if resp.StatusCode == http.StatusOK {
			defer resp.Body.Close()
			var user domain.ExternalUser
			if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
				cancel()
				return nil, fmt.Errorf("failed to decode response JSON: %w", err)
			}
			cancel()
			c.logger.Info(ctx, "HTTP GET request succeeded", "url", url, "user_id", user.ID)
			return &user, nil
		}

		resp.Body.Close()
		cancel()

		if resp.StatusCode >= 500 {
			c.logger.Warn(ctx, "HTTP server error status code, retrying...", "status", resp.StatusCode, "attempt", attempt)
			lastErr = fmt.Errorf("HTTP error status: %d", resp.StatusCode)
			time.Sleep(backoff)
			backoff *= 2
			continue
		}

		return nil, fmt.Errorf("external API returned non-retriable status code: %d", resp.StatusCode)
	}

	return nil, fmt.Errorf("%w: %v", domain.ErrMaxRetriesExceeded, lastErr)
}
