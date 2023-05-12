package repository

import (
	"context"
	"fmt"
	reservation "github.com/sarapetrovic2350/XWS-project/booking/common/proto/reservation-service/pb"
	"log"

	//reservation "github.com/sarapetrovic2350/XWS-project/booking/common/proto/reservation-service/pb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetClient(host, port string) (*mongo.Client, error) {
	uri := fmt.Sprintf("mongodb://%s:%s/", host, port)
	options := options.Client().ApplyURI(uri)
	return mongo.Connect(context.TODO(), options)
}

func NewReservationClient(address string) reservation.ReservationServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Booking service: %v", err)
	}
	return reservation.NewReservationServiceClient(conn)
}

//func getConnection(address string) (*grpc.ClientConn, error) {
//	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
//}
