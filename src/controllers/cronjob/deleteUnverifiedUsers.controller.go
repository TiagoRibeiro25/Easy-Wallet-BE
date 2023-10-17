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
	thresholdTime := time.Now().Add(-24 * time.Hour)

	// Use a single query to fetch user IDs that meet the deletion criteria
	usersToDeleteQuery := db.Table("users").
		Select("id").
		Where("user_verified = ? AND created_at <= ?", false, thresholdTime)

	if err := usersToDeleteQuery.Error; err != nil {
		return "", err
	}

	var usersToDelete []struct {
		ID uint
	}
	if err := usersToDeleteQuery.Find(&usersToDelete).Error; err != nil {
		return "", err
	}

	if len(usersToDelete) == 0 {
		return "No unverified users to delete", nil
	}

	// Extract the IDs of users to delete
	var userIDs []uint
	for _, user := range usersToDelete {
		userIDs = append(userIDs, user.ID)
	}

	// Use a single transaction to delete users and their associated records
	tx := db.Begin()
	if err := tx.Table("users").
		Where("id IN (?)", userIDs).
		Delete(&models.User{}).
		Error; err != nil {
		tx.Rollback()
		return "", err
	}

	// Define associated tables to delete records from
	associatedTables := []struct {
		Name  string
		Model interface{}
	}{
		{"passwords", &models.Password{}},
		{"incomes", &models.Income{}},
		{"expenses", &models.Expense{}},
		{"categories", &models.Category{}},
	}

	// Delete records from associated tables
	for _, table := range associatedTables {
		if err := tx.Table(table.Name).
			Where("user_id IN (?)", userIDs).
			Delete(table.Model).Error; err != nil {
			tx.Rollback()
			return "", err
		}
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return "", err
	}

	return "Successfully deleted " + strconv.Itoa(len(userIDs)) + " unverified users", nil
}
