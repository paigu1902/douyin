package models

import (
	"gorm.io/gorm"
)

type UserInfo struct {
	gorm.Model
	UserName       string `gorm:"not null;unique"`
	Password       string `gorm:"not null"`
	Salt           string `gorm:"default:null"`
	FollowCount    int64  `gorm:"not null;default:0"`
	FollowedCount  int64  `gorm:"not null;default:0"`
	VideoCount     int64  `gorm:"not null;default:0"`
	TotalFavorited string `gorm:"not null;default:0"`
	FavoriteCount  int64  `gorm:"not null;default:0"`
}

func (table *UserInfo) TableName() string {
	return "user_info"
}
