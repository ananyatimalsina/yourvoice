package handlers

import (
	"gorm.io/gorm"
	"html/template"
	"net/http"
	"yourvoice/internal/handlers/routes/expression"
	"yourvoice/internal/handlers/routes/identity"
	"yourvoice/internal/middleware"
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

	expressionRouter.HandleFunc("POST /message", func(w http.ResponseWriter, r *http.Request) {
		expression.Message(w, r, db)
	})

	identityRouter := http.NewServeMux()

	identityRouter.HandleFunc("POST /verifyVote", func(w http.ResponseWriter, r *http.Request) {
		identity.VerifyVote(w, r, db)
	})

	identityRouter.HandleFunc("POST /verifyMessage", func(w http.ResponseWriter, r *http.Request) {
		identity.VerifyMessage(w, r, db)
	})

	router.Handle("/api/expression/", http.StripPrefix("/api/expression", middleware.ContentTypeJson(expressionRouter)))
	router.Handle("/api/identity/", http.StripPrefix("/api/identity", middleware.ContentTypeJson(identityRouter)))
}
