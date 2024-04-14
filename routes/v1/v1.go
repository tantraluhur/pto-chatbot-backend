package v1

import (
	"chatbot-backend/routes/v1/user"

	"github.com/gofiber/fiber/v2"
)

func V1Router(v1Router fiber.Router) {
	user.UserRouter(v1Router)
}
