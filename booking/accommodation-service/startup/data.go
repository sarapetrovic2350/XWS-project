package startup

import (
	"accommodation-service/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var accommodations = []*model.Accommodation{
	{},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
