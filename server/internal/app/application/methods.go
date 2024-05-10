package application

func (app *Application) SaveMessage(message string) error {
	err := app.messageRepository.SaveMessage(message)
	if err != nil {
		return err
	}

	return nil
}
