package models

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
)

type AccountUser struct {
	gorm.Model
	UserName    string         `json:"username" gorm:"column:username;not null;unique;comment:'用户登录名'"`
	PassWord    sql.NullString `json:"-"  gorm:"column:password;not null;comment:'用户登录密码'"`
	CreateIpAt  string         `json:"-"  gorm:"comment:'创建ip'"`
	LastLoginAt time.Time      `json:"-"  gorm:"comment:'最后一次登录时间'"`
	LoginTimes  uint           `json:"-"  gorm:"default:0;comment:'登录次数'"`
	Status      int            `json:"-"  gorm:"default:0;comment:'状态 1:enable, 0:disable, -1:deleted'"`
	Phone       string         `json:"phone" gorm:"not null;unique;comment:'用户手机号'"`
}

type Member struct {
	gorm.Model
	Uid      uint   `json:"uid" gorm:"not null;unique;comment:'账号id'"`
	NickName string `json:"nickname" gorm:"column:nickname;default:'';comment:'昵称'"`
	Avatar   string `json:"avatar" gorm:"varchar(255);default:'';comment:'头像(相对路径)'"`
	Gender   string `json:"gender" gorm:"enum('male','female','unknow');default:'unknow';comment:'性别'"`
	Role     uint   `json:"role" gorm:"default:'0';comment:'角色 0:普通用户 1:vip'"`
}
