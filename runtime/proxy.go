package runtime

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"

	"golang.org/x/net/proxy"
)

// ProxyConfig holds proxy configuration.
type ProxyConfig struct {
	URL string
}

// Validate checks that the proxy URL is valid.
func (p *ProxyConfig) Validate() error {
	if p.URL == "" {
		return &ConfigError{Message: "proxy URL is required"}
	}

	parsed, err := url.Parse(p.URL)
	if err != nil {
		return &ConfigError{Message: fmt.Sprintf("invalid proxy URL: %v", err)}
	}

	switch parsed.Scheme {
	case "http", "https", "socks5":
		// valid
	default:
		return &ConfigError{Message: fmt.Sprintf("unsupported proxy scheme %q: must be http, https, or socks5", parsed.Scheme)}
	}

	if parsed.Host == "" {
		return &ConfigError{Message: "proxy URL must have a host"}
	}

	return nil
}

// Transport creates an *http.Transport configured to use this proxy.
func (p *ProxyConfig) Transport() (*http.Transport, error) {
	if err := p.Validate(); err != nil {
		return nil, err
	}

	parsed, _ := url.Parse(p.URL)

	if parsed.Scheme == "socks5" {
		var auth *proxy.Auth
		if parsed.User != nil {
			password, _ := parsed.User.Password()
			auth = &proxy.Auth{
				User:     parsed.User.Username(),
				Password: password,
			}
		}

		dialer, err := proxy.SOCKS5("tcp", parsed.Host, auth, proxy.Direct)
		if err != nil {
			return nil, &ConfigError{Message: fmt.Sprintf("failed to create SOCKS5 dialer: %v", err)}
		}

		contextDialer, ok := dialer.(proxy.ContextDialer)
		if !ok {
			// Wrap the dialer to support context
			contextDialer = dialerWrapper{dialer}
		}

		return &http.Transport{
			DialContext: contextDialer.DialContext,
		}, nil
	}

	// HTTP/HTTPS proxy
	return &http.Transport{
		Proxy: http.ProxyURL(parsed),
	}, nil
}

// dialerWrapper wraps a proxy.Dialer to implement proxy.ContextDialer.
type dialerWrapper struct {
	dialer proxy.Dialer
}

func (d dialerWrapper) DialContext(_ context.Context, network, addr string) (net.Conn, error) {
	return d.dialer.Dial(network, addr)
}

func (d dialerWrapper) Dial(network, addr string) (net.Conn, error) {
	return d.dialer.Dial(network, addr)
}
