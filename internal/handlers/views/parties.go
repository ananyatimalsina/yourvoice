package views

import (
	"net/http"
	"strconv"
	"yourvoice/internal/database/models"
	"yourvoice/internal/handlers/routes/cud"
	"yourvoice/web/templates/components/input"
	"yourvoice/web/templates/modelmanagement"

	"gorm.io/gorm"
)

func RegisterPartyRoutes(mux *http.ServeMux, db *gorm.DB) {
	modelManagementProps := ModelManagementProps{
		Model:         models.Party{},
		Title:         "Parties",
		PreloadFields: []string{"Candidates"},
		Headers:       []string{"Name", "Candidates", "Created At"},
		MkRow:         mkRow,
		ModalProps: modelmanagement.ModalProps{Title: "Party", FormItemProps: []modelmanagement.FormItemProps{
			{ID: "name", Label: "Name", Placeholder: "Democratic Party", Type: input.TypeText, Required: true, Description: "The name of the party."},
			{ID: "platform", Label: "Platform", Placeholder: "https://democrats.org/where-we-stand/party-platform/", Type: input.TypeURL, Required: true, Description: "Link to the platform of the party."},
		},
		}}

	mux.HandleFunc("POST /parties", func(w http.ResponseWriter, r *http.Request) {
		cud.CreateModel(w, r, db, mkRow, &models.Party{})
	})

	mux.HandleFunc("GET /parties", func(w http.ResponseWriter, r *http.Request) {
		ModelManagement(w, r, db, modelManagementProps)
	})

	mux.HandleFunc("PUT /parties", func(w http.ResponseWriter, r *http.Request) {
		cud.EditModel(w, r, db, mkRow, &models.Party{})
	})

	mux.HandleFunc("DELETE /parties", func(w http.ResponseWriter, r *http.Request) {
		cud.DeleteModel(w, r, db, &models.Party{})
	})
}

func mkRow(model any) modelmanagement.RowProps {
	party := model.(models.Party)
	candidateCount := strconv.Itoa(len(party.Candidates))
	createdAt := party.CreatedAt.Format("Jan 2, 2006")

	return modelmanagement.RowProps{
		Model: party,
		Cells: []string{
			party.Name,
			candidateCount,
			createdAt,
		}}
}
