// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameNpc = "npc"

// Npc mapped from table <npc>
type Npc struct {
	ID           string `gorm:"column:id;type:varchar(255);primaryKey" json:"id"`          // npc的id
	Name         string `gorm:"column:name;type:varchar(255)" json:"name"`                 // 名称
	Image        string `gorm:"column:image;type:varchar(255)" json:"image"`               // 头像
	Introduction string `gorm:"column:introduction;type:varchar(255)" json:"introduction"` // 自我介绍
	Horizon      int64  `gorm:"column:horizon;type:int" json:"horizon"`                    // 横坐标
	Ordinate     int64  `gorm:"column:ordinate;type:int" json:"ordinate"`                  // 纵坐标
	MapID        string `gorm:"column:map_id;type:varchar(255)" json:"map_id"`             // 地图id
	Type         int64  `gorm:"column:type;type:tinyint" json:"type"`                      // 类型
}

// TableName Npc's table name
func (*Npc) TableName() string {
	return TableNameNpc
}
