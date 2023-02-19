package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	FromId  uint64 `gorm:"not null;index"`
	ToId    uint64 `gorm:"not null;index"`
	Content string `gorm:"not null"`
}
