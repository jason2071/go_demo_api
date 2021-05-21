package handler

import (
	"demo_api/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func SuccessResponse(ctx *fiber.Ctx, text string) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": text,
	})
}

func NotFoundResponse(ctx *fiber.Ctx, text string) error {
	return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
		"message": text,
	})
}

func InternalServerError(ctx *fiber.Ctx, err error) error {
	return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
		"error": err,
	})
}

func SuccessWithItem(ctx *fiber.Ctx, data models.Product) error {
	return ctx.Status(http.StatusFound).JSON(fiber.Map{
		"item": data,
	})
}

func SuccessWithItems(ctx *fiber.Ctx, data []models.Product) error {
	return ctx.Status(http.StatusFound).JSON(fiber.Map{
		"item": data,
	})
}

func SuccessWithLogin(ctx *fiber.Ctx, text string) error {
	return ctx.Status(http.StatusFound).JSON(fiber.Map{
		"data": text,
	})
}

func BadRequest(ctx *fiber.Ctx, text string) error {
	return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
		"message": text,
	})
}
