package models

import "yourvoice/internal/utils"

type Feedback struct {
	utils.Expression
	FeedbackSessionID uint   `json:"feedback_session_id" gorm:"not null;unique"`
	Message           string `json:"message" gorm:"not null"`
}
