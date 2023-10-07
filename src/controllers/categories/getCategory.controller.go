package controllers

import "easy-wallet-be/src/models"

// GetCategory retrieves a category from the database by its ID.
// It returns the category and an error if there was any issue with the database query.
func GetCategory(id uint) (models.Category, error) {
	db := models.DB()

	var category models.Category
	err := db.Table("categories").Where("id = ?", id).First(&category).Error

	return category, err
}
