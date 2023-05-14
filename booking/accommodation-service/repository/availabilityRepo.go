package repository

import (
	"accommodation-service/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const (
	AVAILABILITY_DB   = "accommodation"
	AVAILABILITY_COLL = "availability"
)

// UserRepo struct encapsulating Mongo api client
type AvailabilityRepo struct {
	availabilities *mongo.Collection
}

func NewAvailabilityRepo(client *mongo.Client) model.AvailabilityStore {
	availabilities := client.Database(AVAILABILITY_DB).Collection(AVAILABILITY_COLL)
	return &AvailabilityRepo{
		availabilities: availabilities,
	}
}

// Disconnect from database
//func (repo *AvailabilityRepo) Disconnect(ctx context.Context) error {
//	err := repo.client.Disconnect(ctx)
//	if err != nil {
//		return err
//	}
//	return nil
//}

func (repo *AvailabilityRepo) Insert(availability *model.Availability) error {
	result, err := repo.availabilities.InsertOne(context.TODO(), availability)
	if err != nil {
		return err
	}
	availability.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

//func (repo *AvailabilityRepo) getCollection() *mongo.Collection {
//	bookingDatabase := repo.client.Database("accommodationDB")
//	availabilitiesCollection := bookingDatabase.Collection("availabilities")
//	return availabilitiesCollection
//}

func (repo *AvailabilityRepo) GetAll() (model.Availabilities, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var availabilities model.Availabilities
	availabilitiesCursor, err := repo.availabilities.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err = availabilitiesCursor.All(ctx, &availabilities); err != nil {
		return nil, err
	}
	return availabilities, nil
}

func (repo *AvailabilityRepo) FindAvailabilitiesByAccommodationId(id string) (model.Availabilities, error) {
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()

	availabilitiesCollection, _ := repo.GetAll()
	var availabilities model.Availabilities
	for _, availability := range availabilitiesCollection {
		if availability.AccommodationId == id {
			availabilities = append(availabilities, availability)
		}
	}

	return availabilities, nil
}

func (repo *AvailabilityRepo) GetById(id string) (*model.Availability, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var availability model.Availability
	objID, _ := primitive.ObjectIDFromHex(id)
	err := repo.availabilities.FindOne(ctx, bson.M{"_id": objID}).Decode(&availability)
	if err != nil {
		return nil, err
	}
	return &availability, nil
}
