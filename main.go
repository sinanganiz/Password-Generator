package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/django"
	"github.com/sinanganizz/password-generator/database"
	"github.com/sinanganizz/password-generator/router"
)

func main() {

	app := fiber.New(fiber.Config{
		Views: django.New("./views", ".django"),
	})

	app.Static("/", "./views")

	// initialize database connection
	database.InitDatabase()

	// routing functions
	router.SetupRoutes(app)

	// middleware
	app.Use(recover.New())
	app.Use(logger.New())

	log.Fatal(app.Listen(":3000"))
}
