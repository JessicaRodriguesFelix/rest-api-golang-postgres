package handlers

import (
	"github.com/JessicaRodriguesFelix/crud-api-golang-postgres/database"
	"github.com/JessicaRodriguesFelix/crud-api-golang-postgres/models"
	"github.com/gofiber/fiber/v2"
)

func ListFacts(c *fiber.Ctx) error {
	facts:= []models.Fact{}
	database.DB.Db.Find(&facts)
	return c.Render("index", fiber.Map{
		"Title": "Div Golang API",
		"Subtitle": "Questions and Answers for funtimes with friends!",
	})
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	database.DB.Db.Create(&fact)
	return c.Status(200).JSON(fact)
}