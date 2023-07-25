package main

import (
	"andybrandproject/controllers"
	"andybrandproject/db"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	if db.Con == nil {
		log.Fatal("MongoDB collection is nil")
	}
	// -------------------------------
	// START ROUTING
	// -------------------------------
	app.Get("/", controllers.Home)
	app.Get("/users", controllers.Users)
	app.Post("/users", controllers.CreateUsers)
	app.Put("/users/:id", controllers.UpdateUsers)
	app.Delete("/users/:id", controllers.Delete)
	// -------------------------------
	// END ROUTING
	// -------------------------------

	app.Listen(":4000")
}
