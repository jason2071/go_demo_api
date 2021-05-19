package repository

import (
	"database/sql"
	"demo_api/database"
	"demo_api/models"
)

type ProductRepo interface {
	GetAll() ([]models.Product, error)
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
