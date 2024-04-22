package v1

import (
	"chatbot-backend/routes/v1/auth"
	"chatbot-backend/routes/v1/chat"
	"chatbot-backend/routes/v1/user"

	"github.com/gofiber/fiber/v2"
)

func V1Router(v1Router fiber.Router) {
	user.UserRouter(v1Router)
	auth.AuthRouter(v1Router)
	chat.ChatRouter(v1Router)
}
