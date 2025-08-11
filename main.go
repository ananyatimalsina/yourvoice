package main

import (
	"log"
	"net/http"
	"os"
	"yourvoice/internal/handlers"
	"yourvoice/internal/middleware"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	//db := database.LoadDatabase()
	port := os.Getenv("SERVER_PORT")
	router := http.NewServeMux()

	stack := middleware.CreateStack(
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":" + port,
		Handler: stack(router),
	}

	// Static files
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	handlers.LoadHanders(router)

	log.Println("Server started at http://localhost:" + port)
	log.Fatal(server.ListenAndServe())
}
