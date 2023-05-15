package startup

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reservation-service/model"
	"time"
)

var reservations = []*model.Reservation{
	{
		Id:                getObjectId("test1_id"),
		NumberOfGuests:    5,
		StartDate:         time.Now(),
		EndDate:           time.Now(),
		UserId:            "userId1",
		AccommodationId:   "accommodationId1",
		ReservationStatus: model.ACCEPTED,
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
