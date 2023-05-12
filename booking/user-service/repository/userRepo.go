package repository

import (
	"context"
	"fmt"
	"github.com/sarapetrovic2350/XWS-project/booking/user-service/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

// UserRepo struct encapsulating Mongo api client
type UserRepo struct {
	client *mongo.Client
	logger *log.Logger
}

func NewUserRepo(client *mongo.Client, logger *log.Logger) *UserRepo {
	return &UserRepo{client, logger}
}

// Disconnect from database
func (repo *UserRepo) Disconnect(ctx context.Context) error {
	err := repo.client.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepo) Insert(user *model.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	usersCollection := repo.getCollection()

	result, err := usersCollection.InsertOne(ctx, &user)
	if err != nil {
		repo.logger.Println(err)
		return err
	}
	repo.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}

// Ping check database connection
func (repo *UserRepo) Ping() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check connection -> if no error, connection is established
	err := repo.client.Ping(ctx, readpref.Primary())
	if err != nil {
		repo.logger.Println(err)
	}

	// Print available databases
	databases, err := repo.client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		repo.logger.Println(err)
	}
	fmt.Println(databases)
}

func (repo *UserRepo) getCollection() *mongo.Collection {
	bookingDatabase := repo.client.Database("usersDB")
	usersCollection := bookingDatabase.Collection("users")
	return usersCollection
}

func (repo *UserRepo) GetAll() (model.Users, error) {
	// Initialise context (after 5 seconds timeout, abort operation)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	usersCollection := repo.getCollection()

	var users model.Users
	usersCursor, err := usersCollection.Find(ctx, bson.M{})
	if err != nil {
		repo.logger.Println(err)
		return nil, err
	}
	if err = usersCursor.All(ctx, &users); err != nil {
		repo.logger.Println(err)
		return nil, err
	}
	return users, nil
}
func (repo *UserRepo) FindUserByEmail(email string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	usersCollection := repo.getCollection()

	var user model.User
	filter := bson.M{"email": bson.M{"$eq": email}}
	err := usersCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		repo.logger.Println(err)
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepo) GetById(id string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	usersCollection := repo.getCollection()

	var user model.User
	objID, _ := primitive.ObjectIDFromHex(id)
	err := usersCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		repo.logger.Println(err)
		return nil, err
	}
	return &user, nil
}
