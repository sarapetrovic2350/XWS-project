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
		port = "8091"
	}

	// Initialize context
	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	//Initialize the logger we are going to use, with prefix and datetime for every log
	logger := log.New(os.Stdout, "[airbnb-api] ", log.LstdFlags)
	storeLogger := log.New(os.Stdout, "[airbnb-store] ", log.LstdFlags)

	//dburi := os.Getenv("MONGO_DB_URI")
	accommodationDB := os.Getenv("ACCOMMODATION_DB_DOMAIN")
	accommodationDBPort := os.Getenv("ACCOMMODATION_DB_PORT")
	uri := fmt.Sprintf("mongodb://%s:%s/", accommodationDB, accommodationDBPort)
	options := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), options)
	if err != nil {
		log.Panic(err)
	}

	// NoSQL: Initialize repositories
	accommodationRepository := repository.NewAccommodationRepo(client, storeLogger)
	defer accommodationRepository.Disconnect(timeoutContext)

	availabilityRepository := repository.NewAvailabilityRepo(client, storeLogger)
	defer availabilityRepository.Disconnect(timeoutContext)

	// Initialize services
	accommodationService := service.NewAccommodationService(accommodationRepository)
	availabilityService := service.NewAvailabilityService(availabilityRepository)

	//Initialize handlers and inject said logger
	accommodationHandler := handler.NewAccommodationHandler(logger, accommodationService)
	availabilityHandler := handler.NewAvailabilityHandler(logger, availabilityService)

	//Initialize the router and add a middleware for all the requests
	//router := mux.NewRouter()
	headers := gorillaHandlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Access-Control-Allow-Headers", "text/plain"})
	methods := gorillaHandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	origins := gorillaHandlers.AllowedOrigins([]string{"http://localhost:4200"})
	credential := gorillaHandlers.AllowCredentials()
	h := gorillaHandlers.CORS(headers, methods, origins, credential)

	//router.Use(userHandler.MiddlewareContentTypeSet)
	newRouter := mux.NewRouter().StrictSlash(true)
	newRouter.Use(accommodationHandler.MiddlewareContentTypeSet)
	newRouter.Use(availabilityHandler.MiddlewareContentTypeSet)

	//createUserRouter := router.Methods(http.MethodPost).Subrouter()
	newRouter.HandleFunc("/", accommodationHandler.CreateAccommodation).Methods("POST")
	//createUserRouter.Use(userHandler.MiddlewareUserDeserialization)

	//getUsersRouter := router.Methods(http.MethodGet).Subrouter()
	newRouter.HandleFunc("/", accommodationHandler.GetAllAccommodations).Methods("GET")

	newRouter.HandleFunc("/search", accommodationHandler.SearchAccommodations).Methods("POST")

	newRouter.HandleFunc("/createAvailability", availabilityHandler.CreateAvailability).Methods("POST")

	newRouter.HandleFunc("/availabilities", availabilityHandler.GetAllAvailabilities).Methods("GET")

	//cors := gorillaHandlers.CORS(gorillaHandlers.AllowedOrigins([]string{"*"}))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), h(newRouter)))

}
