package views

import (
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"yourvoice/internal/database/models"
	"yourvoice/web/templates"
)

func RegisterPartyRoutes(mux *http.ServeMux, db *gorm.DB) {

	modelManagementProps := ModelManagementProps{
		Model:         models.Party{},
		Title:         "Parties",
		PreloadFields: []string{"Candidates"},
		Headers:       []string{"Name", "Candidates", "Created At"},
		MkRow:         mkRow,
	}

	mux.HandleFunc("GET /parties", func(w http.ResponseWriter, r *http.Request) {
		ModelManagement(w, r, db, modelManagementProps)
	})
}

func mkRow(model any) templates.RowProps {
	party := model.(models.Party)
	candidateCount := strconv.Itoa(len(party.Candidates))
	createdAt := party.CreatedAt.Format("Jan 2, 2006")

	return templates.RowProps{
		Cells: []string{
			party.Name,
			candidateCount,
			createdAt,
		}}
}
