package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"log"
)

var DB *gorm.DB

type UserFavo struct {
	gorm.Model
	UserId   int64  `gorm:"not null;default:0"` //用户id
	VideoId  int64  `gorm:"not null;default:0"` //点赞视频id
	UserName string `gorm:"not null"`           //用户名
	Status   int32  `gorm:"not null;default:1"` //点赞状态 默认1点赞 0取消
}

func (table *UserFavo) TableName() string {
	return "userFavo"
}

// 根据视频id获取点赞用户id列表
func GetFavoUserId(videoId int64) ([]int64, error) {
	var FavoUserId []int64
	err := DB.Model(&UserFavo{}).Where("VideoId = ? and Status = ?", videoId, 1).
		Pluck("UserId", &FavoUserId).Error
	if err != nil {
		if err.Error() == "record not found" {
			log.Println("The video hasn't been liked.")
			return nil, err
		} else {
			log.Println(err.Error())
			return nil, errors.New("Get FavoUserId Failed")
		}
	} else {
		return FavoUserId, nil
	}
}

// 根据用户id查询其点赞的视频id列表
func GetFavoVideoId(userId int64) ([]int64, error) {
	var FavoVideoId []int64
	err := DB.Model(&UserFavo{}).Where("UserId = ? and Status = ?", userId, 1).
		Pluck("VideoId", &FavoVideoId).Error
	if err != nil {
		if err.Error() == "record not found" {
			log.Println("The user hasn't liked any video.")
			return nil, err
		} else {
			log.Println(err.Error())
			return nil, errors.New("Get FavoVideoId Failed")
		}
	} else {
		return FavoVideoId, nil
	}
}

// 查询用户-视频点赞信息
func GetFavoRecord(userId int64, videoId int64) (UserFavo, error) {
	var favoRecord UserFavo
	err := DB.Model(&UserFavo{}).Where("UserId = ? and VideoId = ?", userId, videoId).
		First(&favoRecord).Error
	if err != nil {
		if err.Error() == "record not found" {
			log.Println("The user hasn't like this video.")
			return favoRecord, err
		} else {
			log.Println(err.Error())
			return favoRecord, errors.New("Get FavoRecord Failed")
		}
	} else {
		return favoRecord, nil
	}
}

// 更新点赞状态 双击取消
func UpdateFavoStatus(userId int64, videoId int64, status int32) error {
	err := DB.Model(&UserFavo{}).Where("UserId = ? and VideoId = ?", userId, videoId).
		Update("Status = ?", status).Error
	if err != nil {
		log.Println(err.Error())
		return errors.New("Update Record Failed")
	} else {
		return nil
	}
}

// 创建点赞记录
func CreateFavoRecord(userId int64, videoId int64) error {
	likeRecord := UserFavo{UserId: userId, VideoId: videoId, Status: 1}
	err := DB.Model(&UserFavo{}).Create(&likeRecord).Error
	if err != nil {
		log.Println(err.Error())
		return errors.New("Create Record Failed")
	} else {
		return nil
	}
}
