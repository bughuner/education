package model_view

import "education/model"

type GetUserReq struct {
	ID       string `gorm:"column:id;type:varchar(255);primaryKey" json:"id"` // 用户id
	Account  string `gorm:"column:account;type:varchar(255)" json:"account"`  // 账号
	Name     string `gorm:"column:name;type:varchar(255)" json:"name"`        // 姓名
	Role     string `gorm:"column:role;type:varchar(255)" json:"role"`        // 角色  用户/管理员
	Level    int64  `gorm:"column:level;type:int" json:"level"`               // 等级
	Career   string `gorm:"column:career;type:varchar(255)" json:"career"`    // 职业
	MapID    string `gorm:"column:map_id;type:varchar(255)" json:"map_id"`    // 地图id
	PageNo   int    `json:"page_no"`
	PageSize int    `json:"page_size"`
}

type GetUserResp struct {
	Total int64 `json:"total"`
	Data  []*model.User
}
