package model_view

import "education/model"

type DeleteDocQuestion struct {
	DocQuestionIds []string `json:"doc_question_ids"` // 文章问题id
}

type GetDocQuestionReq struct {
	PageNo   int    `json:"page_no"`
	PageSize int    `json:"page_size"`
	ID       string `json:"id"`     // 文章问题id
	DocId    string `json:"doc_id"` // 文章id
}

type GetDocQuestionResp struct {
	Total int64 `json:"total"`
	Data  []*GetDocQuestionRes
}

type GetDocQuestionRes struct {
	DocQuestion *model.DocQuestion
	Data        *QuestionResult
}
