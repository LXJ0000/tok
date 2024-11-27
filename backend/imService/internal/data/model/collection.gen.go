// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameCollection = "collection"

// Collection 收藏夹
type Collection struct {
	ID          int64     `gorm:"column:id;primaryKey;comment:收藏夹ID" json:"id"`                                          // 收藏夹ID
	UserID      int64     `gorm:"column:user_id;not null;comment:用户ID" json:"user_id"`                                   // 用户ID
	Title       string    `gorm:"column:title;not null;comment:标题" json:"title"`                                         // 标题
	Description string    `gorm:"column:description;not null;comment:描述" json:"description"`                             // 描述
	IsDeleted   int32     `gorm:"column:is_deleted;not null;comment:是否删除" json:"is_deleted"`                             // 是否删除
	CreateTime  time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"` // 创建时间
	UpdateTime  time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_time"` // 更新时间
}

// TableName Collection's table name
func (*Collection) TableName() string {
	return TableNameCollection
}