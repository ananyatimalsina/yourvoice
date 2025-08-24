package modelmanagement

import (
	"github.com/gorilla/schema"
	"gorm.io/gorm"
	"net/http"
)

type DeleteModelRequest struct {
	ID uint `json:"id" schema:"id,required"`
}

func DeleteModel[T any](w http.ResponseWriter, r *http.Request, db *gorm.DB, model *T) {
	var request DeleteModelRequest
	ctx := r.Context()

	if err := schema.NewDecoder().Decode(&request, r.URL.Query()); err != nil {
		http.Error(w, "Failed to parse request data: "+err.Error(), http.StatusBadRequest)
		return
	}

	if _, err := gorm.G[T](db).Where("id = ?", request.ID).Delete(ctx); err != nil {
		http.Error(w, "Failed to delete model", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Model deleted successfully"))
}
