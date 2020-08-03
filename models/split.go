package models

type Split struct {
	UserId  string
	Content string
}

func (Split) TableName() string {
	return "tb_split"
}
