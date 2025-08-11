package routes

import (
	"html/template"
	"net/http"
	"time"
)

func TimeHandler(w http.ResponseWriter, r *http.Request) {
	now := time.Now()

	tmpl := template.Must(template.ParseFiles("web/templates/time.html"))
	tmpl.Execute(w, map[string]any{
		"Time":          now.Format("15:04:05 MST"),
		"Date":          now.Format("Monday, January 2, 2006"),
		"Timezone":      now.Location().String(),
		"UnixTimestamp": now.Unix(),
	})

}
