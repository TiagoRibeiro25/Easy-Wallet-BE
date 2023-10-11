package controllers

import (
	"easy-wallet-be/src/models"
)

// AddCategory adds a new category to the database for the given user.
// It takes the user ID, new category name and new category icon ID as input parameters.
// It returns an error if there is any issue while creating the new category.
func AddCategory(userID uint, newCategoryName string, newCategoryIconID uint) error {
	db := models.DB()

	// Create new category
	newCategory := models.Category{
		Name:   newCategoryName,
		IconID: newCategoryIconID,
		UserID: userID,
	}

	if err := db.Create(&newCategory).Error; err != nil {
		return err
	}

	return nil
}
