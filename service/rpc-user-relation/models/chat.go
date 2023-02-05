package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	From_id uint64 `gorm:"not null"`
	To_id  uint64 `gorm:"not null"`
	Content string `gorm:"not null"`
}

