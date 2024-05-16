package application

import (
	"context"
	"github.com/ternaryinvalid/safenet/server/internal/app/domain/entity"
)

func (app *Application) SaveMessage(ctx context.Context, request entity.MessageSaveDTO) (id int64, err error) {
	id, err = app.messageRepository.SaveMessage(ctx, request)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (app *Application) GetMessages(ctx context.Context, request entity.MessagesGetDTO) (messages []entity.Message, err error) {
	messages, err = app.messageRepository.GetMessages(ctx, request)
	if err != nil {
		return nil, err
	}

	return messages, nil
}
