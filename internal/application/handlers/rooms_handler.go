package handlers

import (
	"context"
	pb "home-manager/server/homemanager/gen"
	"home-manager/server/internal/application/mappers"
	"home-manager/server/internal/core/entities"
	"home-manager/server/internal/infrastructure/services"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RoomsHandler struct {
	roomsService services.RoomService
	pb.UnimplementedRoomServiceServer
}

func NewRoomsServer(grpcServer *grpc.Server, roomService *services.RoomService) {
	roomsGrpc := &RoomsHandler{roomsService: *roomService}
	
	pb.RegisterRoomServiceServer(grpcServer, roomsGrpc)
}

func (rh *RoomsHandler) Create(ctx context.Context, req *pb.CreateRoomRequest) (*pb.CreateRoomResponse, error) {
	newRoom, err := entities.NewRoom(0, req.Name)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	room, err := rh.roomsService.Create(newRoom)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return mappers.ToCreateRoomResponse(room), nil
}
