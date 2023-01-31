package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/get", func(c *fiber.Ctx) error {
		return c.SendString("GET request")
	})

	app.Get("/get/:param", func(c *fiber.Ctx) error {
		return c.SendString("param: " + c.Params("param"))
	})

	app.Post("/", func(c *fiber.Ctx) error {
		return c.SendString("POST request")
	})

	log.Fatal(app.Listen(":3000"))
}
