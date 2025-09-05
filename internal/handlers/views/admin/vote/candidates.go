package admin_vote

import (
	"context"
	"github.com/a-h/templ"
	"github.com/ananyatimalsina/schema"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
	"yourvoice/internal/database/models"
	"yourvoice/internal/handlers/routes/modelmanagement"
	"yourvoice/internal/handlers/views/admin"
	"yourvoice/internal/utils"
	"yourvoice/web/templates/admin/components"
)

// SafeVoteEvent represents a VoteEvent without sensitive information
type SafeCandidate struct {
	ID         uint                `json:"id"`
	CreatedAt  time.Time           `json:"created_at"`
	UpdatedAt  time.Time           `json:"updated_at"`
	VoteEvents []*models.VoteEvent `json:"vote_events"`
	PartyID    uint                `json:"party_id"`
	Name       string              `json:"name"`
	Campaign   string              `json:"campaign"`
}

func sanitizeCandidate(candidate models.Candidate) SafeCandidate {
	return SafeCandidate{
		ID:         candidate.ID,
		CreatedAt:  candidate.CreatedAt,
		UpdatedAt:  candidate.UpdatedAt,
		VoteEvents: candidate.VoteEvents,
		PartyID:    candidate.PartyID,
		Name:       candidate.Name,
		Campaign:   candidate.Campaign,
	}
}

func RegisterCandidateRoutes(mux *http.ServeMux, db *gorm.DB, decoder *schema.Decoder) {
	ctx := context.Background()

	modelManagementProps := admin.ModelManagerProps{
		Model:         models.Candidate{},
		SafeModel:     SafeCandidate{},
		SearchFields:  []string{"name"},
		PreloadFields: []string{"Votes"},
		Headers:       []string{"Candidate Name", "Party", "Votes", "Created At", "Status"},
		MkRow: func(model any) components.RowProps {
			candidate := model.(models.Candidate)

			safeCandidate := sanitizeCandidate(candidate)
			voteCount := strconv.Itoa(len(candidate.Votes))
			createdAt := safeCandidate.CreatedAt.Format("Jan 2, 2006")
			status := getCandidateStatus(safeCandidate.VoteEvents)
			party, err := gorm.G[models.Party](db).Where("id = ?", safeCandidate.PartyID).First(ctx)
			if err != nil {
				party = models.Party{Name: "Unknown"}
			}

			return components.RowProps{Model: safeCandidate, Cells: []string{
				safeCandidate.Name,
				party.Name,
				voteCount,
				createdAt,
				status,
			}}
		},
		Title:         "Candidates",
		SingularTitle: "Candidate",
		Icon:          "👥",
		RelationshipFields: []components.RelationshipField{{
			Name:    "party_id",
			Label:   "Party",
			Type:    utils.InputTypeSelect,
			Options: utils.BuildRelationshipFieldInputOptions(db, models.Party{}),
		},
			{
				Name:    "vote_events",
				Label:   "Vote Events",
				Type:    utils.InputTypeMultiple,
				Options: utils.BuildRelationshipFieldInputOptions(db, models.VoteEvent{}),
			},
		},
		Actions: []templ.Component{},
		Options: [2]bool{true, true},
	}

	mux.HandleFunc("POST /candidates", func(w http.ResponseWriter, r *http.Request) {
		modelmanagement.CreateModel(w, r, db, decoder, &models.Candidate{}, modelManagementProps.MkRow, modelManagementProps.Actions, modelManagementProps.Options)
	})

	mux.HandleFunc("GET /candidates", func(w http.ResponseWriter, r *http.Request) {
		admin.ModelManager(w, r, db, modelManagementProps)
	})

	mux.HandleFunc("PUT /candidates", func(w http.ResponseWriter, r *http.Request) {
		modelmanagement.EditModel(w, r, db, decoder, &models.Candidate{}, modelManagementProps.MkRow, modelManagementProps.Actions, modelManagementProps.Options)
	})

	mux.HandleFunc("DELETE /candidates", func(w http.ResponseWriter, r *http.Request) {
		modelmanagement.DeleteModel(w, r, db, &models.Candidate{})
	})

}

func getCandidateStatus(events []*models.VoteEvent) string {
	// Check if any event is active
	for _, event := range events {
		if GetEventStatus(*event) == "Ongoing" || GetEventStatus(*event) == "Upcoming" {
			return "Active"
		}
	}

	return "Inactive"
}
