package entity

type SaveMessageDTO struct {
	MessageFrom string `json:"message_from"`
	MessageTo   string `json:"message_to"`
	MessageData string `json:"message_data"`
}

type GetMessagesDTO struct {
	MessageTo  string `json:"message_to"`
	Deciphered *bool  `json:"deciphered"`
	Limit      *int   `json:"limit"`
}

type MessageDTO struct {
	Messages []Message `json:"messages"`
}

type Message struct {
	MessageFrom string `json:"message_from"`
	MessageData string `json:"message_data"`
}

type GenerateKeysDTO struct {
	PublicKey string `json:"public_key"`
}
