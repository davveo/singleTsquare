package models

type User struct {
	BaseModel
	Uid      uint   `json:"uid" gorm:"not null;unique;comment:'账号id'"`
	NickName string `json:"nickname" gorm:"column:nickname;default:'';comment:'昵称'"`
	Avatar   string `json:"avatar" gorm:"varchar(255);default:'';comment:'头像(相对路径)'"`
	Gender   string `json:"gender" gorm:"enum('male','female','unknow');default:'unknow';comment:'性别'"`
	Role     uint   `json:"role" gorm:"default:'0';comment:'角色 0:普通用户 1:vip'"`
}
