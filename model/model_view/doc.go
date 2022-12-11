package model_view

import "education/model"

type GetDocReq struct {
	ID       string `gorm:"column:id;type:varchar(255);primaryKey" json:"id"` // 文档
	Author   string `gorm:"column:author;type:varchar(255)" json:"author"`    // 作者
	PageNo   int    `json:"page_no"`
	PageSize int    `json:"page_size"`
}

type GetDocResp struct {
	Total int64 `json:"total"`
	Data  []*model.Doc
}

type DeleteDocReq struct {
	DocIds []string `json:"doc_ids"`
}
