package main

import (
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/FumingPower3925/HackUPC2024-backend/pkg/ai"
	"github.com/FumingPower3925/HackUPC2024-backend/pkg/airport"
)

var users map[string]User

type User struct {
	target airport.Point
	//rotation  int
	lastPoint airport.Point
}

func init() {
	users = make(map[string]User)
}

func main() {
	app := fiber.New()

	app.Get("/test", func(c *fiber.Ctx) error {

		output := "true"

		if !airport.TestLogic() {
			output = "false"
		}

		return c.SendString(output)
	})

	app.Get("/start", func(c *fiber.Ctx) error {
		//rot := c.Params("rotation")
		//roti, _ := strconv.Atoi(rot)
		uid := uuid.NewString()
		target, err := airport.GetTarget(uid)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "client has no flight availiable")
		}
		users[uid] = User{target: target, lastPoint: airport.Point{}}
		return c.SendString(uid)
	})

	app.Post("/gps", func(c *fiber.Ctx) error {
		clientId := c.Params("clientId") // clientId obtained when start
		newPos := c.Params("location")   // position of the client
		usr, notFound := users[clientId]
		if notFound {
			return fiber.NewError(fiber.StatusUnauthorized, "clientId does not exist")
		}
		point, err := airport.Gps2D(newPos)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "gps wrong coords")
		}
		if reflect.DeepEqual(usr.lastPoint, airport.Point{}) {
			usr.lastPoint = point
		}
		nstep, err := airport.NextStep(point, usr.target, usr.lastPoint)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "gps wrong coords")
		}
		if nstep == "ARRIVED" {
			c.Set("Arrived", "true")
		} else {
			c.Set("Arrived", "false")
		}
		path, err := ai.GetCommandVoice(nstep)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Internal server error")
		}
		return c.SendFile(path)
	})

	app.Listen(":8080")
}
