package controller

import (
	"github.com/gofiber/fiber/v2"
	"masasdani.com/http-fiber/config"
)

func DefaultController(c *fiber.Ctx) error {
	return c.JSON(config.Response{
		Code:    200,
		Message: "Hello, World!",
	})
}
