package routes

import (
	"demo_api/controllers"

	"github.com/gofiber/fiber/v2"
)

type UserRoutes interface {
	UserInstall(app *fiber.App)
}

type userRoutes struct {
	userController controllers.UserController
}

func UserNewRoutes(userController controllers.UserController) UserRoutes {
	return &userRoutes{userController}
}

func (r *userRoutes) UserInstall(app *fiber.App) {
	app.Post("/api/login", r.userController.GetUserByUsername)
}
