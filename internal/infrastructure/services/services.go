package services

import (
	"home-manager/server/internal/core/entities"
	"home-manager/server/internal/core/valueobjects"
)

type ChoresService interface {
	GetAll() ([]*valueobjects.Chore, error)
	GetById(id int32) (valueobjects.Chore, error)
	Create(chore *valueobjects.Chore) (valueobjects.Chore, error)
}

type UsersService interface {
	GetAll() ([]*entities.User, error)
	GetById(id int32) (*entities.User, error)
	Create(user *entities.User) (*entities.User, error)
}

type RoomService interface {
	GetAll() ([]*entities.Room, error)
	GetById(id int32) (*entities.Room, error)
	Create(user *entities.Room) (*entities.Room, error)
}
