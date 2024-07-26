package repositories

import (
	"home-manager/server/internal/infrastructure/db/models"

	"github.com/google/uuid"
)

type ChoreRepo interface {
	GetAll() ([]*models.Chore, error)
	GetById(id int32) (models.Chore, error)
	Create(chore *models.Chore) (models.Chore, error)
}

type UserRepo interface {
	GetAll() ([]*models.User, error)
	GetById(id uuid.UUID) (models.User, error)
	Create(user *models.User) (models.User, error)
}

type RoomRepo interface {
	GetAll() ([]*models.Room, error)
	GetById(id int32) (models.Room, error)
	Create(room *models.Room) (models.Room, error)
}