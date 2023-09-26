package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Session struct {
	gorm.Model
	SessionID string    `gorm:"type:varchar(100);unique_index;not null"`
	UserID    uint      `gorm:"not null"`
	ExpiresAt time.Time `gorm:"not null"`
}
