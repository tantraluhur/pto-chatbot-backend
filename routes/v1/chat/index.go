package chat

import (
	"chatbot-backend/middlewares"

	"github.com/gofiber/fiber/v2"
)

func ChatRouter(chatRouter fiber.Router) {
	chatRouter = chatRouter.Group("/chat")
	chatRouter.Get("/chat-session", middlewares.RequireLoggedIn(), GetAllChatSession)
	chatRouter.Post("/chat-session", middlewares.RequireLoggedIn(), CreateChatSession)

	chatRouter.Get("/:id", middlewares.RequireLoggedIn(), GetMessageByChatSession)
	chatRouter.Post("/:id", middlewares.RequireLoggedIn(), CreateMessage)
}
