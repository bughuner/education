// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUserQuestion = "user_question"

// UserQuestion mapped from table <user_question>
type UserQuestion struct {
	ID         string `gorm:"column:id;type:varchar(255);primaryKey" json:"id"`        // 玩家问题表id
	UserID     string `gorm:"column:user_id;type:varchar(255)" json:"user_id"`         // 用户id
	QuestionID string `gorm:"column:question_id;type:varchar(255)" json:"question_id"` // 问题id
}

// TableName UserQuestion's table name
func (*UserQuestion) TableName() string {
	return TableNameUserQuestion
}
