package startup

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rating-service/model"
	"time"
)

var ratingsHost = []*model.RatingHost{
	{
		Id:      getObjectId("test1_id"),
		HostId:  "host1Id",
		GuestId: "guest1Id",
		Rate:    5,
		Date:    time.Now(),
	},
}
var ratingsAccommodation = []*model.RatingAccommodation{
	{
		Id:              getObjectId("test1_id"),
		AccommodationId: "accommodationId1",
		GuestId:         "guest1Id",
		Rate:            5,
		Date:            time.Now(),
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
