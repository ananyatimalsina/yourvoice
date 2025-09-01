package models

import (
	"encoding/json"
	"github.com/ananyatimalsina/schema"
	"gorm.io/gorm"
	"reflect"
	"yourvoice/internal/utils"
)

type Feedback struct {
	utils.Expression
	FeedbackSessionID uint   `json:"feedback_session_id" schema:"feedback_session_id,required" gorm:"not null;unique"`
	Message           string `json:"message" schema:"message,required" gorm:"not null"`
}

func MigrateFeedback(db *gorm.DB, decoder *schema.Decoder) error {
	decoder.RegisterConverter([]Feedback{}, func(s string) reflect.Value {
		if s == "" {
			return reflect.ValueOf([]Feedback{})
		}
		var feedback []Feedback
		err := json.Unmarshal([]byte(s), &feedback)
		if err != nil {
			return reflect.Value{} // invalid
		}
		return reflect.ValueOf(feedback)
	})

	return db.AutoMigrate(&Feedback{})
}
