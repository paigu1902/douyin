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

// 通过视频id获得视频的所有评论id列表

//func GetCommentList(VideoId int64) ([]string, error) {
//	log.Println("Running-Get CommentList By Video_Id: ", VideoId)
//	var CommentList []string
//	err := Db.Model(UserComm{}).Select("ID").Where("Video_Id = ? and Status = ?", VideoId, 1).Find(&CommentList).Error
//	if err != nil {
//		log.Println("GetCommentList: ", err)
//		return nil, err
//	}
//	return CommentList, nil
//}

func GetCommentList(VideoId int64, CommentList *[]int64) error {
	log.Println("Running-Get CommentList By Video_Id: ", VideoId)
	return DB.Select("id").Where("video_id=? and status=?", VideoId, 1).Find(CommentList).Order("created_at").Error
}

// 发表评论

//func InsertComment(Comment *UserComm) error {
//	log.Println("running-insert comment")
//	log.Println("running-Post Comment:", Comment)
//	err := Db.Create(Comment).Error
//	if err != nil {
//		log.Println("Comment Insert failed")
//		return errors.New("Insert Failed")
//	}
//	log.Println("Post Comment success")
//	return nil
//}

func InsertComment(comment *UserComm) error {
	log.Println("running-Post Comment:", comment)
	return DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(comment).Error; err != nil {
			return err
		}
		return nil
	})
}

// 通过评论id 删除评论

//func DeleteComment(CommentId int64) error {
//	log.Println("running-Delete Comment:", CommentId)
//	var ResultComment UserComm
//	comm := Db.Model(UserComm{}).Where("ID = ? and Status = ?", CommentId, 1).First(&ResultComment)
//	//comm := Db.Model(UserComm{}).Where(map[string]interface{}{"id": CommentId, "Status": 1}).First(&ResultComment)
//	if comm.RowsAffected == 0 {
//		log.Println("comment is not exist") //函数返回提示错误信息
//		return errors.New("del comment is not exist")
//	}
//	err := Db.Model(UserComm{}).Where("ID = ?", CommentId).Update("Status", 0).Error
//	if err != nil {
//		log.Println("Delete comment failed")
//		return errors.New("del comment failed")
//	}
//	log.Println("Delete Comment success")
//	return nil
//}

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
		log.Println("Delete Comment success", CommentId)
		return nil
	})
}

// 通过VideoId 获取评论列表

//func GetCommentsByVideo(VideoId int64) ([]UserComm, error) {
//	log.Println("running-Get Comment by Video_Id:", VideoId)
//	var CommentList []UserComm
//	comm := Db.Model(UserComm{}).Where("Video_Id = ? and Status = ?", VideoId, 1).Order("Created_At desc").Find(&CommentList)
//	if comm.RowsAffected == 0 {
//		log.Println("There are no Comment in this Video")
//		return nil, nil
//	}
//	if comm.Error != nil {
//		log.Println(comm.Error.Error())
//		log.Println("Get Comment failed")
//		return CommentList, errors.New("get comment list error")
//	}
//	log.Println("Get Comment success")
//	return CommentList, nil
//}

func GetCommentsByVideo(VideoId int64, CommentList *[]UserComm) error {
	log.Println("running-Get Comment by Video_Id:", VideoId)
	return DB.Model(UserComm{}).Where("video_id=? and status=?", VideoId, 1).Find(&CommentList).Error
}

// 获取评论数量

//func GetCommentsNumByVideo(VideoId int64) (int64, error) {
//	log.Println("running-Get Comment by VideoId:", VideoId)
//	var count int64
//	err := Db.Model(UserComm{}).Where("Video_Id = ? and Status = ?", VideoId, 1).Count(&count).Error
//
//	if err != nil {
//		log.Println(err.Error())
//		log.Println("Get Comment number failed")
//		return -1, errors.New("get comment number error")
//	}
//	log.Println("Get Comment number success")
//	return count, nil
//}

func GetCommentsNumByVideo(VideoId int64) (int64, error) {
	log.Println("running-Get Comment Number by VideoId:", VideoId)
	var count int64
	err := DB.Model(UserComm{}).Where("video_id=? and status=?", VideoId, 1).Count(&count).Error
	log.Println(count)
	if err != nil {
		return 0, errors.New("find number error in video")
	}
	return count, nil
}
