package main

import (
	"github.com/gofiber/fiber/v2"
	"home-server/handlers/getAllBoilerData"
	"home-server/handlers/getAllSensorsData"
	"home-server/handlers/latestData"
	"home-server/handlers/saveBoilerData"
	"home-server/handlers/saveSensorData"
)

func main() {
	app := fiber.New()
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("This is main page =)")
	})
	app.Post("/sensor", saveSensorData.SaveSensorData)
	app.Post("/boiler", saveBoilerData.SaveBoilerData)
	app.Get("/allSensors", getAllSensorsData.GetAllSensorsData)
	app.Get("/allBoiler", getAllBoilerData.GetAllBoilerData)
	app.Get("/latestData", latestData.LatestData)
	app.Listen("0.0.0.0:8000")
}
