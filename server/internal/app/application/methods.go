package application

import (
	"context"
	"github.com/ternaryinvalid/safenet/server/internal/app/domain/entity"
	"github.com/ternaryinvalid/safenet/server/internal/pkg/cryptography"
)

const batchSize = 50

func (app *Application) SaveMessage(ctx context.Context, request entity.SaveMessageDTO) (id int64, err error) {
	id, err = app.messageRepository.SaveMessage(ctx, request)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (app *Application) GetMessages(ctx context.Context, request entity.GetMessagesDTO) (messages []entity.Message, err error) {
	if request.Limit == nil {
		*request.Limit = batchSize
	}

	messages, err = app.messageRepository.GetMessages(ctx, request)
	if err != nil {
		return nil, err
	}

	var deciphered bool
	if request.Deciphered == nil {
		*request.Deciphered = false
	}
	deciphered = *request.Deciphered

	if deciphered {
		decipheredMessages := make([]entity.Message, 0)

		sharedKey, err := app.cache.GetShared(request.MessageTo)
		if err != nil {
			return nil, err
		}

		for _, message := range messages {
			decipheredMessage, err := cryptography.Encrypt([]byte(sharedKey), []byte(message.MessageData))
			if err != nil {
				return nil, err
			}

			messageEncrypted := entity.Message{
				MessageFrom: message.MessageFrom,
				MessageData: decipheredMessage,
			}

			decipheredMessages = append(decipheredMessages, messageEncrypted)
		}

		return decipheredMessages, nil
	}

	return messages, nil
}

func (app *Application) GenerateKeys(remotePublicKey []byte) ([]byte, error) {
	localPublicKey, localPrivateKey, err := cryptography.GenerateKeys()
	if err != nil {
		return nil, err
	}

	sharedKey := cryptography.GetSharedKey(remotePublicKey, localPrivateKey)

	app.cache.SetSecret(string(localPublicKey), string(localPrivateKey.D.Bytes()))
	app.cache.SaveShared(string(remotePublicKey), string(sharedKey))

	return localPublicKey, nil
}
