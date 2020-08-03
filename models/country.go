package models

type Country struct {
	ZhName string
	EnName string
}

func (Country) TableName() string {
	return "tb_country"
}
