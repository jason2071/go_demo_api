package product

import (
	"database/sql"
	"demo_api/models"

	"github.com/gofiber/fiber/v2"
)

func GetProduct(c *fiber.Ctx, db *sql.DB) error {
	var product models.Product
	var products []models.Product

	row, err := db.Query("SELECT * FROM products")

	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"msg":     err.Error(),
		})
	}

	for row.Next() {
		row.Scan(&product.Id, &product.Title, &product.Image, &product.Description, &product.Price)
		products = append(products, product)
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    products,
	})
}
