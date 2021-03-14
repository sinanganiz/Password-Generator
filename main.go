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

	// initialize database connection
	database.InitDatabase()

	app := fiber.New(fiber.Config{
		Views: django.New("./views", ".django"),
	})

	// middleware
	app.Use(recover.New())
	app.Use(logger.New())

	app.Static("/", "./views")

	// routing functions
	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
