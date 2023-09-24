package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes a given password using bcrypt algorithm with default cost.
// It returns the hashed password as a string and an error if the hashing process fails.
func HashPassword(password string) (string, error) {
	// bcrypt.GenerateFromPassword generates a hashed password from the given password and cost.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// ComparePassword compares a hashed password with a plain text password and returns a boolean indicating if they match.
// It uses bcrypt to compare the passwords.
// hashedPassword: the hashed password to compare.
// password: the plain text password to compare.
// Returns true if the passwords match, false otherwise.
func ComparePassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
