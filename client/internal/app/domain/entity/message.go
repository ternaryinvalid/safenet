package entity

type Message struct {
	MessageFrom string `json:"message_from"`
	MessageData string `json:"message_data"`
	Dt          string `json:"dt"`
}

type MessageSendDTO struct {
	MessageFrom string `json:"message_from"`
	MessageTo   string `json:"message_to"`
	MessageData string `json:"message_data"`
}

type MessagesGetDTO struct {
	MessageTo string `json:"message_to"`
}
