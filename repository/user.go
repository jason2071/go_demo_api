package repository

import (
	"database/sql"
	"demo_api/database"
	"demo_api/models"
	"errors"
)

type UserRepo interface {
	GetByUsername(username string) (models.User, error)
}

type userRepo struct {
	c *sql.DB
}

func NewUserRepo(conn database.Connection) UserRepo {
	return &userRepo{conn.DB()}
}

func (r *userRepo) GetByUsername(username string) (models.User, error) {

	var result models.User

	r.c.QueryRow("SELECT * FROM account WHERE username = ?", username).Scan(&result.Id, &result.Username, &result.Password, &result.Name)

	if result.Username == "" {
		return result, errors.New("user not found")
	}

	return result, nil
}
