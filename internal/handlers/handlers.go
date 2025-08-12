package handlers

import (
	"gorm.io/gorm"
	"html/template"
	"net/http"
	"yourvoice/internal/handlers/routes/expression"
)

func LoadHanders(router *http.ServeMux, db *gorm.DB) {
	// Index page
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("web/templates/index.html"))
		tmpl.Execute(w, nil)
	})

	expressionRouter := http.NewServeMux()

	expressionRouter.HandleFunc("POST /vote", func(w http.ResponseWriter, r *http.Request) {
		expression.Vote(w, r, db)
	})

	router.Handle("/api/expression/", http.StripPrefix("/api/expression", expressionRouter))
}
