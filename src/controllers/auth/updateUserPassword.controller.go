package controllers

import (
	"easy-wallet-be/src/models"
	"errors"
)

type UpdateUserPasswordTokenResult struct {
	ID                 uint   `json:"id"`
	Email              string `json:"email"`
	DisplayName        string `json:"display_name"`
	ResetPasswordToken string `json:"reset_password_token"`
}

// UpdateUserPasswordToken updates the reset password token for a user with the given email.
// It returns an UpdateUserPasswordTokenResult and an error.
// The UpdateUserPasswordTokenResult contains the user's id, email, display name and the new reset password token.
// If the user is not found, it returns an error.
// If there is an error updating the user's password token, it returns an error.
// If both queries are successful, it returns the updated user and no error.
func UpdateUserPasswordToken(email string, newToken string) (UpdateUserPasswordTokenResult, error) {
	db := models.DB()

	// Begin a transaction
	tx := db.Begin()
	if tx.Error != nil {
		return UpdateUserPasswordTokenResult{}, tx.Error
	}

	var user UpdateUserPasswordTokenResult

	if err := tx.Table("users").
		Select("users.id, users.email, users.display_name, passwords.reset_password_token").
		Joins("JOIN passwords ON users.id = passwords.user_id").
		Where("users.email = ?", email).
		First(&user).
		Error; err != nil {
		tx.Rollback() // Rollback the transaction in case of an error
		return UpdateUserPasswordTokenResult{}, errors.New("user not found")
	}

	// Update the reset_password_token
	user.ResetPasswordToken = newToken

	if err := tx.Table("passwords").
		Where("user_id = ?", user.ID).Update("reset_password_token", newToken).
		Error; err != nil {
		tx.Rollback() // Rollback the transaction in case of an error
		return UpdateUserPasswordTokenResult{}, errors.New("error updating user password token")
	}

	// Commit the transaction if both queries are successful
	if err := tx.Commit().Error; err != nil {
		return UpdateUserPasswordTokenResult{}, err
	}

	return user, nil
}
