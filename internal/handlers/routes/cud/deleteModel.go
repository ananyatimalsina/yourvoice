package cud

import (
	"net/http"
	"strings"

	"gorm.io/gorm"
)

func DeleteModel[T any](w http.ResponseWriter, r *http.Request, db *gorm.DB, model *T) {
	ctx := r.Context()

	id := strings.TrimPrefix(r.Header.Get("AJAX-Target"), "row-")
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
