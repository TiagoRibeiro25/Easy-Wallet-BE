package controllers

import (
	"easy-wallet-be/src/models"
	"errors"
)

// GetResetPasswordTokenByEmail retrieves the reset password token for a user with the given email.
// It returns the reset password token and an error if the user is not found.
func GetResetPasswordTokenByEmail(email string) (string, error) {
	db := models.DB()

	var userWithResetPasswordToken struct {
		ID                 string
		ResetPasswordToken string
	}

	result := db.Table("users").
		Select("passwords.id, passwords.reset_password_token").
		Joins("JOIN passwords ON users.id = passwords.user_id").
		Where("users.email = ?", email).
		First(&userWithResetPasswordToken)

	if result.RowsAffected == 0 {
		return "", errors.New("user not found")
	}

	return userWithResetPasswordToken.ResetPasswordToken, nil
}
