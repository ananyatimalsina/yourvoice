package models

import "gorm.io/gorm"

type Party struct {
	gorm.Model
	Candidates []Candidate `json:"candidates"`
	Name       string      `json:"name" validate:"required" gorm:"not null"`
	Platform   string      `json:"platform" validate:"required,url" gorm:"not null"`
}
