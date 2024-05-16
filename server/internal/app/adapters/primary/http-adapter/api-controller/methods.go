package api_controller

import (
	"encoding/json"
	"fmt"
	"github.com/ternaryinvalid/safenet/server/internal/app/domain/entity"
	"log"
	"net/http"
)

func (ctr *ApiController) SaveMessage(w http.ResponseWriter, r *http.Request) {
	var dto entity.MessageSaveDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		_, writeErr := w.Write([]byte(fmt.Sprintf("error: %v", err)))
		if writeErr != nil {
			log.Println(writeErr)

			return
		}

		return
	}

	ctx := r.Context()

	id, err := ctr.app.SaveMessage(ctx, dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		_, writeErr := w.Write([]byte(fmt.Sprintf("error: %v", err)))
		if writeErr != nil {
			log.Println(writeErr)

			return
		}

		return
	}

	dtoOut := SaveMessageResponse{MessageId: id}

	err = json.NewEncoder(w).Encode(dtoOut)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		_, writeErr := w.Write([]byte(fmt.Sprintf("error: %v", err)))
		if writeErr != nil {
			log.Println(writeErr)

			return
		}

		return
	}
}

func (ctr *ApiController) GetMessages(w http.ResponseWriter, r *http.Request) {
	var dtoIn entity.MessagesGetDTO

	err := json.NewDecoder(r.Body).Decode(&dtoIn)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		_, writeErr := w.Write([]byte(fmt.Sprintf("error: %v", err)))
		if writeErr != nil {
			log.Println(writeErr)

			return
		}

		return
	}

	ctx := r.Context()

	messages, err := ctr.app.GetMessages(ctx, dtoIn)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		_, writeErr := w.Write([]byte(fmt.Sprintf("error: %v", err)))
		if writeErr != nil {
			log.Println(writeErr)

			return
		}

		return
	}

	dtoOut := toResponse(messages)

	err = json.NewEncoder(w).Encode(dtoOut)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		_, writeErr := w.Write([]byte(fmt.Sprintf("error: %v", err)))
		if writeErr != nil {
			log.Println(writeErr)

			return
		}

		return
	}
}

type GetMessagesResponse struct {
	Messages []MessageResponse `json:"messages"`
}

type MessageResponse struct {
	MessageFrom string `json:"message_from"`
	MessageData string `json:"message_data"`
	Dt          string `json:"dt"`
}

func toResponse(messages []entity.Message) GetMessagesResponse {
	var (
		resp    GetMessagesResponse
		message MessageResponse
	)

	for _, m := range messages {
		message = MessageResponse{
			MessageFrom: m.MessageFrom,
			MessageData: m.MessageData,
			Dt:          m.Dt.Format("2006-01-12"),
		}

		resp.Messages = append(resp.Messages, message)
	}

	return resp
}

type SaveMessageResponse struct {
	MessageId int64 `json:"message_id"`
}
