package middlewares

import (
	"io"
	"net/http"
	"strings"

	"easy-wallet-be/src/utils"

	"github.com/labstack/echo/v4"
	"github.com/xeipuuv/gojsonschema"
)

// ValidateJSONSchema is a middleware function that validates the JSON schema of the request body against a given schema file.
// It takes a schemaPath string as input, which is the path to the schema file relative to the "src/data/schemas" directory.
// It returns an echo.MiddlewareFunc that can be used as middleware in an Echo web server.
// If the schema is invalid or the request body is invalid JSON data, it returns an error response to the client.
func ValidateJSONSchema(schemaPath string) echo.MiddlewareFunc {
	loader := gojsonschema.NewReferenceLoader("file://src/data/schemas/" + schemaPath + ".schema.json")
	schema, err := gojsonschema.NewSchema(loader)
	utils.HandleError(err, "Failed to load schema", true)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Read the request body and convert it to a string
			bodyBytes, err := io.ReadAll(c.Request().Body)
			if err != nil {
				return utils.HandleResponse(
					c,
					http.StatusInternalServerError,
					"Failed to read request body",
					nil,
				)
			}
			requestBody := string(bodyBytes)

			documentLoader := gojsonschema.NewStringLoader(requestBody)
			result, err := schema.Validate(documentLoader)
			if err != nil || !result.Valid() {
				return utils.HandleResponse(
					c,
					http.StatusBadRequest,
					"Invalid JSON data",
					nil,
				)
			}

			// Reset the request body so that it can be read again by subsequent handlers
			c.Request().Body = io.NopCloser(strings.NewReader(requestBody))

			return next(c)
		}
	}
}
