package message_repository

func (repo *MessageRepository) SaveMessage(message string) error {
	return nil
}

func (repo *MessageRepository) GetMessages(key []byte) ([][]byte, error) {
	return [][]byte{}, nil
}
