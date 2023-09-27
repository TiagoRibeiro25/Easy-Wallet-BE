package utils

import "strconv"

// ConvertStringToUint converts a string to an unsigned integer.
// It takes a string as input and returns an unsigned integer and an error.
// If the string cannot be converted to an unsigned integer, it returns an error.
func ConvertStringToUint(str string) (uint, error) {
	nUint64, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0, err
	}

	return uint(nUint64), nil
}
