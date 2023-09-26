package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Session struct {
	gorm.Model
	SessionID           string    `gorm:"type:varchar(100);unique_index;not null"`
	ConfirmSessionToken string    `gorm:"type:varchar(100);not null"`
	UserID              uint      `gorm:"unique_index;not null"`
	ExpiresAt           time.Time `gorm:"default:current_timestamp"`
}
