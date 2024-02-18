package controllers

import (
	"easy-wallet-be/src/models"
)

type NewExpense = struct {
	Name        string
	Cost        int
	Date        string
	Description string
	CategoryID  uint
}

// AddExpense adds a new expense for a given user.
// It takes the user ID and new expense data as parameters.
// Returns the created expense and an error, if any.
func AddExpense(userID uint, newExpenseData NewExpense) (models.Expense, error) {
	db := models.DB()

	newExpense := models.Expense{
		Name:        newExpenseData.Name,
		UserID:      userID,
		Cost:        newExpenseData.Cost,
		Date:        newExpenseData.Date,
		Description: newExpenseData.Description,
		CategoryID:  newExpenseData.CategoryID,
	}

	if err := db.Create(&newExpense).Error; err != nil {
		return models.Expense{}, err
	}

	return newExpense, nil
}
