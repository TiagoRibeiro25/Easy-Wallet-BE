package templates

import "os"

// VerifyUser generates an HTML string containing a link to the frontend route for verifying a user's account.
// The link includes a token parameter that is used to identify the user.
// The frontend route is obtained from the FRONTEND_URL environment variable.
func VerifyUser(token string) string {
	frontendRoute := os.Getenv("FRONTEND_URL") + "/auth/verify-user/" + token

	return "<h4>Verify User Account</h4><a href=" + frontendRoute + ">" + frontendRoute + "</a>"
}
