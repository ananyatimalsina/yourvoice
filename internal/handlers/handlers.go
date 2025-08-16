package handlers

import (
	"github.com/a-h/templ"
	"gorm.io/gorm"
	"net/http"
	"yourvoice/internal/handlers/routes/expression"
	"yourvoice/internal/handlers/routes/identity"
	"yourvoice/internal/middleware"
	"yourvoice/web/templates/admin/pages"
)

func LoadHanders(router *http.ServeMux, db *gorm.DB) {
	// Index page
	//router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	tmpl := template.Must(template.ParseFiles("web/templates/index.html"))
	//	tmpl.Execute(w, nil)
	//})

	expressionRouter := http.NewServeMux()

	expressionRouter.HandleFunc("POST /vote", func(w http.ResponseWriter, r *http.Request) {
		expression.Vote(w, r, db)
	})

	expressionRouter.HandleFunc("POST /message", func(w http.ResponseWriter, r *http.Request) {
		expression.Feedback(w, r, db)
	})

	identityRouter := http.NewServeMux()

	identityRouter.HandleFunc("POST /verifyVote", func(w http.ResponseWriter, r *http.Request) {
		identity.VerifyVote(w, r, db)
	})

	identityRouter.HandleFunc("POST /verifyMessage", func(w http.ResponseWriter, r *http.Request) {
		identity.VerifyFeedback(w, r, db)
	})

	adminRouter := http.NewServeMux()

	adminRouter.Handle("/", templ.Handler(pages.Dashboard()))
	adminRouter.Handle("/data/logs", templ.Handler(pages.SystemLogs()))

	router.Handle("/api/expression/", http.StripPrefix("/api/expression", middleware.ContentTypeJson(expressionRouter)))
	router.Handle("/api/identity/", http.StripPrefix("/api/identity", middleware.ContentTypeJson(identityRouter)))

	router.Handle("/admin/", http.StripPrefix("/admin", adminRouter))
}
