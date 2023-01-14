package getAllSensorsData

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"home-server/clients"
)

func GetAllSensorsData(c *fiber.Ctx) error {
	result := clients.ReadSensorsData()
	jsonData, err := json.Marshal(result)
	if err != nil {
		return err
	}
	c.SendString(string(jsonData))
	return nil
}
