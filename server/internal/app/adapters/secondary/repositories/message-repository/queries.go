package message_repository

import (
	"fmt"

	"github.com/ternaryinvalid/safenet/server/internal/app/domain/entity"
)

func (repo *MessageRepository) createQuerySaveMessage(req entity.SaveMessageDTO) (string, []interface{}) {
	procedure := repo.cfg.Procedures["saveMessage"]

	return fmt.Sprintf(`SELECT * FROM %s($1, $2, $3)`, procedure), []interface{}{req.MessageFrom, req.MessageTo, req.MessageData}
}

func (repo *MessageRepository) createQueryGetMessages(req entity.GetMessagesDTO) (string, []interface{}) {
	procedure := repo.cfg.Procedures["getMessages"]

	return fmt.Sprintf(`SELECT * FROM %s($1, $2)`, procedure), []interface{}{req.MessageTo, req.Limit}
}
