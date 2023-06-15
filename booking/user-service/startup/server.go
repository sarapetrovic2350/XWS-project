package startup

import (
	user "common/proto/user-service/pb"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"log"
	"net"
	"user-service/handler"
	"user-service/model"
	"user-service/repository"
	"user-service/service"
	"user-service/startup/config"
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
	ratingEndpoint := fmt.Sprintf("%s:%s", server.config.RatingDomain, server.config.RatingPort)
	userRepo := server.initUserRepository(mongoClient)
	userService := server.initUserService(userRepo, reservationEndpoint, accommodationEndpoint, ratingEndpoint)
	userHandler := server.initUserHandler(userService)
	server.startGrpcServer(userHandler)

}
func (server *Server) initMongoClient() *mongo.Client {
	client, err := repository.GetClient(server.config.UserDBDomain, server.config.UserDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initUserRepository(client *mongo.Client) model.UserStore {
	store := repository.NewUserRepo(client)
	store.DeleteAll()
	for _, User := range users {
		err := store.Insert(User)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initUserService(store model.UserStore, reservationClientAddress string, accommodationClientAddress string, ratingClientAddress string) *service.UserService {
	return service.NewUserService(store, reservationClientAddress, accommodationClientAddress, ratingClientAddress)
}

// func (server *Server) initSubscriber(subject, queueGroup string) saga.Subscribe

func (server *Server) initUserHandler(service *service.UserService) *handler.UserHandler {
	return handler.NewUserHandler(service)
}

func (server *Server) startGrpcServer(userHandler *handler.UserHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	user.RegisterUserServiceServer(grpcServer, userHandler)
	fmt.Println("Serving...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
