package saveSensorData

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"home-server/clients"
	"home-server/models"
	"time"
)

func SaveSensorData(c *fiber.Ctx) error {
	data := new(models.SensorData)
	if err := c.QueryParser(data); err != nil {
		return err
	}
	fmt.Printf("%s, temp: %f, hum: %f", time.Now(), data.Temperature, data.Humidity)
	fmt.Println()
	resultId, err := clients.WriteSensorData(data)
	err = c.SendString("Data is not saved ((")
	if err != nil {
		return err
	}
	result := fmt.Sprintf("Data is saved, id: %s", resultId)
	return c.SendString(result)
}
