package models

import "yourvoice/internal/utils"

type Message struct {
	utils.Expression
	MessageEventID uint   `json:"message_event_id" gorm:"not null;unique"`
	Message        string `json:"message" gorm:"not null"`
}
