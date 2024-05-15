package api_controller

import (
	"encoding/json"
	"fmt"
	"github.com/ternaryinvalid/safenet/client/internal/app/domain/entity"
	"log"
	"net/http"
)

func (ctr *ApiController) SendMessage(w http.ResponseWriter, r *http.Request) {
	var dto entity.SaveMessageDTO

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

	id, err := ctr.app.SendMessage(ctx, dto)
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
	var dtoIn entity.GetMessagesDTO

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

	dtoOut := GetMessagesResponse{Messages: messages}

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
	Messages []entity.Message `json:"messages"`
}

type SaveMessageResponse struct {
	MessageId int64 `json:"message_id"`
}
