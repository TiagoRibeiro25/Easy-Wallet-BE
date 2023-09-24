package controllers

import (
	schemas "easy-wallet-be/src/data/schemas/auth/register"
	"easy-wallet-be/src/models"
	"easy-wallet-be/src/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// DB_ERROR is a constant that holds the error message for database errors (used multiple times)
const DB_ERROR = "An error occured with our database"
func Register(c echo.Context, bodyData schemas.BodyData) error {
	// Generate tokens
	verifyUserToken, verifyUserTokenErr := utils.GenerateToken()
	resetPasswordToken, resetPasswordTokenErr := utils.GenerateToken()

	// Hash the password
	hashedPassword, hashedPasswordErr := utils.HashPassword(bodyData.Password)

	// Check if there was an error while generating tokens or hashing the password
	if verifyUserTokenErr != nil || resetPasswordTokenErr != nil || hashedPasswordErr != nil {
		return utils.HandleResponse(
			c,
			http.StatusInternalServerError,
			"An error occured while creating the user",
			nil,
		)
	}

	// Get the database instance
	db := models.DB()

	// Create the user and password in a single transaction
	tx := db.Begin()
	if tx.Error != nil {
		return utils.HandleResponse(
			c,
			http.StatusInternalServerError,
			DB_ERROR,
			nil,
		)
	}

	user := models.User{
		Email:           bodyData.Email,
		DisplayName:     bodyData.DisplayName,
		VerifyUserToken: verifyUserToken,
	}

	// Create both user and password in a single transaction
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return utils.HandleResponse(
			c,
			http.StatusInternalServerError,
			DB_ERROR,
			nil,
		)
	}

	password := models.Password{
		Password:           hashedPassword,
		ResetPasswordToken: resetPasswordToken,
		UserID:             user.ID,
	}

	if err := tx.Create(&password).Error; err != nil {
		tx.Rollback()
		return utils.HandleResponse(
			c,
			http.StatusInternalServerError,
			DB_ERROR,
			nil,
		)
	}

	tx.Commit() // Commit the transaction if all operations are successful

	//TODO: Send an email to the user with the verification token
	//TODO: Add a cron job to delete the token after 15 minutes and create a new one
	//TODO: Add a cron job to delete the user after 24 hours if the user is not verified

	return utils.HandleResponse(
		c,
		http.StatusCreated,
		"User created successfully",
		schemas.ResponseData{
			ID:           user.ID,
			Email:        user.Email,
			DisplayName:  user.DisplayName,
			Currency:     user.Currency,
			UserVerified: user.UserVerified,
			CreatedAt:    user.CreatedAt.String(),
		},
	)
}
