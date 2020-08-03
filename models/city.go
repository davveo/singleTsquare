package models

type City struct {
	ZhName    string
	EnName    string
	CountryId string
}

func (City) TableName() string {
	return "tb_city"
}
