package routes

import (
	v1 "chatbot-backend/routes/v1"

	"github.com/gofiber/fiber/v3"
)

func RouteInit(app *fiber.App) {
	//Grouping Router API
	api := app.Group("/api")

	//Grouping V1 API
	version1 := api.Group("/v1")

	//INIT Router Version 1
	v1.V1Router(version1)
}
