package models

import "yourvoice/internal/utils"

type Message struct {
	utils.Expression
	MessageEventID string `json:"message_event_id" gorm:"not null;unique"`
	Message        string `json:"message" gorm:"not null"`
}
