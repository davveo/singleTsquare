package models

type QuestionTag struct {
	TagId      string
	QuestionId string
}

func (QuestionTag) TableName() string {
	return "tb_question_tag"
}
