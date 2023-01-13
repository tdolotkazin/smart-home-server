package saveSensorData

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

type SensorData struct {
	Temperature float32 `json:"temp" xml:"temp" form:"temp" query:"temp"`
	Humidity    float32 `json:"hum" xml:"hum" form:"hum" query:"hum"`
}

func SaveSensorData(c *fiber.Ctx) error {
	data := new(SensorData)
	if err := c.QueryParser(data); err != nil {
		return err
	}
	fmt.Printf("%s, temp: %f, hum: %f", time.Now(), data.Temperature, data.Humidity)
	fmt.Println()
	return c.SendString("Data is saved")
}
