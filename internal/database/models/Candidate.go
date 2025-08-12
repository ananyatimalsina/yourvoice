package models

import "gorm.io/gorm"

type Candidate struct {
	gorm.Model
	Votes      []Vote       `json:"votes"`
	VoteEvents []*VoteEvent `json:"vote_events" gorm:"many2many:vote_event_candidates;"`
	PartyID    uint         `json:"party_id" gorm:"not null"`
	Name       string       `json:"name" gorm:"not null"`
	Campaign   string       `json:"campaign" gorm:"not null"`
}
