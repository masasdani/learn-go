package controller

import (
	"github.com/gofiber/fiber/v2"
	"masasdani.com/http-fiber/pkg"
)

func DefaultController(c *fiber.Ctx) error {
	return c.JSON(pkg.Response{
		Code:    200,
		Message: "Hello, World!",
	})
}
