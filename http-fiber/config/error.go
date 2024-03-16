package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func ErrorHandlers(c *fiber.Ctx, err error) error {
	log.Warn(err)
	code := fiber.StatusInternalServerError
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	resp := Response{
		Code:    code,
		Message: err.Error(),
	}

	return c.Status(code).JSON(resp)
}
