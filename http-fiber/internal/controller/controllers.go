package controller

import "github.com/gofiber/fiber/v2"

func ApplyController(app *fiber.App) {
	app.Get("/", DefaultController)
}
