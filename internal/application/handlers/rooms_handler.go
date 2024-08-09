package handlers

import (
	"context"
	"errors"
	pb "home-manager/server/homemanager/gen"
	"home-manager/server/internal/application/mappers"
	"home-manager/server/internal/core/entities"
	"home-manager/server/internal/core/internalerrors"
	"home-manager/server/internal/infrastructure/services"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type roomsHandler struct {
	roomsService services.RoomService
	pb.UnimplementedRoomServiceServer
}

func NewRoomsServer(grpcServer *grpc.Server, roomService *services.RoomService) {
	roomsGrpc := &roomsHandler{roomsService: *roomService}

	pb.RegisterRoomServiceServer(grpcServer, roomsGrpc)
}

func (rh *roomsHandler) Create(ctx context.Context, req *pb.CreateRoomRequest) (*pb.CreateRoomResponse, error) {
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

func (rh *roomsHandler) Get(ctx context.Context, req *pb.GetRoomRequest) (*pb.GetRoomResponse, error) {
	room, err := rh.roomsService.GetById(req.GetId())
	if errors.Is(err, internalerrors.ErrNotFound) {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	return mappers.ToGetRoomResponse(room), nil
}

func (rh *roomsHandler) GetAll(ctx context.Context, req *pb.Empty) (*pb.GetAllRoomsResponse, error) {
	rooms, err := rh.roomsService.GetAll()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return mappers.ToGetAllRoomsResponse(rooms), nil
}
