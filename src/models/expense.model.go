package models

import "github.com/jinzhu/gorm"

type Expense struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100);not null"`
	Cost        int    `gorm:"not null"`
	Description string `gorm:"type:varchar(100);not null"`
	UserID      uint   `gorm:"unique_index;not null"`
	CategoryID  uint   `gorm:"unique_index;not null"`
}
