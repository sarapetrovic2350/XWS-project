package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reservation-service/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "reservation"
	COLLECTION = "reservation"
)

type ReservationRepo struct {
	reservations *mongo.Collection
}

func NewReservationRepo(client *mongo.Client) model.ReservationStore {
	reservations := client.Database(DATABASE).Collection(COLLECTION)
	return &ReservationRepo{
		reservations: reservations,
	}
}

func (repo *ReservationRepo) Insert(reservation *model.Reservation) error {
	result, err := repo.reservations.InsertOne(context.TODO(), reservation)
	if err != nil {
		return err
	}
	reservation.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (repo *ReservationRepo) GetReservationsByUserId(userId string) (model.Reservations, error) {
	fmt.Println("reservations repo get by id")
	fmt.Println(userId)
	reservations, err := repo.GetAll()
	var retReservations model.Reservations
	for _, itr := range reservations {
		if itr.UserId == userId {
			retReservations = append(retReservations, itr)
			fmt.Println(retReservations)
		}
	}
	if err != nil {
		return nil, err
	}
	return retReservations, nil
}

func (repo *ReservationRepo) GetReservationsByAccommodationId(accommodationId string) (model.Reservations, error) {
	reservations, err := repo.GetAll()
	var retReservations model.Reservations
	for _, itr := range reservations {
		if itr.AccommodationId == accommodationId {
			retReservations = append(retReservations, itr)
		}
	}
	if err != nil {
		return nil, err
	}
	return retReservations, nil
}

func (repo *ReservationRepo) GetAll() (model.Reservations, error) {
	// Initialise context (after 5 seconds timeout, abort operation)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var reservations model.Reservations
	reservationsCursor, err := repo.reservations.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err = reservationsCursor.All(ctx, &reservations); err != nil {
		return nil, err
	}
	return reservations, nil
}

func (repo *ReservationRepo) DeleteAll() {
	repo.reservations.DeleteMany(context.TODO(), bson.D{{}})
}

func (repo *ReservationRepo) GetById(id string) (*model.Reservation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var reservation model.Reservation
	objID, _ := primitive.ObjectIDFromHex(id)
	err := repo.reservations.FindOne(ctx, bson.M{"_id": objID}).Decode(&reservation)
	if err != nil {
		return nil, err
	}
	return &reservation, nil
}

func (repo *ReservationRepo) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	_, err := repo.reservations.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
