package main

import (
	"Rest/handler"
	"Rest/repository"
	"Rest/service"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	//Reading from environment, if not set we will default it to 8080.
	//This allows flexibility in different environments (for eg. when running multiple docker api's and want to override the default port)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8090"
	}

	// Initialize context
	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	//Initialize the logger we are going to use, with prefix and datetime for every log
	logger := log.New(os.Stdout, "[airbnb-api] ", log.LstdFlags)
	storeLogger := log.New(os.Stdout, "[airbnb-store] ", log.LstdFlags)

	//dburi := os.Getenv("MONGO_DB_URI")
	userDB := os.Getenv("USER_DB_DOMAIN")
	userDBport := os.Getenv("USER_DB_PORT")
	uri := fmt.Sprintf("mongodb://%s:%s/", userDB, userDBport)
	options := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), options)
	if err != nil {
		log.Panic(err)
	}

	// NoSQL: Initialize repositories
	userRepository := repository.NewUserRepo(client, storeLogger)
	defer userRepository.Disconnect(timeoutContext)
	userRepository.Ping()

	// Initialize services
	userService := service.NewUserService(userRepository)

	//Initialize handlers and inject said logger
	userHandler := handler.NewUserHandler(logger, userService)

	//Initialize the router and add a middleware for all the requests
	router := mux.NewRouter()
	headers := gorillaHandlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Access-Control-Allow-Headers", "text/plain"})
	methods := gorillaHandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	origins := gorillaHandlers.AllowedOrigins([]string{"http://localhost:8080"})
	credential := gorillaHandlers.AllowCredentials()
	h := gorillaHandlers.CORS(headers, methods, origins, credential)

	router.Use(userHandler.MiddlewareContentTypeSet)
	newRouter := mux.NewRouter().StrictSlash(true)

	//createUserRouter := router.Methods(http.MethodPost).Subrouter()
	newRouter.HandleFunc("/", userHandler.CreateUser).Methods("POST")
	//createUserRouter.Use(userHandler.MiddlewareUserDeserialization)

	//getUsersRouter := router.Methods(http.MethodGet).Subrouter()
	newRouter.HandleFunc("/", userHandler.GetAllUsers).Methods("GET")

	LoginRouter := router.Methods(http.MethodPost).Subrouter()
	LoginRouter.HandleFunc("/users/login", userHandler.Login)

	//cors := gorillaHandlers.CORS(gorillaHandlers.AllowedOrigins([]string{"*"}))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), h(newRouter)))

}
