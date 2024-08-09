package application

import (
	"home-manager/server/internal/application/handlers"
	"home-manager/server/internal/infrastructure/db"
	"home-manager/server/internal/infrastructure/repositories"
	"home-manager/server/internal/infrastructure/services"

	"google.golang.org/grpc"
)

func SetupServers(s *grpc.Server, db *db.Database) {
	SetupUserServer(s, db)
	SetupRoomsServer(s, db)
}

func SetupUserServer(s *grpc.Server, db *db.Database) {
	repo := repositories.NewSqliteUserRepo(db)
	us, _ := services.NewUsersService(repo)

	handlers.NewUsersServer(s, &us)
}

func SetupRoomsServer(s *grpc.Server, db *db.Database) {
	repo := repositories.NewSqliteRoomRepo(db)
	rs, _ := services.NewRoomService(repo)

	handlers.NewRoomsServer(s, &rs)
}
