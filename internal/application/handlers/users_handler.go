package handlers

import (
	"context"
	pb "home-manager/server/homemanager/gen"

	"google.golang.org/grpc"
)

type UsersHandler struct {
	pb.UnimplementedUsersServiceServer
}

func NewUsersServer(grpcServer *grpc.Server) {
	userGrpc := &UsersHandler{}
	pb.RegisterUsersServiceServer(grpcServer, userGrpc)
}

func (*UsersHandler) Create(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	test := pb.CreateUserResponse{
		Id:          "1",
		Name:        req.GetName(),
		TotalPoints: 1,
		Role:        req.GetRole(),
	}
	return &test, nil
}
