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
	SendMessage(ctx context.Context, request entity.SaveMessageDTO) (int64, error)
	GetMessages(ctx context.Context, request entity.GetMessagesDTO) ([]entity.Message, error)
	GenerateKeys(ctx context.Context, request entity.GenerateKeysDTO) (entity.GenerateKeysDTO, error)
}

type cacheRepository interface {
	SaveShared(sharedSecret string)
	GetShared() (string, error)
	SetSecret(localPublicKey, localPrivateKey string)
	IsEmpty() bool
	Public() string
}

func New(cfg config.Application, serverProvider serverProvider, cacheRepository cacheRepository) *Application {
	return &Application{
		cfg:             cfg,
		serverProvider:  serverProvider,
		cacheRepository: cacheRepository,
	}
}
