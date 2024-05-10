package server_provider

func (prv *ServerProvider) SendMessage(key []byte, message []byte) error {
	return nil
}

func (prv *ServerProvider) GetMessages(key []byte) ([][]byte, error) {
	return [][]byte{}, nil
}
