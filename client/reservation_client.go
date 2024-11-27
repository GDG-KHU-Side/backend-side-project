package client

import (
	"context"
	"time"

	pb "github.com/GDG-KHU-Side/backend-side-project/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetReservations(restaurantID int64) ([]*pb.Reservation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := pb.NewReservationServiceClient(conn)
	resp, err := client.GetRestaurantReservations(ctx, &pb.GetReservationsRequest{
		RestaurantId: restaurantID,
	})
	if err != nil {
		return nil, err
	}

	return resp.Reservations, nil
}
