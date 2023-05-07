package repository

import (
	"Rest/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

// UserRepo struct encapsulating Mongo api client
type AvailabilityRepo struct {
	client *mongo.Client
	logger *log.Logger
}

func NewAvailabilityRepo(client *mongo.Client, logger *log.Logger) *AvailabilityRepo {
	return &AvailabilityRepo{client, logger}
}

// Disconnect from database
func (repo *AvailabilityRepo) Disconnect(ctx context.Context) error {
	err := repo.client.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (repo *AvailabilityRepo) Insert(availability *model.Availability) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	availabilityCollection := repo.getCollection()

	result, err := availabilityCollection.InsertOne(ctx, &availability)
	if err != nil {
		repo.logger.Println(err)
		return err
	}
	repo.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}

func (repo *AvailabilityRepo) getCollection() *mongo.Collection {
	bookingDatabase := repo.client.Database("accommodationDB")
	availabilitiesCollection := bookingDatabase.Collection("availabilities")
	return availabilitiesCollection
}

func (repo *AvailabilityRepo) GetAll() (model.Availabilities, error) {
	// Initialise context (after 5 seconds timeout, abort operation)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	availabilitiesCollection := repo.getCollection()

	var accommodations model.Availabilities
	accommodationsCursor, err := availabilitiesCollection.Find(ctx, bson.M{})
	if err != nil {
		repo.logger.Println(err)
		return nil, err
	}
	if err = accommodationsCursor.All(ctx, &accommodations); err != nil {
		repo.logger.Println(err)
		return nil, err
	}
	return accommodations, nil
}

func (repo *AvailabilityRepo) GetById(id string) (*model.Availability, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	availabilitiesCollection := repo.getCollection()

	var availability model.Availability
	objID, _ := primitive.ObjectIDFromHex(id)
	err := availabilitiesCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&availability)
	if err != nil {
		repo.logger.Println(err)
		return nil, err
	}
	return &availability, nil
}
