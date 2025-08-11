package handlers

import (
	"gorth/internal/handlers/routes"
	"html/template"
	"net/http"
)

func LoadHanders(router *http.ServeMux) {
	// Index page
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("web/templates/index.html"))
		tmpl.Execute(w, nil)
	})

	apiRouter := http.NewServeMux()

	apiRouter.HandleFunc("GET /greeting", routes.GreetingHandler)
	apiRouter.HandleFunc("GET /time", routes.TimeHandler)
	apiRouter.HandleFunc("GET /stats", routes.StatsHandler)

	router.Handle("/api/", http.StripPrefix("/api", apiRouter))
}
