package model_view

import "education/model"

type GetNpcTaskReq struct {
	ID       string `gorm:"column:id;type:varchar(255);primaryKey" json:"id"` // npc持有任务id
	NpcID    string `gorm:"column:npc_id;type:varchar(255)" json:"npc_id"`    // npc id
	PageNo   int    `json:"page_no"`
	PageSize int    `json:"page_size"`
}

type GetNpcTaskResp struct {
	Total int64 `json:"total"`
	Data  []*model.NpcTask
}