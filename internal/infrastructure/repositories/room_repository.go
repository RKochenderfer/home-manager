package repositories

import (
	"home-manager/server/internal/infrastructure/db"
	"home-manager/server/internal/infrastructure/db/models"
)

type sqliteRoomRepo struct {
	db *db.Database
}

func NewSqliteRoomRepo(db *db.Database) RoomRepo {
	return &sqliteRoomRepo{db}
}

// Create implements RoomRepo.
func (s *sqliteRoomRepo) Create(user *models.Room) (models.Room, error) {
	panic("unimplemented")
}

// GetAll implements RoomRepo.
func (s *sqliteRoomRepo) GetAll() ([]*models.Room, error) {
	panic("unimplemented")
}

// GetById implements RoomRepo.
func (s *sqliteRoomRepo) GetById(id int32) (models.Room, error) {
	panic("unimplemented")
}
