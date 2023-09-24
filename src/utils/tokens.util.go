package utils

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateToken generates a random token of specified length or default length of 32.
// It returns the generated token as a string and an error if any.
func GenerateToken(length ...int) (string, error) {
	// Set default length to 32 if no argument is passed
	tokenLength := 32
	if len(length) > 0 {
		tokenLength = length[0]
	}

	// Create a byte slice to store random bytes
	tokenBytes := make([]byte, tokenLength)

	// Read random bytes from the crypto/rand package
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", err
	}

	// Encode the random bytes to base64 to make it a string
	token := base64.URLEncoding.EncodeToString(tokenBytes)

	return token, nil
}
