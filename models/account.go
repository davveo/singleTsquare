package models

import (
	"database/sql"
	"time"
)

type Account struct {
	BaseModel
	UserName      string         `json:"username" gorm:"column:username;not null;unique;comment:'用户登录名'"`
	PassWord      sql.NullString `json:"-"  gorm:"column:password;not null;comment:'用户登录密码'"`
	Email         string         `json:"-"  gorm:"default:'';comment:'用户邮箱'"`
	CreateIpAt    string         `json:"-"  gorm:"comment:'创建ip'"`
	LastLoginIpAt string         `json:"-"  gorm:"comment:'最后登录ip'"`
	LastLoginAt   time.Time      `json:"-"  gorm:"comment:'最后一次登录时间'"`
	LoginTimes    uint           `json:"-"  gorm:"default:0;comment:'登录次数'"`
	Status        int            `json:"-"  gorm:"default:0;comment:'状态 1:enable, 0:disable, -1:deleted'"`
	Phone         string         `json:"phone" gorm:"not null;unique;comment:'用户手机号'"`
}
