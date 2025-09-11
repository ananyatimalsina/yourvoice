package cud

import (
	"encoding/json"
	"github.com/a-h/templ"
	"gorm.io/gorm"
	"net/http"
	"yourvoice/web/templates/modelmanagement"
)

func EditModel[T any](w http.ResponseWriter, r *http.Request, db *gorm.DB, mkRow func(model any) modelmanagement.RowProps, model *T) {
	ctx := r.Context()
	var request T

	ids, err := GetIdsFromAjax(r)

	if err != nil {
		http.Error(w, "Failed to parse IDs from request: "+err.Error(), http.StatusBadRequest)
		return
	}

	skipUnique := len(ids) > 1

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Failed to parse request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	validationErrors, err := ValidateStruct(request, skipUnique)
	if err != nil {
		http.Error(w, "Failed to marshal validation errors: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if validationErrors != "" {
		http.Error(w, validationErrors, http.StatusBadRequest)
		return
	}

	if _, err := gorm.G[T](db).Where("id IN ?", ids).Updates(ctx, request); err != nil {
		gormErrors, err := ValidateGorm(request, err)
		if err != nil {
			http.Error(w, "Failed to marshal gorm errors: "+err.Error(), http.StatusInternalServerError)
			return
		}

		if gormErrors != "" {
			http.Error(w, gormErrors, http.StatusBadRequest)
			return
		}

		http.Error(w, "Failed to edit model", http.StatusInternalServerError)
		return
	}

	updatedModels, err := gorm.G[T](db).Where("id IN ?", ids).Find(ctx)
	if err != nil {
		http.Error(w, "Failed to retrieve updated model", http.StatusInternalServerError)
		return
	}

	if len(updatedModels) == 1 {
		modelmanagement.ModelRow(mkRow(updatedModels[0])).Render(ctx, w)
	} else {
		rows := make([]templ.Component, len(updatedModels))
		for i, m := range updatedModels {
			rows[i] = modelmanagement.ModelRow(mkRow(m))
		}

		// wrap inside of <table></table> to ensure valid HTML
		w.Write([]byte("<table><tbody>"))
		templ.Join(rows...).Render(ctx, w)
		w.Write([]byte("</tbody></table>"))
	}
}
