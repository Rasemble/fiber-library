package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateNewAccesToken() (string, error) {
	// Set secret key from .env
	secret := os.Getenv("JWT_SECRET_KEY")

	// Set expires secret count minutes from .env
	minutesCount, _ := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRES"))

	// Create a new claims
	claims := jwt.MapClaims{}

	// Set public claims
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()

	// Create a new JWT access token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate token
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, nil
}
