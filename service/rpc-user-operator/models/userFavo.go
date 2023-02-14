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

// 根据视频id获取点赞用户id列表
func GetLikeUserId(videoId int64) ([]int64, error) {
	var LikeUserId []int64
	err := DB.Model(&UserLike{}).Where("VideoId = ? and Status = ?", videoId, 1).
		Pluck("UserId", &LikeUserId).Error
	if err != nil {
        if err.Error() == "record not found" {
            log.Println("The video hasn't been liked.")
            return nil, err
        } else {
		    log.Println(err.Error())
		    return nil, errors.New("Get LikeUserId Failed")
        }
    } else {
		return LikeUserId, nil
	}
}

// 根据用户id查询其点赞的视频id列表
func GetLikeVideoId(userId int64) ([]int64, error) {
	var LikeVideoId []int64
	err := DB.Model(&UserLike{}).Where("UserId = ? and Status = ?", videoId, 1).
		Pluck("VideoId", &LikeVideoId).Error
	if err != nil {
        if err.Error() == "record not found" {
            log.Println("The user hasn't liked any video.")
            return nil, err
        } else {
		    log.Println(err.Error())
		    return nil, errors.New("Get LikeVideoId Failed")
        }
    } else {
		return LikeVideoId, nil
	}
}

// 查询用户-视频点赞信息
func GetLikeRecord(userId int64, videoId int64) (UserLike, error) {
	var likeRecord UserLike
	err := DB.Model(&UserLike{}).Where("UserId = ? and VideoId = ?", userId, videoId).
		First(&likeRecord).Error
	if err != nil {
        if err.Error() == "record not found" {
            log.Println("The user hasn't like this video.")
            return nil, err
        } else {
            log.Println(err.Error())
		    return nil, errors.New("Get LikeRecord Failed")
	    }
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
		return errors.New("Update Record Failed")
	} else {
		return nil
	}
}

// 创建点赞记录
func CreateLikeRecord(userId int64, videoId int64) error {
	likeRecord := UserLike{UserId: userId, VideoId: videoId, Status:1}
	err := DB.Model(&UserLike{}).Create(&likeRecord).Error
	if err != nil {
		log.Println(err.Error())
		return errors.New("Create Record Failed")
	} else {
		return nil
	}
}
