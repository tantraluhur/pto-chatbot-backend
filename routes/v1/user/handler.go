package user

import (
	"chatbot-backend/services/commons"
	"chatbot-backend/services/user"

	"github.com/gofiber/fiber/v2"
)

func GetUserHandler(ctx *fiber.Ctx) error {
	user, err := user.GetLoggedInUser(ctx)

	if err != nil {
		return ctx.Status(err.Code).JSON(commons.HTTPErrorResponse(err.Message))
	}

	return ctx.Status(200).JSON(commons.HTTPResponse(user))
}
