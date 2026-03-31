package runtime

import (
	"context"
	"sync"
	"time"
)

// RateLimitConfig configures the rate limiter.
type RateLimitConfig struct {
	RequestsPerMinute       int
	SearchRequestsPerMinute int
}

// DefaultRateLimitConfig returns sensible defaults.
func DefaultRateLimitConfig() *RateLimitConfig {
	return &RateLimitConfig{
		RequestsPerMinute:       300,
		SearchRequestsPerMinute: 20,
	}
}

// tokenBucket implements a token bucket rate limiter with smooth refill.
type tokenBucket struct {
	mu         sync.Mutex
	tokens     float64
	maxTokens  float64
	refillRate float64 // tokens per second
	lastRefill time.Time
}

func newTokenBucket(requestsPerMinute int) *tokenBucket {
	rate := float64(requestsPerMinute) / 60.0
	return &tokenBucket{
		tokens:     float64(requestsPerMinute),
		maxTokens:  float64(requestsPerMinute),
		refillRate: rate,
		lastRefill: time.Now(),
	}
}

func (b *tokenBucket) refill() {
	now := time.Now()
	elapsed := now.Sub(b.lastRefill).Seconds()
	b.tokens += elapsed * b.refillRate
	if b.tokens > b.maxTokens {
		b.tokens = b.maxTokens
	}
	b.lastRefill = now
}

// timeUntilToken returns the duration until a token is available.
func (b *tokenBucket) timeUntilToken() time.Duration {
	if b.tokens >= 1.0 {
		return 0
	}
	needed := 1.0 - b.tokens
	return time.Duration(needed / b.refillRate * float64(time.Second))
}

func (b *tokenBucket) take() bool {
	if b.tokens >= 1.0 {
		b.tokens--
		return true
	}
	return false
}

// RateLimiter manages rate limits for both regular and search requests.
type RateLimiter struct {
	general *tokenBucket
	search  *tokenBucket
}

// NewRateLimiter creates a new rate limiter with the given config.
func NewRateLimiter(cfg *RateLimitConfig) *RateLimiter {
	if cfg == nil {
		cfg = DefaultRateLimitConfig()
	}
	rpm := cfg.RequestsPerMinute
	if rpm <= 0 {
		rpm = 300
	}
	srpm := cfg.SearchRequestsPerMinute
	if srpm <= 0 {
		srpm = 20
	}
	return &RateLimiter{
		general: newTokenBucket(rpm),
		search:  newTokenBucket(srpm),
	}
}

// Wait blocks until a token is available or the context is cancelled.
// If isSearch is true, both the search bucket and general bucket are checked.
func (r *RateLimiter) Wait(ctx context.Context, isSearch bool) error {
	for {
		if err := ctx.Err(); err != nil {
			return &NetworkError{Original: err, Transient: false}
		}

		var waitTime time.Duration

		// Lock general bucket
		r.general.mu.Lock()
		r.general.refill()
		generalWait := r.general.timeUntilToken()
		r.general.mu.Unlock()

		waitTime = generalWait

		if isSearch {
			r.search.mu.Lock()
			r.search.refill()
			searchWait := r.search.timeUntilToken()
			r.search.mu.Unlock()

			if searchWait > waitTime {
				waitTime = searchWait
			}
		}

		if waitTime <= 0 {
			// Try to take tokens
			r.general.mu.Lock()
			r.general.refill()
			gotGeneral := r.general.take()
			r.general.mu.Unlock()

			if !gotGeneral {
				continue
			}

			if isSearch {
				r.search.mu.Lock()
				r.search.refill()
				gotSearch := r.search.take()
				r.search.mu.Unlock()

				if !gotSearch {
					// Return general token
					r.general.mu.Lock()
					r.general.tokens++
					r.general.mu.Unlock()
					continue
				}
			}

			return nil
		}

		// Wait for the required duration or context cancellation
		timer := time.NewTimer(waitTime)
		select {
		case <-ctx.Done():
			timer.Stop()
			return &NetworkError{Original: ctx.Err(), Transient: false}
		case <-timer.C:
			// Loop back to try again
		}
	}
}
