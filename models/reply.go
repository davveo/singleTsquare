package models

type SplitRelpy struct {
	UserId  string
	SplitId string
	Content string
}

func (SplitRelpy) TableName() string {
	return "tb_split_reply"
}
