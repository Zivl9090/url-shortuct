package main

import (
	"url-shortcut/api/routes"
	"url-shortcut/cache"
	"url-shortcut/db"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func main() {
	db.Connect()
	db.Init()
	cache.Connect()

	app := fiber.New(fiber.Config{
		ServerHeader: "Url shortcut v1.0.0",
	})

	routes.R_login(app.Group("/login"))

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))

	routes.R_url(app.Group("/url"))
	routes.R_user(app.Group("/user"))

	app.Listen(":3005")
}

func handler(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
