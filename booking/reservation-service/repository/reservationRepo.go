package repository

import (
	"context"
	"log"
	"reservation-service/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserRepo struct encapsulating Mongo api client
type ReservationRepo struct {
	client *mongo.Client
	logger *log.Logger
}

func NewReservationRepo(client *mongo.Client, logger *log.Logger) *ReservationRepo {
	return &ReservationRepo{client, logger}
}

// Disconnect from database
func (repo *ReservationRepo) Disconnect(ctx context.Context) error {
	err := repo.client.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (repo *ReservationRepo) Insert(user *model.Reservation) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	reservationsCollection := repo.getCollection()

	result, err := reservationsCollection.InsertOne(ctx, &user)
	if err != nil {
		repo.logger.Println(err)
		return err
	}
	repo.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}

func (repo *ReservationRepo) getCollection() *mongo.Collection {
	bookingDatabase := repo.client.Database("reservationDB")
	reservationsCollection := bookingDatabase.Collection("reservations")
	return reservationsCollection
}

func (repo *ReservationRepo) GetAll() (model.Reservations, error) {
	// Initialise context (after 5 seconds timeout, abort operation)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	reservationsCollection := repo.getCollection()

	var reservations model.Reservations
	reservationsCursor, err := reservationsCollection.Find(ctx, bson.M{})
	if err != nil {
		repo.logger.Println(err)
		return nil, err
	}
	if err = reservationsCursor.All(ctx, &reservations); err != nil {
		repo.logger.Println(err)
		return nil, err
	}
	return reservations, nil
}

func (repo *ReservationRepo) GetById(id string) (*model.Reservation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	reservationsCollection := repo.getCollection()

	var reservation model.Reservation
	objID, _ := primitive.ObjectIDFromHex(id)
	err := reservationsCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&reservation)
	if err != nil {
		repo.logger.Println(err)
		return nil, err
	}
	return &reservation, nil
}
