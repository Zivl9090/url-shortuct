package routes

import (
	"url-shortcut/login"

	"github.com/gofiber/fiber/v2"
)

func R_login(router fiber.Router) {

	router.Post("/", login.Login)

}
