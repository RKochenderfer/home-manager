package handlers

import (
	"context"
	pb "home-manager/server/homemanager/gen"

	"google.golang.org/grpc"
)

type RoomsHandler struct {
	pb.UnimplementedRoomServiceServer
}

func NewRoomsServer(grpcServer *grpc.Server) {
	roomsHandler := &RoomsHandler{}
	pb.RegisterRoomServiceServer(grpcServer, roomsHandler)
}

func (*RoomsHandler) Create(ctx context.Context, req *pb.CreateRoomRequest) (*pb.CreateRoomResponse, error) {
	test := pb.CreateRoomResponse{
		Id:   1,
		Name: req.GetName(),
	}
	return &test, nil
}
