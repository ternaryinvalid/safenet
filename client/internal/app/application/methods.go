package application

import (
	"context"
	"github.com/ternaryinvalid/safenet/client/internal/app/domain/entity"
	"github.com/ternaryinvalid/safenet/client/internal/pkg/cryptography"
)

func (app *Application) SendMessage(ctx context.Context, request entity.SaveMessageDTO) (_ int64, err error) {
	sharedKey := make([]byte, 0)

	if app.cacheRepository.IsEmpty() {
		sharedKey, err = app.generateKeys(ctx)
		if err != nil {
			return 0, err
		}
	}

	cipheredMessage, err := cryptography.Decrypt(sharedKey, []byte(request.MessageData))
	if err != nil {
		return 0, err
	}

	cipheredDto := entity.SaveMessageDTO{
		MessageTo:   request.MessageTo,
		MessageFrom: request.MessageFrom,
		MessageData: cipheredMessage,
	}

	id, err := app.serverProvider.SendMessage(ctx, cipheredDto)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (app *Application) GetMessages(ctx context.Context, request entity.GetMessagesDTO) ([]entity.Message, error) {
	messages, err := app.serverProvider.GetMessages(ctx, request)
	if err != nil {
		return nil, err
	}

	var deciphered bool
	if request.Deciphered == nil {
		*request.Deciphered = false
	}
	deciphered = *request.Deciphered

	if !deciphered {
		messagesDeciphered := make([]entity.Message, 0)

		sharedKey, err := app.cacheRepository.GetShared(request.MessageTo)
		if err != nil {
			return nil, err
		}

		var encryptedData string

		for _, message := range messages {
			err := cryptography.Encrypt([]byte(sharedKey), []byte(message.MessageData), &encryptedData)
			if err != nil {
				return nil, err
			}

			messageDeciphered := entity.Message{
				MessageFrom: message.MessageFrom,
				MessageData: encryptedData,
			}

			messagesDeciphered = append(messagesDeciphered, messageDeciphered)
		}

		return messagesDeciphered, nil
	}

	return messages, nil
}

func (app *Application) generateKeys(ctx context.Context) ([]byte, error) {
	localPublicKey, localPrivateKey, err := cryptography.GenerateKeys()
	if err != nil {
		return nil, err
	}

	dto := entity.GenerateKeysDTO{
		PublicKey: string(localPublicKey),
	}

	remotePublicKey, err := app.serverProvider.GenerateKeys(ctx, dto)
	if err != nil {
		return nil, err
	}

	shared := cryptography.GetSharedKey([]byte(remotePublicKey.PublicKey), localPrivateKey)

	app.cacheRepository.SaveShared(remotePublicKey.PublicKey, string(shared))
	app.cacheRepository.SetSecret(string(localPublicKey), string(localPrivateKey.D.Bytes()))

	return shared, nil
}
