package runtime

import (
	"context"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// RetryConfig controls the retry behavior.
type RetryConfig struct {
	MaxRetries    int
	BaseDelay     time.Duration
	MaxDelay      time.Duration
	RetryStatuses map[int]bool
	OnRetry       func(attempt int, delay time.Duration, err error)
}

// DefaultRetryConfig returns sensible defaults.
func DefaultRetryConfig() *RetryConfig {
	return &RetryConfig{
		MaxRetries: 3,
		BaseDelay:  1 * time.Second,
		MaxDelay:   30 * time.Second,
		RetryStatuses: map[int]bool{
			429: true,
			502: true,
			503: true,
			504: true,
		},
	}
}

// fillDefaults fills zero-valued fields with defaults.
func (c *RetryConfig) fillDefaults() {
	if c.MaxRetries == 0 {
		c.MaxRetries = 3
	}
	if c.BaseDelay == 0 {
		c.BaseDelay = 1 * time.Second
	}
	if c.MaxDelay == 0 {
		c.MaxDelay = 30 * time.Second
	}
	if c.RetryStatuses == nil {
		c.RetryStatuses = map[int]bool{429: true, 502: true, 503: true, 504: true}
	}
}

// ExecuteWithRetry executes fn with retries according to config.
func ExecuteWithRetry(ctx context.Context, fn func() error, config *RetryConfig) error {
	if config == nil {
		config = DefaultRetryConfig()
	} else {
		// Copy to avoid mutating caller's config
		copy := *config
		copy.fillDefaults()
		config = &copy
	}

	var lastErr error
	for attempt := 0; attempt <= config.MaxRetries; attempt++ {
		if err := ctx.Err(); err != nil {
			return &NetworkError{Original: err, Transient: false}
		}

		lastErr = fn()
		if lastErr == nil {
			return nil
		}

		// Don't retry on last attempt
		if attempt == config.MaxRetries {
			break
		}

		if !isRetryable(lastErr, config.RetryStatuses) {
			return lastErr
		}

		delay := computeDelay(attempt, config, lastErr)

		if config.OnRetry != nil {
			config.OnRetry(attempt+1, delay, lastErr)
		}

		timer := time.NewTimer(delay)
		select {
		case <-ctx.Done():
			timer.Stop()
			return &NetworkError{Original: ctx.Err(), Transient: false}
		case <-timer.C:
		}
	}

	return &RetryExhaustedError{
		Attempts:  config.MaxRetries + 1,
		LastError: lastErr,
	}
}

// computeDelay calculates the delay for a given attempt, accounting for Retry-After.
func computeDelay(attempt int, config *RetryConfig, err error) time.Duration {
	// Check for Retry-After from HttpError
	var httpErr *HttpError
	if asHTTP(err, &httpErr) && httpErr.RetryAfter != nil {
		delay := time.Duration(*httpErr.RetryAfter * float64(time.Second))
		if delay > config.MaxDelay {
			delay = config.MaxDelay
		}
		if delay > 0 {
			return delay
		}
	}

	// Exponential backoff with full jitter
	base := float64(config.BaseDelay) * math.Pow(2, float64(attempt))
	if base > float64(config.MaxDelay) {
		base = float64(config.MaxDelay)
	}
	// Full jitter: random value in [0, base]
	jittered := time.Duration(rand.Float64() * base)
	if jittered < time.Millisecond {
		jittered = time.Millisecond
	}
	return jittered
}

// ParseRetryAfter parses the Retry-After header value.
// It handles both seconds (integer/float) and HTTP-date formats.
func ParseRetryAfter(header http.Header) *float64 {
	val := header.Get("Retry-After")
	if val == "" {
		return nil
	}
	val = strings.TrimSpace(val)

	// Try numeric (seconds)
	if seconds, err := strconv.ParseFloat(val, 64); err == nil {
		return &seconds
	}

	// Try HTTP-date
	formats := []string{
		time.RFC1123,
		time.RFC1123Z,
		time.RFC850,
		time.ANSIC,
	}
	for _, format := range formats {
		if t, err := time.Parse(format, val); err == nil {
			seconds := time.Until(t).Seconds()
			if seconds < 0 {
				seconds = 0
			}
			return &seconds
		}
	}

	return nil
}
