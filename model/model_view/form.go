package model_view

import "education/model"

type DeleteFromReq struct {
	FormIds []string `json:"form_ids"`
}

type GetFormReq struct {
	ID       string `gorm:"column:id;type:varchar(255);primaryKey" json:"id"`
	Type     int64  `gorm:"column:type;type:tinyint" json:"type"`             // 通知类型
	UserID   string `gorm:"column:user_id;type:varchar(255)" json:"user_id"`  // 用户id
	MapID    string `gorm:"column:map_id;type:varchar(255)" json:"map_id"`    // 地图id
	Content  string `gorm:"column:content;type:varchar(1000)" json:"content"` // 通知内容
	Source   int64  `gorm:"column:source;type:tinyint" json:"source"`         // 来源
	PageNo   int    `json:"page_no"`
	PageSize int    `json:"page_size"`
}

type GetFormResp struct {
	Total int64 `json:"total"`
	Data  []*model.Form
}
