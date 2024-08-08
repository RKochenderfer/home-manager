package application

import (
	pb "home-manager/server/homemanager/gen"
	"home-manager/server/internal/application/handlers"
	"home-manager/server/internal/infrastructure/db"
	"home-manager/server/internal/infrastructure/repositories"
	"home-manager/server/internal/infrastructure/services"

	"google.golang.org/grpc"
)

func SetupServers(s *grpc.Server, db *db.Database) {
	SetupUserServer(s, db)
	pb.RegisterRoomServiceServer(s, &handlers.RoomsHandler{})
}

func SetupUserServer(s *grpc.Server, db *db.Database) {
	repo := repositories.NewSqliteUserRepo(db)
	us, _ := services.NewUsersService(repo)

	handlers.NewUsersServer(s, &us)
}
