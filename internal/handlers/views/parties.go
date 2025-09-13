package views

import (
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"yourvoice/internal/database/models"
	"yourvoice/internal/handlers/routes/cud"
	"yourvoice/web/templates/components/input"
	"yourvoice/web/templates/modelmanagement"
)

type SafeParty struct {
	models.Party
	Candidates     int `json:"candidate_count"`
	CandidateCount int `json:"-"`
}

func RegisterPartyRoutes(mux *http.ServeMux, db *gorm.DB) {
	modelManagementProps := ModelManagementProps{
		Model:        models.Party{},
		SafeModel:    SafeParty{},
		Title:        "Parties",
		SearchFields: []string{"name"},
		Headers:      []string{"Name", "Candidates", "Created At"},
		PrepareDB:    prepareDB,
		MkRow:        mkRow,
		ModalProps: modelmanagement.ModalProps{Title: "Party", FormItemProps: []modelmanagement.FormItemProps{
			{ID: "name", Label: "Name", Placeholder: "Sunshine Party", Type: input.TypeText, Required: true, Unique: true, Description: "The name of the party."},
			{ID: "platform", Label: "Platform", Placeholder: "https://sunshine.example.com/platform", Type: input.TypeURL, Required: true, Unique: false, Description: "Link to the platform of the party."},
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

func prepareDB(db *gorm.DB) *gorm.DB {
	return db.Model(&models.Party{}).
		Joins("LEFT JOIN candidates ON candidates.party_id = parties.id").
		Select("parties.*, COUNT(candidates.id) as candidate_count").
		Group("parties.id")
}

func mkRow(model any) modelmanagement.RowProps {
	party := model.(SafeParty)
	candidateCount := strconv.Itoa(party.CandidateCount)
	createdAt := party.CreatedAt.Format("Jan 2, 2006")

	return modelmanagement.RowProps{
		Model: party,
		Cells: []string{
			party.Name,
			candidateCount,
			createdAt,
		}}
}
