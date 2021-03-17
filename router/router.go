package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sinanganizz/password-generator/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", handlers.Homepage)
	app.Get("/MyPasswords", handlers.MyPasswords)
	app.Get("/DeletePassword", handlers.DeletePassword)
	app.Get("/Edit", handlers.Edit)
	app.Post("/EditConfirm", handlers.EditConfirm)
}
