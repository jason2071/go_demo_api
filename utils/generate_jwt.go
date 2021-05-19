package utils

import (
	"demo_api/models"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(user models.User) (string, error) {
	key := []byte(os.Getenv("JWT_SECRET_KEY"))
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["data"] = user
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(key)

	if err != nil {
		return "", err
	}

	return tokenString, err
}
