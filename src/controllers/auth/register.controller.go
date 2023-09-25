package controllers

import (
	schemas "easy-wallet-be/src/data/schemas/auth/register"
	"easy-wallet-be/src/models"
	"easy-wallet-be/src/services"
	"easy-wallet-be/src/utils"
	"errors"

	"github.com/labstack/echo/v4"
)

// DB_ERROR is a constant that holds the error message for database errors (used multiple times)
const DB_ERROR = "an error occured with our database"

// Register creates a new user and password in the database, generates tokens, hashes the password, and sends a verification email.
// It receives an echo context and a BodyData struct containing the email, display name, and password of the user to be created.
// It returns a ResponseData struct containing the user's ID, email, display name, currency, user verification status, and creation date, and an error.
func Register(c echo.Context, bodyData schemas.BodyData) (schemas.ResponseData, error) {
	var responseData schemas.ResponseData
	var err error

	// Generate tokens
	verifyUserToken, verifyUserTokenErr := utils.GenerateToken()
	resetPasswordToken, resetPasswordTokenErr := utils.GenerateToken()

	// Hash password
	hashedPassword, hashedPasswordErr := utils.HashPassword(bodyData.Password)

	// Check if there was an error while generating tokens or hashing the password
	if verifyUserTokenErr != nil || resetPasswordTokenErr != nil || hashedPasswordErr != nil {
		err = errors.New("an error occured while creating the user")
		return responseData, err
	}

	// Get the database instance
	db := models.DB()

	// Create the user and password in a single transaction
	tx := db.Begin()
	if tx.Error != nil {
		err = errors.New(DB_ERROR)
		return responseData, err
	}

	user := models.User{
		Email:           bodyData.Email,
		DisplayName:     bodyData.DisplayName,
		VerifyUserToken: verifyUserToken,
	}

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		err = errors.New(DB_ERROR)
		return responseData, err
	}

	password := models.Password{
		Password:           hashedPassword,
		ResetPasswordToken: resetPasswordToken,
		UserID:             user.ID,
	}

	if err := tx.Create(&password).Error; err != nil {
		tx.Rollback()
		err = errors.New(DB_ERROR)
		return responseData, err
	}

	tx.Commit() // Commit the transaction if all operations are successful

	// Send the verification email
	services.SendEmail(
		bodyData.Email,
		bodyData.DisplayName,
		"Welcome to Easy Wallet",
		"<h4>Verify User Token</h4><p>"+verifyUserToken+"</p>",
	)

	responseData = schemas.ResponseData{
		ID:           user.ID,
		Email:        user.Email,
		DisplayName:  user.DisplayName,
		Currency:     user.Currency,
		UserVerified: user.UserVerified,
		CreatedAt:    user.CreatedAt.String(),
	}

	return responseData, nil
}
