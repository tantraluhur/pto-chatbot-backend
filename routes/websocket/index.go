package websocket

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func WebscoketRouter(websocketRouter fiber.Router) {
	websocketRouter.Get("/ws:id", websocket.New(WebsocketHandler))
}
