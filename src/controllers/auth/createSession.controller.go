package controllers

import (
	"easy-wallet-be/src/models"
	"easy-wallet-be/src/utils"
	"time"
)

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
		SessionID:  sessionID,
		RememberMe: rememberMe,
		UserID:     uint(userID),
		ExpiresAt:  timeToExpire,
	}

	if err := db.Create(&session).Error; err != nil {
		return "", err
	}

	return sessionID, nil
}
