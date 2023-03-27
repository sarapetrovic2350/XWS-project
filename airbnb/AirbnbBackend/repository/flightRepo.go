package repository

import (
	"Rest/model"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type FlightRepo struct {
	client *mongo.Client
	logger *log.Logger
}

func NewFlightRepo(client *mongo.Client, logger *log.Logger) *FlightRepo {
	return &FlightRepo{client, logger}
}

func (repo *FlightRepo) Disconnect(ctx context.Context) error {
	err := repo.client.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (repo *FlightRepo) Insert(flight *model.Flight) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	flightsCollection := repo.getCollection()

	result, err := flightsCollection.InsertOne(ctx, &flight)
	if err != nil {
		repo.logger.Println(err)
		return err
	}
	repo.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}

func (repo *FlightRepo) getCollection() *mongo.Collection {
	airbnbDatabase := repo.client.Database("airbnbDB")
	flightsCollection := airbnbDatabase.Collection("flights")
	return flightsCollection
}

func (repo *FlightRepo) GetAll() (model.Flights, error) {
	// Initialise context (after 5 seconds timeout, abort operation)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	flightsCollection := repo.getCollection()

	var flights model.Flights
	flightsCursor, err := flightsCollection.Find(ctx, bson.M{})
	if err != nil {
		repo.logger.Println(err)
		return nil, err
	}
	if err = flightsCursor.All(ctx, &flights); err != nil {
		repo.logger.Println(err)
		return nil, err
	}
	return flights, nil
}

func (repo *FlightRepo) GetById(id string) (*model.Flight, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	flightsCollection := repo.getCollection()

	var flight model.Flight
	objID, _ := primitive.ObjectIDFromHex(id)
	err := flightsCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&flight)
	if err != nil {
		repo.logger.Println(err)
		return nil, err
	}
	return &flight, nil
}

func (repo *FlightRepo) Update(id string, flight *model.Flight) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	flightsCollection := repo.getCollection()

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{
		"date_time":             flight.DateTime,
		"departure":             flight.Departure,
		"arrival":               flight.Arrival,
		"price":                 flight.Price,
		"total_number_of_seats": flight.TotalNumberOfSeats,
		"available_seats":       flight.AvailableSeats,
	}}
	result, err := flightsCollection.UpdateOne(ctx, filter, update)
	repo.logger.Printf("Documents matched: %v\n", result.MatchedCount)
	repo.logger.Printf("Documents updated: %v\n", result.ModifiedCount)

	if err != nil {
		repo.logger.Println(err)
		return err
	}
	return nil
}

func (repo *FlightRepo) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	flightsCollection := repo.getCollection()

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	result, err := flightsCollection.DeleteOne(ctx, filter)
	if err != nil {
		repo.logger.Println(err)
		return err
	}
	repo.logger.Printf("Documents deleted: %v\n", result.DeletedCount)
	return nil
}
