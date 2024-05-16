package server_provider

import (
	"context"
	"time"

	"github.com/ternaryinvalid/safenet/client/internal/app/domain/entity"
	"github.com/ternaryinvalid/safenet/client/internal/pkg/providerhelpers"
)

func (prv *ServerProvider) SendMessage(ctx context.Context, request entity.MessageSendDTO) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*60)
	defer cancel()

	req := providerhelpers.CreateRequest(ctx, prv.client, prv.cfg.Endpoints.Send)

	var dto messageSendDto

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

func (prv *ServerProvider) GetMessages(ctx context.Context, request entity.MessagesGetDTO) ([]entity.Message, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*60)
	defer cancel()

	req := providerhelpers.CreateRequest(ctx, prv.client, prv.cfg.Endpoints.Get)

	var dto messageGetDto

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

type messageGetDto struct {
	Messages []entity.Message `json:"messages"`
}

type messageSendDto struct {
	MessageId int64 `json:"message_id"`
}
