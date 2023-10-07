package controllers

import "easy-wallet-be/src/models"

// UpdateCategory updates a category with the given id, name and iconID in the database.
// It returns an error if the update operation fails.
func UpdateCategory(id uint, name string, iconID uint) error {
	db := models.DB()

	category := models.Category{
		Name:   name,
		IconID: iconID,
	}

	if err := db.Model(&models.Category{}).Where("id = ?", id).Updates(category).Error; err != nil {
		return err
	}

	return nil
}
