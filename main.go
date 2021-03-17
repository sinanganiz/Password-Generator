package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
	"github.com/markbates/pkger"

	"github.com/sinanganizz/password-generator/database"
	"github.com/sinanganizz/password-generator/router"
)

func main() {
	engine := html.NewFileSystem(pkger.Dir("/views"), ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use("/assets", filesystem.New(filesystem.Config{
		Root: pkger.Dir("/views/assets"),
	}))
	// app.Static("/", "./views")

	// initialize database connection
	database.InitDatabase()

	// routing functions
	router.SetupRoutes(app)

	// middleware
	app.Use(recover.New())
	app.Use(logger.New())

	log.Fatal(app.Listen(":3000"))
}
