package controllers

import (
	schemas "easy-wallet-be/src/data/schemas/user/getUser"
	"easy-wallet-be/src/models"
	"errors"
)

// GetUser retrieves a user from the database by their ID and returns a ResponseData struct and an error.
// The ResponseData struct contains the user's ID, email, display name, currency, user verification status, and creation date.
// If the user is not found in the database, an error is returned.
func GetUser(userID uint) (schemas.ResponseData, error) {
	db := models.DB()

	var user models.User
	result := db.Table("users").
		Select("users.id, users.email, users.display_name, users.user_verified, users.currency, users.created_at").
		Where("users.id = ?", userID).
		First(&user)

	if result.RowsAffected == 0 {
		return schemas.ResponseData{}, errors.New("user not found")
	}

	responseData := schemas.ResponseData{
		ID:           user.ID,
		Email:        user.Email,
		DisplayName:  user.DisplayName,
		Currency:     user.Currency,
		UserVerified: user.UserVerified,
		CreatedAt:    user.CreatedAt.String(),
	}

	return responseData, nil
}
