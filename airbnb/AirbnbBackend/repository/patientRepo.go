package repository

import (
	"Rest/model"
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	// NoSQL: module containing Mongo api client
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// PatientRepo struct encapsulating Mongo api client
type PatientRepo struct {
	client *mongo.Client
	logger *log.Logger
}

func NewPatientRepo(client *mongo.Client, logger *log.Logger) *PatientRepo {
	return &PatientRepo{client, logger}
}

// Disconnect from database
func (repo *PatientRepo) Disconnect(ctx context.Context) error {
	err := repo.client.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Ping check database connection
func (repo *PatientRepo) Ping() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check connection -> if no error, connection is established
	err := repo.client.Ping(ctx, readpref.Primary())
	if err != nil {
		repo.logger.Println(err)
	}

	// Print available databases
	databases, err := repo.client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		repo.logger.Println(err)
	}
	fmt.Println(databases)
}

func (repo *PatientRepo) GetAll() (model.Patients, error) {
	// Initialise context (after 5 seconds timeout, abort operation)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	patientsCollection := repo.getCollection()

	var patients model.Patients
	patientsCursor, err := patientsCollection.Find(ctx, bson.M{})
	if err != nil {
		repo.logger.Println(err)
		return nil, err
	}
	if err = patientsCursor.All(ctx, &patients); err != nil {
		repo.logger.Println(err)
		return nil, err
	}
	return patients, nil
}

func (repo *PatientRepo) GetById(id string) (*model.Patient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	patientsCollection := repo.getCollection()

	var patient model.Patient
	objID, _ := primitive.ObjectIDFromHex(id)
	err := patientsCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&patient)
	if err != nil {
		repo.logger.Println(err)
		return nil, err
	}
	return &patient, nil
}

func (repo *PatientRepo) GetByName(name string) (model.Patients, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	patientsCollection := repo.getCollection()

	var patients model.Patients
	patientsCursor, err := patientsCollection.Find(ctx, bson.M{"name": name})
	if err != nil {
		repo.logger.Println(err)
		return nil, err
	}
	if err = patientsCursor.All(ctx, &patients); err != nil {
		repo.logger.Println(err)
		return nil, err
	}
	return patients, nil
}

func (repo *PatientRepo) Insert(patient *model.Patient) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	patientsCollection := repo.getCollection()

	result, err := patientsCollection.InsertOne(ctx, &patient)
	if err != nil {
		repo.logger.Println(err)
		return err
	}
	repo.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}

func (repo *PatientRepo) Update(id string, patient *model.Patient) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	patientsCollection := repo.getCollection()

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{
		"name":    patient.Name,
		"surname": patient.Surname,
	}}
	result, err := patientsCollection.UpdateOne(ctx, filter, update)
	repo.logger.Printf("Documents matched: %v\n", result.MatchedCount)
	repo.logger.Printf("Documents updated: %v\n", result.ModifiedCount)

	if err != nil {
		repo.logger.Println(err)
		return err
	}
	return nil
}

func (repo *PatientRepo) AddPhoneNumber(id string, phoneNumber string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	patientsCollection := repo.getCollection()

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}
	update := bson.M{"$push": bson.M{
		"phoneNumbers": phoneNumber,
	}}
	result, err := patientsCollection.UpdateOne(ctx, filter, update)
	repo.logger.Printf("Documents matched: %v\n", result.MatchedCount)
	repo.logger.Printf("Documents updated: %v\n", result.ModifiedCount)

	if err != nil {
		repo.logger.Println(err)
		return err
	}
	return nil
}

func (repo *PatientRepo) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	patientsCollection := repo.getCollection()

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	result, err := patientsCollection.DeleteOne(ctx, filter)
	if err != nil {
		repo.logger.Println(err)
		return err
	}
	repo.logger.Printf("Documents deleted: %v\n", result.DeletedCount)
	return nil
}

