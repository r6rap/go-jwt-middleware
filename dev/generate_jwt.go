package main

import (
	"fmt"
	"time"

	"go-jwtx/config"
	"github.com/golang-jwt/jwt/v5"
)

func main() {
	secret := config.JWTKey

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    "12345",
		"username": "Rapox",
		"email": "rafif@example.com",
		"role":  "admin",
		"exp":   time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString(secret)

	if err != nil {
		panic(err)
	}

	fmt.Println("Token:", tokenString)
}