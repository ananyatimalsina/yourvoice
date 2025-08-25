package models

import "gorm.io/gorm"

type Party struct {
	gorm.Model
	Candidates []Candidate `json:"candidates" schema:"candidates"`
	Name       string      `json:"name" schema:"name,required" gorm:"not null"`
	Platform   string      `json:"platform" schema:"platform,required" gorm:"not null"`
}
