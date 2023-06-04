package models

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Store interface {
	GetUsers() (*mongo.Cursor, error)
}
