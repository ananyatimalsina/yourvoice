package routes

import (
	"html/template"
	"net/http"
	"time"
)

func GreetingHandler(w http.ResponseWriter, r *http.Request) {
	greetings := []string{
		"Hello from the Gorth Stack! 🚀",
		"Greetings from your Go server! 👋",
		"Welcome to the future of web development!",
		"HTMX + Go = Pure Magic ✨",
		"Server-side rendering at its finest!",
	}

	currentTime := time.Now()
	greeting := greetings[currentTime.Second()%len(greetings)]

	tmpl := template.Must(template.ParseFiles("web/templates/greeting.html"))
	tmpl.Execute(w, map[string]any{
		"Message":   greeting,
		"Timestamp": currentTime.Format("15:04:05"),
	})
}
