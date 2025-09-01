package main

import (
	"github.com/ananyatimalsina/schema"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"yourvoice/internal/database"
	"yourvoice/internal/handlers"
	"yourvoice/internal/middleware"
)

var decoder *schema.Decoder

func main() {
	godotenv.Load(".env")
	decoder = schema.NewDecoder()
	decoder.ZeroEmpty(true)
	db := database.LoadDatabase(decoder)
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

	handlers.LoadHanders(router, db, decoder)

	log.Println("Server started at http://localhost:" + port)
	log.Fatal(server.ListenAndServe())
}
