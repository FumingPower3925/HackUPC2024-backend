package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/FumingPower3925/HackUPC2024-backend/pkg/airport"
)

func main() {
	app := fiber.New()

	app.Get("/test", func(c *fiber.Ctx) error {

		output := "true"

		if !airport.TestLogic() {
			output = "false"
		}

		return c.SendString(output)
	})

	app.Listen(":8080")
}
