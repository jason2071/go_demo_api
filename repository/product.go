package repository

import (
	"database/sql"
	"demo_api/database"
	"demo_api/models"
	"fmt"
)

type ProductRepo interface {
	GetAll() ([]models.Product, error)
	Search(q string) ([]models.Product, error)
}

type productRepo struct {
	c *sql.DB
}

func NewProductRepo(conn database.Connection) ProductRepo {
	return &productRepo{conn.DB()}
}

func (r *productRepo) GetAll() ([]models.Product, error) {

	var result models.Product
	var results []models.Product

	row, err := r.c.Query("SELECT * FROM products")

	if err != nil {
		return nil, err
	}

	for row.Next() {
		row.Scan(&result.Id, &result.Title, &result.Description, &result.Image, &result.Price)
		results = append(results, result)
	}

	return results, nil

}

func (r *productRepo) Search(q string) ([]models.Product, error) {
	var result models.Product
	var results []models.Product

	sql := "SELECT * FROM products"

	if q != "" {
		sql = fmt.Sprintf("%s WHERE title LIKE '%%%s%%'", sql, q)
	}

	row, err := r.c.Query(sql)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		row.Scan(&result.Id, &result.Title, &result.Description, &result.Image, &result.Price)
		results = append(results, result)
	}

	return results, nil
}
