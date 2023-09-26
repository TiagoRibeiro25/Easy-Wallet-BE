package controllers

import (
	"easy-wallet-be/src/models"
	"strconv"
	"time"
)

// DeleteUnverifiedUsers deletes users that have not been verified and were created more than 24 hours ago.
// It returns a string indicating the number of unverified users deleted or an error if the operation fails.
func DeleteUnverifiedUsers() (string, error) {
	db := models.DB()

	// Find users created more than 24 hours ago and not verified
	var usersIdsToDelete []uint
	thresholdTime := time.Now().Add(-24 * time.Hour)

	// Use a single query to fetch user IDs that meet the deletion criteria
	if err := db.Table("users").
		Where("user_verified = ? AND created_at <= ?", false, thresholdTime).
		Pluck("id", &usersIdsToDelete).Error; err != nil {
		return "", err
	}

	// Delete the unverified users with IDs in usersIdsToDelete from the "users" table in a single query
	if err := db.Table("users").Unscoped().
		Where("id IN (?)", usersIdsToDelete).
		Delete(&models.User{}).Error; err != nil {
		return "", err
	}

	nUsersDeleted := len(usersIdsToDelete)

	if nUsersDeleted == 0 {
		return "No unverified users to delete", nil
	}

	return "Successfully deleted " + strconv.Itoa(nUsersDeleted) + " unverified users", nil
}
