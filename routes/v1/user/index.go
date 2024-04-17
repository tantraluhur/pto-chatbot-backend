package user

import (
	"chatbot-backend/middlewares"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(userRouter fiber.Router) {
	userRouter.Get("/user", middlewares.RequireLoggedIn(), GetUserHandler)
}
