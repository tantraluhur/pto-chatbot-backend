package user

import (
	"github.com/gofiber/fiber/v3"
)

func UserRouter(userRouter fiber.Router) {
	userRouter.Get("/user", GetUserHandler)
}
