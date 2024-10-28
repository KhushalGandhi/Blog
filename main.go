package main

import (
	"Blog/database"
	"Blog/migrations"
	"Blog/routes"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	// Initialize Database
	database.Connect()
	migrations.Migrate(database.DB)

	app := fiber.New()

	// Setup routes
	routes.Setup(app)

	log.Fatal(app.Listen(":8080"))
}
