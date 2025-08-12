package expression

import (
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
	"yourvoice/internal/database/models"
	"yourvoice/internal/utils"
)

type VoteRequest struct {
	models.Vote
	ID     struct{} `json:"-"`
	Digest string   `json:"digest"`
}

func Vote(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	ctx := r.Context()

	var vote VoteRequest
	err := json.NewDecoder(r.Body).Decode(&vote)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	event, err := gorm.G[models.VoteEvent](db).Where("id = ?", vote.VoteEventID).First(ctx)
	if err != nil {
		http.Error(w, "Vote event not found", http.StatusNotFound)
		return
	}

	if !utils.VerifySignature(vote.Digest, vote.Data, event.PrivateKey.Key.PublicKey) {
		http.Error(w, "Invalid signature", http.StatusUnauthorized)
		return
	}

	// TODO: Process the vote
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Vote received successfully"))
}
