package auth

import "github.com/gofiber/fiber/v2"

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func Login(ctx *fiber.Ctx) error {
	var body = LoginRequest{}

	if err := ctx.BodyParser(body); err != nil {
		return err
	}
	return ctx.Status(200).JSON(fiber.Map{
		"Message": "Hello from User!",
	})
}
