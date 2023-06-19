package startup

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"user-service/model"
)

var users = []*model.User{
	{
		Id:        getObjectId("test1_id"),
		Email:     "host1Id",
		FirstName: "Name1",
		LastName:  "LastName1",
		Password:  "password1",
		Address: model.Address{
			Street:       "Ulica",
			StreetNumber: "5",
			City:         "Novi Sad",
			Country:      "Srbija",
		},
		Role: model.GuestUserRole,
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
