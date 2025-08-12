package models

import "yourvoice/internal/utils"

type VoteEvent struct {
	utils.Event
	Votes      []Vote       `json:"votes"`
	Candidates []*Candidate `json:"candidates" gorm:"many2many:vote_event_candidates;"`
}
