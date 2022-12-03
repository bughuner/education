package model_view

import "education/model"

type UserQuestionResp struct {
	UserQuestion *model.UserQuestion
	Data         []*QuestionResp
}
