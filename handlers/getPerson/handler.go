package getPerson

import (
	"github.com/gofiber/fiber/v2"
)

func GetPerson(c *fiber.Ctx) error {
	return c.SendString("Hello, world!")
}
