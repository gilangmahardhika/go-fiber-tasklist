package tasks

import (

	"github.com/gofiber/fiber/v2"
)

func GetHelloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
