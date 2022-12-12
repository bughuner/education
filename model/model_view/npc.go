package model_view

import "education/model"

type GetNpcReq struct {
	ID       string `gorm:"column:id;type:varchar(255);primaryKey" json:"id"` // npc的id
	Name     string `gorm:"column:name;type:varchar(255)" json:"name"`        // 名称
	PlatId   string `gorm:"column:plat_id;type:varchar(255)" json:"plat_id"`  // 地图id
	Type     int64  `gorm:"column:type;type:tinyint" json:"type"`             // 类型
	PageNo   int    `json:"page_no"`
	PageSize int    `json:"page_size"`
}

type GetNpcResp struct {
	Total int64 `json:"total"`
	Data  []*model.Npc
}

type DeleteNpcReq struct {
	NpcIds []string `json:"npc_ids"`
}
