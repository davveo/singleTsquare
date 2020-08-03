package models

type Tag struct {
	BaseModel
	ZhName string `json:"zh_name" gorm:"column:zh_name;not null;unique;comment:'标签中文名'"`
	EnName string `json:"en_name" gorm:"column:en_name;not null;unique;comment:'标签英文名'"`
	IsHot  int    `json:"-"  gorm:"default:0;comment:'是否为热门标签 0:不是, 1:是'"`
}

func (Tag) TableName() string {
	return "tb_tag"
}
