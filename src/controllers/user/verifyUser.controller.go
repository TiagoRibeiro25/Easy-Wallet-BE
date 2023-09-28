package controllers

import (
	"easy-wallet-be/src/models"
	"errors"
)

// VerifyUserByToken updates the user_verified field to true for the user with the given verify_user_token.
// It returns an error if the update operation fails or if no user was updated.
func VerifyUserByToken(token string) error {
	db := models.DB()

	// Use the Update method and capture the result to check the number of affected rows.
	result := db.Model(&models.User{}).
		Where("verify_user_token = ?", token).
		Update("user_verified", true)

	if result.Error != nil {
		return result.Error
	}

	// Check the number of affected rows to see if any user was updated.
	if result.RowsAffected == 0 {
		return errors.New("no user was updated")
	}

	return nil
}
