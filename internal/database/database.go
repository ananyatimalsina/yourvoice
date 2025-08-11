package database

import (
	"gorm.io/gorm"
)

func LoadDatabase() *gorm.DB {
	db, err := gorm.Open(nil, &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}

	migrateModels(db)

	return db
}

func migrateModels(db *gorm.DB) {
}
