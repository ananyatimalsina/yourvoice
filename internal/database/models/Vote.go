package models

import "yourvoice/internal/utils"

type Vote struct {
	utils.Expression
	VoteEventID uint `json:"vote_event_id" gorm:"not null;unique"`
	CandidateID uint `json:"candidate_id" gorm:"not null"`
}
