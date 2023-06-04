package storage

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/kayraberktuncer/portfolion/pkg/common/lib"
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

func (s *Storage) GetUsers() (*mongo.Cursor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := s.collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return cursor, nil
}
