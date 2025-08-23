package modelmanagement

import (
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
)

type DeleteModelRequest struct {
	ID uint `json:"id" gorm:"primaryKey"`
}

func DeleteModel[T any](w http.ResponseWriter, r *http.Request, db *gorm.DB, model *T) {
	var request DeleteModelRequest
	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Failed to read request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	if _, err := gorm.G[T](db).Where("id = ?", request.ID).Delete(ctx); err != nil {
		http.Error(w, "Failed to delete model", http.StatusInternalServerError)
		return
	}

	// TODO: Process the vote
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("Model deleted successfully"))
}
