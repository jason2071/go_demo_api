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
	app.Get("/api/products", r.productController.GetAllProduct)
}
