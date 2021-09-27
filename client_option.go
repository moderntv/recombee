package recombee

import (
	"time"
)

type ClientOption func(c *Client)

func WithRequestTimeout(t time.Duration) ClientOption {
	return func(c *Client) {
		c.requestTimeout = t
	}
}

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
