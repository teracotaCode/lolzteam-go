package runtime

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"strings"
	"time"
)

const (
	defaultTimeout = 60 * time.Second
	defaultBaseURL = "https://prod-api.lzt.market"
)

// ClientConfig holds all configuration for the HTTP client.
type ClientConfig struct {
	Token     string
	BaseURL   string
	Proxy     *ProxyConfig
	Retry     *RetryConfig
	RateLimit *RateLimitConfig
	Timeout   time.Duration
}

// RequestOptions holds per-request options.
type RequestOptions struct {
	Params         map[string]interface{}
	JSON           map[string]interface{}
	JSONBody       interface{} // Raw JSON body (for array bodies like batch); takes precedence over JSON if set
	Files          map[string]FileUpload
	Search         bool
	ForceMultipart bool // Send JSON fields as multipart form fields even without files
}

// FileUpload represents a file to upload.
type FileUpload struct {
	Filename string
	Data     []byte
}

// Requester is the interface implemented by Client for making API requests.
type Requester interface {
	Request(ctx context.Context, method, path string, opts RequestOptions) (json.RawMessage, error)
}

// Client is the HTTP client for the Lolzteam API.
type Client struct {
	httpClient  *http.Client
	baseURL     string
	token       string
	retryConfig *RetryConfig
	rateLimiter *RateLimiter
}

// NewClient creates a new API client.
func NewClient(cfg ClientConfig) (*Client, error) {
	if cfg.Token == "" {
		return nil, &ConfigError{Message: "token is required"}
	}

	baseURL := cfg.BaseURL
	if baseURL == "" {
		baseURL = defaultBaseURL
	}
	baseURL = strings.TrimRight(baseURL, "/")

	timeout := cfg.Timeout
	if timeout == 0 {
		timeout = defaultTimeout
	}

	var transport http.RoundTripper
	if cfg.Proxy != nil {
		t, err := cfg.Proxy.Transport()
		if err != nil {
			return nil, err
		}
		transport = t
	}

	httpClient := &http.Client{
		Timeout:   timeout,
		Transport: transport,
	}

	retryConfig := cfg.Retry
	if retryConfig == nil {
		retryConfig = DefaultRetryConfig()
	}

	rateLimiter := NewRateLimiter(cfg.RateLimit)

	return &Client{
		httpClient:  httpClient,
		baseURL:     baseURL,
		token:       cfg.Token,
		retryConfig: retryConfig,
		rateLimiter: rateLimiter,
	}, nil
}

// Request makes an HTTP request to the API and returns the raw JSON response.
func (c *Client) Request(ctx context.Context, method, path string, opts RequestOptions) (json.RawMessage, error) {
	// Wait for rate limit
	if err := c.rateLimiter.Wait(ctx, opts.Search); err != nil {
		return nil, err
	}

	var result json.RawMessage

	err := ExecuteWithRetry(ctx, func() error {
		res, err := c.doRequest(ctx, method, path, opts)
		if err != nil {
			return err
		}
		result = res
		return nil
	}, c.retryConfig)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// doRequest performs a single HTTP request.
func (c *Client) doRequest(ctx context.Context, method, path string, opts RequestOptions) (json.RawMessage, error) {
	// Build URL
	fullURL := c.baseURL + path

	// Add query params
	if len(opts.Params) > 0 {
		qs := BuildQueryString(opts.Params)
		if qs != "" {
			if strings.Contains(fullURL, "?") {
				fullURL += "&" + qs
			} else {
				fullURL += "?" + qs
			}
		}
	}

	// Build body
	var body io.Reader
	var contentType string

	if len(opts.Files) > 0 || opts.ForceMultipart {
		// Multipart form
		var buf bytes.Buffer
		writer := multipart.NewWriter(&buf)

		// Add JSON fields as form fields
		if opts.JSON != nil {
			for k, v := range opts.JSON {
				if v == nil {
					continue
				}
				val := formatFormValue(v)
				if err := writer.WriteField(k, val); err != nil {
					return nil, &NetworkError{Original: fmt.Errorf("failed to write form field: %w", err), Transient: false}
				}
			}
		}

		// Add files
		for fieldName, file := range opts.Files {
			part, err := writer.CreateFormFile(fieldName, file.Filename)
			if err != nil {
				return nil, &NetworkError{Original: fmt.Errorf("failed to create form file: %w", err), Transient: false}
			}
			if _, err := part.Write(file.Data); err != nil {
				return nil, &NetworkError{Original: fmt.Errorf("failed to write file data: %w", err), Transient: false}
			}
		}

		if err := writer.Close(); err != nil {
			return nil, &NetworkError{Original: fmt.Errorf("failed to close multipart writer: %w", err), Transient: false}
		}

		body = &buf
		contentType = writer.FormDataContentType()
	} else if opts.JSONBody != nil {
		// Raw JSON body (e.g. array body for batch endpoints)
		jsonData, err := json.Marshal(opts.JSONBody)
		if err != nil {
			return nil, &NetworkError{Original: fmt.Errorf("failed to marshal JSON: %w", err), Transient: false}
		}
		body = bytes.NewReader(jsonData)
		contentType = "application/json"
	} else if len(opts.JSON) > 0 {
		// JSON body — strip nil values
		cleaned := stripNilValues(opts.JSON)
		if len(cleaned) > 0 {
			jsonData, err := json.Marshal(cleaned)
			if err != nil {
				return nil, &NetworkError{Original: fmt.Errorf("failed to marshal JSON: %w", err), Transient: false}
			}
			body = bytes.NewReader(jsonData)
			contentType = "application/json"
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, fullURL, body)
	if err != nil {
		return nil, &NetworkError{Original: err, Transient: false}
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer "+c.token)
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	req.Header.Set("Accept", "application/json")

	// Execute
	resp, err := c.httpClient.Do(req)
	if err != nil {
		// Classify network errors
		transient := isTransientNetworkError(err)
		return nil, &NetworkError{Original: err, Transient: transient}
	}
	defer resp.Body.Close()

	// Read body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, &NetworkError{Original: fmt.Errorf("failed to read response body: %w", err), Transient: true}
	}

	// Check for errors
	if resp.StatusCode >= 400 {
		retryAfter := ParseRetryAfter(resp.Header)
		return nil, NewHttpError(resp.StatusCode, string(respBody), retryAfter)
	}

	// Return raw JSON bytes
	if len(respBody) == 0 {
		return json.RawMessage("null"), nil
	}
	return json.RawMessage(respBody), nil
}

// stripNilValues removes nil values from a map recursively.
func stripNilValues(m map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{}, len(m))
	for k, v := range m {
		if v == nil {
			continue
		}
		if sub, ok := v.(map[string]interface{}); ok {
			cleaned := stripNilValues(sub)
			if len(cleaned) > 0 {
				result[k] = cleaned
			}
			continue
		}
		result[k] = v
	}
	return result
}

// formatFormValue converts a value to a string suitable for form fields.
func formatFormValue(v interface{}) string {
	switch val := v.(type) {
	case bool:
		if val {
			return "true"
		}
		return "false"
	case string:
		return val
	default:
		return fmt.Sprintf("%v", v)
	}
}

// isTransientNetworkError checks if a network error is transient.
func isTransientNetworkError(err error) bool {
	if err == nil {
		return false
	}
	// DNS errors, connection refused, timeouts
	if netErr, ok := err.(net.Error); ok {
		return netErr.Timeout()
	}
	// Check wrapped errors
	if opErr, ok := err.(*net.OpError); ok {
		return opErr.Temporary()
	}
	return false
}
