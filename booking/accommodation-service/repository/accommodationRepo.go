package repository

import (
	"accommodation-service/model"
	accommodationGw "common/proto/accommodation-service/pb"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

const (
	ACCOMMODATION_DB   = "accommodation"
	ACCOMMODATION_COLL = "accommodation"
)

// UserRepo struct encapsulating Mongo api client
type AccommodationRepo struct {
	accommodations *mongo.Collection
}

func NewAccommodationRepo(client *mongo.Client) model.AccommodationStore {
	accommodations := client.Database(ACCOMMODATION_DB).Collection(ACCOMMODATION_COLL)
	return &AccommodationRepo{
		accommodations: accommodations,
	}
}

//// Disconnect from database
//func (repo *AccommodationRepo) Disconnect(ctx context.Context) error {
//	err := repo.client.Disconnect(ctx)
//	if err != nil {
//		return err
//	}
//	return nil
//}

func (repo *AccommodationRepo) Insert(accommodation *model.Accommodation) error {
	result, err := repo.accommodations.InsertOne(context.TODO(), accommodation)
	if err != nil {
		return err
	}
	accommodation.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

//func (repo *AccommodationRepo) getCollection() *mongo.Collection {
//	bookingDatabase := repo.client.Database("accommodationDB")
//	accommodationsCollection := bookingDatabase.Collection("accommodations")
//	return accommodationsCollection
//}

func (repo *AccommodationRepo) GetAll() (model.Accommodations, error) {
	// Initialise context (after 5 seconds timeout, abort operation)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var accommodations model.Accommodations
	accommodationsCursor, err := repo.accommodations.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err = accommodationsCursor.All(ctx, &accommodations); err != nil {
		return nil, err
	}
	return accommodations, nil
}

func (repo *AccommodationRepo) DeleteAll() {
	repo.accommodations.DeleteMany(context.TODO(), bson.D{{}})
}

func (repo *AccommodationRepo) FindAccommodationByEmail(email string) (*model.Accommodation, error) {
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//
	//accommodationsCollection := repo.getCollection()
	//
	//var accommodation model.Accommodation
	//filter := bson.M{"email": bson.M{"$eq": email}}
	//err := accommodationsCollection.FindOne(ctx, filter).Decode(&accommodation)
	//if err != nil {
	//	repo.logger.Println(err)
	//	return nil, err
	//}
	//return &accommodation, nil

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var accommodation model.Accommodation
	filter := bson.M{"email": bson.M{"$eq": email}}
	err := repo.accommodations.FindOne(ctx, filter).Decode(&accommodation)
	if err != nil {
		return nil, err
	}
	return &accommodation, nil
}

func (repo *AccommodationRepo) GetById(id string) (*model.Accommodation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var accommodation model.Accommodation
	objID, _ := primitive.ObjectIDFromHex(id)
	err := repo.accommodations.FindOne(ctx, bson.M{"_id": objID}).Decode(&accommodation)
	if err != nil {
		return nil, err
	}
	return &accommodation, nil
}

func (repo *AccommodationRepo) SearchAccommodation(searchCriteria accommodationGw.GetAccommodationsByParamsRequest) model.Accommodations {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"country": searchCriteria.GetSearchParams().Country, "city": searchCriteria.GetSearchParams().City}

	var accommodations model.Accommodations
	accommodationsCollection := repo.accommodations
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

func (accommmodationRepo *AccommodationRepo) AddAvailabilityForAccommodation(accommodationID primitive.ObjectID, availability *model.Availability) error {

	// Add the new availability to the availability array
	update := bson.M{
		"$push": bson.M{
			"availabilities": availability,
		},
	}
	var filter = bson.M{"_id": accommodationID}
	_, err := accommmodationRepo.accommodations.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return status.Errorf(
			codes.Internal,
			"Error updating document: %v",
			err,
		)
	}
	return nil
}
