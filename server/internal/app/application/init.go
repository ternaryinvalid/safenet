package application

import (
	"context"

	"github.com/ternaryinvalid/safenet/server/internal/app/domain/config"
	"github.com/ternaryinvalid/safenet/server/internal/app/domain/entity"
)

type Application struct {
	cfg               config.Application
	messageRepository messageRepository
	cache             cache
}

type messageRepository interface {
	SaveMessage(ctx context.Context, request entity.SaveMessageDTO) (id int64, err error)
	GetMessages(ctx context.Context, request entity.GetMessagesDTO) (messages []entity.Message, err error)
}

type cache interface {
	GetShared() (string, error)
	SaveShared(sharedSecret string)
	SetSecret(localPublicKey, localPrivateKey string)
	IsEmpty() bool
	GetSecret() string
}

func New(cfg config.Application, messageRepository messageRepository, cache cache) *Application {
	return &Application{
		cfg:               cfg,
		messageRepository: messageRepository,
		cache:             cache,
	}
}
