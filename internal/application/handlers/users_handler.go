package handlers

import (
	"context"
	"errors"
	pb "home-manager/server/homemanager/gen"
	"home-manager/server/internal/application/mappers"
	"home-manager/server/internal/core/entities/useraggregate"
	"home-manager/server/internal/core/internalerrors"
	"home-manager/server/internal/infrastructure/services"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type usersHandler struct {
	usersService services.UsersService
	pb.UnimplementedUsersServiceServer
}

func NewUsersServer(grpcServer *grpc.Server, userService *services.UsersService) {
	userGrpc := &usersHandler{usersService: *userService}

	pb.RegisterUsersServiceServer(grpcServer, userGrpc)
}

func (uh *usersHandler) Create(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	newUser, err := useraggregate.NewUser(req.GetName(), 0, useraggregate.Role(req.GetRole()))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	user, err := uh.usersService.Create(newUser)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return mappers.ToCreateUserResponse(user), nil
}

func (uh *usersHandler) Get(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	parsed, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	user, err := uh.usersService.GetById(parsed)

	if errors.Is(err, internalerrors.ErrNotFound) {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return mappers.ToGetUserResponse(user), nil
}

func (uh *usersHandler) GetAll(ctx context.Context, req *pb.Empty) (*pb.GetAllUsersResponse, error) {
	users, err := uh.usersService.GetAll()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return mappers.ToGetAllUsersResponse(&users), nil
}
