package main

import (
	"github.com/JessicaRodriguesFelix/crud-api-golang-postgres/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListFacts)

	app.Get("fact", handlers.NewFactView)

	//new 
	app.Delete("/fact/:id", handlers.DeleteFact)

	app.Post("/fact", handlers.CreateFact)
}