package utils

import (
	"github.com/fatih/color"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// The function `HandleResponse` handles the response for an HTTP request by constructing a JSON
// response object and returning it.
func HandleResponse(context echo.Context, statusCode int, message string, data interface{}) error {
	if statusCode == 204 {
		return context.NoContent(statusCode)
	}

	response := Response{
		Success: statusCode >= 200 && statusCode < 300,
		Message: message,
		Data:    data,
	}

	// Print the response message depending on the status code of the response.
	if statusCode >= 500 {
		color.New(color.FgHiYellow).Print("Server Error: ")
		color.New(color.FgHiRed).Println(message)
	}

	return context.JSON(statusCode, response)
}
