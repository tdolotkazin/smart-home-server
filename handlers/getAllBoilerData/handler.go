package getAllBoilerData

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"home-server/clients"
	"time"
)

func GetAllBoilerData(c *fiber.Ctx) error {
	from := c.Query("from")
	fromDate, err := time.Parse("2006-01-02", from)
	if err != nil {
		return err
	}
	to := c.Query("to")
	toDate, err := time.Parse("2006-01-02", to)
	if err != nil {
		return err
	}
	result := clients.ReadBoilerData(fromDate, toDate)

	jsonData, err := json.Marshal(result)
	if err != nil {
		return err
	}
	c.SendString(string(jsonData))
	return nil
}
