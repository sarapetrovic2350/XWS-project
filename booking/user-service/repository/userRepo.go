package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"user-service/model"
)

const (
	DATABASE   = "user"
	COLLECTION = "user"
)

// UserRepo struct encapsulating Mongo api client
type UserRepo struct {
	users *mongo.Collection
}

func NewUserRepo(client *mongo.Client) model.UserStore {
	users := client.Database(DATABASE).Collection(COLLECTION)
	return &UserRepo{
		users: users,
	}
}
func (repo *UserRepo) Get(id primitive.ObjectID) (*model.User, error) {
	filter := bson.M{"_id": id}
	return repo.filterOne(filter)
}

func (repo *UserRepo) GetAll() (model.Users, error) {
	// Initialise context (after 5 seconds timeout, abort operation)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var users model.Users
	reservationsCursor, err := repo.users.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err = reservationsCursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}
func (repo *UserRepo) FindUserByEmail(email string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user model.User
	filter := bson.M{"email": bson.M{"$eq": email}}
	err := repo.users.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepo) Insert(user *model.User) error {
	result, err := repo.users.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	user.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (repo *UserRepo) GetById(id string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user model.User
	objID, _ := primitive.ObjectIDFromHex(id)
	err := repo.users.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepo) DeleteAll() {
	repo.users.DeleteMany(context.TODO(), bson.D{{}})
}

func (repo *UserRepo) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	fmt.Print(id)
	objID, _ := primitive.ObjectIDFromHex(id)
	fmt.Print(objID)
	filter := bson.D{{Key: "_id", Value: objID}}
	_, err := repo.users.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
func (repo *UserRepo) filterOne(filter interface{}) (User *model.User, err error) {
	result := repo.users.FindOne(context.TODO(), filter)
	err = result.Decode(&User)
	return
}
