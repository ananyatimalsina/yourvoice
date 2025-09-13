package models

import (
	"yourvoice/internal/utils"
)

type Candidate struct {
	utils.Model
	Votes      []Vote       `json:"votes"`
	VoteEvents []*VoteEvent `json:"vote_events" gorm:"many2many:vote_event_candidates"`
	PartyID    uint         `json:"party_id"`
	Name       string       `json:"name" validate:"required" gorm:"not null"`
	Campaign   string       `json:"campaign" validate:"required,url" gorm:"not null"`
}
