package startup

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rating-service/model"
)

var ratingsHost = []*model.RatingHost{
	{
		Id:      getObjectId("test1_id"),
		HostId:  "host1Id",
		GuestId: "guest1Id",
		Rate:    5,
	},
}
var ratingsAccommodation = []*model.RatingAccommodation{
	{
		Id:              getObjectId("test1_id"),
		AccommodationId: "accommodationId1",
		GuestId:         "guest1Id",
		Rate:            5,
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
