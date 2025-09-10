package cud

import (
	"net/http"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

func DeleteModel[T any](w http.ResponseWriter, r *http.Request, db *gorm.DB, model *T) {
	ids := []uint64{}

	if r.Header.Get("AJAX-Targets") != "" {
		targets := strings.SplitSeq(r.Header.Get("AJAX-Targets"), ",")
		for t := range targets {
			id := strings.TrimPrefix(t, "row-")
			parseUint, err := strconv.ParseUint(id, 10, 64)
			if err != nil {
				http.Error(w, "Invalid ID format: "+id, http.StatusBadRequest)
				return
			}
			ids = append(ids, parseUint)
		}
	}

	if r.Header.Get("AJAX-Target") != "" {
		id := strings.TrimPrefix(r.Header.Get("AJAX-Target"), "row-")
		parseUint, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			http.Error(w, "Invalid ID format: "+id, http.StatusBadRequest)
			return
		}
		ids = append(ids, parseUint)
	}

	if len(ids) == 0 {
		http.Error(w, "At least one ID is required for deleting a model", http.StatusBadRequest)
		return
	}

	result := db.Delete(model, ids)
	if result.Error != nil {
		http.Error(w, "Failed to delete model(s): "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Model deleted successfully"))
}
