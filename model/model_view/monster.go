package model_view

import "education/model"

type GetMonsterReq struct {
	ID              string `gorm:"column:id;type:varchar(255);primaryKey" json:"id"`         // 怪物的id
	Type            int64  `gorm:"column:type;type:tinyint" json:"type"`                     // 类型
	Name            string `gorm:"column:name;type:varchar(255)" json:"name"`                // 名称
	Level           int64  `gorm:"column:level;type:int" json:"level"`                       // 等级
	ExperienceValue int64  `gorm:"column:experience_value;type:int" json:"experience_value"` // 经验值
	Coin            int64  `gorm:"column:coin;type:int" json:"coin"`                         // 金币
	Hp              int64  `gorm:"column:hp;type:int" json:"hp"`                             // 生命值
	PlatID          string `gorm:"column:plat_id;type:varchar(255)" json:"plat_id"`          // 地图id
	PageNo          int    `json:"page_no"`
	PageSize        int    `json:"page_size"`
}

type GetMonsterResp struct {
	Total int64 `json:"total"`
	Data  []*model.Monster
}

type DeleteMonsterReq struct {
	MonsterIds []string `json:"monster_ids"`
}
