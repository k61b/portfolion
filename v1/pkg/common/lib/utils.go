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
	secret := []byte(GoDotEnvVariable("JWT_SECRET"))
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		Issuer:    "portfolion",
		Subject:   username,
	})
	token, err := claims.SignedString(secret)
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseJWT(token string) (string, error) {
	secret := []byte(GoDotEnvVariable("JWT_SECRET"))
	claims := jwt.StandardClaims{}
	parsedToken, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return "", err
	}
	if !parsedToken.Valid {
		return "", jwt.ErrSignatureInvalid
	}
	return claims.Subject, nil
}
