package database

import (
	"fmt"
	"github.com/ananyatimalsina/schema"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"yourvoice/internal/database/models"
)

func LoadDatabase(decoder *schema.Decoder) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PORT"), os.Getenv("TZ"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database, " + err.Error())
	}

	migrateModels(db, decoder)

	return db
}

func migrateModels(db *gorm.DB, decoder *schema.Decoder) {
	models.MigrateCandidate(db, decoder)
	err := db.AutoMigrate(
		&models.Party{},
		&models.FeedbackSession{},
	)
	models.MigrateVote(db, decoder)
	models.MigrateVoteEvent(db, decoder)
	models.MigrateFeedback(db, decoder)

	if err != nil {
		panic("failed to migrate models: " + err.Error())
	}
}
