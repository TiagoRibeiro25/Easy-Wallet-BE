package utils

import "os"

// The function GetEnv retrieves the value of an environment variable given its key
func GetEnv(key string) string {
	return os.Getenv(key)
}

// The function SetEnv sets the value of an environment variable
func SetEnv(key string, value string) {
	os.Setenv(key, value)
}
