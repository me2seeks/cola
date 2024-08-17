package models

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	CreatedAt time.Time      `gorm:"column:create_at;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"column:update_at;not null;default:CURRENT_TIMESTAMP;autoUpdateTime"`
	DeleteAt  gorm.DeletedAt `gorm:"column:delete_at"`
}
