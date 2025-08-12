package utils

import (
	"gorm.io/gorm"
	"time"
)

type Expression struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Data string `json:"data" gorm:"not null"`
}

type Event struct {
	gorm.Model
	Name       string    `json:"name" gorm:"not null"`
	StartDate  time.Time `json:"start_date" gorm:"not null"`
	EndDate    time.Time `json:"end_date" gorm:"not null"`
	PrivateKey []byte    `json:"private_key" gorm:"not null"`
	PublicKey  []byte    `json:"public_key" gorm:"not null"`
}
