package controllers

import (
	schemas "easy-wallet-be/src/data/schemas/categories/getCategories"
	"easy-wallet-be/src/models"
	"errors"
)

// GetCategories retrieves all categories for a given user ID from the database.
// It returns a slice of ResponseData structs and an error if the user is not found.
func GetCategories(userID uint) ([]schemas.ResponseData, error) {
	db := models.DB()

	var responseData []schemas.ResponseData
	err := db.Table("categories").Where("user_id = ?", userID).Find(&responseData).Error
	if err != nil {
		return responseData, errors.New("user not found")
	}

	return responseData, nil
}
