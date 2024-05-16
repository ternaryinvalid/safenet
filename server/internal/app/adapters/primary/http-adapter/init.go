package http_adapter

import (
	"net/http"
	"time"

	api_controller "github.com/ternaryinvalid/safenet/server/internal/app/adapters/primary/http-adapter/api-controller"
	"github.com/ternaryinvalid/safenet/server/internal/app/adapters/primary/http-adapter/router"
	"github.com/ternaryinvalid/safenet/server/internal/app/application"
	"github.com/ternaryinvalid/safenet/server/internal/app/domain/config"
)

const (
	_defaultReadTimeout       = 5 * time.Second
	_defaultWriteTimeout      = 500 * time.Second
	_defaultReadHeaderTimeout = 5 * time.Second
	_defaultShutdownTimeout   = 3 * time.Second
)

type HttpAdapter struct {
	config          config.HttpAdapter
	server          *http.Server
	router          http.Handler
	shutdownTimeout time.Duration
	notify          chan error
}

func New(config config.HttpAdapter, app *application.Application) *HttpAdapter {
	r := router.New()

	controller := api_controller.New(app)

	r.AppendRoutesV1(controller)

	router := r.Router()

	httpServer := &http.Server{
		Handler:           router,
		ReadTimeout:       _defaultReadTimeout,
		WriteTimeout:      _defaultWriteTimeout,
		ReadHeaderTimeout: _defaultReadHeaderTimeout,
		Addr:              config.Server.Port,
	}

	return &HttpAdapter{
		config:          config,
		server:          httpServer,
		router:          router,
		shutdownTimeout: _defaultShutdownTimeout,
		notify:          make(chan error, 1),
	}
}
