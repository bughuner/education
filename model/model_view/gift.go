package model_view

import "education/model"

type GetGiftReq struct {
	ID       string `gorm:"column:id;type:varchar(255);primaryKey" json:"id"` // 礼物的id
	Name     string `gorm:"column:name;type:varchar(255)" json:"name"`        // 礼物名称
	Coin     int64  `gorm:"column:coin;type:int" json:"coin"`                 // 价值金币数
	PageNo   int    `json:"page_no"`
	PageSize int    `json:"page_size"`
}

type GetGiftResp struct {
	Total int64 `json:"total"`
	Data  []*model.Gift
}

type DeleteGiftReq struct {
	GiftIds []string `json:"gift_ids"`
}
