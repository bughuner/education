package model_view

import "education/model"

type GetPlatReq struct {
	ID       string `gorm:"column:id;type:varchar(255);primaryKey" json:"id"`
	Name     string `gorm:"column:name;type:varchar(255)" json:"name"`   // 名称
	Image    string `gorm:"column:image;type:varchar(255)" json:"image"` // 图片
	Width    int64  `gorm:"column:width;type:int" json:"width"`          // 宽
	Height   int64  `gorm:"column:height;type:int" json:"height"`        // 高
	PageNo   int    `json:"page_no"`
	PageSize int    `json:"page_size"`
}

type GetPlatResp struct {
	Total int64 `json:"total"`
	Data  []*model.Plat
}

type DeletePlatReq struct {
	PlatIds []string `json:"plat_ids"`
}
