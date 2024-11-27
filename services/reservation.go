package services

import (
	"context"

	pb "github.com/GDG-KHU-Side/backend-side-project/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb" // timestamppb 패키지 추가
)

type ReservationServer struct {
	pb.UnimplementedReservationServiceServer
	service *ReservationService
}

func NewReservationServer(service *ReservationService) *ReservationServer {
	return &ReservationServer{service: service}
}

func (s *ReservationServer) GetRestaurantReservations(ctx context.Context, req *pb.GetReservationsRequest) (*pb.GetReservationsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be nil")
	}

	reservations, err := s.service.GetAllRestaurants(req.RestaurantId)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get reservations: "+err.Error())
	}

	protoReservations := make([]*pb.Reservation, 0, len(reservations))
	for _, r := range reservations {
		protoReservations = append(protoReservations, &pb.Reservation{
			Id:               r.ID,
			RestId:           r.RestID,
			Count:            int32(r.Count),
			CustomerPhoneNum: r.CustomerPhoneNum,
			CreatedAt:        timestamppb.New(r.CreatedAt),
			IsEntry:          int32(r.IsEntry),
		})
	}

	return &pb.GetReservationsResponse{
		Reservations: protoReservations,
	}, nil
}
