package models

import (
	"encoding/json"
	"github.com/ananyatimalsina/schema"
	"gorm.io/gorm"
	"reflect"
)

type Candidate struct {
	gorm.Model
	Votes      []Vote       `json:"votes" schema:"votes"`
	VoteEvents []*VoteEvent `json:"vote_events" schema:"vote_events" gorm:"many2many:vote_event_candidates;"`
	PartyID    uint         `json:"party_id" schema:"party_id"`
	Name       string       `json:"name" schema:"name,required" gorm:"not null"`
	Campaign   string       `json:"campaign" schema:"campaign,required" gorm:"not null"`
}

func MigrateCandidate(db *gorm.DB, decoder *schema.Decoder) error {
	decoder.RegisterConverter([]Candidate{}, func(s string) reflect.Value {
		if s == "" {
			return reflect.ValueOf([]Candidate{})
		}
		var candidates []Candidate
		err := json.Unmarshal([]byte(s), &candidates)
		if err != nil {
			return reflect.Value{} // invalid
		}
		return reflect.ValueOf(candidates)
	})

	return db.AutoMigrate(&Candidate{})
}
