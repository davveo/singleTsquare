package models

import (
	"database/sql"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	gorm.Model
	Uuid     uuid.UUID      `json:"uuid" gorm: "comment: '用户uuid'"`
	Username string         `json:"userName" gorm:"comment:'用户登录名'"`
	Password sql.NullString `json:"-"  gorm:"comment:'用户登录密码'"`
	Phone    string         `json:"phone" gorm:"comment:'用户昵称'" `
}
