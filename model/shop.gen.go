// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameShop = "shop"

// Shop mapped from table <shop>
type Shop struct {
	ID     string `gorm:"column:id;type:varchar(255);primaryKey" json:"id"` // 商城id
	GiftID string `gorm:"column:gift_id;type:varchar(255)" json:"gift_id"`  // 礼物id
	Count  int64  `gorm:"column:count;type:int" json:"count"`               // 礼物数量
}

// TableName Shop's table name
func (*Shop) TableName() string {
	return TableNameShop
}
