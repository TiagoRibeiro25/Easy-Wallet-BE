package utils

import (
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
