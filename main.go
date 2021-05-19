package main

import (
	"demo_api/controllers"
	"demo_api/database"
	"demo_api/repository"
	"demo_api/routes"
	"fmt"

	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Panicln(err)
	}
}

func main() {
	conn := database.InitDB()
	defer conn.Close()

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	productRepo := repository.NewProductRepo(conn)
	productController := controllers.NewProductController(productRepo)
	productRoutes := routes.NewRoutes(productController)
	productRoutes.Install(app)

	userRepo := repository.NewUserRepo(conn)
	userController := controllers.NewUserController(userRepo)
	userRoutes := routes.UserNewRoutes(userController)
	userRoutes.UserInstall(app)

	fmt.Print(productRepo)
	app.Listen(":5000")
}
