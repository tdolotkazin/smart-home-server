package saveBoilerData

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"home-server/clients"
	"home-server/models"
	"time"
)

func SaveBoilerData(c *fiber.Ctx) error {
	data := new(models.BoilerDataIn)
	if err := c.QueryParser(data); err != nil {
		return err
	}
	fmt.Printf("%s, waterTemp: %f, isRunning: %t", time.Now(), data.WaterTemperature, data.IsRunning)
	fmt.Println()
	dataWithTime := models.BoilerDataOut{
		WaterTemperature: data.WaterTemperature,
		IsRunning:        data.IsRunning,
		Time:             time.Now(),
	}
	resultId, err := clients.WriteBoilerData(&dataWithTime)
	err = c.SendString("Data is not saved ((")
	if err != nil {
		return err
	}
	result := fmt.Sprintf("Data is saved, id: %s", resultId)
	return c.SendString(result)
}
