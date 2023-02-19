package models

import (
	"time"
)

type Relation struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	FromId    uint64 `gorm:"not null;index:idx_fromTo,unique;index"`
	ToId      uint64 `gorm:"not null;index:idx_fromTo,unique;index"`
}
