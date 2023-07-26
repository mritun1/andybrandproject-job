package main

import (
	"andybrandproject/controllers"
	"andybrandproject/db"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	err_env := godotenv.Load()
	if err_env != nil {
		log.Fatal(".env not available")
	}

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
	app.Post("/login", controllers.LoginUsers)
	// -------------------------------
	// END ROUTING
	// -------------------------------

	app.Listen(":4000")
}
