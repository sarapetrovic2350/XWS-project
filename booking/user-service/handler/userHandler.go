package handler

import (
	user "common/proto/user-service/pb"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"user-service/service"
)

type KeyProduct struct{}

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
	fmt.Println("In Delete grpc api")
	fmt.Print("Request.Id: ")
	fmt.Println(request.Id)
	formatedId, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}
	deletedUser, err := handler.userService.Get(formatedId)
	fmt.Print("deletedUser after mapping: ")
	fmt.Println(deletedUser)
	err = handler.userService.Delete(request)
	if err != nil {
		return nil, err
	}
	return &user.DeleteUserResponse{
		User: mapUser(deletedUser)}, nil
}
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

//func (handler *UserHandler) CreateUser(rw http.ResponseWriter, h *http.Request) {
//	fmt.Println("creating")
//	var user model.User
//	err := json.NewDecoder(h.Body).Decode(&user)
//	if err != nil {
//		rw.WriteHeader(http.StatusBadRequest)
//		return
//	}
//	fmt.Println(user)
//	existingUser, _ := handler.userService.FindUserByEmail(user.Email)
//	if existingUser != nil {
//		rw.WriteHeader(http.StatusBadRequest)
//		return
//	}
//	err = handler.userService.CreateUser(&user)
//	if err != nil {
//		fmt.Println(err)
//		rw.WriteHeader(http.StatusExpectationFailed)
//	}
//	rw.WriteHeader(http.StatusCreated)
//	rw.Header().Set("Content-Type", "application/json")
//}

//func (handler *UserHandler) Login(rw http.ResponseWriter, h *http.Request) {
//	var userLogin dto.Login
//	err := json.NewDecoder(h.Body).Decode(&userLogin)
//	token, err := handler.userService.Login(&userLogin)
//	fmt.Println(token)
//	if err != nil {
//		rw.WriteHeader(http.StatusBadRequest)
//		return
//	}
//	rw.WriteHeader(http.StatusOK)
//	rwToken, err := json.Marshal(token)
//	rw.Write(rwToken)
//}
//func (handler *UserHandler) GetUserByEmail(rw http.ResponseWriter, h *http.Request) {
//	vars := mux.Vars(h)
//	email := vars["email"]
//	fmt.Println(vars)
//	user, err := handler.userService.FindUserByEmail(email)
//	fmt.Println(user)
//	if err != nil {
//		handler.logger.Print("Database exception: ", err)
//	}
//	if user == nil {
//		http.Error(rw, "User with given email not found", http.StatusNotFound)
//		handler.logger.Printf("User with email: '%s' not found", email)
//		return
//	}
//	err = user.ToJSON(rw)
//	if err != nil {
//		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
//		handler.logger.Fatal("Unable to convert to json :", err)
//		return
//	}
//}
