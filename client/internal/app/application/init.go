package application

import "github.com/ternaryinvalid/safenet/client/internal/app/domain/config"

type Application struct {
	cfg            config.Application
	serverProvider serverProvider
}

type serverProvider interface {
	SendMessage(key []byte, message []byte) error
	GetMessages(key []byte) ([][]byte, error)
}

func New(cfg config.Application, serverProvider serverProvider) *Application {
	return &Application{
		cfg:            cfg,
		serverProvider: serverProvider,
	}
}
