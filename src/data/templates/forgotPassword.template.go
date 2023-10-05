package templates

import "os"

// ForgotPassword generates a HTML template for the forgot password email.
// It takes a token string as input and returns a string containing the HTML template.
// The HTML template contains a link to the frontend route for changing the password.
// The frontend route is constructed using the FRONTEND_URL environment variable and the token string.
func ForgotPassword(token string) string {
	frontendRoute := os.Getenv("FRONTEND_URL") + "/auth/change-password/" + token

	return "<h4>Change password</h4><a href=" + frontendRoute + " >" + frontendRoute + "</a>"
}
