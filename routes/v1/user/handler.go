package user

import (
	user "chatbot-backend/services/auth"
	"chatbot-backend/services/commons"

	"github.com/gofiber/fiber/v2"
)

func GetUserHandler(ctx *fiber.Ctx) error {
	user, err := user.GetLoggedInUser(ctx)

	if err != nil {
		return ctx.Status(err.Code).JSON(commons.HTTPErrorResponse(err.Message))
	}

	return ctx.Status(200).JSON(fiber.Map{
		"Message": "Hello from " + user.Name,
	})
}
