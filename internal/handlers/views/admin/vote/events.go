package admin_vote

import (
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
	"yourvoice/internal/database/models"
	"yourvoice/internal/handlers/views/admin"
	"yourvoice/internal/utils"
	"yourvoice/web/templates/admin/components"
)

// SafeVoteEvent represents a VoteEvent without sensitive information
type SafeVoteEvent struct {
	ID         uint                `json:"id"`
	CreatedAt  time.Time           `json:"created_at"`
	UpdatedAt  time.Time           `json:"updated_at"`
	Name       string              `json:"name"`
	StartDate  time.Time           `json:"start_date"`
	EndDate    time.Time           `json:"end_date"`
	Candidates []*models.Candidate `json:"candidates"`
}

func sanitizeVoteEvent(event models.VoteEvent) SafeVoteEvent {
	return SafeVoteEvent{
		ID:         event.ID,
		CreatedAt:  event.CreatedAt,
		UpdatedAt:  event.UpdatedAt,
		Name:       event.Name,
		StartDate:  event.StartDate,
		EndDate:    event.EndDate,
		Candidates: event.Candidates,
	}
}

func Events(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	admin.ModelManager(w, r, db, admin.ModelManagerProps{
		Model:         models.VoteEvent{},
		SafeModel:     SafeVoteEvent{},
		SearchFields:  []string{"name"},
		PreloadFields: []string{"Candidates", "Votes"},
		Headers:       []string{"Event Name", "Election Date", "Status", "Candidates", "Total Votes"},
		MkRow: func(model any) components.RowProps {
			event := model.(models.VoteEvent)

			safeEvent := sanitizeVoteEvent(event)

			electionDate := safeEvent.EndDate.Format("Jan 2, 2006")
			status := GetEventStatus(event)
			candidateCount := strconv.Itoa(len(safeEvent.Candidates))
			voteCount := strconv.Itoa(len(event.Votes))

			return components.RowProps{Model: safeEvent, Cells: []string{
				safeEvent.Name,
				electionDate,
				status,
				candidateCount,
				voteCount,
			}}
		},
		Title:         "Vote Events",
		SingularTitle: "Vote Event",
		Icon:          "🗳️",
		RelationshipFields: []components.RelationshipField{{
			Name:    "candidates",
			Label:   "Candidates",
			Type:    utils.InputTypeMultiple,
			Options: utils.BuildRelationshipFieldInputOptions(db, models.Candidate{}),
		},
		},
	})
}

func GetEventStatus(event models.VoteEvent) string {
	if event.EndDate.Before(time.Now()) {
		return "Ended"
	} else if event.StartDate.After(time.Now()) {
		return "Upcoming"
	}
	return "Ongoing"
}
