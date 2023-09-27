package controllers

import "easy-wallet-be/src/models"

// VerifyUserByToken updates the user_verified field to true for the user with the given verify_user_token.
// It takes a token string as input and returns an error if the update operation fails.
func VerifyUserByToken(token string) error {
	db := models.DB()

	err := db.Model(&models.User{}).
		Where("verify_user_token = ?", token).
		Update("user_verified", true).
		Error

	if err != nil {
		return err
	}

	return nil
}
