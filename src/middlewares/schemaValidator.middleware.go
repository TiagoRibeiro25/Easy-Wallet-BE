package middlewares

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"easy-wallet-be/src/utils"

	"github.com/labstack/echo/v4"
	"github.com/xeipuuv/gojsonschema"
)

// ValidateJSONSchema returns an echo middleware function that validates the request body against a JSON schema.
// It takes a schemaPath string parameter that specifies the path to the JSON schema file.
// The function reads the request body and converts it to a string, then verifies if it is valid JSON data.
// If the request body is valid JSON data, it validates it against the specified JSON schema.
// If the validation fails, it returns an error response with a description of the validation error.
// If the validation succeeds, it resets the request body so that it can be read again by subsequent handlers.
func ValidateJSONSchema(schemaPath string) echo.MiddlewareFunc {
	loader := gojsonschema.NewReferenceLoader("file://src/data/schemas/" + schemaPath + "/schema.json")
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

			// Verify if the request body is valid JSON data
			if !isValidJSON(requestBody) {
				return utils.HandleResponse(
					c,
					http.StatusBadRequest,
					"Invalid JSON data in the request body",
					nil,
				)
			}

			documentLoader := gojsonschema.NewStringLoader(requestBody)
			result, err := schema.Validate(documentLoader)
			if err != nil || !result.Valid() {
				return utils.HandleResponse(
					c,
					http.StatusBadRequest,
					result.Errors()[0].Description()+" (at "+result.Errors()[0].Field()+")",
					nil,
				)
			}

			// Reset the request body so that it can be read again by subsequent handlers
			c.Request().Body = io.NopCloser(strings.NewReader(requestBody))

			return next(c)
		}
	}
}

// isValidJSON checks if a given string represents valid JSON data.
func isValidJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}
