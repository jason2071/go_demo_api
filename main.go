package main

import (
	"demo_api/controllers"
	"demo_api/database"
	"demo_api/models"
	"demo_api/repository"
	"demo_api/routes"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/joho/godotenv"
)

var conn database.Connection

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Panicln(err)
	}
}

func handler(ctx *fiber.Ctx) error {

	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var result models.User
	var body request
	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})

	}

	conn.DB().QueryRow("SELECT * FROM account WHERE username = ?", body.Email).Scan(&result.Id, &result.Username, &result.Password, &result.Name)

	if result.Username == "" {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "user not found",
		})
	}

	if body.Password != result.Password {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "password is correct",
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["data"] = result
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	s, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"token": s,
		"data":  result,
	})
}

func restricted(c *fiber.Ctx) error {

	tokenString := c.Get("Authorization")
	words := strings.Fields(tokenString)
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(words[1], claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	for key, val := range claims {
		fmt.Printf("Key: %v, value: %v\n", key, val)
	}

	return c.JSON(fiber.Map{
		"success": true,
		"token":   token.Claims,
	})

}

func main() {
	// db
	conn = database.InitDB()
	defer conn.Close()

	// fiber
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	// product
	productRepo := repository.NewProductRepo(conn)
	productController := controllers.NewProductController(productRepo)
	productRoutes := routes.NewRoutes(productController)
	productRoutes.Install(app)

	// Login route
	app.Post("/api/login", handler)

	// Restricted Routes
	app.Use(jwtware.New(jwtware.Config{
		ErrorHandler: func(c *fiber.Ctx, e error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
		SigningKey: []byte(os.Getenv("JWT_SECRET_KEY")),
	}))
	app.Get("/api/profile", restricted)

	app.Listen(":5000")
}
