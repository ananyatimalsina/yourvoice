package models

import "yourvoice/internal/utils"

type FeedbackSession struct {
	utils.Event
	Feedback []Feedback `json:"feedback" schema:"feedback"`
}
