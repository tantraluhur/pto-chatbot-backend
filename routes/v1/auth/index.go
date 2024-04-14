package auth

import (
	"github.com/gofiber/fiber/v3"
)

func AuthRouter(authRouter fiber.Router) {
	authRouter.Post("/login", Login)
}
