package comhttp

import (
	"context"
	"net/http"
	"time"

	"github.com/yunerou/oauth2-client/singleton"
)

// NewHttpClient ..if timeout == nil -> use default value
func NewHttpClient(ctx context.Context, timeout *time.Duration) *http.Client {
	var timeoutValue time.Duration
	if timeout != nil {
		timeoutValue = *timeout
	} else {
		timeoutValue = time.Duration(singleton.GetViper().GetDuration("DEFAULT_HTTP_CLIENT_TIMEOUT"))
	}

	return &http.Client{
		Transport: LoggingRoundTripper{Ctx: ctx, Proxied: http.DefaultTransport},
		Timeout:   timeoutValue,
	}
}
