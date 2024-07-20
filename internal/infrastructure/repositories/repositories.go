package repositories

import "home-manager/server/internal/core/entities"

type ChoreRepo interface {
	GetAll() ([]*entities.Chore, error)
	GetById(id int32) (entities.Chore, error)
	Create(chore *entities.Chore) (entities.Chore, error)
}

type UserRepo interface {
	GetAll() ([]*entities.User, error)
	GetById(id int32) (entities.User, error)
	Create(user *entities.User) (entities.User, error)
}