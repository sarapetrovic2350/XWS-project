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

type TicketRepo struct {
	client *mongo.Client
	logger *log.Logger
}

func NewTicketRepo(client *mongo.Client, logger *log.Logger) *TicketRepo {
	return &TicketRepo{client, logger}
}

func (repo *TicketRepo) Disconnect(ctx context.Context) error {
	err := repo.client.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (repo *TicketRepo) Insert(ticket *model.Ticket) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ticketsCollection := repo.getCollection()

	result, err := ticketsCollection.InsertOne(ctx, &ticket)
	if err != nil {
		repo.logger.Println(err)
		return err
	}
	repo.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}

func (repo *TicketRepo) getCollection() *mongo.Collection {
	airbnbDatabase := repo.client.Database("airbnbDB")
	ticketsCollection := airbnbDatabase.Collection("tickets")
	return ticketsCollection
}

func (repo *TicketRepo) GetAll() (model.Tickets, error) {
	// Initialise context (after 5 seconds timeout, abort operation)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ticketsCollection := repo.getCollection()

	var tickets model.Tickets
	ticketsCursor, err := ticketsCollection.Find(ctx, bson.M{})
	if err != nil {
		repo.logger.Println(err)
		return nil, err
	}
	if err = ticketsCursor.All(ctx, &tickets); err != nil {
		repo.logger.Println(err)
		return nil, err
	}
	return tickets, nil
}

func (repo *TicketRepo) GetById(id string) (*model.Ticket, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ticketsCollection := repo.getCollection()

	var ticket model.Ticket
	objID, _ := primitive.ObjectIDFromHex(id)
	err := ticketsCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&ticket)
	if err != nil {
		repo.logger.Println(err)
		return nil, err
	}
	return &ticket, nil
}

func (repo *TicketRepo) Update(id string, ticket *model.Ticket) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ticketsCollection := repo.getCollection()

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{
		"date_of_purchase":  ticket.DateOfPurchase,
		"number_of_tickets": ticket.NumberOfTickets,
		"total_price":       ticket.TotalPrice,
		"id_user":           ticket.IdUser,
		"id_flight":         ticket.IdFlight,
	}}
	result, err := ticketsCollection.UpdateOne(ctx, filter, update)
	repo.logger.Printf("Documents matched: %v\n", result.MatchedCount)
	repo.logger.Printf("Documents updated: %v\n", result.ModifiedCount)

	if err != nil {
		repo.logger.Println(err)
		return err
	}
	return nil
}

func (repo *TicketRepo) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ticketsCollection := repo.getCollection()

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	result, err := ticketsCollection.DeleteOne(ctx, filter)
	if err != nil {
		repo.logger.Println(err)
		return err
	}
	repo.logger.Printf("Documents deleted: %v\n", result.DeletedCount)
	return nil
}
