package services

import (
	"home-manager/server/internal/core/entities"
	"home-manager/server/internal/infrastructure/repositories"
)

type usersService struct {
	repo repositories.UserRepo
}

func NewUsersService(repo repositories.UserRepo) (UsersService, error) {
	return &usersService{repo}, nil
}

func (u *usersService) Create(user *entities.User) (entities.User, error) {
	return u.repo.Create(user)
}

func (u *usersService) GetAll() ([]*entities.User, error) {
	return u.repo.GetAll()
}

func (u *usersService) GetById(id int32) (entities.User, error) {
	return u.repo.GetById(id)
}
