package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"yourvoice/internal/database/models"
)

func LoadDatabase() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), os.Getenv("TZ"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}

	migrateModels(db)

	return db
}

func migrateModels(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.Candidate{},
		&models.Party{},
		&models.Vote{},
		&models.VoteEvent{},
		&models.Message{},
		&models.MessageEvent{},
	)

	if err != nil {
		panic("failed to migrate models: " + err.Error())
	}
}
