package models

type Answer struct {
	BaseModel
	QuestionId string
	UserId     string
	Content    string
}

func (Answer) TableName() string {
	return "tb_answer"
}
