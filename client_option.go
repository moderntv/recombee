package recombee

import (
	"time"
)

// ClientOption for http client to be configured.
type ClientOption func(c *Client)

// WithRequestTimeout specifies default http request timeout.
func WithRequestTimeout(t time.Duration) ClientOption {
	return func(c *Client) {
		c.requestTimeout = t
	}
}

// WithMaxBatchSize specifies max batch size that is send by http client as single http request to recombee API.
func WithMaxBatchSize(max int) ClientOption {
	return func(c *Client) {
		c.maxBatchSize = max
	}
}

// WithDistinctRecomms when set to true, makes recommended items for certain user distinct among multiple recommendations.
func WithDistinctRecomms(v bool) ClientOption {
	return func(c *Client) {
		c.distinctRecomms = v
	}
}
