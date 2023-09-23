package models

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Name      string `gorm:"type:varchar(100);not null"`
	IconID    int    `gorm:"unique_index;not null"`
	TextColor string `gorm:"type:varchar(100);not null"`
	UserID    uint   `gorm:"unique_index;not null"`

	Expenses []Expense `gorm:"foreignkey:CategoryID"`
}
