package application

import (
	"context"
	"fmt"
	"github.com/ternaryinvalid/safenet/client/internal/app/domain/entity"
	"github.com/ternaryinvalid/safenet/client/internal/pkg/cryptography"
	"log"
)

func (app *Application) SendMessage(ctx context.Context, message entity.MessageSendDTO) (_ int64, err error) {
	account, err := app.cacheRepository.LoadAccount()
	if err != nil {
		err = fmt.Errorf("аккаунт не создан")

		return 0, err
	}

	sharedKey, err := cryptography.GetTransportKey(account.PrivateKey, message.MessageTo)
	if err != nil {
		return 0, err
	}

	cipheredMessage, err := cryptography.Encrypt(message.MessageData, string(sharedKey))
	if err != nil {
		return 0, err
	}

	cipheredDto := entity.MessageSendDTO{
		MessageTo:   message.MessageTo,
		MessageFrom: message.MessageFrom,
		MessageData: cipheredMessage,
	}

	id, err := app.serverProvider.SendMessage(ctx, cipheredDto)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (app *Application) GetMessages(ctx context.Context) ([]entity.Message, error) {
	account, err := app.cacheRepository.LoadAccount()
	if err != nil {
		err = fmt.Errorf("аккаунт не создан")

		return nil, err
	}

	log.Println(account.PublicKey)

	dto := entity.MessagesGetDTO{
		MessageTo: account.PublicKey,
	}

	messages, err := app.serverProvider.GetMessages(ctx, dto)
	if err != nil {
		return nil, err
	}

	decipheredMessages := make([]entity.Message, 0)

	for _, message := range messages {
		sharedKey, err := cryptography.GetTransportKey(account.PrivateKey, message.MessageFrom)
		if err != nil {
			return nil, err
		}

		decryptedData, err := cryptography.Decrypt(message.MessageData, string(sharedKey))
		if err != nil {
			return nil, err
		}

		decipheredMessage := entity.Message{
			MessageFrom: message.MessageFrom,
			MessageData: decryptedData,
			Dt:          message.Dt,
		}

		decipheredMessages = append(decipheredMessages, decipheredMessage)
	}

	return decipheredMessages, nil
}

func (app *Application) CreateAccount(req entity.AccountCreateDTO) (entity.AccountResponseDTO, error) {
	acc, err := app.cacheRepository.LoadAccount()
	if err == nil {
		if acc.Name == req.Name {
			err = fmt.Errorf("аккаунт уже создан.")

			return entity.AccountResponseDTO{}, err
		}
	}

	if req.PrivateKey == nil {
		private, public, err := cryptography.GenerateKeys()
		if err != nil {
			return entity.AccountResponseDTO{}, err
		}

		account := entity.Account{
			Name:       req.Name,
			PublicKey:  public,
			PrivateKey: private,
		}

		response, err := app.createAccount(account)
		if err != nil {
			return entity.AccountResponseDTO{}, err
		}

		return response, nil
	}

	public, err := cryptography.GetPublicKeyAndAddressByPrivateKey(*req.PrivateKey)
	if err != nil {
		return entity.AccountResponseDTO{}, err
	}

	account := entity.Account{
		Name:       req.Name,
		PublicKey:  public,
		PrivateKey: *req.PrivateKey,
	}

	response, err := app.createAccount(account)
	if err != nil {
		return entity.AccountResponseDTO{}, err
	}

	return response, nil
}

func (app *Application) createAccount(account entity.Account) (entity.AccountResponseDTO, error) {
	err := app.cacheRepository.SaveAccount(&account)
	if err != nil {
		return entity.AccountResponseDTO{}, err
	}

	accountDto := entity.AccountResponseDTO{
		Name:      account.Name,
		PublicKey: account.PublicKey,
	}

	return accountDto, nil
}
