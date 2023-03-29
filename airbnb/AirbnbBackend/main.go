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
		port = "8080"
	}

	// Initialize context
	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	//Initialize the logger we are going to use, with prefix and datetime for every log
	logger := log.New(os.Stdout, "[airbnb-api] ", log.LstdFlags)
	storeLogger := log.New(os.Stdout, "[airbnb-store] ", log.LstdFlags)

	//dburi := os.Getenv("MONGO_DB_URI")

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:pass@mongo:27017"))
	if err != nil {
		log.Panic(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Panic(err)
	}

	// NoSQL: Initialize repositories
	repos := repository.InitRepositories(client, storeLogger)
	defer repos.PatientRepo.Disconnect(timeoutContext)
	defer repos.UserRepo.Disconnect(timeoutContext)
	defer repos.FlightRepo.Disconnect(timeoutContext)
	defer repos.TicketRepo.Disconnect(timeoutContext)

	// NoSQL: Checking if the connection was established
	repos.PatientRepo.Ping()

	// Initialize services
	services := service.InitServices(repos)

	//Initialize handlers and inject said logger
	handlers := handler.InitHandlers(logger, services)

	//Initialize the router and add a middleware for all the requests
	router := mux.NewRouter()
	headers := gorillaHandlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Access-Control-Allow-Headers", "text/plain"})
	methods := gorillaHandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	origins := gorillaHandlers.AllowedOrigins([]string{"http://localhost:4200"})
	credential := gorillaHandlers.AllowCredentials()
	h := gorillaHandlers.CORS(headers, methods, origins, credential)

	router.Use(handlers.UserHandler.MiddlewareContentTypeSet)
	router.Use(handlers.FlightHandler.MiddlewareContentTypeSet)
	router.Use(handlers.TicketHandler.MiddlewareContentTypeSet)

	createUserRouter := router.Methods(http.MethodPost).Subrouter()
	createUserRouter.HandleFunc("/users/", handlers.UserHandler.CreateUser)
	createUserRouter.Use(handlers.UserHandler.MiddlewareUserDeserialization)

	getUsersRouter := router.Methods(http.MethodGet).Subrouter()
	getUsersRouter.HandleFunc("/users/", handlers.UserHandler.GetAllUsers)

	//flights

	createFlightRouter := router.Methods(http.MethodPost).Subrouter()
	createFlightRouter.HandleFunc("/flights/createFlight", handlers.FlightHandler.CreateFlight)
	createFlightRouter.Use(handlers.FlightHandler.MiddlewareUserDeserialization)

	getFlightsRouter := router.Methods(http.MethodGet).Subrouter()
	getFlightsRouter.HandleFunc("/flights/getAllFlights", handlers.FlightHandler.GetAllFlights)

	deleteFlightRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteFlightRouter.HandleFunc("/flights/deleteFlight/{id}", handlers.FlightHandler.DeleteFlight)

	updateFlightRouter := router.Methods(http.MethodPut).Subrouter()
	updateFlightRouter.HandleFunc("/flights/updateFlight/{id}", handlers.FlightHandler.UpdateFlight)

	getFlightByIdRouter := router.Methods(http.MethodGet).Subrouter()
	getFlightByIdRouter.HandleFunc("/flights/{id}", handlers.FlightHandler.GetFlightById)

	//tickets

	createTicketRouter := router.Methods(http.MethodPost).Subrouter()
	createTicketRouter.HandleFunc("/tickets/createTicket", handlers.TicketHandler.CreateTicket)
	createTicketRouter.Use(handlers.TicketHandler.MiddlewareTicketDeserialization)

	getTicketRouter := router.Methods(http.MethodGet).Subrouter()
	getTicketRouter.HandleFunc("/tickets/getAllTickets", handlers.TicketHandler.GetAllTickets)

	getTicketByIdRouter := router.Methods(http.MethodGet).Subrouter()
	getTicketByIdRouter.HandleFunc("/tickets/{id}", handlers.TicketHandler.GetTicketById)

	getTicketsByUserRouter := router.Methods(http.MethodGet).Subrouter()
	getTicketsByUserRouter.HandleFunc("/tickets/getTicketsByUser/{id}", handlers.TicketHandler.GetTicketsByUserId)

	//cors := gorillaHandlers.CORS(gorillaHandlers.AllowedOrigins([]string{"*"}))
	LoginRouter := router.Methods(http.MethodPost).Subrouter()
	LoginRouter.HandleFunc("/users/login", handlers.UserHandler.Login)

	//cors := gorillaHandlers.CORS(gorillaHandlers.AllowedOrigins([]string{"*"}))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), h(router)))

	//Initialize the server
	/*server := http.Server{
		Addr:         ":" + port,
		Handler:      cors(router),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	logger.Println("Server listening on port", port)
	//Distribute all the connections to goroutines
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt)
	signal.Notify(sigCh, os.Kill)

	sig := <-sigCh
	logger.Println("Received terminate, graceful shutdown", sig)

	//Try to shutdown gracefully
	if server.Shutdown(timeoutContext) != nil {
		logger.Fatal("Cannot gracefully shutdown...")
	}
	logger.Println("Server stopped")
	*/
}
