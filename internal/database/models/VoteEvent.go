package models

import "yourvoice/internal/utils"

type VoteEvent struct {
	utils.Event
	Votes      []Vote       `json:"votes", schema:"votes"`
	Candidates []*Candidate `json:"candidates" schema:"candidates" gorm:"many2many:vote_event_candidates;"`
}
