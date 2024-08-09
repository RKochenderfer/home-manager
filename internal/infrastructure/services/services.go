package services

import (
	"home-manager/server/internal/core/entities"
	"home-manager/server/internal/core/entities/useraggregate"
	"home-manager/server/internal/core/valueobjects"

	"github.com/google/uuid"
)

type ChoresService interface {
	GetAll() ([]valueobjects.Chore, error)
	GetById(id int32) (valueobjects.Chore, error)
	Create(chore *valueobjects.Chore) (valueobjects.Chore, error)
}

type UsersService interface {
	GetAll() ([]useraggregate.User, error)
	GetById(id uuid.UUID) (*useraggregate.User, error)
	Create(user *useraggregate.User) (*useraggregate.User, error)
}

type RoomService interface {
	GetAll() ([]*entities.Room, error)
	GetById(id int32) (*entities.Room, error)
	Create(user *entities.Room) (*entities.Room, error)
}
