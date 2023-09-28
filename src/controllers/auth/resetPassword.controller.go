package controllers

import (
	"easy-wallet-be/src/models"
	"easy-wallet-be/src/utils"
)

// ResetPassword updates the password of a user with the given resetToken to the provided newPassword.
// It returns an error if the update operation fails.
func ResetPassword(resetToken string, newPassword string) error {
	newToken, err := utils.GenerateToken()
	if err != nil {
		return err
	}

	db := models.DB()

	// Update the password and reset password token
	err = db.Model(&models.Password{}).
		Where("reset_password_token = ?", resetToken).
		Updates(map[string]interface{}{
			"password":             newPassword,
			"reset_password_token": newToken,
		}).Error

	if err != nil {
		return err
	}

	return nil
}
