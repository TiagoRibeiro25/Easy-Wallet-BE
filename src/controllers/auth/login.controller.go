package controllers

import (
	"easy-wallet-be/src/models"
	"easy-wallet-be/src/utils"
	"errors"
	"time"

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

// CreateSession creates a new session for a user.
// It takes a userID uint and a rememberMe bool as parameters.
// It returns a sessionID string and an error (if any).
// The function generates a sessionID using the GenerateToken function from the utils package.
// If rememberMe is true, the session will expire in 30 days, otherwise it will expire in 1 day.
// The function creates a new session in the database and returns the sessionID.
func CreateSession(userID uint, rememberMe bool) (string, error) {
	db := models.DB()
	sessionID, err := utils.GenerateToken()
	if err != nil {
		return "", err
	}

	// If the user wants to be remembered, set the expiration date to 30 days from now, otherwise set it to 1 day from now
	daysToExpire := 1
	if rememberMe {
		daysToExpire = 30
	}

	timeToExpire := time.Now().AddDate(0, 0, daysToExpire)

	session := models.Session{
		SessionID: sessionID,
		UserID:    uint(userID),
		ExpiresAt: timeToExpire,
	}

	if err := db.Create(&session).Error; err != nil {
		return "", err
	}

	return sessionID, nil
}
