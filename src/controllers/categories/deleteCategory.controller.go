package controllers

import "easy-wallet-be/src/models"

// DeleteCategory deletes a category from the database by its ID.
// It takes an ID of type uint as a parameter and returns an error if any.
// It returns nil if the category was deleted successfully.
func DeleteCategory(id uint) error {
	db := models.DB()

	err := db.Table("categories").Where("id = ?", id).Delete(&models.Category{}).Error

	return err
}
