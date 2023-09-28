package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	DisplayName     string `gorm:"type:varchar(100);not null"`
	Email           string `gorm:"type:varchar(100);not null"`
	UserVerified    bool   `gorm:"type:boolean;not null;default:false"`
	VerifyUserToken string `gorm:"unique_index;type:varchar(100);not null"`
	Currency        string `gorm:"type:varchar(100);not null;default:'EUR'"`

	Password Password  `gorm:"foreignkey:UserID;on_delete:CASCADE"`
	Sessions []Session `gorm:"foreignkey:UserID;on_delete:CASCADE"`
	Incomes  []Income  `gorm:"foreignkey:UserID;on_delete:CASCADE"`
	Expenses []Expense `gorm:"foreignkey:UserID;on_delete:CASCADE"`
}
