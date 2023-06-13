package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"rating-service/model"
	"time"
)

const (
	DATABASE   = "rating"
	COLLECTION = "rating"
)

// RatingRepo struct encapsulating Mongo api client
type RatingRepo struct {
	ratings *mongo.Collection
}

func NewRatingRepo(client *mongo.Client) model.RatingStore {
	ratings := client.Database(DATABASE).Collection(COLLECTION)
	return &RatingRepo{
		ratings: ratings,
	}
}
func (repo *RatingRepo) GetRatingHost(id primitive.ObjectID) (*model.RatingHost, error) {
	filter := bson.M{"_id": id}
	return repo.filterOne(filter)
}

func (repo *RatingRepo) GetAllRatingsHost() (model.RatingsHost, error) {
	// Initialise context (after 5 seconds timeout, abort operation)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var ratingsHost model.RatingsHost
	ratingsCursor, err := repo.ratings.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err = ratingsCursor.All(ctx, &ratingsHost); err != nil {
		return nil, err
	}
	return ratingsHost, nil
}

func (repo *RatingRepo) InsertRatingHost(rh *model.RatingHost) error {
	result, err := repo.ratings.InsertOne(context.TODO(), rh)
	if err != nil {
		return err
	}
	rh.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (repo *RatingRepo) DeleteAll() {
	repo.ratings.DeleteMany(context.TODO(), bson.D{{}})
}

func (repo *RatingRepo) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	fmt.Print(id)
	objID, _ := primitive.ObjectIDFromHex(id)
	fmt.Print(objID)
	filter := bson.D{{Key: "_id", Value: objID}}
	_, err := repo.ratings.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (repo *RatingRepo) filterOne(filter interface{}) (RatingHost *model.RatingHost, err error) {
	result := repo.ratings.FindOne(context.TODO(), filter)
	err = result.Decode(&RatingHost)
	return
}
