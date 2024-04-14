package auth

import (
	"github.com/gofiber/fiber/v2"
)

func AuthRouter(authRouter fiber.Router) {
	authRouter.Post("/login", Login)
}
