package identity

import (
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
	"yourvoice/internal/database/models"
	"yourvoice/internal/utils"
)

type VerifyRequest struct {
	Digest  string `json:"digest"`
	EventID uint   `json:"event_id"`
}

func Verify() bool {
	// Implement your verification logic here
	return true
}

func VerifyVote(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	ctx := r.Context()

	var verify VerifyRequest
	err := json.NewDecoder(r.Body).Decode(&verify)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	event, err := gorm.G[models.VoteEvent](db).Where("id = ?", verify.EventID).First(ctx)
	if err != nil {
		http.Error(w, "Vote event not found", http.StatusNotFound)
		return
	}

	if !Verify() {
		http.Error(w, "Verification failed", http.StatusUnauthorized)
		return
	}

	signature, err := utils.GenerateSignature(verify.Digest, event.PrivateKey.Key)
	if err != nil {
		http.Error(w, "Failed to generate signature", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"signature": signature,
	})
}

func VerifyMessage(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	ctx := r.Context()

	var verify VerifyRequest
	err := json.NewDecoder(r.Body).Decode(&verify)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	event, err := gorm.G[models.MessageEvent](db).Where("id = ?", verify.EventID).First(ctx)
	if err != nil {
		http.Error(w, "Message event not found", http.StatusNotFound)
		return
	}

	if !Verify() {
		http.Error(w, "Verification failed", http.StatusUnauthorized)
		return
	}

	signature, err := utils.GenerateSignature(verify.Digest, event.PrivateKey.Key)
	if err != nil {
		http.Error(w, "Failed to generate signature", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"signature": signature,
	})
}
