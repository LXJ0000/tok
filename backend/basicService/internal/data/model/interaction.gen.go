// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameInteraction = "interaction"

// Interaction mapped from table <interaction>
type Interaction struct {
	ID         int64     `gorm:"column:id;primaryKey" json:"id"`
	BizID      int64     `gorm:"column:biz_id" json:"biz_id"`
	Biz        string    `gorm:"column:biz" json:"biz"`
	ReadCnt    int64     `gorm:"column:read_cnt;not null" json:"read_cnt"`
	LikeCnt    int64     `gorm:"column:like_cnt;not null" json:"like_cnt"`
	CollectCnt int64     `gorm:"column:collect_cnt;not null" json:"collect_cnt"`
	IsDeleted  int32     `gorm:"column:is_deleted;not null;comment:是否删除" json:"is_deleted"`                             // 是否删除
	CreateTime time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_time"` // 更新时间
}

// TableName Interaction's table name
func (*Interaction) TableName() string {
	return TableNameInteraction
}
