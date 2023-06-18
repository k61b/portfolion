package storage

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/kayraberktuncer/portfolion/pkg/common/lib"
	"github.com/kayraberktuncer/portfolion/pkg/common/models"
)

type Storage struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewStorage() (*Storage, error) {
	clientOptions := options.Client().ApplyURI(lib.GoDotEnvVariable("MONGO_URI"))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	db := client.Database(lib.GoDotEnvVariable("MONGO_DB"))
	collection := db.Collection(lib.GoDotEnvVariable("MONGO_COLLECTION"))

	return &Storage{
		client:     client,
		collection: collection,
	}, nil
}

func (s *Storage) CreateUser(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if _, err := s.collection.InsertOne(ctx, user); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (s *Storage) GetUserByUsername(username string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	if err := s.collection.FindOne(ctx, bson.M{"username": username}).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		log.Fatal(err)
		return nil, err
	}

	return &user, nil
}

func (s *Storage) CreateBookmark(username string, bookmark *models.Bookmark) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if _, err := s.collection.UpdateOne(ctx, bson.M{"username": username}, bson.M{"$push": bson.M{"bookmarks": bookmark}}); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (s *Storage) GetBookmarks(username string) ([]models.Bookmark, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	if err := s.collection.FindOne(ctx, bson.M{"username": username}).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		log.Fatal(err)
		return nil, err
	}

	return user.Bookmarks, nil
}

func (s *Storage) UpdateBookmark(username string, symbol string, bookmark *models.Bookmark) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if _, err := s.collection.UpdateOne(ctx, bson.M{"username": username, "bookmarks.symbol": symbol}, bson.M{"$set": bson.M{"bookmarks.$": bookmark}}); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (s *Storage) DeleteBookmark(username string, symbol string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if _, err := s.collection.UpdateOne(ctx, bson.M{"username": username}, bson.M{"$pull": bson.M{"bookmarks": bson.M{"symbol": symbol}}}); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
