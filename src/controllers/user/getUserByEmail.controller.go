package controllers

import (
	"easy-wallet-be/src/models"
	"errors"

	"github.com/labstack/echo/v4"
)

const SELECT_QUERY = "users.id, users.email, users.display_name, users.user_verified, users.verify_user_token ,users.currency, users.created_at, passwords.password"

type UserWithPassword struct {
	ID              string
	Email           string
	DisplayName     string
	UserVerified    bool
	VerifyUserToken string
	Currency        string
	CreatedAt       string
	Password        string
}

// FetchUserByEmail fetches a user with their password from the database by email.
// It takes an echo context and an email string as parameters.
// It returns a UserWithPassword struct and an error (if any).
// The function queries the database for a user with the given email and joins the passwords table to retrieve the user's password.
// If the user is not found, it returns an empty UserWithPassword struct and an error.
func FetchUserByEmail(c echo.Context, email string) (UserWithPassword, error) {
	db := models.DB()

	var userWithPassword UserWithPassword
	result := db.Table("users").
		Select(SELECT_QUERY).
		Joins("JOIN passwords ON users.id = passwords.user_id").
		Where("users.email = ?", email).
		First(&userWithPassword)

	if result.RowsAffected == 0 {
		return UserWithPassword{}, errors.New("user not found")
	}

	return *result.Value.(*UserWithPassword), nil
}
