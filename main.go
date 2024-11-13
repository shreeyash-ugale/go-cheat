package main

import (
	"github.com/atotto/clipboard"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New(".", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/send", htmlForm)

	app.Post("/form", formHandler)

	app.Get("/clip", clipHandler)

	app.Listen(":3000")
}

func htmlForm(c *fiber.Ctx) error {
	return c.SendFile("form.html")
}

func formHandler(c *fiber.Ctx) error {
	data := c.FormValue("data")
	err := clipboard.WriteAll(data)
	if err != nil {
		return c.Status(500).SendString("Failed to write to clipboard")

	}
	return c.SendString("Received data: " + data)
}

func clipHandler(c *fiber.Ctx) error {
	clipboardContent, err := clipboard.ReadAll()
	if err != nil {
		return c.Status(500).SendString("Failed to get clipboard content")
	}
	return c.Render("disp", fiber.Map{
		"clipboardContent": clipboardContent,
	})
}
