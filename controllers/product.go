package controllers

import (
	"demo_api/handler"
	"demo_api/repository"

	"github.com/gofiber/fiber/v2"
)

type ProductController interface {
	GetAllProduct(ctx *fiber.Ctx) error
}

type productController struct {
	productRepo repository.ProductRepo
}

func NewProductController(productRepo repository.ProductRepo) ProductController {
	return &productController{productRepo}
}

func (c *productController) GetAllProduct(ctx *fiber.Ctx) error {

	result, err := c.productRepo.GetAll()

	if err != nil {
		return handler.InternalServerError(nil, err)
	}

	return handler.SuccessWithItems(ctx, result)
}
