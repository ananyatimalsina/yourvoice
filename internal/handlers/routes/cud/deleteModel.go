package cud

import (
	"gorm.io/gorm"
	"net/http"
)

func DeleteModel[T any](w http.ResponseWriter, r *http.Request, db *gorm.DB, model *T) {
	ids, err := GetIdsFromAjax(r)

	if err != nil {
		http.Error(w, "Failed to parse IDs from request: "+err.Error(), http.StatusBadRequest)
		return
	}

	result := db.Delete(model, ids)
	if result.Error != nil {
		http.Error(w, "Failed to delete model(s): "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Model(s) deleted successfully"))
}
