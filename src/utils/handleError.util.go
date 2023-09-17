package utils

import (
	"os"

	"github.com/fatih/color"
)

// The function `HandleError` prints an error message and optionally exits the program if an error occurs.
func HandleError(err error, message string, exit bool) {
	if err != nil {
		color.New(color.FgHiYellow).Print("Error: ")

		if message != "" {
			color.New(color.FgHiRed).Println(message)
		} else {
			color.New(color.FgHiRed).Println(err.Error())
		}

		if exit {
			os.Exit(1)
		}
	}
}
