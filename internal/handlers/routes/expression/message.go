package expression

import (
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
	"yourvoice/internal/database/models"
	"yourvoice/internal/utils"
)

type MessageRequest struct {
	models.Message
	ID     struct{} `json:"-"`
	Digest string   `json:"digest"`
}

func Message(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	ctx := r.Context()

	var message MessageRequest
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	event, err := gorm.G[models.MessageEvent](db).Where("id = ?", message.MessageEventID).First(ctx)
	if err != nil {
		http.Error(w, "Message event not found", http.StatusNotFound)
		return
	}

	if !utils.VerifySignature(message.Digest, message.Data, event.PrivateKey.Key.PublicKey) {
		http.Error(w, "Invalid signature", http.StatusUnauthorized)
		return
	}

	// TODO: Process the message
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Vote received successfully"))
}
