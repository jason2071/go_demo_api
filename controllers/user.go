package controllers

import (
	"demo_api/handler"
	"demo_api/models"
	"demo_api/repository"
	"demo_api/utils"

	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	GetUserByUsername(c *fiber.Ctx) error
}

type userController struct {
	userRepo repository.UserRepo
}

func NewUserController(userRepo repository.UserRepo) UserController {
	return &userController{userRepo}
}

func (c *userController) GetUserByUsername(ctx *fiber.Ctx) error {

	body := models.Login{}
	errBody := ctx.BodyParser(&body)

	if errBody != nil {
		return handler.NotFoundResponse(ctx, errBody.Error())
	}

	var data models.User
	var err error

	data, err = c.userRepo.GetByUsername(body.Username)

	if err != nil {
		return handler.NotFoundResponse(ctx, err.Error())
	}

	if data.Password != body.Password {
		return handler.NotFoundResponse(ctx, "password is incorrect")
	}

	tokenString, errorToken := utils.GenerateJWT(data)

	if errorToken != nil {
		return handler.NotFoundResponse(ctx, errorToken.Error())
	}

	return handler.SuccessWithLogin(ctx, tokenString)
}
