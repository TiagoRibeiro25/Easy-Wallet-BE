package models

import (
	"encoding/json"
	"os"

	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Name   string `gorm:"type:varchar(100);not null"`
	IconID uint   `gorm:"not null"`
	UserID uint   `gorm:"index;not null"`

	Expenses []Expense `gorm:"foreignkey:CategoryID;constraint:OnDelete:CASCADE;"`
}


// AddDefaultCategories adds default categories to the database for a given user ID.
// It reads the categories from a JSON file and creates them in a single transaction.
// If any error occurs, it rolls back the transaction and returns the error.
func AddDefaultCategories(userID uint) error {
	file, err := os.Open("src/data/categories.json")
	if err != nil {
		return err
	}

	defer file.Close()

	var categories []struct {
		Name   string `json:"name"`
		IconID uint   `json:"icon_id"`
	}

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&categories); err != nil {
		return err
	}

	db := DB()

	// Start a transaction
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Create all the categories in a single transaction
	for _, category := range categories {
		err = tx.Create(&Category{
			Name:   category.Name,
			IconID: category.IconID,
			UserID: userID,
		}).Error

		if err != nil {
			// Rollback the transaction and return the error
			tx.Rollback()
			return err
		}
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
