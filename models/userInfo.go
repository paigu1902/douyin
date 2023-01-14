package models

import "gorm.io/gorm"

type UserInfo struct {
	gorm.Model
	UserName      string `gorm:"not null;unique"`
	Password      string `gorm:"not null"`
	FollowCount   int64  `gorm:"not null;default:0"`
	FollowedCount int64  `gorm:"not null;default:0"`
	VedioCount    int64  `gorm:"not null;default:0"`
}
