package controllers

import (
	schemas "easy-wallet-be/src/data/schemas/auth/register"
	"easy-wallet-be/src/models"
	"easy-wallet-be/src/utils"
	"errors"

	"github.com/labstack/echo/v4"
)

// DB_ERROR is a constant that holds the error message for database errors (used multiple times)
const DB_ERROR = "an error occured with our database"

// Register creates a new user and password in the database, and returns a response data, a verification token and an error.
// It takes an echo context and a BodyData struct as input parameters.
// The function generates two tokens, one for user verification and another for password reset, and hashes the password.
// It then creates the user and password in a single transaction in the database.
// If any error occurs while generating tokens or hashing the password, the function returns an error.
// If any error occurs while creating the user or password in the database, the function rolls back the transaction and returns an error.
// If all operations are successful, the function commits the transaction and returns a response data and a verification token.
func Register(c echo.Context, bodyData schemas.BodyData) (schemas.ResponseData, string, error) {
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
		return responseData, "", err
	}

	// Get the database instance
	db := models.DB()

	// Create the user and password in a single transaction
	tx := db.Begin()
	if tx.Error != nil {
		err = errors.New(DB_ERROR)
		return responseData, "", err
	}

	user := models.User{
		Email:           bodyData.Email,
		DisplayName:     bodyData.DisplayName,
		VerifyUserToken: verifyUserToken,
	}

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		err = errors.New(DB_ERROR)
		return responseData, "", err
	}

	password := models.Password{
		Password:           hashedPassword,
		ResetPasswordToken: resetPasswordToken,
		UserID:             user.ID,
	}

	if err := tx.Create(&password).Error; err != nil {
		tx.Rollback()
		err = errors.New(DB_ERROR)
		return responseData, "", err
	}

	tx.Commit() // Commit the transaction if all operations are successful

	responseData = schemas.ResponseData{
		ID:           user.ID,
		Email:        user.Email,
		DisplayName:  user.DisplayName,
		Currency:     user.Currency,
		UserVerified: user.UserVerified,
		CreatedAt:    user.CreatedAt.String(),
	}

	return responseData, verifyUserToken, nil
}

// DoesUserExist checks if a user with the given email already exists in the database.
// It returns true if the user exists, false otherwise.
func DoesUserExist(email string) bool {
	db := models.DB()

	user := db.Where("email = ?", email).First(&models.User{})

	return user.RowsAffected > 0
}
