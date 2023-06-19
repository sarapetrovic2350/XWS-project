package startup

import (
	"accommodation-service/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var accommodations = []*model.Accommodation{
	{
		Id:   getObjectId("test1_id"),
		Name: "Accommodation1",
		Availabilities: []*model.Availability{
			{
				Id:             getObjectId("test2_id"),
				StartDate:      time.Now().AddDate(0, 0, 5),
				EndDate:        time.Now().AddDate(0, 0, 15),
				Price:          2000.5,
				PriceSelection: model.PER_PERSON,
			},
		},
		MinNumberOfGuests: 2,
		MaxNumberOfGuests: 6,
		HostID:            "hostId1",
		Benefits:          []string{"WIFI", "PARKING"},
		Address: model.Address{
			Street:       "Ulica",
			StreetNumber: "5",
			City:         "Novi Sad",
			Country:      "Srbija",
		},
		Images: []model.Image{},
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
