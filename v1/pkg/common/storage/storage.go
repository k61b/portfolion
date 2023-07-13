package storage

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/kayraberktuncer/portfolion/pkg/common/lib"
	"github.com/kayraberktuncer/portfolion/pkg/common/models"
	log "github.com/sirupsen/logrus"
)

type Storage struct {
	client            *mongo.Client
	usersCollection   *mongo.Collection
	symbolsCollection *mongo.Collection
}

func NewStorage() (*Storage, error) {
	clientOptions := options.Client().ApplyURI(lib.GoDotEnvVariable("MONGO_URI"))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Error("Error connecting to MongoDB:", err)
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Error("Error pinging MongoDB:", err)
		return nil, err
	}

	db := client.Database(lib.GoDotEnvVariable("MONGO_DB"))
	usersCollection := db.Collection(lib.GoDotEnvVariable("USERS_COLLECTION"))
	symbolsCollection := db.Collection(lib.GoDotEnvVariable("SYMBOLS_COLLECTION"))

	return &Storage{
		client:            client,
		usersCollection:   usersCollection,
		symbolsCollection: symbolsCollection,
	}, nil
}

func (s *Storage) CreateUser(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if _, err := s.usersCollection.InsertOne(ctx, user); err != nil {
		log.Error("Error inserting user:", err)
		return err
	}

	return nil
}

func (s *Storage) GetUserByUsername(username string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	if err := s.usersCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		log.Error("Error retrieving user:", err)
		return nil, err
	}

	return &user, nil
}

func (s *Storage) CreateBookmark(username string, bookmark *models.Bookmark) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if _, err := s.usersCollection.UpdateOne(
		ctx,
		bson.M{"username": username},
		bson.M{"$push": bson.M{"bookmarks": bookmark}},
	); err != nil {
		log.Error("Error creating bookmark:", err)
		return err
	}

	return nil
}

func (s *Storage) GetBookmarks(username string) ([]models.Bookmark, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	if err := s.usersCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		log.Error("Error retrieving user:", err)
		return nil, err
	}

	return user.Bookmarks, nil
}

func (s *Storage) UpdateBookmark(username string, symbol string, bookmark *models.Bookmark) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if _, err := s.usersCollection.UpdateOne(
		ctx,
		bson.M{"username": username, "bookmarks.symbol": symbol},
		bson.M{"$set": bson.M{"bookmarks.$": bookmark}},
	); err != nil {
		log.Error("Error updating bookmark:", err)
		return err
	}

	return nil
}

func (s *Storage) DeleteBookmark(username string, symbol string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if _, err := s.usersCollection.UpdateOne(
		ctx,
		bson.M{"username": username},
		bson.M{"$pull": bson.M{"bookmarks": bson.M{"symbol": symbol}}},
	); err != nil {
		log.Error("Error deleting bookmark:", err)
		return err
	}

	return nil
}

func (s *Storage) CreateOrUpdateSymbol(symbol *models.Symbol) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"symbol": symbol.Symbol}
	update := bson.M{"$set": symbol}
	opts := options.Update().SetUpsert(true)

	_, err := s.symbolsCollection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		log.Error("Error creating or updating symbol:", err)
		return err
	}

	return nil
}

func (s *Storage) GetSymbolValue(symbol string) (*models.Symbol, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var symbolData models.Symbol
	err := s.symbolsCollection.FindOne(ctx, bson.M{"symbol": symbol}).Decode(&symbolData)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, mongo.ErrNoDocuments
		}

		log.Error("Error retrieving symbol:", err)
		return nil, err
	}

	return &symbolData, nil
}

func (s *Storage) GetSymbols() ([]models.Symbol, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var symbols []models.Symbol
	cursor, err := s.symbolsCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Error("Error retrieving symbols:", err)
		return nil, err
	}

	if err = cursor.All(ctx, &symbols); err != nil {
		log.Error("Error retrieving symbols:", err)
		return nil, err
	}

	return symbols, nil
}
