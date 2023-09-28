package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	DisplayName     string `gorm:"type:varchar(100);not null"`
	Email           string `gorm:"type:varchar(100);not null"`
	UserVerified    bool   `gorm:"type:boolean;not null;default:false"`
	VerifyUserToken string `gorm:"unique_index;type:varchar(100);not null"`
	Currency        string `gorm:"type:varchar(100);not null;default:'EUR'"`

	Password Password  `gorm:"foreignkey:UserID;constraint:OnDelete:CASCADE;"`
	Sessions []Session `gorm:"foreignkey:UserID;constraint:OnDelete:CASCADE;"`
	Incomes  []Income  `gorm:"foreignkey:UserID;constraint:OnDelete:CASCADE;"`
	Expenses []Expense `gorm:"foreignkey:UserID;constraint:OnDelete:CASCADE;"`
}
