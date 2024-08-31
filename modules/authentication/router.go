package authentication

import (
	"github.com/gofiber/fiber/v2"
)

func Register(app fiber.Router) {
	route := app.Group("/auth")
	{
		route.Post("/login", LoginHandler)
		route.Post("/new-password", NewPasswordHandler)
	}
}
