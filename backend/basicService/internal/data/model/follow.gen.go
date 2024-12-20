// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameFollow = "follow"

// Follow mapped from table <follow>
type Follow struct {
	ID           int64     `gorm:"column:id;primaryKey;comment:关注ID" json:"id"`                                           // 关注ID
	UserID       int64     `gorm:"column:user_id;not null;comment:用户ID" json:"user_id"`                                   // 用户ID
	TargetUserID int64     `gorm:"column:target_user_id;not null;comment:被关注的用户ID" json:"target_user_id"`                 // 被关注的用户ID
	IsDeleted    int32     `gorm:"column:is_deleted;not null;comment:是否删除" json:"is_deleted"`                             // 是否删除
	CreateTime   time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"` // 创建时间
	UpdateTime   time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_time"` // 更新时间
}

// TableName Follow's table name
func (*Follow) TableName() string {
	return TableNameFollow
}
