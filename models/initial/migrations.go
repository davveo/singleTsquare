package initial

import (
	"fmt"

	"github.com/davveo/singleTsquare/models"
	"github.com/davveo/singleTsquare/utils/migrations"
	"github.com/jinzhu/gorm"
)

var (
	list = []migrations.MigrationStage{
		{
			Name:     "initial",
			Function: migrate0001,
		},
	}
)

func MigrateAll(db *gorm.DB) error {
	return migrations.Migrate(db, list)
}

func migrate0001(db *gorm.DB, name string) error {
	if err := db.CreateTable(new(models.Account)).Error; err != nil {
		return fmt.Errorf("Error creating AccountUser table: %s", err)
	}

	if err := db.CreateTable(new(models.User)).Error; err != nil {
		return fmt.Errorf("Error creating Member table: %s", err)
	}

	//err := db.Model(new(OauthUser)).AddForeignKey(
	//	"role_id", "oauth_roles(id)",
	//	"RESTRICT", "RESTRICT",
	//).Error
	//if err != nil {
	//	return fmt.Errorf("Error creating foreign key on "+
	//		"oauth_users.role_id for oauth_roles(id): %s", err)
	//}
	return nil
}
