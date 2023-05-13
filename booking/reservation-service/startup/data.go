package startup

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reservation-service/model"
)

var reservations = []*model.Reservation{
	{},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
