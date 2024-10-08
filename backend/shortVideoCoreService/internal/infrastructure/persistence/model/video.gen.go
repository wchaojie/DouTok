// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameVideo = "video"

// Video mapped from table <video>
type Video struct {
	ID           int64     `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true" json:"id"`
	UserID       int64     `gorm:"column:user_id;type:bigint;not null" json:"user_id"`
	Title        string    `gorm:"column:title;type:varchar(50);not null" json:"title"`
	Description  string    `gorm:"column:description;type:varchar(255);not null" json:"description"`
	VideoURL     string    `gorm:"column:video_url;type:varchar(255);not null" json:"video_url"`
	CoverURL     string    `gorm:"column:cover_url;type:varchar(255);not null" json:"cover_url"`
	LikeCount    int64     `gorm:"column:like_count;type:bigint;not null" json:"like_count"`
	CommentCount int64     `gorm:"column:comment_count;type:bigint;not null" json:"comment_count"`
	CreatedAt    time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName Video's table name
func (*Video) TableName() string {
	return TableNameVideo
}
