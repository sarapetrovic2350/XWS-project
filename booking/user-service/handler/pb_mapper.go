package handler

import (
	user "common/proto/user-service/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"user-service/model"
)

func mapNewUser(userPb *user.NewUser) *model.User {
	user := &model.User{

		Id:        primitive.NewObjectID(),
		Role:      userPb.Role,
		Password:  userPb.Password,
		Email:     userPb.Email,
		FirstName: userPb.FirstName,
		LastName:  userPb.LastName,
		Address: model.Address{
			Country:      userPb.Address.Country,
			City:         userPb.Address.City,
			Street:       userPb.Address.Street,
			StreetNumber: userPb.Address.Number,
		},
	}
	return user
}

func mapUpdatedUser(userPb *user.User) *model.User {
	UserId, _ := primitive.ObjectIDFromHex(userPb.Id)
	user := &model.User{
		Id:        UserId,
		Role:      userPb.Role,
		Password:  userPb.Password,
		Email:     userPb.Email,
		FirstName: userPb.FirstName,
		LastName:  userPb.LastName,
		Address: model.Address{
			Country:      userPb.Address.Country,
			City:         userPb.Address.City,
			Street:       userPb.Address.Street,
			StreetNumber: userPb.Address.Number,
		},
		IsSuperHost: userPb.IsSuperHost,
	}
	return user
}

func mapUser(modelUser *model.User) *user.User {
	userPb := &user.User{
		Id:        modelUser.Id.Hex(),
		Password:  modelUser.Password,
		Email:     modelUser.Email,
		FirstName: modelUser.FirstName,
		LastName:  modelUser.LastName,
		Role:      modelUser.Role,
		Address: &user.AddressDTO{
			Country: modelUser.Address.Country,
			City:    modelUser.Address.City,
			Street:  modelUser.Address.Street,
			Number:  modelUser.Address.StreetNumber,
		},
		IsSuperHost: modelUser.IsSuperHost,
	}
	return userPb
}
