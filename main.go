package main

import (
	"github.com/gofiber/fiber/v2"
	"home-server/handlers/getAllSensorsData"
	"home-server/handlers/saveSensorData"
)

func main() {
	app := fiber.New()
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("This is main page =)")
	})
	app.Post("/sensor", saveSensorData.SaveSensorData)
	app.Get("/allSensors", getAllSensorsData.GetAllSensorsData)
	app.Post("/person", createPerson)
	app.Put("/person/:id", updatePerson)
	app.Delete("/person/:id", deletePerson)
	app.Listen("0.0.0.0:8000")
}

func createPerson(c *fiber.Ctx) error {
	return c.SendString("Hello, world!")
}
func updatePerson(c *fiber.Ctx) error {
	return c.SendString("Hello, world!")
}
func deletePerson(c *fiber.Ctx) error {
	return c.SendString("Hello, world!")
}
