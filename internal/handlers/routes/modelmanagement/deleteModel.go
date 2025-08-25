package modelmanagement

import (
	"gorm.io/gorm"
	"io"
	"net/http"
	"net/url"
)

type DeleteModelRequest struct {
	ID uint `json:"id" schema:"id,required"`
}

func DeleteModel[T any](w http.ResponseWriter, r *http.Request, db *gorm.DB, model *T) {
	ctx := r.Context()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read body: "+err.Error(), http.StatusBadRequest)
		return
	}
	values, err := url.ParseQuery(string(body))
	if err != nil {
		http.Error(w, "Failed to parse form data: "+err.Error(), http.StatusBadRequest)
		return
	}

	id := values.Get("id")
	if id == "" {
		http.Error(w, "ID is required for deleting a model", http.StatusBadRequest)
		return
	}

	if _, err := gorm.G[T](db).Where("id = ?", id).Delete(ctx); err != nil {
		http.Error(w, "Failed to delete model", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Model deleted successfully"))
}
