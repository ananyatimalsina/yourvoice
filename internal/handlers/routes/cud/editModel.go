package cud

import (
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
	"yourvoice/internal/utils"
	"yourvoice/web/templates/modelmanagement"
)

func EditModel[T any](w http.ResponseWriter, r *http.Request, db *gorm.DB, mkRow func(model any) modelmanagement.RowProps, model *T) {
	var request T
	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Failed to parse request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	id := utils.GetModelID(request)

	if id == 0 {
		http.Error(w, "ID is required for editing a model", http.StatusBadRequest)
		return
	}

	if _, err := gorm.G[T](db).Where("id = ?", id).Updates(ctx, request); err != nil {
		http.Error(w, "Failed to edit model", http.StatusInternalServerError)
		return
	}

	updatedModel, err := gorm.G[T](db).Where("id = ?", id).First(ctx)
	if err != nil {
		http.Error(w, "Failed to retrieve updated model", http.StatusInternalServerError)
		return
	}

	modelmanagement.ModelRow(mkRow(updatedModel)).Render(ctx, w)
}
