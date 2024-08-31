package category

import (
	"github.com/gofiber/fiber/v2"
)

func Register(app fiber.Router) {
	app.Get("/", ListHandler)
	app.Post("/", InsertHandler)
	app.Put("/:id", UpdateHandler)
	app.Delete("/:id", DeleteHandler)
}
