package recombee

import (
	"time"
)

// Http client options to be configured.
type ClientOption func(c *Client)

// Specifies default http request timeout.
func WithRequestTimeout(t time.Duration) ClientOption {
	return func(c *Client) {
		c.requestTimeout = t
	}
}

// Specifies max batch size that is send into as single http request to recombee API.
func WithMaxBatchSize(max int) ClientOption {
	return func(c *Client) {
		c.maxBatchSize = max
	}
}

func WithDistinctRecomms(v bool) ClientOption {
	return func(c *Client) {
		c.distinctRecomms = v
	}
}
