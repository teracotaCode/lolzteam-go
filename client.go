// Package lolzteam provides a Go client for the Lolzteam Forum and Market APIs.
package lolzteam

import (
	"context"
	"encoding/json"
	"time"

	"github.com/teracotaCode/lolzteam-go/forum"
	"github.com/teracotaCode/lolzteam-go/market"
	"github.com/teracotaCode/lolzteam-go/runtime"
)

const (
	forumBaseURL  = "https://prod-api.lolz.live"
	marketBaseURL = "https://prod-api.lzt.market"
)

// ClientConfig holds configuration for constructing clients.
type ClientConfig struct {
	Token     string
	BaseURL   string
	Proxy     *runtime.ProxyConfig
	Retry     *runtime.RetryConfig
	RateLimit *runtime.RateLimitConfig
	Timeout   time.Duration
}

// Option is a functional option for configuring clients.
type Option func(*ClientConfig)

// WithProxy sets the proxy configuration.
func WithProxy(url string) Option {
	return func(cfg *ClientConfig) {
		cfg.Proxy = &runtime.ProxyConfig{URL: url}
	}
}

// WithRetry sets the retry configuration.
func WithRetry(retryCfg runtime.RetryConfig) Option {
	return func(cfg *ClientConfig) {
		cfg.Retry = &retryCfg
	}
}

// WithRateLimit sets the rate limit configuration.
func WithRateLimit(rlCfg runtime.RateLimitConfig) Option {
	return func(cfg *ClientConfig) {
		cfg.RateLimit = &rlCfg
	}
}

// WithTimeout sets the HTTP client timeout.
func WithTimeout(d time.Duration) Option {
	return func(cfg *ClientConfig) {
		cfg.Timeout = d
	}
}

// WithBaseURL overrides the base URL.
func WithBaseURL(url string) Option {
	return func(cfg *ClientConfig) {
		cfg.BaseURL = url
	}
}

// ForumClient wraps the generated Forum API client.
type ForumClient struct {
	*forum.ForumClient
}

// MarketClient wraps the generated Market API client.
type MarketClient struct {
	*market.MarketClient
}

// NewForumClient creates a new Forum API client.
func NewForumClient(token string, opts ...Option) (*ForumClient, error) {
	cfg := &ClientConfig{
		Token:   token,
		BaseURL: forumBaseURL,
	}
	for _, opt := range opts {
		opt(cfg)
	}

	c, err := runtime.NewClient(runtime.ClientConfig{
		Token:     cfg.Token,
		BaseURL:   cfg.BaseURL,
		Proxy:     cfg.Proxy,
		Retry:     cfg.Retry,
		RateLimit: cfg.RateLimit,
		Timeout:   cfg.Timeout,
	})
	if err != nil {
		return nil, err
	}

	return &ForumClient{ForumClient: forum.New(&forumAdapter{c: c})}, nil
}

// NewMarketClient creates a new Market API client.
func NewMarketClient(token string, opts ...Option) (*MarketClient, error) {
	cfg := &ClientConfig{
		Token:   token,
		BaseURL: marketBaseURL,
	}
	for _, opt := range opts {
		opt(cfg)
	}

	c, err := runtime.NewClient(runtime.ClientConfig{
		Token:     cfg.Token,
		BaseURL:   cfg.BaseURL,
		Proxy:     cfg.Proxy,
		Retry:     cfg.Retry,
		RateLimit: cfg.RateLimit,
		Timeout:   cfg.Timeout,
	})
	if err != nil {
		return nil, err
	}

	return &MarketClient{MarketClient: market.New(&marketAdapter{c: c})}, nil
}

// forumAdapter adapts runtime.Client to forum.Requester.
type forumAdapter struct {
	c *runtime.Client
}

func (a *forumAdapter) Request(ctx context.Context, method, path string, opts forum.RequestOptions) (json.RawMessage, error) {
	rtOpts := runtime.RequestOptions{
		Params: opts.Params,
		JSON:   opts.JSON,
		Search: opts.Search,
	}
	if len(opts.Files) > 0 {
		rtOpts.Files = make(map[string]runtime.FileUpload, len(opts.Files))
		for k, v := range opts.Files {
			rtOpts.Files[k] = runtime.FileUpload{
				Filename: v.Filename,
				Data:     v.Data,
			}
		}
	}
	return a.c.Request(ctx, method, path, rtOpts)
}

// marketAdapter adapts runtime.Client to market.Requester.
type marketAdapter struct {
	c *runtime.Client
}

func (a *marketAdapter) Request(ctx context.Context, method, path string, opts market.RequestOptions) (json.RawMessage, error) {
	rtOpts := runtime.RequestOptions{
		Params: opts.Params,
		JSON:   opts.JSON,
		Search: opts.Search,
	}
	if len(opts.Files) > 0 {
		rtOpts.Files = make(map[string]runtime.FileUpload, len(opts.Files))
		for k, v := range opts.Files {
			rtOpts.Files[k] = runtime.FileUpload{
				Filename: v.Filename,
				Data:     v.Data,
			}
		}
	}
	return a.c.Request(ctx, method, path, rtOpts)
}
