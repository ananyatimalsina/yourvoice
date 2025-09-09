package models

import "gorm.io/gorm"

type Party struct {
	gorm.Model
	Candidates []Candidate `json:"candidates"`
	Name       string      `json:"name" gorm:"not null"`
	Platform   string      `json:"platform" gorm:"not null"`
}
