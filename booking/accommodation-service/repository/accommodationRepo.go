package repository

import (
	"Rest/dto"
	"Rest/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

// UserRepo struct encapsulating Mongo api client
type AccommodationRepo struct {
	client *mongo.Client
	logger *log.Logger
}

func NewAccommodationRepo(client *mongo.Client, logger *log.Logger) *AccommodationRepo {
	return &AccommodationRepo{client, logger}
}

// Disconnect from database
func (repo *AccommodationRepo) Disconnect(ctx context.Context) error {
	err := repo.client.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (repo *AccommodationRepo) Insert(user *model.Accommodation) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	accommodationsCollection := repo.getCollection()

	result, err := accommodationsCollection.InsertOne(ctx, &user)
	if err != nil {
		repo.logger.Println(err)
		return err
	}
	repo.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}

func (repo *AccommodationRepo) getCollection() *mongo.Collection {
	bookingDatabase := repo.client.Database("accommodationDB")
	accommodationsCollection := bookingDatabase.Collection("accommodations")
	return accommodationsCollection
}

func (repo *AccommodationRepo) GetAll() (model.Accommodations, error) {
	// Initialise context (after 5 seconds timeout, abort operation)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	accommodationsCollection := repo.getCollection()

	var accommodations model.Accommodations
	accommodationsCursor, err := accommodationsCollection.Find(ctx, bson.M{})
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
func (repo *AccommodationRepo) FindAccommodationByEmail(email string) (*model.Accommodation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	accommodationsCollection := repo.getCollection()

	var accommodation model.Accommodation
	filter := bson.M{"email": bson.M{"$eq": email}}
	err := accommodationsCollection.FindOne(ctx, filter).Decode(&accommodation)
	if err != nil {
		repo.logger.Println(err)
		return nil, err
	}
	return &accommodation, nil
}

func (repo *AccommodationRepo) GetById(id string) (*model.Accommodation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	accommodationsCollection := repo.getCollection()

	var accommodation model.Accommodation
	objID, _ := primitive.ObjectIDFromHex(id)
	err := accommodationsCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&accommodation)
	if err != nil {
		repo.logger.Println(err)
		return nil, err
	}
	return &accommodation, nil
}

func (repo *AccommodationRepo) SearchAccommodation(searchCriteria dto.SearchDTO) model.Accommodations {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//results := []model.Flight{}

	//treba za kasnije
	//end := time.Date(searchCriteria.Date.Year(), searchCriteria.Date.Month(), searchCriteria.Date.Day(), 23, 59, 59, 999999999, time.UTC)
	//fmt.Println(end)
	filter := bson.M{"country": searchCriteria.Country, "city": searchCriteria.City,
		//"departure_date_time": bson.M{"$gte": searchCriteria.Date, "$lt": end},
		"min_number_of_guests": bson.M{"$gte": searchCriteria.NumberOfGuests}, "max_number_of_guests": bson.M{"$gte": searchCriteria.NumberOfGuests}}

	var accommodations model.Accommodations
	accommodationsCollection := repo.getCollection()
	cursor, err := accommodationsCollection.Find(ctx, filter)
	if err != nil {
		log.Panic("Could not find document in database", err.Error())
		return nil
	}
	if err = cursor.All(context.TODO(), &accommodations); err != nil {
		log.Panic("Could not find document in database", err.Error())
		return nil
	}

	fmt.Println("SearchResult:")
	fmt.Println(&accommodations)

	return accommodations
}
