package main

import (
	"log"
	"github.com/Miranlfk/golang-experiments/handlers"
	"github.com/Miranlfk/golang-experiments/db"
	"github.com/gofiber/fiber/v2"
)


func main(){
	if err := db.ConnectDB(); err != nil{
		log.Fatal(err)
	}
	
	// Create a new Fiber instance
	app := fiber.New()

	// Defining the routes of the application
	app.Get("/employees", handlers.GetAllEmployees)
	app.Get("/employees/:id", handlers.GetEmployee)
	app.Post("/employees", handlers.CreateEmployee)
	app.Put("/employees/:id", handlers.UpdateEmployee)
	app.Delete("/employees/:id", handlers.DeleteEmployee)

	(app.Listen(":3000"))
}