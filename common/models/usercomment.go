package models

import (
	"errors"
	"gorm.io/gorm"
	"log"
)

type UserComm struct {
	gorm.Model        // 继承ID, 创建时间, 修改时间
	UserName   string `gorm:"not null;"`          // 评论者
	VideoId    int64  `gorm:"not null"`           // 视频id
	Status     int32  `gorm:"not null;default:1"` // 评论状态 默认1有效
	CommText   string `gorm:"not null"`           // 评论内容
	UserId     uint64 `gorm:"not null"`           //用户id

}

func (t *UserComm) TableName() string {
	// 查询当前表的名称
	return "user_comm"
}

func InsertComment(comment *UserComm) error {
	log.Println("running-Post Comment:", comment)
	return DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(comment).Error; err != nil {
			return err
		}
		if err := tx.Model(&VideoInfo{}).UpdateColumn("comment_count", gorm.Expr("comment_count + ?", 1)).Error; err != nil {
			return err
		}
		return nil
	})
}

func DeleteComment(CommentId int64) error {
	log.Println("running-Delete Comment:", CommentId)
	return DB.Transaction(func(tx *gorm.DB) error {
		result := UserComm{}
		if err := tx.Where("id=? and status=?", CommentId, 1).First(&result).Error; err != nil {
			return err
		}
		if err := tx.Model(&result).UpdateColumn("status", 0).Error; err != nil {
			return err
		}
		if err := tx.Model(&VideoInfo{}).UpdateColumn("comment_count", gorm.Expr("comment_count - ?", 1)).Error; err != nil {
			return err
		}
		log.Println("Delete Comment success", CommentId)
		return nil
	})
}

func GetCommentsByVideo(VideoId int64, CommentList *[]UserComm, LimitCommentIds *[]uint) error {
	log.Println("running-Get Comment by Video_Id:", VideoId)
	where := map[string]interface{}{}
	where["id"] = LimitCommentIds
	return DB.Model(UserComm{}).Where("video_id=? and status=?", VideoId, 1).Not(where).Find(&CommentList).Error
}

func GetCommentsNumByVideo(VideoId int64) (int64, error) {
	log.Println("running-Get Comment Number by VideoId:", VideoId)
	var count int64
	err := DB.Model(UserComm{}).Where("video_id=? and status=?", VideoId, 1).Count(&count).Error
	if err != nil {
		return 0, errors.New("find number error in video")
	}
	return count, nil
}
