package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username  string             `json:"username"`
	Password  string             `json:"password"`
	Bookmarks []Bookmark         `json:"bookmarks"`
}

type Store interface {
	CreateUser(user *User) error
	GetUserByUsername(username string) (*User, error)
	CreateBookmark(username string, bookmark *Bookmark) error
	GetBookmarks(username string) ([]Bookmark, error)
	UpdateBookmark(username string, symbol string, bookmark *Bookmark) error
	DeleteBookmark(username string, symbol string) error
	CreateOrUpdateSymbol(symbol *Symbol) error
	GetSymbolValue(symbol string) (*Symbol, error)
}

type Bookmark struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"added_price"`
	Pieces float64 `json:"pieces"`
}

type Symbol struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
}
