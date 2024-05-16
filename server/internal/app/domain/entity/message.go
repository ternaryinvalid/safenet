package entity

import "time"

type Message struct {
	MessageFrom string    `db:"public_key_from" json:"message_from"`
	MessageData string    `db:"message_data" json:"message_data"`
	Dt          time.Time `db:"dt" json:"dt"`
}

type MessageSaveDTO struct {
	MessageFrom string `json:"message_from"`
	MessageTo   string `json:"message_to"`
	MessageData string `json:"message_data"`
}

type MessagesGetDTO struct {
	MessageTo string `json:"message_to"`
}
