package model_view

import "education/model"

type ExchangeGiftReq struct {
	UserID   string `gorm:"column:user_id;type:varchar(255)" json:"user_id"` // 用户id
	GiftID   string `gorm:"column:gift_id;type:varchar(255)" json:"gift_id"` // 物品id
	Count    int64  `gorm:"column:count;type:int" json:"count"`              // 兑换数量
	PageNo   int    `json:"page_no"`
	PageSize int    `json:"page_size"`
}

type ExchangeGiftAndGift struct {
	ExchangeGift *model.ExchangeGift
	Gift         *model.Gift
}

type ExchangeGiftResp struct {
	Data  []*ExchangeGiftAndGift
	Total int64 `json:"total"`
}
