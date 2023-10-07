package controllers

import (
	"easy-wallet-be/src/models"
	"errors"
)

// GetSession retrieves a session from the database by session ID.
// It returns a Session model and an error if the session is not found.
func GetSession(sessionID string) (models.Session, error) {
	db := models.DB()

	var session models.Session
	db.Where("session_id = ?", sessionID).First(&session)

	// Check if the session exists
	if session.SessionID != sessionID {
		return session, errors.New("session not found")
	}

	return session, nil
}
