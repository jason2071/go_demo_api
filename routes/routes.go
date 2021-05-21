package routes

import (
	"demo_api/controllers"

	"github.com/gofiber/fiber/v2"
)

type Routes interface {
	Install(app *fiber.App)
}

type productRoutes struct {
	productController controllers.ProductController
}

func NewRoutes(productController controllers.ProductController) Routes {
	return &productRoutes{productController}
}

func (r *productRoutes) Install(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1", func(c *fiber.Ctx) error {
		c.Set("Version", "v1")
		return c.Next()
	})
	v1.Get("/products", r.productController.GetAllProduct)
	v1.Get("/search", r.productController.SearchProduct)
	v1.Post("/product/create", r.productController.InsertProduct)
	v1.Put("/product/edit", r.productController.UpdateProduct)
}
