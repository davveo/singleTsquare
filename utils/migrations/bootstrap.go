package migrations

import (
	"fmt"

	log "github.com/davveo/singleTsquare/utils"
	"github.com/jinzhu/gorm"
)

func Bootstrap(db *gorm.DB) error {
	migrationName := "bootstrap_migrations"

	migration := new(Migration)

	exists := nil == db.Where("name = ?", migrationName).First(migration).Error
	if exists {
		log.INFO.Printf("Skipping %s migration", migrationName)
		return nil
	}
	log.INFO.Printf("Skipping %s migration", migrationName)
	if err := db.CreateTable(new(Migration)).Error; err != nil {
		return fmt.Errorf("Error creating migrations table: %s", db.Error)
	}
	if err := db.Create(migration).Error; err != nil {
		return fmt.Errorf("Error saving record to migrations table: %s", err)
	}

	return nil
}
