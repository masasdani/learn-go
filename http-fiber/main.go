package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"masasdani.com/http-fiber/internal/controller"
	"masasdani.com/http-fiber/pkg"
)

func main() {
	// ...
	log.SetLevel(log.LevelInfo)

	pkg.InitConfig()

	app := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: pkg.ErrorHandlers,
	})

	controller.ApplyController(app)

	app.Listen(fmt.Sprintf(":%v", pkg.AppConfig.Server.Port))
}
