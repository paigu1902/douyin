package models

import (
	"time"
)

type Relation struct{
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	From_id uint64 `gorm:"not null;index:idx_fromTo,unique"`
	To_id  uint64 `gorm:"not null;index:idx_fromTo,unique"`
}