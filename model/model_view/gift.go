package model_view

// ShopGiftCount 商店中每种礼物的数量和详细信息
type ShopGiftCount struct {
	ID     string `gorm:"column:id;type:varchar(255);primaryKey" json:"id"` // 商城id
	GiftID string `gorm:"column:gift_id;type:varchar(255)" json:"gift_id"`  // 礼物id
	Count  int64  `gorm:"column:count;type:int" json:"count"`               // 礼物数量
	Name         string `gorm:"column:name;type:varchar(255)" json:"name"`                 // 礼物名称
	Introduction string `gorm:"column:introduction;type:varchar(255)" json:"introduction"` // 礼物介绍
	Image        string `gorm:"column:image;type:varchar(255)" json:"image"`               // 图像链接
	Coin         int64  `gorm:"column:coin;type:int" json:"coin"`                          // 价值金币数
}
