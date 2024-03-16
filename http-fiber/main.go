package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"masasdani.com/http-fiber/config"
	"masasdani.com/http-fiber/controller"
)

func main() {
	// ...
	log.SetLevel(log.LevelInfo)
	app := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: config.ErrorHandlers,
	})

	applyController(app)

	app.Listen(":8000")
}

func applyController(app *fiber.App) {
	app.Get("/", controller.DefaultController)
}
