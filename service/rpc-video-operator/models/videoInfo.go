package models

import "gorm.io/gorm"

type VideoInfo struct {
	gorm.Model
	AuthorId      int64  `gorm:"not null"`           //作者Id
	Title         string `gorm:"not null"`           //视频标题
	PlayUrl       string `gorm:"not null"`           //视频地址
	CoverUrl      string `gorm:"not null"`           //封面地址
	FavoriteCount int64  `gorm:"not null;default:0"` //点赞数
	CommentCount  int64  `gorm:"not null;default:0"` //评论数
}

func (t *VideoInfo) TableName() string {
	return "video_info"
}

func CreateVideoInfo(videoInfo *VideoInfo) error {
	return DB.Create(videoInfo).Error
}

func UpdateVideoInfo(info *VideoInfo) error {
	return DB.Model(info).Update("play_url", info.PlayUrl).Error
}
