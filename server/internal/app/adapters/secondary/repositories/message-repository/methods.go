package message_repository

import (
	"context"

	"github.com/ternaryinvalid/safenet/server/internal/app/domain/entity"
)

func (repo *MessageRepository) SaveMessage(ctx context.Context, request entity.SaveMessageDTO) (id int64, err error) {
	query, args := repo.createQuerySaveMessage(request)

	err = repo.DB.SelectContext(ctx, &id, query, args)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (repo *MessageRepository) GetMessages(ctx context.Context, request entity.GetMessagesDTO) (messages []entity.Message, err error) {
	query, args := repo.createQueryGetMessages(request)

	err = repo.DB.SelectContext(ctx, &messages, query, args)
	if err != nil {
		return nil, err
	}

	return messages, nil
}
