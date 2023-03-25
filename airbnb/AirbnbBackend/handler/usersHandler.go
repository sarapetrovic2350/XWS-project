package handler

import (
	"Rest/model"
	"Rest/repository"
	"context"
	"log"
	"net/http"
)

type UsersHandler struct {
	logger *log.Logger
	// NoSQL: injecting users repository
	repo *repository.UsersRepo
}

// Injecting the logger makes this code much more testable.
func NewUsersHandler(l *log.Logger, r *repository.UsersRepo) *UsersHandler {
	return &UsersHandler{l, r}
}

func (uh *UsersHandler) CreateUser(rw http.ResponseWriter, h *http.Request) {
	user := h.Context().Value(KeyProduct{}).(*model.User)
	uh.repo.Insert(user)
	rw.WriteHeader(http.StatusCreated)
}

func (uh *UsersHandler) GetAllUsers(rw http.ResponseWriter, h *http.Request) {
	users, err := uh.repo.GetAll()
	if err != nil {
		uh.logger.Print("Database exception: ", err)
	}

	if users == nil {
		return
	}

	err = users.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		uh.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

func (uh *UsersHandler) MiddlewareUserDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		user := &model.User{}
		err := user.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			uh.logger.Fatal(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyProduct{}, user)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}

func (uh *UsersHandler) MiddlewareContentTypeSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		uh.logger.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)

		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, h)
	})
}
