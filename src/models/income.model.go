package models

import "github.com/jinzhu/gorm"

type Income struct {
	gorm.Model
	Month  int  `gorm:"not null"`
	Year   int  `gorm:"not null"`
	Amount uint `gorm:"not null"`
	UserID uint `gorm:"index;not null"`
}
