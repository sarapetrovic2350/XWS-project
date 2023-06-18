package startup

import (
	user "common/proto/user-service/pb"
	saga "common/saga/messaging"
	"common/saga/messaging/nats"
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

const (
	QueueGroup = "user_service"
)

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	userRepo := server.initUserRepository(mongoClient)

	reservationEndpoint := fmt.Sprintf("%s:%s", server.config.ReservationDomain, server.config.ReservationPort)
	accommodationEndpoint := fmt.Sprintf("%s:%s", server.config.AccommodationDomain, server.config.AccommodationPort)
	ratingEndpoint := fmt.Sprintf("%s:%s", server.config.RatingDomain, server.config.RatingPort)
	commandPublisher := server.initPublisher(server.config.DeleteUserCommandSubject)
	replySubscriber := server.initSubscriber(server.config.DeleteUserReplySubject, QueueGroup)
	deleteUserOrchestrator := server.initDeleteUserOrchestrator(commandPublisher, replySubscriber)
	userService := server.initUserService(userRepo, reservationEndpoint, accommodationEndpoint, ratingEndpoint, deleteUserOrchestrator)

	commandSubscriber := server.initSubscriber(server.config.DeleteUserCommandSubject, QueueGroup)
	replyPublisher := server.initPublisher(server.config.DeleteUserReplySubject)
	server.initDeleteUserHandler(userService, replyPublisher, commandSubscriber)

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

func (server *Server) initPublisher(subject string) saga.Publisher {
	publisher, err := nats.NewNATSPublisher(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject)
	if err != nil {
		log.Fatal(err)
	}
	return publisher
}

func (server *Server) initSubscriber(subject, queueGroup string) saga.Subscriber {
	subscriber, err := nats.NewNATSSubscriber(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject, queueGroup)
	if err != nil {
		log.Fatal(err)
	}
	return subscriber
}

func (server *Server) initDeleteUserOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) *service.DeleteUserOrchestrator {
	orchestrator, err := service.NewDeleteUserOrchestrator(publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
	return orchestrator
}

func (server *Server) initUserService(store model.UserStore, reservationClientAddress string, accommodationClientAddress string, ratingClientAddress string, orchestrator *service.DeleteUserOrchestrator) *service.UserService {
	return service.NewUserService(store, reservationClientAddress, accommodationClientAddress, ratingClientAddress, orchestrator)
}

func (server *Server) initDeleteUserHandler(service *service.UserService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := handler.NewDeleteUserCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

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
