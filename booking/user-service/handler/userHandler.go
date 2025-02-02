package handler

import (
	user "common/proto/user-service/pb"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"user-service/service"
)

type UserHandler struct {
	user.UnimplementedUserServiceServer
	userService *service.UserService
}

// NewUserHandler Injecting the logger makes this code much more testable.
func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{
		userService: s,
	}
}
func (handler *UserHandler) GetAll(ctx context.Context, request *user.GetAllRequest) (*user.GetAllResponse, error) {
	fmt.Println("In GetAll grpc api")
	users, err := handler.userService.GetAllUsers()
	if err != nil {
		return nil, err
	}
	response := &user.GetAllResponse{
		Users: []*user.User{},
	}
	for _, modelUser := range users {
		current := mapUser(modelUser)
		response.Users = append(response.Users, current)
	}
	return response, nil
}
func (handler *UserHandler) GetAllHosts(ctx context.Context, request *user.GetAllHostsRequest) (*user.GetAllHostsResponse, error) {
	fmt.Println("In GetAllHosts grpc api")
	hostUsers, err := handler.userService.GetAllHostUsers()
	if err != nil {
		return nil, err
	}
	response := &user.GetAllHostsResponse{
		Users: []*user.User{},
	}
	for _, modelUser := range hostUsers {
		current := mapUser(modelUser)
		response.Users = append(response.Users, current)
	}
	return response, nil
}
func (handler *UserHandler) CreateUser(ctx context.Context, request *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	fmt.Println("In CreateUser grpc api")
	fmt.Print("Request.User: ")
	fmt.Println(request.User)
	modelUser := mapNewUser(request.User)
	fmt.Print("user after mapping: ")
	fmt.Println(modelUser)
	err := handler.userService.CreateUser(modelUser)
	if err != nil {
		return nil, err
	}
	return &user.CreateUserResponse{
		User: mapUser(modelUser),
	}, nil
}
func (handler *UserHandler) DeleteUser(ctx context.Context, request *user.DeleteUserRequest) (*user.DeleteUserResponse, error) {
	fmt.Println("In Delete Guest grpc api")
	fmt.Print("Request.Id: ")
	fmt.Println(request.Id)
	formattedId, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}

	err = handler.userService.DeleteUser(formattedId)
	if err != nil {
		return nil, err
	}
	return &user.DeleteUserResponse{}, nil
}

//func (handler *UserHandler) DeleteGuestUser(ctx context.Context, request *user.DeleteUserRequest) (*user.DeleteUserResponse, error) {
//	fmt.Println("In Delete Guest grpc api")
//	fmt.Print("Request.Id: ")
//	fmt.Println(request.Id)
//	formatedId, err := primitive.ObjectIDFromHex(request.Id)
//	if err != nil {
//		return nil, err
//	}
//	deletedUser, err := handler.userService.Get(formatedId)
//	fmt.Print("deletedUser after mapping: ")
//	fmt.Println(deletedUser)
//	err = handler.userService.DeleteGuestUser(request)
//	if err != nil {
//		return nil, err
//	}
//	return &user.DeleteUserResponse{
//		User: mapUser(deletedUser)}, nil
//}
//
//func (handler *UserHandler) DeleteHostUser(ctx context.Context, request *user.DeleteUserRequest) (*user.DeleteUserResponse, error) {
//	fmt.Println("In Delete Host grpc api")
//	fmt.Print("Request.Id: ")
//	fmt.Println(request.Id)
//	formatedId, err := primitive.ObjectIDFromHex(request.Id)
//	if err != nil {
//		return nil, err
//	}
//	deletedUser, err := handler.userService.Get(formatedId)
//	fmt.Print("deletedUser after mapping: ")
//	fmt.Println(deletedUser)
//	err = handler.userService.DeleteHostUser(request)
//	if err != nil {
//		return nil, err
//	}
//	return &user.DeleteUserResponse{
//		User: mapUser(deletedUser)}, nil
//}

func (handler *UserHandler) Login(ctx context.Context, request *user.LoginRequest) (*user.LoginResponse, error) {
	fmt.Println("In Login grpc api")
	fmt.Print("Request.Login.Email: ")
	fmt.Println(request.Login.Email)
	fmt.Print("Request.Login.Password: ")
	fmt.Println(request.Login.Password)
	jwt, err := handler.userService.Login(request.Login.Email, request.Login.Password)
	if err != nil {
		return nil, err
	}
	retVal := &user.LoginResponse{Jwt: jwt}
	return retVal, nil
}

func (handler *UserHandler) GetUserByEmail(ctx context.Context, request *user.GetUserByEmailRequest) (*user.GetUserByEmailResponse, error) {
	fmt.Println("In GetUserByEmail grpc api")
	fmt.Print("Request.Email: ")
	modelUser, err := handler.userService.FindUserByEmail(request.Email)
	if err != nil {
		return nil, err
	}
	return &user.GetUserByEmailResponse{
		User: mapUser(modelUser),
	}, nil
}

func (handler *UserHandler) GetUserById(ctx context.Context, request *user.GetUserByIdRequest) (*user.GetUserByIdResponse, error) {
	fmt.Println("In GetUserById grpc api")
	modelUser, err := handler.userService.FindUserById(request.Id)
	if err != nil {
		return nil, err
	}
	return &user.GetUserByIdResponse{
		User: mapUser(modelUser),
	}, nil
}

func (handler *UserHandler) UpdateUser(ctx context.Context, request *user.UpdateUserRequest) (*user.UpdateUserResponse, error) {
	fmt.Println("In UpdateUser grpc api")
	fmt.Print("Request.User: ")
	fmt.Println(request.User)
	modelUser := mapUpdatedUser(request.User)
	fmt.Print("user after mapping: ")
	fmt.Println(modelUser)
	err := handler.userService.Update(modelUser)
	if err != nil {
		return nil, err
	}
	return &user.UpdateUserResponse{
		User: mapUser(modelUser),
	}, nil
}

func (handler *UserHandler) mustEmbedUnimplementedUserServiceServer() {
	//TODO implement me
	panic("implement me")
}

func (handler *UserHandler) GetIfHostIsSuperHost(ctx context.Context, request *user.GetIfHostIsSuperHostRequest) (*user.GetIfHostIsSuperHostResponse, error) {
	fmt.Println("In GetUserByEmail grpc api")
	fmt.Print("Request.Email: ")
	isSuperHost, err := handler.userService.GetIfHostIsSuperHost(request.Id)
	if err != nil {
		return nil, err
	}
	return &user.GetIfHostIsSuperHostResponse{
		IsSuperHost: isSuperHost,
	}, nil
}
