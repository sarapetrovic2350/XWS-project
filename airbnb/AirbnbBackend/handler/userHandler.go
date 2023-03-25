package handler

import (
	"Rest/model"
	"Rest/service"
	"context"
	"log"
	"net/http"
)

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
	user := h.Context().Value(KeyProduct{}).(*model.User)
	if err := handler.userService.CreateUser(user); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	rw.WriteHeader(http.StatusCreated)
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
