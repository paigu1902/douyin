package models

import (
	"errors"
	"gorm.io/gorm"
	"log"
)

type UserFavo struct {
	gorm.Model
	UserId   uint64 `gorm:"not null;default:0;index:idx_user_id;index:idx_fromTo,unique"`  //用户id
	VideoId  uint64 `gorm:"not null;default:0;index:idx_video_id;index:idx_fromTo,unique"` //点赞视频id
	UserName string `gorm:"not null"`                                                      //用户名
	Status   uint32 `gorm:"not null;default:1"`                                            //点赞状态 默认1点赞 0取消
}

func (table *UserFavo) TableName() string {
	return "user_favo"
}

// GetFavoUserId 根据视频id获取点赞用户id列表
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
			return nil, errors.New("get favoUserId Failed")
		}
	}
	return favoUserId, nil
}

// GetFavoVideoId 根据用户id查询其点赞的视频id列表
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
			return nil, errors.New("get favoVideoId Failed")
		}
	}
	return favoVideoId, nil
}

// GetFavoRecord 查询用户-视频点赞信息
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
			return favoRecord, errors.New("get LikeRecord Failed")
		}
	}
	return favoRecord, nil

}

// UpdateFavoStatus 更新点赞状态 双击取消
func UpdateFavoStatus(favoRecord *UserFavo) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		result := UserFavo{}
		log.Println(favoRecord)
		if err := tx.Where("user_id = ? and video_id = ?", favoRecord.UserId, favoRecord.VideoId).First(&result).Error; err != nil {
			return err
		}
		if err := tx.Model(&result).UpdateColumn("status", favoRecord.Status).Error; err != nil {
			return err
		}
		if favoRecord.Status == 1 {
			if err := DB.Model(&VideoInfo{}).UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
				return err
			}
		} else {
			if err := DB.Model(&VideoInfo{}).UpdateColumn("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

// CreateFavoRecord 创建点赞记录
func CreateFavoRecord(favoRecord *UserFavo) error {
	err := DB.Model(&UserFavo{}).Create(&favoRecord).Error
	if err != nil {
		log.Println(err.Error())
		return errors.New("create Record Failed")
	}
	if err := DB.Model(&VideoInfo{}).UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
		return err
	}
	return nil
}
