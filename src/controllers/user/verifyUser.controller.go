package controllers

import (
	"easy-wallet-be/src/models"
	"errors"
)

// VerifyUserByToken updates the user_verified field to true for the user with the given verify_user_token.
// It returns an error if the update operation fails, if no user was updated, or if the user is already verified.
func VerifyUserByToken(token string) error {
	db := models.DB()

	// Start a database transaction
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Check if the user is already verified within the transaction
	var existingUser models.User
	if err := tx.Where("verify_user_token = ?", token).First(&existingUser).Error; err != nil {
		// Roll back the transaction and return an error
		tx.Rollback()
		return err
	}

	if existingUser.UserVerified {
		// Roll back the transaction and return an error
		tx.Rollback()
		return errors.New("user is already verified")
	}

	// Update the user_verified field
	result := tx.Model(&models.User{}).
		Where("verify_user_token = ?", token).
		Updates(map[string]interface{}{"user_verified": true})

	if result.Error != nil {
		// Roll back the transaction and return an error
		tx.Rollback()
		return result.Error
	}

	// Check the number of affected rows to see if any user was updated
	if result.RowsAffected == 0 {
		// Roll back the transaction and return an error
		tx.Rollback()
		return errors.New("no user was updated")
	}

	// Commit the transaction if everything was successful
	return tx.Commit().Error
}
