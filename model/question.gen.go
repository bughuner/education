// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameQuestion = "question"

// Question mapped from table <question>
type Question struct {
	ID       string `gorm:"column:id;type:varchar(255);primaryKey" json:"id"`  // 问答题id
	Level    int64  `gorm:"column:level;type:int" json:"level"`                // 题目等级
	Type     int64  `gorm:"column:type;type:int" json:"type"`                  // 题目分类 1-判断 2-单选 3-多选
	Question string `gorm:"column:question;type:varchar(255)" json:"question"` // 问题
	Damage   int64  `gorm:"column:damage;type:int" json:"damage"`              // 题目伤害
	Status   int64  `gorm:"column:status;type:tinyint" json:"status"`          // 状态 0-待审核 1-审核
}

// TableName Question's table name
func (*Question) TableName() string {
	return TableNameQuestion
}
