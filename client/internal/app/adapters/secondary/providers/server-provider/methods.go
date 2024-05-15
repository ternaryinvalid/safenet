package server_provider

import (
	"context"
	api_controller "github.com/ternaryinvalid/safenet/client/internal/app/adapters/primary/http-adapter/api-controller"
	"time"

	"github.com/ternaryinvalid/safenet/client/internal/app/domain/entity"
	"github.com/ternaryinvalid/safenet/client/internal/pkg/providerhelpers"
)

func (prv *ServerProvider) SendMessage(ctx context.Context, request entity.SaveMessageDTO) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*60)
	defer cancel()

	req := providerhelpers.CreateRequest(ctx, prv.client, prv.cfg.Endpoints.Send)

	var dto api_controller.SaveMessageResponse

	req.
		SetBody(request).
		ForceContentType("application/json").
		SetResult(&dto)

	_, err := req.Send()
	if err != nil {
		return 0, err
	}

	return dto.MessageId, nil
}

func (prv *ServerProvider) GetMessages(ctx context.Context, request entity.GetMessagesDTO) ([]entity.Message, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*60)
	defer cancel()

	req := providerhelpers.CreateRequest(ctx, prv.client, prv.cfg.Endpoints.Send)

	var dto entity.MessageDTO

	req.
		SetBody(request).
		ForceContentType("application/json").
		SetResult(&dto)

	_, err := req.Send()
	if err != nil {
		return nil, err
	}

	return dto.Messages, nil
}

func (prv *ServerProvider) GenerateKeys(ctx context.Context, request entity.GenerateKeysDTO) (entity.GenerateKeysDTO, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*60)
	defer cancel()

	req := providerhelpers.CreateRequest(ctx, prv.client, prv.cfg.Endpoints.GenKeys)

	var dto entity.GenerateKeysDTO

	req.
		SetBody(request).
		ForceContentType("application/json").
		SetResult(&dto)

	_, err := req.Send()
	if err != nil {
		return entity.GenerateKeysDTO{}, err
	}

	return dto, nil
}
