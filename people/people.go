package people

import (
	"github.com/MateuszGorski1/API-Fiber/database"
	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int64  `json:"age"`
}

func ShowPeople(c *fiber.Ctx) {
	db := database.DBConn
	var people []Person
	db.Find(&people)
	c.JSON(people)
}

func ShowPerson(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var person Person
	db.Find(&person, id)
	c.JSON(person)
}

func AddPerson(c *fiber.Ctx) {
	db := database.DBConn
	person := new(Person)
	if err := c.BodyParser(person); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&person)
	c.JSON(person)

}

func DeletePerson(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var person Person
	db.First(&person, id)
	if person.Name == "" {
		c.Status(500).Send("No person found with id")
	}
	db.Delete(&person)
	c.Send("Succesfully deleted")

}
