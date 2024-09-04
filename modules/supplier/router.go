package supplier

import (
	"github.com/gofiber/fiber/v2"
)

func Register(app fiber.Router) {
	route := app.Group("/branch")
	{
		route.Get("/", IndexHandler)
		route.Get("/:id", DetailHandler)
		route.Post("/", InsertHandler)
		route.Put("/:id", UpdateHandler)
	}
}
