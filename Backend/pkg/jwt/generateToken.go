package jwt

import (
	"log"
	"time"

	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/structs"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(env structs.Env, username string) (string, error) {
	secretKey := []byte(env.JwtSecretKey)
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Println("Failed to signed with secret key.")
		return "", nil
	}
	return tokenString, nil
}
