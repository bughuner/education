package model_view

import "education/model"

type QuestionReq struct {
	ID       string   `gorm:"column:id;type:varchar(255);primaryKey" json:"id"`  // 问答题id
	Level    int64    `gorm:"column:level;type:int" json:"level"`                // 题目等级
	Type     int64    `gorm:"column:type;type:int" json:"type"`                  // 题目分类
	Question string   `gorm:"column:question;type:varchar(255)" json:"question"` // 问题
	Damage   int64    `gorm:"column:damage;type:int" json:"damage"`              // 题目伤害
	Select   []string `json:"select"`                                            // 选项内容
	Answer   []string `json:"answer"`                                            // 答案
	PageNo   int      `json:"page_no"`
	PageSize int      `json:"page_size"`
}

type QuestionResp struct {
	Total int64 `json:"total"`
	Data  []*QuestionResult
}

type QuestionResult struct {
	Question       *model.Question
	QuestionSelect []*model.QuestionSelect
	QuestionAnswer []*model.QuestionAnswer
}
