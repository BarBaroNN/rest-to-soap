package transport

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

// Client is a custom HTTP client with logging
type Client struct {
	client *http.Client
	logger *zap.Logger
}

// NewClient creates a new HTTP client with the given timeout
func NewClient(timeout time.Duration, logger *zap.Logger) *Client {
	return &Client{
		client: &http.Client{
			Timeout: timeout,
		},
		logger: logger,
	}
}

// Do sends an HTTP request and returns the response
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	start := time.Now()
	resp, err := c.client.Do(req)
	if err != nil {
		c.logger.Error("HTTP request failed",
			zap.String("method", req.Method),
			zap.String("url", req.URL.String()),
			zap.Error(err),
		)
		return nil, err
	}

	c.logger.Info("HTTP request completed",
		zap.String("method", req.Method),
		zap.String("url", req.URL.String()),
		zap.Int("status", resp.StatusCode),
		zap.Duration("duration", time.Since(start)),
	)

	return resp, nil
}
