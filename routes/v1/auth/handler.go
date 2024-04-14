package auth

import "github.com/gofiber/fiber/v3"

func Login(ctx fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{
		"Message": "Hello from User!",
	})
}
