package models

import (
	"encoding/json"
	"github.com/ananyatimalsina/schema"
	"gorm.io/gorm"
	"reflect"
	"yourvoice/internal/utils"
)

type VoteEvent struct {
	utils.Event
	Votes      []Vote       `json:"votes" schema:"votes"`
	Candidates []*Candidate `json:"candidates" schema:"candidates" gorm:"many2many:vote_event_candidates;"`
}

func MigrateVoteEvent(db *gorm.DB, decoder *schema.Decoder) error {
	decoder.RegisterConverter([]VoteEvent{}, func(s string) reflect.Value {
		if s == "" {
			return reflect.ValueOf([]Candidate{})
		}
		var voteEvents []VoteEvent
		err := json.Unmarshal([]byte(s), &voteEvents)
		if err != nil {
			return reflect.Value{} // invalid
		}
		return reflect.ValueOf(voteEvents)
	})

	return db.AutoMigrate(&VoteEvent{})
}
