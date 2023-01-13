package getAllSensorsData

import (
	"github.com/gofiber/fiber/v2"
	"home-server/clients"
)

func GetAllSensorsData(c *fiber.Ctx) error {
	result := clients.ReadSensorsData()
	// TODO: - Convert result to JSON
	return nil
}
