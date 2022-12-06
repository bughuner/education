// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameMonster = "monster"

// Monster mapped from table <monster>
type Monster struct {
	ID              string `gorm:"column:id;type:varchar(255);primaryKey" json:"id"`         // 怪物的id
	Type            int64  `gorm:"column:type;type:tinyint" json:"type"`                     // 类型
	Name            string `gorm:"column:name;type:varchar(255)" json:"name"`                // 名称
	Level           int64  `gorm:"column:level;type:int" json:"level"`                       // 等级
	ExperienceValue int64  `gorm:"column:experience_value;type:int" json:"experience_value"` // 经验值
	Coin            int64  `gorm:"column:coin;type:int" json:"coin"`                         // 金币
	Hp              int64  `gorm:"column:hp;type:int" json:"hp"`                             // 生命值
	Image           string `gorm:"column:image;type:varchar(255)" json:"image"`              // 头像
}

// TableName Monster's table name
func (*Monster) TableName() string {
	return TableNameMonster
}
