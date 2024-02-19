package controllers

import (
	"easy-wallet-be/src/models"
)

func GetExpenses(userID uint, date string, page int, limit int) ([]models.Expense, error) {
	db := models.DB()

	expenses := []models.Expense{}

	err := db.Where("user_id = ?", userID).
		Where("date = ?", date).
		Offset((page - 1) * limit).
		Limit(limit).
		Find(&expenses).Error

	if err != nil {
		return nil, err
	}

	return expenses, nil
}
