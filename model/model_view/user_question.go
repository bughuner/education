package model_view

import "education/model"

type UserQuestionResp struct {
	Total int64 `json:"total"`
	Data  []*UserQuestionRes
}

type UserQuestionRes struct {
	UserQuestion *model.UserQuestion
	Data         *QuestionResult
}
