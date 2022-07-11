package server

import (
	"fmt"

	"github.com/MateuszGorski1/API-Fiber/database"
	"github.com/MateuszGorski1/API-Fiber/people"
	"github.com/gofiber/fiber"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/showPeople", people.ShowPeople)
	app.Get("/showPerson/:id", people.ShowPerson)
	app.Post("/addPerson", people.AddPerson)
	app.Delete("/deletePerson/:id", people.DeletePerson)
}

func connectToDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("people.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	database.DBConn.AutoMigrate(&people.Person{})
	fmt.Println("Succesfully connected")
}

func StartServer() {
	app := fiber.New()
	connectToDatabase()
	setupRoutes(app)
	app.Listen(8080)
}
