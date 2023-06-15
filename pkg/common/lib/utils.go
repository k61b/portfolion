package lib

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func GoDotEnvVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func GenerateJWT(username string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		Issuer:    "portfolion",
		Subject:   username,
	})

	return claims.SignedString([]byte(GoDotEnvVariable("JWT_SECRET")))
}

func ParseJWT(token string) (string, error) {
	claims := jwt.StandardClaims{}
	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(GoDotEnvVariable("JWT_SECRET")), nil
	})

	if err != nil {
		return "", err
	}

	return claims.Subject, nil
}
