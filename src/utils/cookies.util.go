package utils

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/labstack/echo/v4"
)

// CookiesConfig represents the configuration for cookies.
type CookiesConfig struct {
	Domain   string // Domain represents the domain name for the cookie.
	MaxAge   int    // MaxAge represents the maximum age of the cookie.
	Secure   bool   // Secure represents whether the cookie is secure or not.
	HttpOnly bool   // HttpOnly represents whether the cookie is accessible via HTTP only or not.
	Path     string // Path represents the path for the cookie.
}

// GetCookiesConfig returns the default configuration for cookies.
// If the app is running in a production environment, it updates the domain and secure fields accordingly.
func getCookiesConfig() CookiesConfig {
	cookiesConfig := CookiesConfig{
		Domain:   os.Getenv("DOMAIN"),
		MaxAge:   0,
		Secure:   false,
		HttpOnly: true,
		Path:     "/",
	}

	if IsProduction() {
		cookiesConfig.Secure = true
	}

	return cookiesConfig
}

// WriteCookie writes a cookie to the response.
// It takes the following parameters:
// context: an echo.Context object representing the HTTP request and response.
// name: a string representing the name of the cookie.
// value: a string representing the value of the cookie.
// expires: a time.Time object representing the expiration time of the cookie.
// It applies the default configuration for cookies and then applies the custom configuration for cookies.
func WriteCookie(context echo.Context, name string, value string, expires time.Time) {
	cookiesOptions := getCookiesConfig()

	cookie := new(http.Cookie)

	cookie.Domain = cookiesOptions.Domain
	cookie.MaxAge = cookiesOptions.MaxAge
	cookie.Secure = cookiesOptions.Secure
	cookie.HttpOnly = cookiesOptions.HttpOnly
	cookie.Path = cookiesOptions.Path

	cookie.Name = name
	cookie.Value = value
	cookie.Expires = expires

	context.SetCookie(cookie)
}

// DeleteCookie deletes a cookie with the given name from the context's response.
// It takes an echo.Context and the name of the cookie as input parameters.
// The function retrieves the cookie configuration options, creates a new cookie with the given name,
// sets its properties to delete the cookie, and adds it to the context's response.
func DeleteCookie(context echo.Context, name string) {
	cookiesOptions := getCookiesConfig()

	cookie := new(http.Cookie)

	cookie.Domain = cookiesOptions.Domain
	cookie.MaxAge = -1 // This will delete the cookie
	cookie.Secure = cookiesOptions.Secure
	cookie.HttpOnly = cookiesOptions.HttpOnly
	cookie.Path = cookiesOptions.Path
	cookie.Expires = time.Now().Add(-1 * time.Hour) // This will delete the cookie

	cookie.Name = name

	context.SetCookie(cookie)
}

// ReadCookie reads a cookie with the given name from the provided echo.Context.
// It returns the cookie value as a string and an error if the cookie is not found.
func ReadCookie(context echo.Context, name string) (string, error) {
	cookie, err := context.Cookie(name)
	if err != nil {
		return "", err
	}

	return cookie.Value, nil
}

// GetCookieExpiration returns the expiration time for a cookie based on the rememberMe flag.
// If rememberMe is true, the cookie will expire in 30 days. Otherwise, it will expire in 1 day.
func GetCookieExpiration(rememberMe ...bool) time.Time {
	if len(rememberMe) > 0 && rememberMe[0] {
		return time.Now().Add(720 * time.Hour) // 30 days
	}

	return time.Now().Add(24 * time.Hour) // 1 day
}

// ReadAllCookies reads all cookies from the provided echo.Context and prints their name and value to the console.
// It takes in an echo.Context as a parameter.
// Used for debugging purposes.
func ReadAllCookies(context echo.Context) {
	for _, cookie := range context.Cookies() {
		color.New(color.FgHiBlue).Print("COOKIE: ")
		color.New(color.FgHiWhite).Print(cookie.Name)
		fmt.Print(" = ")
		color.New(color.FgHiYellow).Print(cookie.Value)
		fmt.Println()
	}
}
