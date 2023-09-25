package controllers

import (
	schemas "easy-wallet-be/src/data/schemas/auth/register"
	"easy-wallet-be/src/models"
	"easy-wallet-be/src/services"
	"easy-wallet-be/src/utils"
	"errors"

	"github.com/labstack/echo/v4"
)

const DB_ERROR = "an error occured with our database"

func Register(c echo.Context, bodyData schemas.BodyData) (schemas.ResponseData, error) {
	var responseData schemas.ResponseData
	var err error

	verifyUserToken, verifyUserTokenErr := utils.GenerateToken()
	resetPasswordToken, resetPasswordTokenErr := utils.GenerateToken()

	hashedPassword, hashedPasswordErr := utils.HashPassword(bodyData.Password)

	if verifyUserTokenErr != nil || resetPasswordTokenErr != nil || hashedPasswordErr != nil {
		err = errors.New("an error occured while creating the user")
		return responseData, err
	}

	db := models.DB()

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

	tx.Commit()

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
