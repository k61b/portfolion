package lib

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func GoDotEnvVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Error("Error loading .env file")
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
		log.Error("Error generating JWT:", err)
		return "", err
	}
	return token, nil
}

func ParseJWT(token string) (string, error) {
	secret := []byte(GoDotEnvVariable("JWT_SECRET"))
	claims := jwt.StandardClaims{}
	parsedToken, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		log.Info("Parsing JWT")
		return secret, nil
	})
	if err != nil {
		log.Error("Error parsing JWT:", err)
		return "", err
	}
	if !parsedToken.Valid {
		log.Error("Invalid JWT")
		return "", jwt.ErrSignatureInvalid
	}
	return claims.Subject, nil
}

func Logger() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	file, err := os.OpenFile("portfolion.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		log.SetOutput(file)
	}
}
