package repository

import (
	"database/sql"
	"demo_api/database"
	"demo_api/models"
	"errors"
	"fmt"
)

type ProductRepo interface {
	GetAll() ([]models.Product, error)
	GetProductByID(id string) (models.Product, error)
	Search(q string) ([]models.Product, error)
	CreateProduct(item models.Product) (string, error)
	PutProduct(item models.Product) (string, error)
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

	row, err := r.c.Query("SELECT products.id,products.title,products.description,products.image,products.price,category.title as category FROM products JOIN category ON category.id = products.category")

	if err != nil {
		return nil, err
	}

	for row.Next() {
		row.Scan(&result.Id, &result.Title, &result.Description, &result.Image, &result.Price, &result.Category)
		results = append(results, result)
	}

	return results, nil
}

func (r *productRepo) GetProductByID(id string) (models.Product, error) {
	var result models.Product
	r.c.QueryRow("SELECT * FROM products WHERE id = ?", id).Scan(&result.Id, &result.Title, &result.Description, &result.Image, &result.Price, &result.Category)
	if result.Title == "" {
		return result, errors.New("product not found")
	}
	return result, nil
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

func (repo *productRepo) CreateProduct(item models.Product) (string, error) {
	_, err := repo.c.Query("INSERT INTO products (title, description, image, price, category) VALUE (?,?,?,?,?)", item.Title, item.Description, item.Image, item.Price, item.Category)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("product %s create success", item.Title), nil
}

func (r *productRepo) PutProduct(item models.Product) (string, error) {
	_, err := r.c.Query("UPDATE products SET title=?, description=?, image=?, price=?, category=? WHERE id=?", item.Title, item.Description, item.Image, item.Price, item.Category, item.Id)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("product %s update success", item.Title), nil
}
