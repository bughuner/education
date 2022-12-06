package model_view

import "education/model"

type DeleteMonsterQuestion struct {
	MonsterQuestionIds []string `json:"monster_question_ids"` // 怪物问题ID
}

type GetMonsterQuestionReq struct {
	PageNo    int    `json:"page_no"`
	PageSize  int    `json:"page_size"`
	ID        string `gorm:"column:id;type:varchar(255);primaryKey" json:"id"`      // 怪物问题Id
	MonsterID string `gorm:"column:monster_id;type:varchar(255)" json:"monster_id"` // 怪物id
}

type GetMonsterQuestionResp struct {
	Total int64 `json:"total"`
	Data  []*GetMonsterQuestionRes
}

type GetMonsterQuestionRes struct {
	MonsterQuestion *model.MonsterQuestion
	Data            *QuestionResult
}
