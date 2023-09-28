package utils

import (
	"crypto/rand"
	"encoding/base64"
)

const DEFAULT_TOKEN_LENGTH = 32

// GenerateToken generates a random token of specified length or default length.
// It returns the generated token as a string and an error if any.
func GenerateToken(length ...int) (string, error) {
	// Set default length if no argument is passed
	tokenLength := DEFAULT_TOKEN_LENGTH
	if len(length) > 0 {
		tokenLength = length[0]
	}

	// Create a byte slice to store random bytes
	tokenBytes := make([]byte, tokenLength)

	// Read random bytes from the crypto/rand package
	if _, err := rand.Read(tokenBytes); err != nil {
		return "", err
	}

	// Encode the random bytes to base64 to make it a string
	token := base64.URLEncoding.EncodeToString(tokenBytes)

	return token, nil
}

// ValidateToken validates a token by attempting to decode it from base64 to bytes using Standard Encoding.
// If decoding with Standard Encoding fails, it tries with URLEncoding.
// Returns true if the length of the decoded bytes is not zero, false otherwise.
func ValidateToken(token string) bool {
	// Attempt to decode the token from base64 to bytes using Standard Encoding
	tokenBytes, err := base64.StdEncoding.DecodeString(token)

	if err != nil {
		// If decoding with Standard Encoding fails, try with URLEncoding
		tokenBytes, err = base64.URLEncoding.DecodeString(token)
		if err != nil {
			return false
		}
	}

	// Check if the length of the decoded bytes is zero
	return len(tokenBytes) != 0
}
