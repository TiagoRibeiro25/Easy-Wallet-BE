package controllers

import (
	"easy-wallet-be/src/models"
	"errors"
)

// DeleteSession deletes a session from the database by session ID.
// It returns an error if the session is not found or if there is an error while deleting it.
func DeleteSession(sessionID string) error {
	db := models.DB()

	result := db.Where("session_id = ?", sessionID).Delete(&models.Session{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("session not found")
	}

	return nil
}
