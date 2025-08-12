package models

import "yourvoice/internal/utils"

type MessageEvent struct {
	utils.Event
	Messages []Message `json:"messages"`
}
