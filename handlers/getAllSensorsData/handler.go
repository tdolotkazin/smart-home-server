package getAllSensorsData

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"home-server/clients"
	"strconv"
)

func GetAllSensorsData(c *fiber.Ctx) error {
	days, err := strconv.Atoi(c.Query("days"))
	result := clients.ReadSensorsData(days)
	jsonData, err := json.Marshal(result)
	if err != nil {
		return err
	}
	c.SendString(string(jsonData))
	return nil
}
