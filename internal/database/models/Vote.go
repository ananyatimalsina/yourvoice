package models

import "yourvoice/internal/utils"

type Vote struct {
	utils.Expression
	VoteEventID uint `json:"vote_event_id" schema:"vote_event_id,required" gorm:"not null;unique"`
	CandidateID uint `json:"candidate_id" schema:"candidate_id,required" gorm:"not null"`
}
