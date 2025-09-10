package cud

import (
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
	"yourvoice/web/templates/modelmanagement"
)

func CreateModel[T any](w http.ResponseWriter, r *http.Request, db *gorm.DB, mkRow func(model any) modelmanagement.RowProps, model *T) {
	var request T
	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Failed to parse request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	validationErrors, err := ValidateStruct(request)
	if err != nil {
		http.Error(w, "Failed to marshal validation errors: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if validationErrors != "" {
		http.Error(w, validationErrors, http.StatusBadRequest)
		return
	}

	if err := gorm.G[T](db).Create(ctx, &request); err != nil {
		http.Error(w, "Failed to create model", http.StatusInternalServerError)
		return
	}

	modelmanagement.ModelRow(mkRow(request)).Render(ctx, w)
}
