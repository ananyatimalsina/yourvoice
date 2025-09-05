package handlers

import (
	"net/http"
	"yourvoice/internal/database/models"
	"yourvoice/internal/handlers/routes/expression"
	"yourvoice/internal/handlers/routes/identity"
	"yourvoice/internal/handlers/routes/modelmanagement"
	"yourvoice/internal/handlers/views/admin/vote"
	"yourvoice/internal/middleware"
	"yourvoice/web/templates/admin/pages"

	"github.com/a-h/templ"
	"github.com/ananyatimalsina/schema"
	"gorm.io/gorm"
)

func LoadHanders(router *http.ServeMux, db *gorm.DB, decoder *schema.Decoder) {
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

	router.Handle("/api/expression/", http.StripPrefix("/api/expression", middleware.ContentTypeJson(expressionRouter)))

	identityRouter := http.NewServeMux()

	identityRouter.HandleFunc("POST /verifyVote", func(w http.ResponseWriter, r *http.Request) {
		identity.VerifyVote(w, r, db)
	})

	identityRouter.HandleFunc("POST /verifyMessage", func(w http.ResponseWriter, r *http.Request) {
		identity.VerifyFeedback(w, r, db)
	})

	router.Handle("/api/identity/", http.StripPrefix("/api/identity", middleware.ContentTypeJson(identityRouter)))

	adminRouter := http.NewServeMux()

	adminRouter.Handle("/", templ.Handler(pages.Dashboard()))
	adminRouter.Handle("/data/logs", templ.Handler(pages.SystemLogs()))

	adminVoteRouter := http.NewServeMux()

	// Event management
	adminVoteRouter.HandleFunc("GET /events", func(w http.ResponseWriter, r *http.Request) {
		admin_vote.Events(w, r, db)
	})
	adminVoteRouter.HandleFunc("DELETE /events", func(w http.ResponseWriter, r *http.Request) {
		modelmanagement.DeleteModel(w, r, db, &models.VoteEvent{})
	})

	// Party management
	admin_vote.RegisterPartyRoutes(adminVoteRouter, db, decoder)

	// Candidate management
	admin_vote.RegisterCandidateRoutes(adminVoteRouter, db, decoder)

	adminRouter.Handle("/vote/", http.StripPrefix("/vote", adminVoteRouter))

	router.Handle("/admin/", http.StripPrefix("/admin", adminRouter))
}
