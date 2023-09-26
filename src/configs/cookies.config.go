package configs

type CookiesData struct {
	AuthCookieName string
}

// GetCookiesConfig returns the names of all the cookies used in the app
func GetCookiesConfig() CookiesData {
	return CookiesData{
		AuthCookieName: "easywallet-session-id", // Cookie used for session management
	}
}
