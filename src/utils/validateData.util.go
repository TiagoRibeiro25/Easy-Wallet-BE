package utils

import (
	"strconv"
	"time"
)

// Validates if the given string represents a valid date in the format "YYYY-MM-DD".
// It returns true if the date is valid, otherwise it returns false.
func IsDateValid(date string) bool {
	_, err := time.Parse("2006-01-02", date)
	return err == nil
}

// IsDateInFuture checks if the given date is in the future.
// It takes a string representation of a date in the format "2006-01-02" and returns a boolean value.
// The function returns true if the parsed date is after the current time, and false otherwise.
func IsDateInFuture(date string) bool {
	parsedDate, _ := time.Parse("2006-01-02", date)
	return parsedDate.After(time.Now())
}

// ArePageAndLimitFromQueryValid checks if the page and limit values from the query are valid.
// It takes in two strings, page and limit, and returns the parsed integers for page and limit,
// along with a boolean indicating if the values are valid.
// If either page or limit is not a valid integer or is less than 1, it returns false.
// Otherwise, it returns the parsed integers and true.
func ArePageAndLimitFromQueryValid(page, limit string) (int, int, bool) {
	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < 1 {
		return 0, 0, false
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil || limitInt < 1 {
		return 0, 0, false
	}

	return pageInt, limitInt, true
}
