package expression

import (
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
	"yourvoice/internal/database/models"
	"yourvoice/internal/utils"
)

type FeedbackRequest struct {
	models.Feedback
	ID     struct{} `json:"-"`
	Digest string   `json:"digest"`
}

func Feedback(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	ctx := r.Context()

	var feedback FeedbackRequest
	err := json.NewDecoder(r.Body).Decode(&feedback)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	event, err := gorm.G[models.FeedbackSession](db).Where("id = ?", feedback.FeedbackSessionID).First(ctx)
	if err != nil {
		http.Error(w, "Message event not found", http.StatusNotFound)
		return
	}

	if !utils.VerifySignature(feedback.Digest, feedback.Data, event.PrivateKey.Key.PublicKey) {
		http.Error(w, "Invalid signature", http.StatusUnauthorized)
		return
	}

	// TODO: Process the message
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Vote received successfully"))
}
