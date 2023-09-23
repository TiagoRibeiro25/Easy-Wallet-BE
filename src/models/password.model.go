package models

import "github.com/jinzhu/gorm"

type Password struct {
	gorm.Model
	Password           string `gorm:"type:varchar(100);not null"`
	ResetPasswordToken string `gorm:"type:varchar(100);not null"`
	UserID             uint   `gorm:"unique_index;not null"`
}
