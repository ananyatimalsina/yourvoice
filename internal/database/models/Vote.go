package models

import "yourvoice/internal/utils"

type Vote struct {
	utils.Expression
	VoteEventID uint `json:"vote_event_id" validate:"required" gorm:"not null"`
	CandidateID uint `json:"candidate_id" validate:"required" gorm:"not null"`
}
