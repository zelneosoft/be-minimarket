package products

import (
	"backend/modules/products/brand"
	"backend/modules/products/category"

	"github.com/gofiber/fiber/v2"
)

func Register(app fiber.Router) {
	route := app.Group("/product")
	{
		categories := route.Group("/category")
		category.Register(categories)

		brands := route.Group("/brand")
		brand.Register(brands)

		route.Get("/", IndexHandler)
		route.Get("/:id", DetailHandler)
		route.Post("/", InsertHandler)
		route.Put("/:id", UpdateHandler)
	}
}
