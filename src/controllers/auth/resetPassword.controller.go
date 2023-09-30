package controllers

import (
	"easy-wallet-be/src/models"
	"easy-wallet-be/src/utils"
)

// ResetPassword updates the password and reset password token for a user with the provided resetToken.
// It generates a new reset password token and returns an error if the update fails.
// It returns an error if the provided resetToken does not exist in the database.
// resetToken: string - the reset password token for the user.
// newPassword: string - the new password for the user.
// Returns an error if any operation fails.
func ResetPassword(resetToken string, newPassword string) error {
	newToken, err := utils.GenerateToken()
	if err != nil {
		return err
	}

	db := models.DB()

	// Start a database transaction
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Check if a record with the provided resetToken exists
	var existingRecord models.Password
	if err := tx.Where("reset_password_token = ?", resetToken).First(&existingRecord).Error; err != nil {
		// Roll back the transaction and return an error
		tx.Rollback()
		return err
	}

	// Update the password and reset password token
	if err := tx.Model(&models.Password{}).
		Where("reset_password_token = ?", resetToken).
		Updates(map[string]interface{}{
			"password":             newPassword,
			"reset_password_token": newToken,
		}).Error; err != nil {
		// Roll back the transaction and return an error
		tx.Rollback()
		return err
	}

	// Commit the transaction if everything was successful
	return tx.Commit().Error
}
