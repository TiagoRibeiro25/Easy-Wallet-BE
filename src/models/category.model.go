package models

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Name   string `gorm:"type:varchar(100);not null"`
	IconID int    `gorm:"not null"`
	UserID uint   `gorm:"index;not null"`

	Expenses []Expense `gorm:"foreignkey:CategoryID"`
}
