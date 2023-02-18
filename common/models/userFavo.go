package models

import (
	"errors"
	"gorm.io/gorm"
	"log"
)

type UserFavo struct {
	gorm.Model
	UserId   uint64 `gorm:"not null;default:0"` //用户id
	VideoId  uint64 `gorm:"not null;default:0"` //点赞视频id
	UserName string `gorm:"not null"`           //用户名
	Status   uint32 `gorm:"not null;default:1"` //点赞状态 默认1点赞 0取消
}

func (table *UserFavo) TableName() string {
	return "user_favo"
}

// 根据视频id获取点赞用户id列表
func GetFavoUserId(videoId uint64) ([]uint64, error) {
	var favoUserId []uint64
	err := DB.Model(&UserFavo{}).Where("video_id = ? and status = ?", videoId, 1).
		Pluck("user_id", &favoUserId).Error
	if err != nil {
		if err.Error() == "record not found" {
			log.Println("The video hasn't been liked.")
			return nil, err
		} else {
			log.Println(err.Error())
			return nil, errors.New("Get favoUserId Failed")
		}
	}
	return favoUserId, nil
}

// 根据用户id查询其点赞的视频id列表
func GetFavoVideoId(userId uint64) ([]uint64, error) {
	var favoVideoId []uint64
	err := DB.Model(&UserFavo{}).Where("user_id = ? and status = ?", userId, 1).
		Pluck("video_id", &favoVideoId).Error
	if err != nil {
		if err.Error() == "record not found" {
			log.Println("The user hasn't liked any video.")
			return nil, err
		} else {
			log.Println(err.Error())
			return nil, errors.New("Get favoVideoId Failed")
		}
	}
	return favoVideoId, nil
}

// 查询用户-视频点赞信息
func GetFavoRecord(userId uint64, videoId uint64) (UserFavo, error) {
	var favoRecord UserFavo
	err := DB.Model(&UserFavo{}).Where("user_id = ? and video_id = ?", userId, videoId).
		First(&favoRecord).Error
	if err != nil {
		if err.Error() == "record not found" {
			log.Println("The user hasn't like this video.")
			return favoRecord, err
		} else {
			log.Println(err.Error())
			return favoRecord, errors.New("Get LikeRecord Failed")
		}
	}
	return favoRecord, nil

}

// 更新点赞状态 双击取消
func UpdateFavoStatus(userId uint64, videoId uint64, status uint32) error {
	err := DB.Model(&UserFavo{}).Where("user_id = ? and video_id = ?", userId, videoId).
		Update("status = ?", status).Error
	if err != nil {
		log.Println(err.Error())
		return errors.New("Update Record Failed")
	}
	return nil
}

// 创建点赞记录
func CreateFavoRecord(favoRecord *UserFavo) error {
	err := DB.Model(&UserFavo{}).Create(&favoRecord).Error
	if err != nil {
		log.Println(err.Error())
		return errors.New("Create Record Failed")
	}
	return nil
}

//func CreateFavoRecord(userId uint64, videoId uint64) error {
//	favoRecord := UserFavo{UserId: userId, VideoId: videoId, Status: 1}
//	err := DB.Model(&UserFavo{}).Create(&favoRecord).Error
//	if err != nil {
//		log.Println(err.Error())
//		return errors.New("Create Record Failed")
//	}
//	return nil
//}
