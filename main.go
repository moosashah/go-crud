package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/moosashah/go-crud/initializers"
	"github.com/moosashah/go-crud/tournament"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	app := fiber.New()

	tournament.Routes(app, "/tournament")

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"mesage": "pong"})
	})

	app.Listen(":4200")
}
