package sandbox

import (
	"context"
	"sync"
	"time"
)

const expiryMargin = 30 * time.Second

// Grant is a short-lived sandbox access token together with the URL it unlocks.
type Grant struct {
	URL       string
	Token     string
	ExpiresAt time.Time
}

// AccessProvider fetches a fresh Grant, typically from the sandbox machine access endpoints.
type AccessProvider func(ctx context.Context) (*Grant, error)

type accessCache struct {
	provider AccessProvider

	mu    sync.Mutex
	grant *Grant
}

func newAccessCache(provider AccessProvider) *accessCache {
	return &accessCache{provider: provider}
}

func (c *accessCache) usable(grant *Grant) bool {
	return grant.ExpiresAt.IsZero() || time.Until(grant.ExpiresAt) > expiryMargin
}

func (c *accessCache) get(ctx context.Context) (*Grant, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.grant != nil && c.usable(c.grant) {
		return c.grant, nil
	}

	grant, err := c.provider(ctx)
	if err != nil {
		return nil, err
	}

	c.grant = grant
	return grant, nil
}

func (c *accessCache) invalidate() {
	c.mu.Lock()
	c.grant = nil
	c.mu.Unlock()
}
