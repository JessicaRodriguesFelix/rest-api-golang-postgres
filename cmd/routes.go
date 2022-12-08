package main

import (
	"github.com/JessicaRodriguesFelix/crud-api-golang-postgres/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListFacts)

	app.Get("fact", handlers.NewFactView)

	app.Post("/fact", handlers.CreateFact)
}