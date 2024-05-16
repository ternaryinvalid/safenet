package main

import (
	"log"

	http_adapter "github.com/ternaryinvalid/safenet/client/internal/app/adapters/primary/http-adapter"
	os_signal_adapter "github.com/ternaryinvalid/safenet/client/internal/app/adapters/primary/os-signal-adapter"
	"github.com/ternaryinvalid/safenet/client/internal/app/adapters/secondary/cache"
	server_provider "github.com/ternaryinvalid/safenet/client/internal/app/adapters/secondary/providers/server-provider"
	"github.com/ternaryinvalid/safenet/client/internal/app/application"
	"github.com/ternaryinvalid/safenet/client/internal/pkg/config"
)

func main() {
	config := config.New()

	serverProvider := server_provider.New(config.Adapters.Secondary.Providers.ServerProvider)

	cacheRepository := cache.New(config.Adapters.Secondary.Repositories.Cache)

	app := application.New(config.Application, serverProvider, cacheRepository)

	httpAdapter := http_adapter.New(config.Adapters.Primary.HttpAdapter, app)

	go httpAdapter.Start()

	osSignal := os_signal_adapter.New()

	go osSignal.Start()

	// Graceful shutdown
	select {
	case err := <-httpAdapter.Notify():
		log.Println(err.Error(), "main")
	case err := <-osSignal.Notify():
		log.Println(err.Error(), "main")
	}
}
