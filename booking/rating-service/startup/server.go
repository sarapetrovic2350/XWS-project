package startup

import (
	rating "common/proto/rating-service/pb"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"log"
	"net"
	"rating-service/handler"
	"rating-service/model"
	"rating-service/repository"
	"rating-service/service"
	"rating-service/startup/config"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	reservationEndpoint := fmt.Sprintf("%s:%s", server.config.ReservationDomain, server.config.ReservationPort)
	accommodationEndpoint := fmt.Sprintf("%s:%s", server.config.AccommodationDomain, server.config.AccommodationPort)
	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserDomain, server.config.UserPort)
	ratingRepo := server.initRatingRepository(mongoClient)
	ratingService := server.initRatingService(ratingRepo, reservationEndpoint, accommodationEndpoint, userEndpoint)
	ratingHandler := server.initRatingHandler(ratingService)
	server.startGrpcServer(ratingHandler)

}
func (server *Server) initMongoClient() *mongo.Client {
	client, err := repository.GetClient(server.config.RatingDBDomain, server.config.RatingDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initRatingRepository(client *mongo.Client) model.RatingStore {
	store := repository.NewRatingRepo(client)
	store.DeleteAll()
	for _, RatingHost := range ratingsHost {
		err := store.InsertRatingHost(RatingHost)
		if err != nil {
			log.Fatal(err)
		}
	}
	for _, RatingAccommodation := range ratingsAccommodation {
		err := store.InsertRatingAccommodation(RatingAccommodation)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initRatingService(store model.RatingStore, reservationClientAddress string, accommodationClientAddress string, userClientAddress string) *service.RatingService {
	return service.NewRatingService(store, reservationClientAddress, accommodationClientAddress, userClientAddress)
}

// func (server *Server) initSubscriber(subject, queueGroup string) saga.Subscribe

func (server *Server) initRatingHandler(service *service.RatingService) *handler.RatingHandler {
	return handler.NewRatingHandler(service)
}

func (server *Server) startGrpcServer(ratingHandler *handler.RatingHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	rating.RegisterRatingServiceServer(grpcServer, ratingHandler)
	fmt.Println("Serving...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
