package middlewares

import (
	response "chatbot-backend/services/commons"
	"net/http"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func RequireLoggedIn() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte("secret"),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(http.StatusUnauthorized).JSON(response.HTTPErrorResponse("Missing or malformed authentication token."))

	}

	if err.Error() == "signature is invalid" {
		return c.Status(http.StatusUnauthorized).JSON(response.HTTPErrorResponse("Invalid access token."))
	}

	return c.Status(http.StatusUnauthorized).JSON(response.HTTPErrorResponse("Access token already expired."))
}
