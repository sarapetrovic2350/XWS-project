package startup

import (
	cfg "api-gateway/startup/config"
	accommodationGw "common/proto/accommodation-service/pb"
	ratingGw "common/proto/rating-service/pb"
	reservationGw "common/proto/reservation-service/pb"
	userGw "common/proto/user-service/pb"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	config *cfg.Config
	mux    *runtime.ServeMux
}

func NewServer(config *cfg.Config) *Server {
	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
	}
	server.initHandlers()
	return server
}
func (server *Server) initHandlers() {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	//NOTE: My endpoints
	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserDomain, server.config.UserPort)
	fmt.Print("userEndpoint: ")
	fmt.Println(userEndpoint)
	err := userGw.RegisterUserServiceHandlerFromEndpoint(context.TODO(), server.mux, userEndpoint, opts)
	if err != nil {
		panic(err)
	}

	reservationEndpoint := fmt.Sprintf("%s:%s", server.config.ReservationDomain, server.config.ReservationPort)
	fmt.Print("reservationEndpoint: ")
	fmt.Println(reservationEndpoint)
	err = reservationGw.RegisterReservationServiceHandlerFromEndpoint(context.TODO(), server.mux, reservationEndpoint, opts)
	if err != nil {
		panic(err)
	}
	fmt.Println("Handlers initalized")

	accommodationEndpoint := fmt.Sprintf("%s:%s", server.config.AccommodationDomain, server.config.AccommodationPort)
	fmt.Print("accommodationEndpoint: ")
	fmt.Println(accommodationEndpoint)
	err = accommodationGw.RegisterAccommodationServiceHandlerFromEndpoint(context.TODO(), server.mux, accommodationEndpoint, opts)
	if err != nil {
		panic(err)
	}
	fmt.Println("Handlers initalized")

	ratingEndpoint := fmt.Sprintf("%s:%s", server.config.RatingDomain, server.config.RatingPort)
	fmt.Print("ratingEndpoint: ")
	fmt.Println(ratingEndpoint)
	err = ratingGw.RegisterRatingServiceHandlerFromEndpoint(context.TODO(), server.mux, ratingEndpoint, opts)
	if err != nil {
		panic(err)
	}
	fmt.Println("Handlers initalized")

}

func (server *Server) Start() {
	handler := server.getHandlerCORSWrapped()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), handler))
}

func (server *Server) getHandlerCORSWrapped() http.Handler {
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{server.config.AllowedCorsOrigin},
	})
	handler := corsMiddleware.Handler(server.mux)
	return handler
}