func (repo *PatientRepo) AddAnamnesis(id string, anamnesis *model.Anamnesis) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	patientsCollection := repo.getCollection()

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	update := bson.M{"$push": bson.M{
		"anamnesis": anamnesis,
	}}
	result, err := patientsCollection.UpdateOne(ctx, filter, update)
	repo.logger.Printf("Documents matched: %v\n", result.MatchedCount)
	repo.logger.Printf("Documents updated: %v\n", result.ModifiedCount)

	if err != nil {
		repo.logger.Println(err)
		return err
	}
	return nil
}

func (repo *PatientRepo) AddTherapy(id string, therapy *model.Therapy) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	patientsCollection := repo.getCollection()

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	update := bson.M{"$push": bson.M{
		"therapy": therapy,
	}}
	result, err := patientsCollection.UpdateOne(ctx, filter, update)
	repo.logger.Printf("Documents matched: %v\n", result.MatchedCount)
	repo.logger.Printf("Documents updated: %v\n", result.ModifiedCount)

	if err != nil {
		repo.logger.Println(err)
		return err
	}
	return nil
}

func (repo *PatientRepo) UpdateAddress(id string, address *model.Address) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	patientsCollection := repo.getCollection()

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	update := bson.M{"$set": bson.M{
		"address": address,
	}}
	result, err := patientsCollection.UpdateOne(ctx, filter, update)
	repo.logger.Printf("Documents matched: %v\n", result.MatchedCount)
	repo.logger.Printf("Documents updated: %v\n", result.ModifiedCount)

	if err != nil {
		repo.logger.Println(err)
		return err
	}
	return nil
}

func (repo *PatientRepo) ChangePhone(id string, index int, phoneNumber string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	patientsCollection := repo.getCollection()

	// What happens if set value for index=10, but we only have 3 phone numbers?
	// -> Every value in between will be set to an empty string
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	update := bson.M{"$set": bson.M{
		"phoneNumbers." + strconv.Itoa(index): phoneNumber,
	}}
	result, err := patientsCollection.UpdateOne(ctx, filter, update)
	repo.logger.Printf("Documents matched: %v\n", result.MatchedCount)
	repo.logger.Printf("Documents updated: %v\n", result.ModifiedCount)

	if err != nil {
		repo.logger.Println(err)
		return err
	}
	return nil
}

// BONUS

func (repo *PatientRepo) Receipt(id string) (float64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	patientsCollection := repo.getCollection()

	objID, _ := primitive.ObjectIDFromHex(id)
	matchStage := bson.D{{"$match", bson.D{{"_id", objID}}}}
	sumStage := bson.D{{"$addFields",
		bson.D{{"total", bson.D{{"$add",
			bson.D{{"$sum", "$therapy.price"}},
		}},
		}},
	}}
	projectStage := bson.D{{"$project",
		bson.D{{"total", 1}},
	}}

	cursor, err := patientsCollection.Aggregate(ctx, mongo.Pipeline{matchStage, sumStage, projectStage})
	if err != nil {
		repo.logger.Println(err)
		return -1, err
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		repo.logger.Println(err)
		return -1, err
	}
	for _, result := range results {
		repo.logger.Println(result)
		return result["total"].(float64), nil
	}
	return -1, nil
}

func (repo *PatientRepo) Report() (map[string]float64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	patientsCollection := repo.getCollection()

	sumStage := bson.D{{"$addFields",
		bson.D{{"total", bson.D{{"$add",
			bson.D{{"$sum", "$therapy.price"}},
		}},
		}},
	}}
	projectStage := bson.D{{"$project",
		bson.D{{"name", 1}, {"surname", 1}, {"total", 1}},
	}}

	cursor, err := patientsCollection.Aggregate(ctx, mongo.Pipeline{sumStage, projectStage})
	if err != nil {
		repo.logger.Println(err)
		return nil, err
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		repo.logger.Println(err)
		return nil, err
	}
	report := make(map[string]float64)
	for _, result := range results {
		repo.logger.Println(result)
		report[result["_id"].(primitive.ObjectID).Hex()] = result["total"].(float64)
	}
	return report, nil
}

func (repo *PatientRepo) getCollection() *mongo.Collection {
	patientDatabase := repo.client.Database("patientsDB")
	patientsCollection := patientDatabase.Collection("patients")
	return patientsCollection
}
