package models

type Count struct {
	SplitId      string `json:"split_id" gorm:"default:0"`
	QuestionId   string `json:"question_id" gorm:"default:0"`
	UpCount      int64  `json:"up_count" gorm:"default:0"`
	CollectCount int64  `json:"collect_count" gorm:"default:0"`
	ShareCount   int64  `json:"share_count" gorm:"default:0"`
	ReplyCount   int64  `json:"reply_count" gorm:"default:0"`
}

func (Count) TableName() string {
	return "tb_count"
}
