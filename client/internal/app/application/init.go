package application

import (
	"context"
	"github.com/ternaryinvalid/safenet/client/internal/app/domain/config"
	"github.com/ternaryinvalid/safenet/client/internal/app/domain/entity"
)

type Application struct {
	cfg             config.Application
	serverProvider  serverProvider
	cacheRepository cacheRepository
}

type serverProvider interface {
	SendMessage(ctx context.Context, message entity.MessageSendDTO) (int64, error)
	GetMessages(ctx context.Context, address entity.MessagesGetDTO) ([]entity.Message, error)
}

type cacheRepository interface {
	LoadAccount() (*entity.Account, error)
	SaveAccount(account *entity.Account) error
}

func New(cfg config.Application, serverProvider serverProvider, cacheRepository cacheRepository) *Application {
	return &Application{
		cfg:             cfg,
		serverProvider:  serverProvider,
		cacheRepository: cacheRepository,
	}
}
