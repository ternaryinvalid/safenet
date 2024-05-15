package providerhelpers

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/ternaryinvalid/safenet/client/internal/app/domain/config"
)

func CreateRequest(ctx context.Context, client *resty.Client, endpoint config.Endpoint) *resty.Request {
	req := client.R()

	req.Method = endpoint.Method
	req.URL = endpoint.Path

	req.SetContext(ctx)

	return req
}
