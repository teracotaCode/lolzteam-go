package runtime

import "fmt"

// LolzteamError is the base error type for all library errors.
type LolzteamError struct {
	Message string
}

func (e *LolzteamError) Error() string {
	return fmt.Sprintf("lolzteam: %s", e.Message)
}

// HttpError represents an HTTP-level error response.
type HttpError struct {
	StatusCode int
	Body       string
	RetryAfter *float64
}

func (e *HttpError) Error() string {
	return fmt.Sprintf("HTTP %d: %s", e.StatusCode, e.Body)
}

// AuthError represents a 401 Unauthorized response.
type AuthError struct {
	HttpError
}

func (e *AuthError) Error() string {
	return fmt.Sprintf("authentication error: %s", e.HttpError.Error())
}

func (e *AuthError) Unwrap() error {
	return &e.HttpError
}

// ForbiddenError represents a 403 Forbidden response.
type ForbiddenError struct {
	HttpError
}

func (e *ForbiddenError) Error() string {
	return fmt.Sprintf("forbidden: %s", e.HttpError.Error())
}

func (e *ForbiddenError) Unwrap() error {
	return &e.HttpError
}

// NotFoundError represents a 404 Not Found response.
type NotFoundError struct {
	HttpError
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("not found: %s", e.HttpError.Error())
}

func (e *NotFoundError) Unwrap() error {
	return &e.HttpError
}

// RateLimitError represents a 429 Too Many Requests response.
type RateLimitError struct {
	HttpError
}

func (e *RateLimitError) Error() string {
	return fmt.Sprintf("rate limited: %s", e.HttpError.Error())
}

func (e *RateLimitError) Unwrap() error {
	return &e.HttpError
}

// ServerError represents a 5xx server error response.
type ServerError struct {
	HttpError
}

func (e *ServerError) Error() string {
	return fmt.Sprintf("server error: %s", e.HttpError.Error())
}

func (e *ServerError) Unwrap() error {
	return &e.HttpError
}

// ValidationError represents a 400/422 validation error response.
type ValidationError struct {
	HttpError
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error: %s", e.HttpError.Error())
}

func (e *ValidationError) Unwrap() error {
	return &e.HttpError
}

// NetworkError represents a network-level error (DNS, timeout, connection refused, etc.).
type NetworkError struct {
	Original  error
	Transient bool
}

func (e *NetworkError) Error() string {
	return fmt.Sprintf("network error: %v", e.Original)
}

func (e *NetworkError) Unwrap() error {
	return e.Original
}

// ConfigError represents a configuration validation error.
type ConfigError struct {
	Message string
}

func (e *ConfigError) Error() string {
	return fmt.Sprintf("config error: %s", e.Message)
}

// RetryExhaustedError is returned when all retry attempts have been exhausted.
type RetryExhaustedError struct {
	Attempts  int
	LastError error
}

func (e *RetryExhaustedError) Error() string {
	return fmt.Sprintf("retry exhausted after %d attempts: %v", e.Attempts, e.LastError)
}

func (e *RetryExhaustedError) Unwrap() error {
	return e.LastError
}

// isRetryable checks if an error should be retried.
func isRetryable(err error, retryStatuses map[int]bool) bool {
	var httpErr *HttpError
	if asHTTP(err, &httpErr) {
		return retryStatuses[httpErr.StatusCode]
	}
	var netErr *NetworkError
	if asNetwork(err, &netErr) {
		return netErr.Transient
	}
	return false
}

// asHTTP extracts an HttpError from err chain.
func asHTTP(err error, target **HttpError) bool {
	// Check direct HttpError
	var he *HttpError
	var ae *AuthError
	var fe *ForbiddenError
	var ne *NotFoundError
	var re *RateLimitError
	var se *ServerError
	var ve *ValidationError

	switch {
	case errorAs(err, &he):
		*target = he
		return true
	case errorAs(err, &ae):
		*target = &ae.HttpError
		return true
	case errorAs(err, &fe):
		*target = &fe.HttpError
		return true
	case errorAs(err, &ne):
		*target = &ne.HttpError
		return true
	case errorAs(err, &re):
		*target = &re.HttpError
		return true
	case errorAs(err, &se):
		*target = &se.HttpError
		return true
	case errorAs(err, &ve):
		*target = &ve.HttpError
		return true
	}
	return false
}

// asNetwork extracts a NetworkError from err chain.
func asNetwork(err error, target **NetworkError) bool {
	var ne *NetworkError
	if errorAs(err, &ne) {
		*target = ne
		return true
	}
	return false
}

// errorAs is a thin wrapper around type assertion in the error chain.
// We use errors.As from the standard library via an import-free approach.
func errorAs[T error](err error, target *T) bool {
	for err != nil {
		if e, ok := err.(T); ok {
			*target = e
			return true
		}
		u, ok := err.(interface{ Unwrap() error })
		if !ok {
			return false
		}
		err = u.Unwrap()
	}
	return false
}

// NewHttpError creates the appropriate typed error for an HTTP status code.
func NewHttpError(statusCode int, body string, retryAfter *float64) error {
	base := HttpError{
		StatusCode: statusCode,
		Body:       body,
		RetryAfter: retryAfter,
	}
	switch {
	case statusCode == 401:
		return &AuthError{HttpError: base}
	case statusCode == 403:
		return &ForbiddenError{HttpError: base}
	case statusCode == 404:
		return &NotFoundError{HttpError: base}
	case statusCode == 429:
		return &RateLimitError{HttpError: base}
	case statusCode == 400 || statusCode == 422:
		return &ValidationError{HttpError: base}
	case statusCode >= 500:
		return &ServerError{HttpError: base}
	default:
		return &base
	}
}
