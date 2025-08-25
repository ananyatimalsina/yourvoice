package models

import "gorm.io/gorm"

type Candidate struct {
	gorm.Model
	Votes      []Vote       `json:"votes" schema:"votes"`
	VoteEvents []*VoteEvent `json:"vote_events" schema:"vote_events" gorm:"many2many:vote_event_candidates;"`
	PartyID    uint         `json:"party_id" schema:"party_id"`
	Name       string       `json:"name" schema:"name,required" gorm:"not null"`
	Campaign   string       `json:"campaign" schema:"campaign,required" gorm:"not null"`
}
