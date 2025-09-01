package models

import "yourvoice/internal/utils"

type Feedback struct {
	utils.Expression
	FeedbackSessionID uint   `json:"feedback_session_id" schema:"feedback_session_id,required" gorm:"not null;unique"`
	Message           string `json:"message" schema:"message,required" gorm:"not null"`
}
