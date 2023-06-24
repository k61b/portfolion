package lib_test

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"github.com/kayraberktuncer/portfolion/pkg/common/lib"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func setup() {
	rootDir, err := filepath.Abs("../../..")
	if err != nil {
		panic(err)
	}

	err = os.Chdir(rootDir)
	if err != nil {
		panic(err)
	}

	err = godotenv.Load(filepath.Join(rootDir, ".env.example"))
	if err != nil {
		panic(err)
	}
}

func TestGenerateJWT(t *testing.T) {
	username := "testuser"
	token, err := lib.GenerateJWT(username)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	parsedUsername, err := lib.ParseJWT(token)
	assert.NoError(t, err)
	assert.Equal(t, username, parsedUsername)
}

func TestParseJWT(t *testing.T) {
	username := "testuser"
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		Issuer:    "portfolion",
		Subject:   username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("testsecret"))
	assert.NoError(t, err)

	err = os.Setenv("JWT_SECRET", "testsecret")
	assert.NoError(t, err)

	parsedUsername, err := lib.ParseJWT(tokenString)
	assert.NoError(t, err)
	assert.Equal(t, username, parsedUsername)
}
