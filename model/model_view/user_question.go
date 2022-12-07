package model_view

import "education/model"

type GetUserQuestionReq struct {
	ID       string `gorm:"column:id;type:varchar(255);primaryKey" json:"id"` // 玩家问题表id
	UserID   string `gorm:"column:user_id;type:varchar(255)" json:"user_id"`  // 用户id
	PageNo   int    `json:"page_no"`
	PageSize int    `json:"page_size"`
}

type GetUserQuestionResp struct {
	Total int64 `json:"total"`
	Data  []*UserQuestionRes
}

type UserQuestionRes struct {
	UserQuestion *model.UserQuestion
	Data         *QuestionResult
}
