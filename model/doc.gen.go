// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameDoc = "doc"

// Doc mapped from table <doc>
type Doc struct {
	ID      string `gorm:"column:id;type:varchar(255);primaryKey" json:"id"` // 文档
	Link    string `gorm:"column:link;type:varchar(255)" json:"link"`        // 链接
	Author  string `gorm:"column:author;type:varchar(255)" json:"author"`    // 作者
	Content string `gorm:"column:content;type:longtext" json:"content"`      // 内容
	Status  int64  `gorm:"column:status;type:tinyint" json:"status"`         // 状态 0-待审核 1-审核
	Type    string `gorm:"column:type;type:varchar(255)" json:"type"`        // 类型
	Label   string `gorm:"column:label;type:varchar(255)" json:"label"`      // 标签
	Title   string `gorm:"column:title;type:varchar(255)" json:"title"`      // 标题
}

// TableName Doc's table name
func (*Doc) TableName() string {
	return TableNameDoc
}
