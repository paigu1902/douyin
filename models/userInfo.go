package models

import (
	"fmt"
	"gorm.io/gorm"
	"paigu1902/douyin/utils"
	"time"
)

type UserInfo struct {
	gorm.Model
	UserName      string `gorm:"not null;unique"`
	Password      string `gorm:"not null"`
	Salt          string `gorm:"default:null"`
	FollowCount   int64  `gorm:"not null;default:0"`
	FollowedCount int64  `gorm:"not null;default:0"`
	VideoCount    int64  `gorm:"not null;default:0"`
}

func (table *UserInfo) TableName() string {
	return "user_info"
}

func FindUserByName(name string) UserInfo {
	user := UserInfo{}
	fmt.Println("first", user)
	DB.Where("user_name = ?", name).First(&user)
	return user
}

func FindUserBynameAndPwd(name string, password string) UserInfo {
	user := UserInfo{}
	fmt.Println("first", user)
	DB.Where("name = ? and password=?", name, password).First(&user)
	//token 加密
	str := fmt.Sprintf("%d", time.Now().Unix())
	temp := utils.MD5Encode(str)
	DB.Model(&user).Where("id = ?", user.ID).Update("identity", temp)
	return user
}

func CreateUser(user UserInfo) *gorm.DB {
	return DB.Create(&user)
}
