package service

import "github.com/gofiber/fiber/v2"

func StartServer() {
	app := fiber.New()

	app.Get("/hello-world", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "hello world!",
		})
	})

	app.Listen(":8080")
}
