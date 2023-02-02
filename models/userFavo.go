package models

import (
	"errors"
	"gorm.io/gorm"
	"log"
)

type UserLike struct {
	gorm.Model
	UserId   int64  `gorm:"not null;default:0"` //用户id
	VideoId  int64  `gorm:"not null;default:0"` //点赞视频id
	UserName string `gorm:"not null"`           //用户名
	Status   int32  `gorm:"not null;default:1"` //点赞状态 默认1点赞 0取消
}

func (table *UserLike) TableName() string {
	return "user_like"
}

// 根据视频id获取点赞用户id
// 待改进：区分 no record 和 get record failed
func GetLikeUserId(videoId int64) ([]int64, error) {
	var LikeUserId []int64
	err := DB.Model(&UserLike{}).Where("VideoId = ? and Status = ?", videoId, 1).
		Pluck("UserId", &LikeUserId).Error
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("get LikeUserId failed")
	} else {
		return LikeUserId, nil
	}
}

// 根据用户id查询其点赞的视频id
// 待改进：区分 no record 和 get record failed
func GetLikeVideoId(userId int64) ([]int64, error) {
	var LikeVideoId []int64
	err := DB.Model(&UserLike{}).Where("UserId = ? and Status = ?", videoId, 1).
		Pluck("VideoId", &LikeVideoIdId).Error
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("get LikeVideoId failed")
	} else {
		return LikeVideoId, nil
	}
}

// 查询用户-视频点赞信息
// 待改进：区分 no record 和 get record failed
func GetLikeRecord(userId int64, videoId int64) (UserLike, error) {
	var likeRecord UserLike
	err := DB.Model(&UserLike{}).Where("UserId = ? and VideoId = ?", userId, videoId).
		First(&likeRecord).Error
	if err != nil {
		log.Println(err.Error())
		return likeRecord, errors.New("get LikeRecord failed")
	} else {
		return likeRecord, nil
	}
}

// 更新点赞状态 双击取消
func UpdateLikeStatus(userId int64, videoId int64, status int32) error {
	err := DB.Model(&UserLike{}).Where("UserId = ? and VideoId = ?", userId, videoId).
		Update("Status = ?", status).Error
	if err != nil {
		log.Println(err.Error())
		return errors.New("update record failed")
	} else {
		return nil
	}
}

// 创建点赞记录
func CreateLikeRecord(likeRecord UserLike) error {
	err := DB.Model(&UserLike{}).Create(&likeRecord).Error
	if err != nil {
		log.Println(err.Error())
		return errors.New("create record failed")
	} else {
		return nil
	}
}
