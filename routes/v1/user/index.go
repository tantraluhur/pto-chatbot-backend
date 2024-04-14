package user

import (
	"github.com/gofiber/fiber/v2"
)

func UserRouter(userRouter fiber.Router) {
	userRouter.Get("/user", GetUserHandler)
}
