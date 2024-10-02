package purchase

import (
	"github.com/gofiber/fiber/v2"
)

func Register(app fiber.Router) {
	route := app.Group("/purchase")
	{

		route.Get("/", ListHandler)
		route.Get("/detail/:id", DetailHandler)
		route.Post("/", InsertHandler)
		route.Put("/:id", UpdateHandler)

		route.Get("/product", ListProductHandler)
	}
}
