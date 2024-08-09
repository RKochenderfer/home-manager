package repositories

import (
	"home-manager/server/internal/core/internalerrors"
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
func (s *sqliteRoomRepo) Create(toAdd *models.Room) (*models.Room, error) {
	if result := s.db.Connection().Create(&toAdd); result.Error != nil {
		return nil, result.Error
	}

	return s.GetById(int32(toAdd.ID))
}

// GetAll implements RoomRepo.
func (s *sqliteRoomRepo) GetAll() ([]*models.Room, error) {
	var rooms []*models.Room

	if result := s.db.Connection().Find(&rooms); result.Error != nil {
		return rooms, result.Error
	}

	return rooms, nil
}

// GetById implements RoomRepo.
func (s *sqliteRoomRepo) GetById(id int32) (*models.Room, error) {
	var room models.Room
	if result := s.db.Connection().First(&room, id); result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, internalerrors.ErrNotFound
		}
		return nil, result.Error
	}

	return &room, nil
}
