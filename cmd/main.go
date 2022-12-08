package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/JessicaRodriguesFelix/crud-api-golang-postgres/database"
	"github.com/gofiber/template/html"
)

func main() {
	database.ConnectDb()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config {
		Views: engine,
	})

	setupRoutes(app)

	app.Listen(":3000")
}