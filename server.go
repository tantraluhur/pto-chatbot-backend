package main

import (
	"chatbot-backend/routes"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	//ROUTER INIT
	routes.RouteInit(app)

	app.Listen(":8000")
}
