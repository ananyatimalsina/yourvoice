package modelmanagement

import (
	"github.com/a-h/templ"
	"github.com/ananyatimalsina/schema"
	"gorm.io/gorm"
	"net/http"
	"yourvoice/web/templates/admin/components"
)

func CreateModel[T any](w http.ResponseWriter, r *http.Request, db *gorm.DB, decoder *schema.Decoder, model *T, mkRow func(model any) components.RowProps, actions []templ.Component, options [2]bool) {
	var request T
	ctx := r.Context()

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form data: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := decoder.Decode(&request, r.PostForm); err != nil {
		http.Error(w, "Failed to parse request data: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := gorm.G[T](db).Create(ctx, &request); err != nil {
		http.Error(w, "Failed to create model", http.StatusInternalServerError)
		return
	}

	row := mkRow(request)

	components.TBody(components.TBodyProps{
		Rows: []components.RowProps{row},
	}, [3]bool{options[0], options[1], false}).Render(ctx, w)
}
