package migrations

import "github.com/jinzhu/gorm"

type Migration struct {
	gorm.Model
	Name string `sql: "size:255"`
}
