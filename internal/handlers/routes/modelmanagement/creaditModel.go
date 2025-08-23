package modelmanagement

import (
	"github.com/gorilla/schema"
	"gorm.io/gorm"
	"net/http"
)

func redirectToOriginal(w http.ResponseWriter, r *http.Request) {
	// Get the original URL from the request header (set by our middleware)
	originalURL := r.Header.Get("X-Original-URL")
	if originalURL == "" {
		// Fallback: try to reconstruct from the request URL
		originalURL = r.URL.Path
	}

	// Preserve query parameters
	if r.URL.RawQuery != "" {
		originalURL += "?" + r.URL.RawQuery
	}

	http.Redirect(w, r, originalURL, http.StatusSeeOther)
}

func CreaditModel[T any](w http.ResponseWriter, r *http.Request, db *gorm.DB, model *T) {
	var request T
	ctx := r.Context()

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form data: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := schema.NewDecoder().Decode(&request, r.PostForm); err != nil {
		http.Error(w, "Failed to parse request data: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Check if ID exists
	id := r.PostForm.Get("id")
	if id != "" {
		if _, err := gorm.G[T](db).Where("id = ?", id).Updates(ctx, request); err != nil {
			http.Error(w, "Failed to edit model", http.StatusInternalServerError)
			return
		}

	} else {

		if err := gorm.G[T](db).Create(ctx, &request); err != nil {
			http.Error(w, "Failed to create model", http.StatusInternalServerError)
			return
		}
	}

	// Redirect to original URL to prevent form resubmission
	redirectToOriginal(w, r)
}
