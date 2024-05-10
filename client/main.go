package main

import (
	http_adapter "github.com/ternaryinvalid/safenet/client/internal/app/adapters/primary/http-adapter"
	os_signal_adapter "github.com/ternaryinvalid/safenet/client/internal/app/adapters/primary/os-signal-adapter"
	server_provider "github.com/ternaryinvalid/safenet/client/internal/app/adapters/secondary/providers/server-provider"
	"github.com/ternaryinvalid/safenet/client/internal/app/application"
	"github.com/ternaryinvalid/safenet/client/internal/pkg/config"
	"log"
)

func main() {
	config := config.New()

	serverProvider := server_provider.New(config.Adapters.Secondary.Providers.ServerProvider)

	app := application.New(config.Application, serverProvider)

	httpAdapter := http_adapter.New(config.Adapters.Primary.HttpAdapter, app)

	go httpAdapter.Start()

	osSignal := os_signal_adapter.New()

	go osSignal.Start()

	select {
	case err := <-httpAdapter.Notify():
		log.Println(err.Error(), "main")
	case err := <-osSignal.Notify():
		log.Println(err.Error(), "main")
	}
}
