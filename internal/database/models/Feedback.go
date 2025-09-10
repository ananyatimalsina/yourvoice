package models

import "yourvoice/internal/utils"

type Feedback struct {
	utils.Expression
	FeedbackSessionID uint   `json:"feedback_session_id" validate:"required" gorm:"not null"`
	Message           string `json:"message" validate:"required" gorm:"not null"`
}
