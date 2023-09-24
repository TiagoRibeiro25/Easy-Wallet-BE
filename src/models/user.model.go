package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	DisplayName     string `gorm:"type:varchar(100);not null"`
	Email           string `gorm:"type:varchar(100);not null"`
	UserVerified    bool   `gorm:"type:boolean;not null;default:false"`
	VerifyUserToken string `gorm:"type:varchar(100);not null"`
	Currency        string `gorm:"type:varchar(100);not null;default:'EUR'"`

	Password Password  `gorm:"foreignkey:UserID"`
	Sessions []Session `gorm:"foreignkey:UserID"`
	Incomes  []Income  `gorm:"foreignkey:UserID"`
	Expenses []Expense `gorm:"foreignkey:UserID"`
}
