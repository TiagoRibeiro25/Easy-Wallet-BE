package controllers

import "easy-wallet-be/src/models"

// ValidateSession validates a session by checking if a session with the given sessionID exists in the database.
// Returns true if the session exists, false otherwise.
func ValidateSession(sessionID string) bool {
	db := models.DB()

	session := db.Where("session_id = ?", sessionID).First(&models.Session{})

	return session.RowsAffected > 0
}
