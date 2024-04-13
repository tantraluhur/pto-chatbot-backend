package main

import (
	"chatbot-backend/database"
	"chatbot-backend/routes"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	//Init Database & Migrations
	database.DatabaseInit()
	database.Migration()

	//Init Router
	routes.RouteInit(app)

	app.Listen(":8000")
}
