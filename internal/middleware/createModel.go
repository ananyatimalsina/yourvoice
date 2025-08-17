package middleware

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

func CreateModel(w http.ResponseWriter, r *http.Request, db *gorm.DB, model any) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form data: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := schema.NewDecoder().Decode(model, r.PostForm); err != nil {
		http.Error(w, "Failed to parse request data: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.Create(model).Error; err != nil {
		http.Error(w, "Failed to create model", http.StatusInternalServerError)
		return
	}

	// Redirect to original URL to prevent form resubmission
	redirectToOriginal(w, r)
}
