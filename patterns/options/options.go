package options

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const (
	defaultTimeout = 1 * time.Second
)

func main() {
	fmt.Println("Options pattern")

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	u := NewContext(
		WithUsername("thanhn"),
		WithPassword("password1"),
		WithClient(client),
		WithTimeout(defaultTimeout),
	)

	fmt.Printf("%#v", u)
}

type Context struct {
	ctx      context.Context
	client   *http.Client
	username string
	password string
	timeout  time.Duration
}

// NewContext creates a new Context with the provided options
func NewContext(opts ...ContextOption) *Context {
	// Here we can do any initialization for all options, then the provided parameters can overwrite them.
	uCtx := Context{}
	for _, o := range opts {
		o(&uCtx)
	}
	return &uCtx
}

type ContextOption func(*Context)

// WithTimeout sets the context timeout
func WithTimeout(ttl time.Duration) ContextOption {
	return func(c *Context) {
		c.timeout = ttl
		var cancel func()
		c.ctx, cancel = context.WithTimeout(context.Background(), ttl)
		tick := time.NewTimer(ttl)
		go func() {
			select {
			case <-c.ctx.Done():
				cancel()
			case <-tick.C:
				cancel()
			}
		}()
	}
}

// WithClient set an *http.Client to be used by the context
func WithClient(client *http.Client) ContextOption {
	return func(c *Context) {
		c.client = client
	}
}

// WithUsername is an option that sets the username for remote server collector connections
func WithUsername(username string) ContextOption {
	return func(c *Context) {
		c.username = username
	}
}

// WithPassword is an option that sets the password for remote server collector connections
func WithPassword(password string) ContextOption {
	return func(c *Context) {
		c.password = password
	}
}

// FromContext sets a derived context
func FromContext(ctx context.Context) ContextOption {
	return func(c *Context) {
		c.ctx = context.WithValue(c.ctx, "context", ctx)
	}
}
