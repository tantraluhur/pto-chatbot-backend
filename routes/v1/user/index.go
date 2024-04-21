package user

import (
	"chatbot-backend/middlewares"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(userRouter fiber.Router) {
	userRouter = userRouter.Group("/user")

	userRouter.Get("/", middlewares.RequireLoggedIn(), GetUserHandler)
}
