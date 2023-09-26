package controllers

import (
	"easy-wallet-be/src/models"
	"easy-wallet-be/src/utils"
	"time"

	"github.com/labstack/echo/v4"
)

//TODO: Move all the logic except the database queries to the middlewares package.

func ValidateSession(c echo.Context, sessionID string) bool {
	db := models.DB()

	var session models.Session
	db.Where("session_id = ?", sessionID).First(&session)

	// Check if the session exists
	if session.SessionID != sessionID {
		return false
	}

	// Check if the session is expired
	if session.ExpiresAt.Before(time.Now()) {
		db.Unscoped().Delete(&session) // Delete the session from the database
		return false
	}

	// If the session expires in less than 30 minutes, delete the session and create a new one
	if session.ExpiresAt.Before(time.Now().Add(30 * time.Minute)) {
		db.Unscoped().Delete(&session) // Delete the session from the database
		newSessionID, err := CreateSession(session.UserID, session.RememberMe)
		if err != nil {
			return false
		}

		// Create a session cookie
		utils.WriteCookie(
			c,
			"easywallet-session-id",
			newSessionID,
			utils.GetCookieExpiration(session.RememberMe),
		)
	}

	return true
}
