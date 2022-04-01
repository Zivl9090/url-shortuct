package routes

import (
	"url-shortcut/server/url"

	"github.com/gofiber/fiber/v2"
)

func R_url(router fiber.Router) {
	router.Get("/:shortUrl", url.GetUrl)

	router.Post("/:originalUrl", url.CreateUrl)
}
