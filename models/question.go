package models

type Question struct {
	UserId  string
	Title   string
	Desc    string
	Content string
}

func (Question) TableName() string {
	return "tb_question"
}
