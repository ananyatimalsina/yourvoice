package routes

import (
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
)

type DeleteModelRequest struct {
	ID uint `json:"id" gorm:"primaryKey"`
}

func DeleteModel(w http.ResponseWriter, r *http.Request, db *gorm.DB, model any) {
	var request DeleteModelRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Failed to read request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.Delete(model, request.ID).Error; err != nil {
		http.Error(w, "Failed to delete model", http.StatusInternalServerError)
		return
	}

	// TODO: Process the vote
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("Model deleted successfully"))
}
