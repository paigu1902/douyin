package models

import (
	"gorm.io/gorm"
)

type UserFavo struct {
	gorm.Model
	UserId   int64  `gorm:"not null;default:0"` //用户id
	VideoId  int64  `gorm:"not null;default:0"` //点赞视频id
	UserName string `gorm:"not null"`           //用户名
	Status   int32  `gorm:"not null;default:1"` //点赞状态 默认1点赞 0取消
}
