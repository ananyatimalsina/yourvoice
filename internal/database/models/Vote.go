package models

import "yourvoice/internal/utils"

type Vote struct {
	utils.Expression
	VoteEventID string `json:"vote_event_id" gorm:"not null;unique"`
	CandidateID string `json:"candidate_id" gorm:"not null"`
}
