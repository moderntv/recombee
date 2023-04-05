package recombee

import (
	"time"
)

// ClientOption for http client to be configured.
type ClientOption func(c *Client)

// WithRequestTimeout sets request timeout.
func WithRequestTimeout(t time.Duration) ClientOption {
	return func(c *Client) {
		c.requestTimeout = t
	}
}

// WithMaxBatchSize sets max batch size that is sent to the API in a single request.
func WithMaxBatchSize(max int) ClientOption {
	return func(c *Client) {
		c.maxBatchSize = max
	}
}

// WithDistinctRecomms makes all the recommended items for a certain user distinct among multiple recommendation
// requests in the batch.
func WithDistinctRecomms(v bool) ClientOption {
	return func(c *Client) {
		c.distinctRecomms = v
	}
}
