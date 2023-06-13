package startup

import (
	"accommodation-service/handler"
	"accommodation-service/model"
	"accommodation-service/repository"
	"accommodation-service/service"
	"accommodation-service/startup/config"
	accommodation "common/proto/accommodation-service/pb"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"log"
	"net"
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
	accommodationRepo := server.initAccommodationRepository(mongoClient)
	accommodationService := server.initAccommodationService(accommodationRepo, reservationEndpoint)
	accommodationHandler := server.initAccommodationHandler(accommodationService)
	server.startGrpcServer(accommodationHandler)

}
func (server *Server) initMongoClient() *mongo.Client {
	client, err := repository.GetClient(server.config.AccommodationDBDomain, server.config.AccommodationDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initAccommodationRepository(client *mongo.Client) model.AccommodationStore {
	store := repository.NewAccommodationRepo(client)
	store.DeleteAll()
	for _, Accommodation := range accommodations {
		err := store.Insert(Accommodation)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initAccommodationService(store model.AccommodationStore, reservationClientAddress string) *service.AccommodationService {
	return service.NewAccommodationService(store, reservationClientAddress)
}

// func (server *Server) initSubscriber(subject, queueGroup string) saga.Subscribe

func (server *Server) initAccommodationHandler(service *service.AccommodationService) *handler.AccommodationHandler {
	return handler.NewAccommodationHandler(service)
}

func (server *Server) startGrpcServer(accommodationHandler *handler.AccommodationHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	accommodation.RegisterAccommodationServiceServer(grpcServer, accommodationHandler)
	fmt.Println("Serving...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
