package model_view

import "education/model"

type TaskReq struct {
	Type       int64  `gorm:"column:type;type:int" json:"type"`                  // 任务类型 1-打怪类 2-阅读任务 3-完成文章的题目4-附件任务
	Level      int64  `gorm:"column:level;type:int" json:"level"`                // 任务等级
	Image      string `gorm:"column:image;type:varchar(255)" json:"image"`       // 图片链接
	Experience int64  `gorm:"column:experience;type:int" json:"experience"`      // 经验奖励
	Coin       int64  `gorm:"column:coin;type:int" json:"coin"`                  // 金币奖励
	Num        int64  `gorm:"column:num;type:int" json:"num"`                    // 时间/打怪的个数
	PreTask    string `gorm:"column:pre_task;type:varchar(255)" json:"pre_task"` // 前序任务id
	PageNo     int    `json:"page_no"`
	PageSize   int    `json:"page_size"`
}

type TaskResp struct {
	Total int64 `json:"total"`
	Data  []*model.Task
}

type DeleteTaskReq struct {
	TaskIds []string `json:"task_ids"`
}
