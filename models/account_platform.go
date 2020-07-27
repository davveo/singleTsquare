package models

/*
第三方登录授权后，存储用户信息:
	用户在登录授权后，会要求进行绑定绑定手机号或者邮箱。这样的话，accout表与accountplatform表是一对多的关系。
	如果用户绑定的邮箱或者手机号不存在，则会创建用户。
*/
type AccountPlatform struct {
	BaseModel
	Uid          uint   `json:"uid" gorm:"not null;default'0';comment:'账号id'"`
	IdentifyId   uint   `json:"identify_id" gorm:"not null;unique;comment:'平台唯一id'"`
	Accesstoken  string `json:"access_token" gorm:"varchar(255);default:'';comment:'平台access_token'"`
	NickName     string `json:"nickname" gorm:"column:nickname;default:'';comment:'昵称'"`
	Avatar       string `json:"avatar" gorm:"varchar(255);default:'';comment:'头像(相对路径)'"`
	PlatformType uint   `json:"platform_type" gorm:"enum(0, 1, 2, 3, 4);default:'0';comment:'平台类型 0:未知,1:qq,2:wechat,3:weibo,4:github'"`
}
