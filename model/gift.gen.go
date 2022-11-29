// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameGift = "gift"

// Gift mapped from table <gift>
type Gift struct {
	ID           string `gorm:"column:id;type:varchar(255);primaryKey" json:"id"`          // 礼物的id
	Name         string `gorm:"column:name;type:varchar(255)" json:"name"`                 // 礼物名称
	Introduction string `gorm:"column:introduction;type:varchar(255)" json:"introduction"` // 礼物介绍
	Image        string `gorm:"column:image;type:varchar(255)" json:"image"`               // 图像链接
	Coin         int64  `gorm:"column:coin;type:int" json:"coin"`                          // 价值金币数
}

// TableName Gift's table name
func (*Gift) TableName() string {
	return TableNameGift
}