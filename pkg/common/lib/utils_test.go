package lib_test

import (
	"os"
	"testing"

	"github.com/kayraberktuncer/portfolion/pkg/common/lib"
	"github.com/stretchr/testify/assert"
)

func init() {
	os.Chdir("../../..")
}

func TestGoDotEnvVariable(t *testing.T) {
	os.Create(".env.test")
	os.Setenv("TEST_KEY", "test_value")

	result := lib.GoDotEnvVariable("TEST_KEY")

	assert.Equal(t, "test_value", result)

	os.Remove(".env.test")
}

func TestGenerateJWT(t *testing.T) {
	username := "test_user"

	token, err := lib.GenerateJWT(username)

	assert.NotEmpty(t, token)
	assert.NoError(t, err)

	parsedUsername, err := lib.ParseJWT(token)

	assert.Equal(t, username, parsedUsername)
	assert.NoError(t, err)
}

func TestParseJWT(t *testing.T) {
	username := "test_user"

	token, err := lib.GenerateJWT(username)

	assert.NotEmpty(t, token)
	assert.NoError(t, err)

	parsedUsername, err := lib.ParseJWT(token)

	assert.Equal(t, username, parsedUsername)
	assert.NoError(t, err)
}

func TestParseJWTWithInvalidToken(t *testing.T) {
	username := "test_user"

	token, err := lib.GenerateJWT(username)

	assert.NotEmpty(t, token)
	assert.NoError(t, err)

	parsedUsername, err := lib.ParseJWT(token + "invalid")

	assert.Empty(t, parsedUsername)
	assert.Error(t, err)
}
