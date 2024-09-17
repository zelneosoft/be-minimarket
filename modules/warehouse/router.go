package warehouse

import (
	"github.com/gofiber/fiber/v2"
)

func Register(app fiber.Router) {
	route := app.Group("/warehouse")
	{

		route.Get("/", ListHandler)
		// route.Get("/:id", DetailHandler)
		route.Post("/", InsertHandler)
		route.Put("/:id", UpdateHandler)
	}
}
