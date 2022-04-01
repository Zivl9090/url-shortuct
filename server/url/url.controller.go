package url

import (
	"github.com/gofiber/fiber/v2"
)

func GetUrl(c *fiber.Ctx) error {
	url := Logic_GetUrl(c.Params("shortUrl"))
	return c.SendString(url.ShortcutUrl)
}

func CreateUrl(c *fiber.Ctx) error {
	Logic_CreateUrl(c.Params("originalUrl"))
	return c.SendString("respond with a resource")
}
