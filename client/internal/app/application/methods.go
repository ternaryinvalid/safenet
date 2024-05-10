package application

func (app *Application) SendMessage(pubKey []byte, message string) error {
	bytesMessage := []byte(message)

	err := app.serverProvider.SendMessage(pubKey, bytesMessage)
	if err != nil {
		return err
	}

	return nil
}
