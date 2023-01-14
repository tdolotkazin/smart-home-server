package latestData

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"home-server/clients"
)

func LatestData(c *fiber.Ctx) error {
	result := clients.ReadLatestData()
	jsonData, err := json.Marshal(result)
	if err != nil {
		return err
	}
	c.SendString(string(jsonData))
	return nil
}
