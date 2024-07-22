package services

import "home-manager/server/internal/core/entities"

type ChoresService interface {
	GetAll() ([]*entities.Chore, error)
	GetById(id int32) (entities.Chore, error)
	Create(chore *entities.Chore) (entities.Chore, error)
}

type UsersService interface {
	GetAll() ([]*entities.User, error)
	GetById(id int32) (entities.User, error)
	Create(user *entities.User) (entities.User, error)
}

type RoomService interface {
	GetAll() ([]*entities.Room, error)
	GetById(id int32) (entities.Room, error)
	Create(user *entities.Room) (entities.Room, error)
}
