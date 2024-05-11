package application

import (
	"context"

	"github.com/ternaryinvalid/safenet/server/internal/app/domain/config"
	"github.com/ternaryinvalid/safenet/server/internal/app/domain/entity"
)

type Application struct {
	cfg               config.Application
	messageRepository messageRepository
}

type messageRepository interface {
	SaveMessage(ctx context.Context, request entity.SaveMessageDTO) (id int64, err error)
	GetMessages(ctx context.Context, request entity.GetMessagesDTO) (messages []entity.Message, err error)
}

func New(cfg config.Application, messageRepository messageRepository) *Application {
	return &Application{
		cfg:               cfg,
		messageRepository: messageRepository,
	}
}
