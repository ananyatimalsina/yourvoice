package admin_vote

import (
	"net/http"
	"strconv"
	"yourvoice/internal/database/models"
	"yourvoice/internal/handlers/routes/modelmanagement"
	"yourvoice/internal/handlers/views/admin"
	"yourvoice/web/templates/admin/components"

	"github.com/a-h/templ"
	"github.com/ananyatimalsina/schema"
	"gorm.io/gorm"
)

func RegisterPartyRoutes(mux *http.ServeMux, db *gorm.DB, decoder *schema.Decoder) {
	modelManagementProps := admin.ModelManagerProps{
		Model:         models.Party{},
		SearchFields:  []string{"name"},
		PreloadFields: []string{"Candidates"},
		Headers:       []string{"Party Name", "Candidates", "Created"},
		MkRow: func(model any) components.RowProps {
			party := model.(models.Party)
			candidateCount := strconv.Itoa(len(party.Candidates))
			createdAt := party.CreatedAt.Format("Jan 2, 2006")

			return components.RowProps{Model: party, Cells: []string{
				party.Name,
				candidateCount,
				createdAt,
			}}
		},
		Title:         "Parties",
		SingularTitle: "Party",
		Icon:          "🏛️",
		Actions:       []templ.Component{},
		Options:       [2]bool{true, true},
	}

	mux.HandleFunc("POST /parties", func(w http.ResponseWriter, r *http.Request) {
		modelmanagement.CreateModel(w, r, db, decoder, &models.Party{}, modelManagementProps.MkRow, modelManagementProps.Actions, modelManagementProps.Options)
	})

	mux.HandleFunc("GET /parties", func(w http.ResponseWriter, r *http.Request) {
		admin.ModelManager(w, r, db, modelManagementProps)
	})

	mux.HandleFunc("PUT /parties", func(w http.ResponseWriter, r *http.Request) {
		modelmanagement.EditModel(w, r, db, decoder, &models.Party{}, modelManagementProps.MkRow, modelManagementProps.Actions, modelManagementProps.Options)
	})

	mux.HandleFunc("DELETE /parties", func(w http.ResponseWriter, r *http.Request) {
		modelmanagement.DeleteModel(w, r, db, &models.Party{})
	})

}
