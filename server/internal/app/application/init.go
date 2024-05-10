package application

import "github.com/ternaryinvalid/safenet/server/internal/app/domain/config"

type Application struct {
	cfg               config.Application
	messageRepository messageRepository
}

type messageRepository interface {
	SaveMessage(message string) error
	GetMessages(key []byte) ([][]byte, error)
}

func New(cfg config.Application, messageRepository messageRepository) *Application {
	return &Application{
		cfg:               cfg,
		messageRepository: messageRepository,
	}
}
