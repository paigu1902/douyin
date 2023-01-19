package models

import "gorm.io/gorm"

type VideoInfo struct {
	gorm.Model
	Id            int64  `gorm:"noe null;unique"`    //视频Id
	AuthorId      int64  `gorm:"not null"`           //作者Id
	Title         string `gorm:"not null"`           //视频标题
	PlayUrl       string `gorm:"not null"`           //视频地址
	CoverUrl      string `gorm:"not null"`           //封面地址
	FavoriteCount int64  `gorm:"not null;default:0"` //点赞数
	CommentCount  int64  `gorm:"not null;default:0"` //评论数

}
