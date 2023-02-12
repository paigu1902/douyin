package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	FromId  uint64 `gorm:"not null"`
	ToId    uint64 `gorm:"not null"`
	Content string `gorm:"not null"`
}
