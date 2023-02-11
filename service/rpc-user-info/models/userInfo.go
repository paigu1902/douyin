package models

import (
	"context"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type UserInfo struct {
	gorm.Model
	UserName      string `gorm:"not null;unique"`
	Password      string `gorm:"not null"`
	Salt          string `gorm:"default:null"`
	FollowCount   uint64 `gorm:"not null;default:0"`
	FollowedCount uint64 `gorm:"not null;default:0"`
	VideoCount    uint64 `gorm:"not null;default:0"`
}

func (table *UserInfo) TableName() string {
	return "user_info"
}

func FindUserByName(name string) UserInfo {
	user := UserInfo{}
	DB.Where("user_name = ?", name).First(&user)
	return user
}
func FindUserByID(ID uint64) UserInfo {
	user := UserInfo{}
	DB.Where("id = ?", ID).First(&user)
	return user
}

func CreateUser(user *UserInfo) error {
	return DB.Create(user).Error
}

func Actcion(fromId uint64, toId uint64, actionType string) error {
	var err error
	err = RDB.Del(context.Background(), strconv.Itoa(int(fromId))).Err()
	err = RDB.Del(context.Background(), strconv.Itoa(int(toId))).Err()
	if err != nil {
		return errors.New("删除缓存失败")
	}
	defer func() {
		go func() {
			time.Sleep(time.Second * 3)
			err = RDB.Del(context.Background(), strconv.Itoa(int(fromId))).Err()
			err = RDB.Del(context.Background(), strconv.Itoa(int(toId))).Err()
		}()
	}()

	err = DB.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := tx.Model(&UserInfo{}).Where("id", fromId).Update("follow_count", gorm.Expr("follow_count"+actionType+"?", 1)).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}
		if err := tx.Model(&UserInfo{}).Where("id", toId).Update("followed_count", gorm.Expr("followed_count"+actionType+"?", 1)).Error; err != nil {
			return err
		}
		// 返回 nil 提交事务
		return nil
	})

	return err
}

//func DeleteActcion(fromId uint64, toId uint64) error {
//	return DB.Transaction(func(tx *gorm.DB) error {
//		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
//		if err := tx.Model(&UserInfo{}).Where("id", fromId).Update("follow_count", gorm.Expr("follow_count-?", 1)).Error; err != nil {
//			// 返回任何错误都会回滚事务
//			return err
//		}
//		if err := tx.Model(&UserInfo{}).Where("id", toId).Update("followed_count", gorm.Expr("followed_count-?", 1)).Error; err != nil {
//			return err
//		}
//		// 返回 nil 提交事务
//		return nil
//	})
//}

func InfoRDB(userId uint64) (*UserInfo, error) {
	var userinfo UserInfo
	var err error
	userCache, err := RDB.Do(context.Background(), "get", userId).Text()
	if err == nil {
		err := json.Unmarshal([]byte(userCache), &userinfo)
		if err == nil {
			return &userinfo, nil
		}
	}
	user := FindUserByID(userId)
	u, _ := json.Marshal(user)
	err = RDB.Do(context.Background(), "set", user.ID, string(u)).Err()
	if err != nil {
		return &userinfo, errors.New("写入异常")
	}
	return &user, nil
}
