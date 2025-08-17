package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"yourvoice/internal/database"
	"yourvoice/internal/handlers"
	"yourvoice/internal/middleware"
)

func main() {
	godotenv.Load(".env")
	db := database.LoadDatabase()
	port := os.Getenv("SERVER_PORT")
	router := http.NewServeMux()

	stack := middleware.CreateStack(
		middleware.LogOriginalURL,
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":" + port,
		Handler: stack(router),
	}

	// Static files
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	handlers.LoadHanders(router, db)

	log.Println("Server started at http://localhost:" + port)
	log.Fatal(server.ListenAndServe())
}
