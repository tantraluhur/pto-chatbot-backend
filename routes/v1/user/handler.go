package user

import "github.com/gofiber/fiber/v3"

func GetUserHandler(ctx fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{
		"Message": "Hello from User!",
	})
}
