package model_view

import "education/model"

type GetUserTaskReq struct {
	ID         string `gorm:"column:id;type:varchar(255);primaryKey" json:"id"`   // 用户持有任务id
	UserID     string `gorm:"column:user_id;type:varchar(255)" json:"user_id"`    // 用户id
	IsFinished int64  `gorm:"column:is_finished;type:tinyint" json:"is_finished"` // 是否完成
	PageNo     int    `json:"page_no"`
	PageSize   int    `json:"page_size"`
}

type GetUserTaskResp struct {
	Total int64 `json:"total"`
	Data  []*UserTaskResp
}

type UserTaskResp struct {
	UserTask *model.UserTask
	Task     *model.Task
}
