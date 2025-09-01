package models

import (
	"encoding/json"
	"github.com/ananyatimalsina/schema"
	"gorm.io/gorm"
	"reflect"
	"yourvoice/internal/utils"
)

type Vote struct {
	utils.Expression
	VoteEventID uint `json:"vote_event_id" schema:"vote_event_id,required" gorm:"not null;unique"`
	CandidateID uint `json:"candidate_id" schema:"candidate_id,required" gorm:"not null"`
}

func MigrateVote(db *gorm.DB, decoder *schema.Decoder) error {
	decoder.RegisterConverter([]Vote{}, func(s string) reflect.Value {
		if s == "" {
			return reflect.ValueOf([]Vote{})
		}
		var votes []Vote
		err := json.Unmarshal([]byte(s), &votes)
		if err != nil {
			return reflect.Value{} // invalid
		}
		return reflect.ValueOf(votes)
	})

	return db.AutoMigrate(&Vote{})
}
