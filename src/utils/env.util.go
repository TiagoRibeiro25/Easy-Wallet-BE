package utils

import "os"

func IsProduction() bool {
	return os.Getenv("GO_ENV") == "production" || os.Getenv("GO_ENV") == "prod"
}
