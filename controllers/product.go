package controllers

import (
	"demo_api/handler"
	"demo_api/models"
	"demo_api/repository"

	"github.com/gofiber/fiber/v2"
)

type ProductController interface {
	GetAllProduct(ctx *fiber.Ctx) error
	SearchProduct(ctx *fiber.Ctx) error
	InsertProduct(ctx *fiber.Ctx) error
	UpdateProduct(ctx *fiber.Ctx) error
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

func (c *productController) SearchProduct(ctx *fiber.Ctx) error {
	q := ctx.Query("q")
	data, err := c.productRepo.Search(q)

	if err != nil {
		return handler.NotFoundResponse(ctx, err.Error())
	}

	return handler.SuccessWithItems(ctx, data)
}

func (controller *productController) InsertProduct(ctx *fiber.Ctx) error {

	body := models.Product{}
	err := ctx.BodyParser(&body)

	if err != nil {
		return handler.NotFoundResponse(ctx, err.Error())
	}

	if body.Title == "" {
		return handler.BadRequest(ctx, "title is require")
	}

	if body.Category == "" {
		return handler.BadRequest(ctx, "category is require")
	}

	if body.Price == "" {
		body.Price = "0"
	}

	result, err := controller.productRepo.CreateProduct(body)

	if err != nil {
		return handler.InternalServerError(ctx, err)
	}

	return handler.SuccessResponse(ctx, result)
}

func (c *productController) UpdateProduct(ctx *fiber.Ctx) error {

	body := models.Product{}
	err := ctx.BodyParser(&body)

	if err != nil {
		return handler.NotFoundResponse(ctx, err.Error())
	}

	product, err := c.productRepo.GetProductByID(body.Id)

	if err != nil {
		return handler.NotFoundResponse(ctx, err.Error())
	}

	if body.Title == "" {
		body.Title = product.Title
	}

	if body.Description == "" {
		body.Description = product.Description
	}

	if body.Image == "" {
		body.Image = product.Image
	}

	if body.Price == "" {
		body.Price = product.Price
	}

	if body.Category == "" {
		body.Category = product.Category
	}

	result, err := c.productRepo.PutProduct(body)

	if err != nil {
		return handler.NotFoundResponse(ctx, err.Error())
	}
	return handler.SuccessResponse(ctx, result)
}
