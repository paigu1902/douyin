package models

import (
	"errors"
	"gorm.io/gorm"
	"log"
)

var Db *gorm.DB

type UserComm struct {
	gorm.Model        // 继承ID, 创建时间, 修改时间
	UserName   string `gorm:"not null;"`          // 评论者
	VideoId    int64  `gorm:"not null"`           // 视频id
	Status     int32  `gorm:"not null;default:1"` // 评论状态 默认1有效
	CommText   string `gorm:"not null"`           // 评论内容
}

// 查询当前表的名称
func (UserComm) TableName() string {
	return "user_comm"
}

// 通过视频id获得视频的所有评论id列表

func GetCommentList(VideoId int64) ([]string, error) {
	log.Println("Running-Get CommentList By VideoId: ", VideoId)
	var CommentList []string
	err := Db.Model(UserComm{}).Select("ID").Where("VideoId = ?", VideoId).Find(&CommentList).Error
	if err != nil {
		log.Println("GetCommentList: ", err)
		return nil, err
	}
	return CommentList, nil
}

// 发表评论

func InsertComment(Comment UserComm) (UserComm, error) {
	log.Println("running-Post Comment:", Comment)
	err := Db.Model(UserComm{}).Create(&Comment).Error
	if err != nil {
		log.Println("Comment Insert failed")
		return Comment, errors.New("create comment failed")
	}
	log.Println("Post Comment success")
	return Comment, nil
}

// 通过评论id 删除评论

func DeleteComment(CommentId int64) error {
	log.Println("running-Delete Comment:", CommentId)
	var ResultComment UserComm
	comm := Db.Model(UserComm{}).Where("ID = ? and Status = ?", CommentId, 1).First(&ResultComment)
	//comm := Db.Model(UserComm{}).Where(map[string]interface{}{"id": CommentId, "Status": 1}).First(&ResultComment)
	if comm.RowsAffected == 0 {
		log.Println("comment is not exist") //函数返回提示错误信息
		return errors.New("del comment is not exist")
	}
	err := Db.Model(UserComm{}).Where("ID = ?", CommentId).Update("Status", 0).Error
	if err != nil {
		log.Println("Delete comment failed")
		return errors.New("del comment failed")
	}
	log.Println("Delete Comment success")
	return nil
}

// 通过VideoId 获取评论列表

func GetCommentsByVideo(VideoId int64) ([]UserComm, error) {
	log.Println("running-Get Comment by VideoId:", VideoId)
	var CommentList []UserComm
	comm := Db.Model(UserComm{}).Where("VideoId = ? and Status = ?", VideoId, 1).Order("CreatedAt desc").Find(&CommentList)
	if comm.RowsAffected == 0 {
		log.Println("There are no Comment in this Video")
		return nil, nil
	}
	if comm.Error != nil {
		log.Println(comm.Error.Error())
		log.Println("Get Comment failed")
		return CommentList, errors.New("get comment list error")
	}
	log.Println("Get Comment success")
	return CommentList, nil
}
