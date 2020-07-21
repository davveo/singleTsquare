package migrations

import (
	"fmt"

	log "github.com/davveo/singleTsquare/utils"
	"github.com/jinzhu/gorm"
)

type MigrationStage struct {
	Name     string
	Function func(db *gorm.DB, name string) error
}

func Migrate(db *gorm.DB, migrations []MigrationStage) error {
	for _, m := range migrations {
		if MigrationExists(db, m.Name) {
			continue
		}

		if err := m.Function(db, m.Name); err != nil {
			return err
		}

		if err := SaveMigration(db, m.Name); err != nil {
			return err
		}
	}
	return nil
}

func MigrateAll(db *gorm.DB, migrationFunctions []func(*gorm.DB) error) {
	if err := Bootstrap(db); err != nil {
		log.ERROR.Print(err)
	}

	for _, m := range migrationFunctions {
		if err := m(db); err != nil {
			log.ERROR.Print(err)
		}
	}
}

func SaveMigration(db *gorm.DB, migrationName string) error {
	migration := new(Migration)
	migration.Name = migrationName

	if err := db.Create(migration).Error; err != nil {
		log.ERROR.Printf("Error saving record to migrations table: %s", err)
		return fmt.Errorf("Error saving record to migrations table: %s", err)
	}
	return nil
}

func MigrationExists(db *gorm.DB, migrationName string) bool {
	migration := new(Migration)
	found := !db.Where("name = ?", migrationName).First(migration).RecordNotFound()
	if found {
		log.INFO.Printf("Skipping %s migration", migrationName)
	} else {
		log.INFO.Printf("Running %s migration", migrationName)
	}
	return found
}
