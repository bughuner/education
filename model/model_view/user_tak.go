package model_view

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
	ID           string `gorm:"column:id;type:varchar(255);primaryKey" json:"id"`          // 用户持有任务id
	UserID       string `gorm:"column:user_id;type:varchar(255)" json:"user_id"`           // 用户id
	TaskID       string `gorm:"column:task_id;type:varchar(255)" json:"task_id"`           // 任务id
	Type         int64  `gorm:"column:type;type:tinyint" json:"type"`                      // 任务类型
	IsFinished   int64  `gorm:"column:is_finished;type:tinyint" json:"is_finished"`        // 是否完成
	Count        int64  `gorm:"column:count;type:int" json:"count"`                        // 个数
	Level        int64  `gorm:"column:level;type:int" json:"level"`                        // 任务等级
	Introduction string `gorm:"column:introduction;type:varchar(255)" json:"introduction"` // 任务描述
	Image        string `gorm:"column:image;type:varchar(255)" json:"image"`               // 图片链接
	Experience   int64  `gorm:"column:experience;type:int" json:"experience"`              // 经验奖励
	Coin         int64  `gorm:"column:coin;type:int" json:"coin"`                          // 金币奖励
	Num          int64  `gorm:"column:num;type:int" json:"num"`                            // 时间/打怪的个数
	PreTask      string `gorm:"column:pre_task;type:varchar(255)" json:"pre_task"`         // 前序任务id
	TargetId     string `json:"target_id"`
}
