package auth

import (
	"chatbot-backend/middlewares"

	"github.com/gofiber/fiber/v2"
)

func AuthRouter(authRouter fiber.Router) {
	authRouter = authRouter.Group("/auth")
	authRouter.Post("/login", Login)
	authRouter.Post("/register", Register)
	authRouter.Get("/logout", middlewares.RequireLoggedIn(), Logout)
}
