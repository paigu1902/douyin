package models

import (
	"gorm.io/gorm"
	"time"
)

type UserComm struct {
	gorm.Model
	CommId   int64     `gorm:"not null;unique"`    // 评论id
	UserId   int64     `gorm:"not null;"`          // 评论者id
	UserName string    `gorm:"not null;"`          // 评论者
	VideoId  int64     `gorm:"not null"`           // 视频id
	Status   int32     `gorm:"not null;default:1"` // 评论状态 默认1有效
	CommText string    `gorm:"not null"`           // 评论内容
	CommTime time.Time `gorm:"not null"`           //评论时间
}
