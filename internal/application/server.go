package application

import (
	"flag"
	pb "home-manager/server/homemanager/gen"
	"home-manager/server/internal/application/handlers"
	"home-manager/server/internal/infrastructure/db"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "gRPC server port")
)

func SetupServers(db db.Database, s *grpc.Server) {
	pb.RegisterUsersServiceServer(s, &handlers.UsersHandler{})
	pb.RegisterRoomServiceServer(s, &handlers.RoomsHandler{})
}
