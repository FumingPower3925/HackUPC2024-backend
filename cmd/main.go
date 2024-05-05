package main

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/FumingPower3925/HackUPC2024-backend/pkg/ai"
	"github.com/FumingPower3925/HackUPC2024-backend/pkg/airport"
)

var users map[string]User

type User struct {
	target    airport.Point
	rotation  int
	lastPoint airport.Point
}

func init() {
	users = make(map[string]User)
}

func main() {
	app := fiber.New()

	app.Static("/public/", "./public")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./views/index.html")
	})

	app.Get("/test", func(c *fiber.Ctx) error {

		output := "true"

		if !airport.TestLogic() {
			output = "false"
		}

		return c.SendString(output)
	})

	app.Get("/start", func(c *fiber.Ctx) error {
		rot := c.Get("rotation")
		roti, _ := strconv.Atoi(rot)
		uid := uuid.NewString()
		target, err := airport.GetTarget(uid)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "client has no flight availiable")
		}
		users[uid] = User{target: target, lastPoint: airport.Point{}, rotation: roti}
		return c.SendString(uid)
	})

	app.Post("/gps", func(c *fiber.Ctx) error {
		clientId := c.Get("clientId") // clientId obtained when start
		newPos := c.Get("location")   // position of the client
		usr, ok := users[clientId]
		if !ok {
			return fiber.NewError(fiber.StatusUnauthorized, "clientId does not exist")
		}
		point, err := airport.Gps2D(newPos)
		fmt.Printf("In Post: %d %d\n", usr.target.X, usr.target.Y)
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
			c.Set("X-Arrived", "true")
		} else {
			c.Set("X-Arrived", "false")
		}
		path, err := ai.GetCommandVoice(nstep)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Internal server error")
		}
		return c.SendFile(path)
	})

	// Endpoint to upload and receive a WAV file
	app.Post("/upload", func(c *fiber.Ctx) error {
		// Receive the file
		file, err := c.FormFile("audio")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Upload failed")
		}

		/*
			clientId := c.Get("clientId") // clientId obtained when start
			_, ok := users[clientId]
			if !ok {
				return fiber.NewError(fiber.StatusUnauthorized, "clientId does not exist")
			}
		*/

		path := "./tmp/" + file.Filename
		// Save the file to the server
		err = c.SaveFile(file, path)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Could not save file")
		}

		request, err := ai.Speech2Text(path)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("S2T Failure")
		}

		response, err := ai.GetResponse(request)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("GPT Failure")
		}

		output, err := ai.GetAudio(response)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.SendFile(output)
	})

	ai.GetResponse("I am arriving late to the flight")

	app.Listen(":8080")
}
