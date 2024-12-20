// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameComment = "comment"

// Comment mapped from table <comment>
type Comment struct {
	ID            int64     `gorm:"column:id;primaryKey;comment:评论ID" json:"id"`                                           // 评论ID
	VideoID       int64     `gorm:"column:video_id;not null;comment:视频ID" json:"video_id"`                                 // 视频ID
	UserID        int64     `gorm:"column:user_id;not null;comment:发表评论的用户ID" json:"user_id"`                              // 发表评论的用户ID
	ParentID      int64     `gorm:"column:parent_id;comment:父评论ID" json:"parent_id"`                                       // 父评论ID
	ToUserID      int64     `gorm:"column:to_user_id;comment:评论所回复的用户ID" json:"to_user_id"`                                // 评论所回复的用户ID
	Content       string    `gorm:"column:content;not null;comment:评论内容" json:"content"`                                   // 评论内容
	FirstComments string    `gorm:"column:first_comments;not null;comment:最开始的几条子评论" json:"first_comments"`                // 最开始的几条子评论
	IsDeleted     int32     `gorm:"column:is_deleted;not null;comment:是否删除" json:"is_deleted"`                             // 是否删除
	CreateTime    time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"` // 创建时间
	UpdateTime    time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_time"` // 更新时间
}

// TableName Comment's table name
func (*Comment) TableName() string {
	return TableNameComment
}
