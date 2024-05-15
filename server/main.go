package main

import (
	"github.com/ternaryinvalid/safenet/server/internal/app/adapters/secondary/repositories/cache"
	"log"

	http_adapter "github.com/ternaryinvalid/safenet/server/internal/app/adapters/primary/http-adapter"
	os_signal_adapter "github.com/ternaryinvalid/safenet/server/internal/app/adapters/primary/os-signal-adapter"
	message_repository "github.com/ternaryinvalid/safenet/server/internal/app/adapters/secondary/repositories/message-repository"
	"github.com/ternaryinvalid/safenet/server/internal/app/application"
	"github.com/ternaryinvalid/safenet/server/internal/pkg/config"
)

func main() {
	config := config.New()

	messagesRepository := message_repository.New(config.Adapters.Secondary.Databases.Messages)

	cache := cache.New()

	app := application.New(config.Application, messagesRepository, cache)

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
