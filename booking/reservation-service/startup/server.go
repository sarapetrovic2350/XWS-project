package startup

import (
	reservation "common/proto/reservation-service/pb"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"log"
	"net"
	"reservation-service/handler"
	"reservation-service/model"
	"reservation-service/repository"
	"reservation-service/service"
	"reservation-service/startup/config"
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
	accommodationEndpoint := fmt.Sprintf("%s:%s", server.config.AccommodationDomain, server.config.AccommodationPort)
	reservationRepo := server.initReservationRepository(mongoClient)
	reservationService := server.initReservationService(reservationRepo, accommodationEndpoint)
	reservationHandler := server.initReservationHandler(reservationService)
	server.startGrpcServer(reservationHandler)

}
func (server *Server) initMongoClient() *mongo.Client {
	client, err := repository.GetClient(server.config.ReservationDBDomain, server.config.ReservationDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initReservationRepository(client *mongo.Client) model.ReservationStore {
	store := repository.NewReservationRepo(client)
	store.DeleteAll()
	for _, Reservation := range reservations {
		err := store.Insert(Reservation)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initReservationService(store model.ReservationStore, accommodationClientAddress string) *service.ReservationService {
	return service.NewReservationService(store, accommodationClientAddress)
}

// func (server *Server) initSubscriber(subject, queueGroup string) saga.Subscribe

func (server *Server) initReservationHandler(service *service.ReservationService) *handler.ReservationHandler {
	return handler.NewReservationHandler(service)
}

func (server *Server) startGrpcServer(reservationHandler *handler.ReservationHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	reservation.RegisterReservationServiceServer(grpcServer, reservationHandler)
	fmt.Println("Serving...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
