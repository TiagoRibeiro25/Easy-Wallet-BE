package controllers

import (
	schemas "easy-wallet-be/src/data/schemas/auth/register"
	"easy-wallet-be/src/models"
	"easy-wallet-be/src/utils"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Register(c echo.Context, bodyData schemas.BodyData) error {
	fmt.Println(bodyData)

	// Hash the password
	hashedPassword, err := utils.HashPassword(bodyData.Password)
	if err != nil {
		utils.HandleResponse(
			c,
			http.StatusInternalServerError,
			"Error hashing the password",
			nil,
		)
	}

	db := models.DB()

	// Create the user and password in a single transaction
	tx := db.Begin()
	if tx.Error != nil {
		return utils.HandleResponse(
			c,
			http.StatusInternalServerError,
			"Error starting transaction",
			nil,
		)
	}

	user := models.User{
		Email:           bodyData.Email,
		DisplayName:     bodyData.DisplayName,
		VerifyUserToken: "token", //TODO: Generate a random token
	}

	// Create both user and password in a single transaction
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return utils.HandleResponse(
			c,
			http.StatusInternalServerError,
			"Error creating user",
			nil,
		)
	}

	password := models.Password{
		Password:           hashedPassword,
		ResetPasswordToken: "token", //TODO: Generate a random token
		UserID:             user.ID,
	}

	if err := tx.Create(&password).Error; err != nil {
		tx.Rollback()
		return utils.HandleResponse(
			c,
			http.StatusInternalServerError,
			"Error creating password",
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
