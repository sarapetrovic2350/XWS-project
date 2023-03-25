package repository

import (
	"Rest/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

// NoSQL: UsersRepo struct encapsulating Mongo api client
type UsersRepo struct {
	client *mongo.Client
	logger *log.Logger
}

// NoSQL: Constructor which reads db configuration from environment
func NewUsersRepo(client *mongo.Client, logger *log.Logger) (*UsersRepo, error) {
	return &UsersRepo{
		client: client,
		logger: logger,
	}, nil
}

// Disconnect from database
func (ur *UsersRepo) Disconnect(ctx context.Context) error {
	err := ur.client.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (ur *UsersRepo) Insert(user *model.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	usersCollection := ur.getCollection()

	result, err := usersCollection.InsertOne(ctx, &user)
	if err != nil {
		ur.logger.Println(err)
		return err
	}
	ur.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}

func (ur *UsersRepo) getCollection() *mongo.Collection {
	airbnbDatabase := ur.client.Database("airbnbDB")
	usersCollection := airbnbDatabase.Collection("users")
	return usersCollection
}

func (ur *UsersRepo) GetAll() (model.Users, error) {
	// Initialise context (after 5 seconds timeout, abort operation)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	usersCollection := ur.getCollection()

	var users model.Users
	usersCursor, err := usersCollection.Find(ctx, bson.M{})
	if err != nil {
		ur.logger.Println(err)
		return nil, err
	}
	if err = usersCursor.All(ctx, &users); err != nil {
		ur.logger.Println(err)
		return nil, err
	}
	return users, nil
}
