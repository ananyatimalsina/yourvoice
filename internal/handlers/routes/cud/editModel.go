package cud

import (
	"github.com/a-h/templ"
	"github.com/ananyatimalsina/schema"
	"gorm.io/gorm"
	"net/http"
)

func EditModel[T any](w http.ResponseWriter, r *http.Request, db *gorm.DB, decoder *schema.Decoder, model *T, actions []templ.Component, options [2]bool) {
	var request T
	ctx := r.Context()

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form data: "+err.Error(), http.StatusBadRequest)
		return
	}

	r.PostForm.Del("CreatedAt")
	r.PostForm.Del("UpdatedAt")
	r.PostForm.Del("DeletedAt")

	if err := decoder.Decode(&request, r.PostForm); err != nil {
		http.Error(w, "Failed to parse request data: "+err.Error(), http.StatusBadRequest)
		return
	}

	id := r.PostForm.Get("ID")

	if id == "" {
		http.Error(w, "ID is required for editing a model", http.StatusBadRequest)
		return
	}

	if _, err := gorm.G[T](db).Where("id = ?", id).Updates(ctx, request); err != nil {
		http.Error(w, "Failed to edit model", http.StatusInternalServerError)
		return
	}
}
