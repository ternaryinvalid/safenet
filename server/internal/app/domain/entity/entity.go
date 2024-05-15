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

type Message struct {
	MessageFrom string `json:"message_from"`
	MessageData string `json:"message_data"`
}

type RegisterDTO struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
