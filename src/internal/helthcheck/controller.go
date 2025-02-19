package helthcheck

import "github.com/gofiber/fiber/v2"

func get(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func Register(app *fiber.App) {
	app.Get("/helthcheck", get)
}
