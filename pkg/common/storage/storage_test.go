package storage_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kayraberktuncer/portfolion/pkg/common/models"
	"github.com/kayraberktuncer/portfolion/pkg/common/storage"
)

func init() {
	os.Chdir("../../..")
}

func TestNewStorage(t *testing.T) {
	os.Create(".env.test")
	os.Setenv("MONGO_URI", "mongodb://localhost:27017")
	os.Setenv("MONGO_DB", "portfolion")
	os.Setenv("MONGO_COLLECTION", "users")

	_, err := storage.NewStorage()

	assert.NoError(t, err)

	os.Remove(".env.test")
}

func TestCreateUser(t *testing.T) {
	os.Create(".env.test")
	os.Setenv("MONGO_URI", "mongodb://localhost:27017")
	os.Setenv("MONGO_DB", "portfolion")
	os.Setenv("MONGO_COLLECTION", "users")

	s, err := storage.NewStorage()

	assert.NoError(t, err)

	user := &models.User{
		Username: "test_user",
		Password: "test_password",
	}

	err = s.CreateUser(user)

	assert.NoError(t, err)

	os.Remove(".env.test")
}

func TestGetUserByUsername(t *testing.T) {
	os.Create(".env.test")
	os.Setenv("MONGO_URI", "mongodb://localhost:27017")
	os.Setenv("MONGO_DB", "portfolion")
	os.Setenv("MONGO_COLLECTION", "users")

	s, err := storage.NewStorage()

	assert.NoError(t, err)

	user := &models.User{
		Username: "test_user",
		Password: "test_password",
	}

	err = s.CreateUser(user)

	assert.NoError(t, err)

	foundUser, err := s.GetUserByUsername("test_user")

	assert.NoError(t, err)
	assert.Equal(t, user.Username, foundUser.Username)
	assert.Equal(t, user.Password, foundUser.Password)

	os.Remove(".env.test")
}
