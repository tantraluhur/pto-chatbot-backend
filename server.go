package main

import (
	"chatbot-backend/database"
	"chatbot-backend/routes"

	ws "chatbot-backend/routes/websocket"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	//Init Websocket Router
	app.Use("/ws", ws.WsUpgrade)
	app.Get("/ws/:id", websocket.New(ws.WebsocketHandler))

	//Init Database & Migrations
	database.DatabaseInit()
	database.Migration()

	//Cors Setting
	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin, Content-Type, Accept, Content-Length, Accept-Language, Accept-Encoding, Connection, Access-Control-Allow-Origin",
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	//Init Router
	routes.RouteInit(app)

	app.Listen(":8000")
}
