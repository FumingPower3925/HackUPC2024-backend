package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/FumingPower3925/HackUPC2024-backend/pkg/airport"
)

func main() {
	app := fiber.New()

	app.Get("/test", func(c *fiber.Ctx) error {
		// var aux [][]airport.Entity = airport.DummyMap()
		// airport.PrintMatrix(&aux)
		airport.Test()
		return c.SendString("true")
	})

	app.Listen(":8080")
}
