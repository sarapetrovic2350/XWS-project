package repository

import (
	"context"
	"fmt"
	"rating-service/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE                         = "rating"
	RATINGS_HOST_COLLECTION          = "ratingsHost"
	RATINGS_ACCOMMODATION_COLLECTION = "ratingsAccommodation"
)

// RatingRepo struct encapsulating Mongo api client
type RatingRepo struct {
	ratingsHost          *mongo.Collection
	ratingsAccommodation *mongo.Collection
}

func NewRatingRepo(client *mongo.Client) model.RatingStore {
	ratingsHost := client.Database(DATABASE).Collection(RATINGS_HOST_COLLECTION)
	ratingsAccommodation := client.Database(DATABASE).Collection(RATINGS_ACCOMMODATION_COLLECTION)
	return &RatingRepo{
		ratingsHost:          ratingsHost,
		ratingsAccommodation: ratingsAccommodation,
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
	ratingsCursor, err := repo.ratingsHost.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err = ratingsCursor.All(ctx, &ratingsHost); err != nil {
		return nil, err
	}
	return ratingsHost, nil
}
func (repo *RatingRepo) GetAllRatingsAccommodation() (model.RatingsAccommodation, error) {
	// Initialise context (after 5 seconds timeout, abort operation)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var ratingsAccommodation model.RatingsAccommodation
	ratingsCursor, err := repo.ratingsAccommodation.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err = ratingsCursor.All(ctx, &ratingsAccommodation); err != nil {
		return nil, err
	}
	return ratingsAccommodation, nil
}

func (repo *RatingRepo) GetRatingHostById(id string) (*model.RatingHost, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var ratingHost model.RatingHost
	objID, _ := primitive.ObjectIDFromHex(id)
	err := repo.ratingsHost.FindOne(ctx, bson.M{"_id": objID}).Decode(&ratingHost)
	if err != nil {
		return nil, err
	}
	return &ratingHost, nil
}

func (repo *RatingRepo) GetRatingAccommodationById(id string) (*model.RatingAccommodation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var ratingAccommodation model.RatingAccommodation
	objID, _ := primitive.ObjectIDFromHex(id)
	err := repo.ratingsAccommodation.FindOne(ctx, bson.M{"_id": objID}).Decode(&ratingAccommodation)
	if err != nil {
		return nil, err
	}
	return &ratingAccommodation, nil
}

func (repo *RatingRepo) InsertRatingHost(rh *model.RatingHost) error {
	result, err := repo.ratingsHost.InsertOne(context.TODO(), rh)
	if err != nil {
		return err
	}
	rh.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}
func (repo *RatingRepo) InsertRatingAccommodation(rh *model.RatingAccommodation) error {
	result, err := repo.ratingsAccommodation.InsertOne(context.TODO(), rh)
	if err != nil {
		return err
	}
	rh.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}
func (repo *RatingRepo) DeleteAll() {
	repo.ratingsHost.DeleteMany(context.TODO(), bson.D{{}})
	repo.ratingsAccommodation.DeleteMany(context.TODO(), bson.D{{}})
}

func (repo *RatingRepo) DeleteRatingForHost(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	fmt.Print(id)
	objID, _ := primitive.ObjectIDFromHex(id)
	fmt.Print(objID)
	filter := bson.D{{Key: "_id", Value: objID}}
	_, err := repo.ratingsHost.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (repo *RatingRepo) DeleteRatingForAccommodation(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	fmt.Print(id)
	objID, _ := primitive.ObjectIDFromHex(id)
	fmt.Print(objID)
	filter := bson.D{{Key: "_id", Value: objID}}
	_, err := repo.ratingsAccommodation.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (repo *RatingRepo) filterOne(filter interface{}) (RatingHost *model.RatingHost, err error) {
	result := repo.ratingsHost.FindOne(context.TODO(), filter)
	err = result.Decode(&RatingHost)
	return
}
