package controllers

import (
	"easy-wallet-be/src/models"
	"time"
)

// DeleteExpiredSessions deletes expired sessions from the database.
// It returns the number of deleted sessions and an error (if any).
func DeleteExpiredSessions() (int, error) {
	// Get the database instance
	db := models.DB()

	// Find sessions where the "ExpiresAt" field is less than or equal to the current time
	var expiredSessions []models.Session
	if err := db.Where("expires_at <= ?", time.Now()).Find(&expiredSessions).Error; err != nil {
		return 0, err
	}

	if len(expiredSessions) == 0 {
		return 0, nil
	}

	// Extract the IDs of sessions to delete
	var sessionIDs []uint
	for _, session := range expiredSessions {
		sessionIDs = append(sessionIDs, session.ID)
	}

	// Use a single transaction to delete sessions and their associated records
	tx := db.Begin()
	if err := tx.Where("id IN (?)", sessionIDs).Delete(&models.Session{}).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	tx.Commit()

	return len(expiredSessions), nil
}
