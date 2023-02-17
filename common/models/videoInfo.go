package models

import (
	"gorm.io/gorm"
	"paigu1902/douyin/service/rpc-video-operator/kitex_gen/videoOperatorPb"
)

type VideoInfo struct {
	gorm.Model
	AuthorId      uint64 `gorm:"not null"`           //作者Id
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

func GetVideoInfo(latestTime string, limit int, videoList *[]VideoInfo) error {
	return DB.Where("deleted_at is NULL and created_at<=?", latestTime).Limit(limit).Find(videoList).Error
}

func GetVideoListByAuthorId(authorId uint64, videoList *[]VideoInfo) error {
	return DB.Where("author_id=?", authorId).Find(videoList).Order("created_at").Error
}

func GetVideosByIds(videoIdList []uint64, videoList *[]VideoInfo) error {
	// select * from video_info where in [videoIdList]
	return DB.Find(videoList, videoIdList).Error
}

func (v *VideoInfo) TransToVideoWithAuthor(author *videoOperatorPb.User) *videoOperatorPb.Video {
	video := v.TransToVideo()
	video.Author = author
	return video
}

func (v *VideoInfo) TransToVideo() *videoOperatorPb.Video {
	return &videoOperatorPb.Video{
		Id:            uint64(v.ID),
		PlayUrl:       v.PlayUrl,
		CoverUrl:      v.CoverUrl,
		CommentCount:  v.CommentCount,
		FavoriteCount: v.FavoriteCount,
		Title:         v.Title,
	}
}
