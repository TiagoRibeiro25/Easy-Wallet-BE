package controllers

import (
	schemas "easy-wallet-be/src/data/schemas/auth/register"
	"easy-wallet-be/src/models"
	"errors"

	"github.com/labstack/echo/v4"
)

// DB_ERROR is a constant that holds the error message for database errors (used multiple times)
const DB_ERROR = "an error occured with our database"

// Register creates a new user and password in the database.
// It receives the echo context, the body data containing the email and display name,
// the verify user token, the reset password token, and the hashed password.
// It returns a response data struct and an error.
// The function starts a transaction to create the user and password in a single transaction.
// It returns an error if the transaction fails.
// It creates a new user and password in the database.
// It returns an error if the creation fails.
// It commits the transaction if all operations are successful.
func Register(
	c echo.Context,
	bodyData schemas.BodyData,
	verifyUserToken string,
	resetPasswordToken string,
	hashedPassword string,
) (schemas.ResponseData, error) {
	var responseData schemas.ResponseData
	var err error

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
