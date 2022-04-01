package routes

import "github.com/gofiber/fiber/v2"

func R_user(router fiber.Router) {
	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("respond with a resource")
	})
}
