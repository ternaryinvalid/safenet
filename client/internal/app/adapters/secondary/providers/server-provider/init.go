package server_provider

import (
	"github.com/go-resty/resty/v2"
	"github.com/ternaryinvalid/safenet/client/internal/app/domain/config"
)

type ServerProvider struct {
	cfg    config.ServerProvider
	client *resty.Client
}

func New(cfg config.ServerProvider) *ServerProvider {
	client := resty.New().
		SetBaseURL(cfg.Host).
		SetRetryCount(3)

	return &ServerProvider{
		cfg:    cfg,
		client: client,
	}
}
