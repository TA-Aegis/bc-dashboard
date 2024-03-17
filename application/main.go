package main

import "github.com/gofiber/fiber/v3"

func main() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/ping", func(c fiber.Ctx) error {
		return c.SendString("pong!")
	})

	app.Listen(":8080")
}