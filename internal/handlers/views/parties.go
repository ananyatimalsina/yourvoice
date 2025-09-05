package views

import (
	"context"
	"github.com/a-h/templ"
	"gorm.io/gorm"
	"io"
	"net/http"
	"yourvoice/web/templates"
)

func Parties(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	parties := templates.Party()

	if r.Header.Get("AJAX-Target") == "main" {
		parties.Render(r.Context(), w)
		return
	}

	layout := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return templates.Layout("Parties").Render(templ.WithChildren(ctx, parties), w)
	})
	templ.Handler(layout).ServeHTTP(w, r)

}
