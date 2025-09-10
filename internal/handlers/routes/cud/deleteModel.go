package cud

import (
	"net/http"
	"strings"

	"gorm.io/gorm"
)

func DeleteModel[T any](w http.ResponseWriter, r *http.Request, db *gorm.DB, model *T) {
	ids := []string{}

	if r.Header.Get("AJAX-Targets") != "" {
		targets := strings.SplitSeq(r.Header.Get("AJAX-Targets"), ",")
		for t := range targets {
			id := strings.TrimPrefix(t, "row-")
			ids = append(ids, id)
		}
	}

	if r.Header.Get("AJAX-Target") != "" {
		ids = append(ids, strings.TrimPrefix(r.Header.Get("AJAX-Target"), "row-"))
	}

	if len(ids) == 0 {
		http.Error(w, "At least one ID is required for deleting a model", http.StatusBadRequest)
		return
	}

	db.Delete(model, ids)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Model deleted successfully"))
}
