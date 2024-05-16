package message_repository

import (
	"context"
	"log"
	"time"

	"github.com/ternaryinvalid/safenet/server/internal/app/domain/entity"
)

func (repo *MessageRepository) SaveMessage(ctx context.Context, request entity.MessageSaveDTO) (id int64, err error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*60)
	defer cancel()

	query, args := repo.createQuerySaveMessage(request)

	var dto dtoSave

	err = repo.DB.GetContext(ctx, &dto, query, args...)
	if err != nil {
		log.Println(err)

		return 0, err
	}

	return dto.MessageID, nil
}

func (repo *MessageRepository) GetMessages(ctx context.Context, request entity.MessagesGetDTO) (messages []entity.Message, err error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*60)
	defer cancel()

	query, args := repo.createQueryGetMessages(request)

	log.Println(query)

	err = repo.DB.SelectContext(ctx, &messages, query, args...)
	if err != nil {
		log.Println(err)

		return nil, err
	}

	log.Println(messages)

	return messages, nil
}

type dtoSave struct {
	MessageID int64 `db:"message_id" json:"message_id"`
}
