package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sukenda/go-restful-api/exception"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	}
}
