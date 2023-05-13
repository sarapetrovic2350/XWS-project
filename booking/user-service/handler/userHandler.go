package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"user-service/dto"
	"user-service/model"
	"user-service/service"
)

type KeyProduct struct{}

type UserHandler struct {
	logger *log.Logger
	// NoSQL: injecting user service
	userService *service.UserService
}

// NewUserHandler Injecting the logger makes this code much more testable.
func NewUserHandler(l *log.Logger, s *service.UserService) *UserHandler {
	return &UserHandler{l, s}
}

func (handler *UserHandler) CreateUser(rw http.ResponseWriter, h *http.Request) {
	fmt.Println("creating")
	var user model.User
	err := json.NewDecoder(h.Body).Decode(&user)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(user)
	existingUser, _ := handler.userService.FindUserByEmail(user.Email)
	if existingUser != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.userService.CreateUser(&user)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusExpectationFailed)
	}
	rw.WriteHeader(http.StatusCreated)
	rw.Header().Set("Content-Type", "application/json")
}

func (handler *UserHandler) Login(rw http.ResponseWriter, h *http.Request) {
	var userLogin dto.Login
	err := json.NewDecoder(h.Body).Decode(&userLogin)
	if err != nil {
		handler.logger.Print("Database exception: ", err)
	}
	token, err := handler.userService.Login(&userLogin)
	fmt.Println(token)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	rw.WriteHeader(http.StatusOK)
	rwToken, err := json.Marshal(token)
	rw.Write(rwToken)
}

func (handler *UserHandler) GetAllUsers(rw http.ResponseWriter, h *http.Request) {
	users, err := handler.userService.GetAllUsers()
	if err != nil {
		handler.logger.Print("Database exception: ", err)
	}

	if users == nil {
		return
	}

	err = users.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		handler.logger.Fatal("Unable to convert to json :", err)
		return
	}
}
func (handler *UserHandler) GetUserByEmail(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	email := vars["email"]
	fmt.Println(vars)
	user, err := handler.userService.FindUserByEmail(email)
	fmt.Println(user)
	if err != nil {
		handler.logger.Print("Database exception: ", err)
	}
	if user == nil {
		http.Error(rw, "User with given email not found", http.StatusNotFound)
		handler.logger.Printf("User with email: '%s' not found", email)
		return
	}
	err = user.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		handler.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

func (handler *UserHandler) MiddlewareUserDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		user := &model.User{}
		err := user.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			handler.logger.Fatal(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyProduct{}, user)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}

func (handler *UserHandler) MiddlewareContentTypeSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		handler.logger.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)

		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, h)
	})
}
