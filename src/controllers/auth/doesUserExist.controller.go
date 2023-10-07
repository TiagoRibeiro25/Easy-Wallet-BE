package controllers

import (
	"easy-wallet-be/src/models"
)

// DoesUserExist checks if a user with the given email already exists in the database.
// It returns true if the user exists, false otherwise.
func DoesUserExist(email string) bool {
	db := models.DB()

	user := db.Where("email = ?", email).First(&models.User{})

	return user.RowsAffected > 0
}
