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
		"Title": "Golang- Trivia Game ",
		"Subtitle": "Questions and Answers for funtimes with coworkers!",
		"Facts": facts,
	})
}

func NewFactView(c *fiber.Ctx) error {
	return c.Render("new", fiber.Map{
		"Title" : "New Question",
		"Subtitle": "Add a cool question here!",
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
	return ConfirmationView(c)
}
///changes here
func DeleteFact(c *fiber.Ctx) error {
	facts:= []models.Fact{}
	id := c.Params("id")
	if id == "" {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "id cannot be empty",
		}) 
		return nil
	}
	database.DB.Db.Delete(facts, id)
	return ListFacts(c)
}

func ConfirmationView(c *fiber.Ctx) error {
	return c.Render("confirmation", fiber.Map{
		"Title": "Question added...Yay!!",
		"Subtitle": "Add more fun questions to the list",

	})
} 